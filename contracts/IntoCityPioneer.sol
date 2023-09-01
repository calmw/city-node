// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./IntoCity.sol";
import "./RoleAccess.sol";
import "./IntoCityNodeVote.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract IntoCityPioneer is RoleAccess, Initializable {
    enum LifeTime {
        none,
        firstMonth,
        secondMonth,
        thirdMonth,
        fourToSixMonth,
        moreThanSixMonth
    }

    // 每日新增质押事件
    event DailyIncreaseDelegateRecord(
        address pioneerAddress, // 先锋地址
        bytes32 cityId, //城市ID
        uint256 amount, // 新增质押金额
        uint256 ctime // 新增质押的时间
    );

    // 保证金退还事件
    event EarnestMoneyRecord(
        address pioneerAddress, // 先锋地址
        uint256 amount // 退还保证金的额度
    );

    // 城市先锋奖励事件
    event DailyRewardRecord(
        address pioneerAddress, // 先锋地址
        uint256 toxReward, // 赠送质押包奖励的额度
        uint256 foundsReward, // 社交基金奖励的额度
        uint256 delegateReward // 新增质押奖励的额度
    );

    // 城市先锋提取奖励事件
    event WithdrawalRewardRecord(
        address pioneerAddress, // 先锋地址
        uint256 reward // 提现的奖励数量
    );

    struct Pioneer {
        address pioneerAddress;
        uint256 day; // 成为城市节点的时间(天数)
        uint256 firstMonthDelegate; // 第1个月累计质押量
        uint256 secondMonthDelegate; // 第2个月累计质押量
        uint256 thirdMonthDelegate; // 第3个月累计质押量
        bool firstMonthReturnEarnest; // 第1个月是否退了保证金
        bool secondMonthReturnEarnest; // 第2个月是否退了保证金
        bool thirdMonthReturnEarnest; // 第3个月是否退了保证金
        uint256 cityLevel; // 所在城市等级
        LifeTime lifeTime; // 城市先锋生命周期
        bool assessmentMonthStatus; // 按月考核状态
        bool assessmentStatus; // 最终考核状态
    }

    address public TOX; // TOX代币合约地址
    address public cityAddress; // 城市合约地址
    uint256 public assessmentPassed; // 城市先锋第一个月直接通过考核的条件（点数）

    // 城市等级 => (月份 => 质押点数)
    mapping(uint256 => mapping(uint256 => uint256))public  assessmentCriteria; // 城市先锋考核标准
    // 城市等级 => (月份 => 退还标准)
    mapping(uint256 => mapping(uint256 => uint256))public  assessmentReturnCriteria; // 城市先锋保证金退还标准
    // 城市等级 => (月份 => 退还比例)
    mapping(uint256 => mapping(uint256 => uint256))public  assessmentReturnRate; // 城市先锋保证金退还比例
    // 先锋地址 => 先锋信息， 先锋信息
    mapping(address => Pioneer) public pioneerInfo;
    // 先锋地址 => 先锋福利包收益
    mapping(address => uint256) public benefitPackageReward;
    // 先锋地址 => 先锋社交基金收益
    mapping(address => uint256) public fundsReward;
    // 先锋地址 => 先锋新增质押收益
    mapping(address => uint256) public delegateReward;
    // 先锋地址 => 先锋退还的保证金
    mapping(address => uint256) public earnestMoneyReward;

    // 先锋地址 => 每天新增质押[]
    mapping(address => uint256[]) public pioneerDelegateInfo;

    mapping(bytes32 => uint256)public  rewardsForSocialFunds; // 城市每日社交基金
    mapping(bytes32 => uint256)public  latestDayDelegate; // 城市每日新增质押

    function initialize() public initializer {
        _addAdmin(msg.sender);
    }

    // 管理员设置TOX代币地址
    function adminSetTOXAddress(address TOX_) public onlyAdmin {
        TOX = TOX_;
    }

    // 管理员设置城市合约地址
    function adminSetCityAddress(address cityAddress_) public onlyAdmin {
        cityAddress = cityAddress_;
    }

    // 管理员设置城市先锋第一个月直接通过考核的最大阀值（点数）
    function adminSetAssessmentPassed(uint256 assessmentPassed_) public onlyAdmin {
        assessmentPassed = assessmentPassed_;
    }

    // 管理员设置考核标准,月份为1，2，3
    function adminSetAssessmentCriteria(uint256 cityLevel_, uint256 month_, uint256 point_) public onlyAdmin {
        assessmentCriteria[cityLevel_][month_] = point_;
    }

    // 管理员设置先锋每天新增质押量
    function setPioneerDelegate(address pioneerAddress_, uint256 amount_) public onlyAdmin {
        Pioneer storage pioneer = pioneerInfo[pioneerAddress_];
        uint256 day = getDay() - pioneer.day; // 第几天质押
        pioneerDelegateInfo[pioneerAddress_][day] = amount_;
        pioneer.day = day;
        // 更新城市先锋信息
        if (day > 180) {
            pioneer.lifeTime = LifeTime.moreThanSixMonth;
        } else if (day > 90) {
            pioneer.lifeTime = LifeTime.fourToSixMonth;
        } else if (day > 60) {
            pioneer.lifeTime = LifeTime.thirdMonth;
            pioneer.thirdMonthDelegate += amount_;
        } else if (day > 30) {
            pioneer.lifeTime = LifeTime.secondMonth;
            pioneer.secondMonthDelegate += amount_;
        } else {
            pioneer.lifeTime = LifeTime.firstMonth;
            pioneer.firstMonthDelegate += amount_;
        }

        IntoCity city = IntoCity(cityAddress);
        emit DailyIncreaseDelegateRecord(
            pioneer.pioneerAddress,
            city.pioneerCity(pioneerAddress_),
            amount_,
            block.timestamp
        );
    }

    function getDay() public view returns (uint256){
        uint day = block.timestamp / 86400;
        return uint256(day);
    }

    // 缴纳保证金
    function depositEarnestMoney() public {
        IntoCity city = IntoCity(cityAddress);
        // 查询调用者地址是否是城市先锋
        bytes32 cityId = city.pioneerCity(msg.sender);
        require(cityId != bytes32(0), "you are not pioneer"); // 不是城市先锋

        // 查询该城市先锋对应的城市，根据城市查询需要缴纳的保证金数量
        uint256 earnestMoney = city.earnestMoney(cityId);
        IERC20 TOXContract = IERC20(TOX);
        uint256 userBalance = TOXContract.balanceOf(msg.sender);
        require(userBalance >= earnestMoney, "your balance is insufficient"); // 余额不足
        TOXContract.transfer(address(this), earnestMoney);
        pioneerInfo[msg.sender].pioneerAddress = msg.sender;
        pioneerInfo[msg.sender].assessmentMonthStatus = true;
        pioneerInfo[msg.sender].day = getDay();
        pioneerInfo[msg.sender].cityLevel = city.cityLevel(cityId);
        pioneerInfo[msg.sender].lifeTime = LifeTime.firstMonth;
    }

    // 每天执行一次，计算奖励和考核
    function dailyTask() public onlyAdmin {
        IntoCity city = IntoCity(cityAddress);
        for (uint256 i = 0; i < city.getPioneerCityNumber(); i++) {
            // 考核
            checkPioneer(city.pioneerCityIds(i), city.cityPioneer(city.pioneerCityIds(i)));
            // 奖励
            reward(city.pioneerCityIds(i), city.cityPioneer(city.pioneerCityIds(i)));
            // 更新城市每天累计质押最大值
            updateCityMaxDailyDelegate();
        }
    }

    // 更新城市每天累计最大值
    function updateCityMaxDailyDelegate() public {
        IntoCity city = IntoCity(cityAddress);
        uint256 today = getDay();
        uint256 allCityNumber = city.getAllCityNumber();

        for (uint256 i = 0; i < allCityNumber; i++) {
            uint256 yesterdayDelegate = city.cityDelegateRecord(city.allCityIds(i), today - 1);
            uint256 maxDelegate = city.cityMaxDelegate(city.allCityIds(i), 2);
            if (yesterdayDelegate > maxDelegate) {
                city.setCityMaxDelegate(city.allCityIds(i), yesterdayDelegate, today - 1);
            }
        }
    }

    // 检测考核与保证金退还,每日执行一次,考核失败的城市，可以参与城市节点竞选
    function checkPioneer(bytes32 cityId, address pioneerAddress_) private {
        Pioneer storage pioneer = pioneerInfo[pioneerAddress_];
        if (pioneer.assessmentStatus == true) {
            return;
        }
        IntoCity city = IntoCity(cityAddress);
        uint256 day = getDay() - pioneer.day;

        // 前三个月考核
        assessmentPioneer(cityId, pioneer, city, day);

        // 计算退还保证金额度,并更新退还状态
        uint256 earnestMoneyResult = calculateRefund(cityId, pioneer, city, day);

        // 退还保证金操作
        if (earnestMoneyResult > 0) {
            earnestMoneyReward[pioneer.pioneerAddress] += earnestMoneyResult;
            emit EarnestMoneyRecord(
                pioneerAddress_,
                earnestMoneyResult
            );
        }
    }

    // 前三个月考核
    function assessmentPioneer(bytes32 cityId, Pioneer storage pioneer, IntoCity city, uint256 day) private {
        uint256 assessmentCriteriaThreshold; // 考核标准金额
        if (day == 90) {
            assessmentCriteriaThreshold = assessmentCriteria[pioneer.cityLevel][3] * 100;
            if (pioneer.thirdMonthDelegate < assessmentCriteriaThreshold) {
                pioneer.assessmentMonthStatus = false;
                city.setCityPioneerAssessment(cityId); // 将该城市设置为先锋计划洛选城市
            }
        } else if (day == 60) {
            assessmentCriteriaThreshold = assessmentCriteria[pioneer.cityLevel][2] * 100;
            if (pioneer.secondMonthDelegate < assessmentCriteriaThreshold) {
                pioneer.assessmentMonthStatus = false;
                city.setCityPioneerAssessment(cityId); // 将该城市设置为先锋计划洛选城市
            }
        } else if (day == 30) {
            assessmentCriteriaThreshold = assessmentCriteria[pioneer.cityLevel][1] * 100;
            if (pioneer.firstMonthDelegate < assessmentCriteriaThreshold) {
                pioneer.assessmentMonthStatus = false;
                city.setCityPioneerAssessment(cityId); // 将该城市设置为先锋计划洛选城市
            }
            // 检测是否满足直接考核通过
            if (pioneer.firstMonthDelegate >= assessmentPassed) {
                pioneer.assessmentStatus = true;
            }
        }
    }

    // 计算退还保证金额度,并更新退还状态
    function calculateRefund(bytes32 cityId, Pioneer storage pioneer, IntoCity city, uint256 day) private returns (uint256){

        uint256 cityLevel = city.cityLevel(cityId); // 城市等级
        uint256 earnestMoney = city.earnestMoney(cityId); // 城市保证金
        uint256 earnestMoneyReturn; // 退还保证金的金额
        if (day == 90) {
            if (pioneer.secondMonthDelegate >= assessmentReturnCriteria[cityLevel][3] * 100) {
                uint256 rate = assessmentReturnRate[cityLevel][3];
                if (pioneer.firstMonthReturnEarnest) {
                    rate -= assessmentReturnRate[cityLevel][1];
                }
                if (pioneer.secondMonthReturnEarnest) {
                    rate -= assessmentReturnRate[cityLevel][2];
                }
                earnestMoneyReturn = earnestMoney * rate / 100;
                pioneer.thirdMonthReturnEarnest = true;
            }
        } else if (day == 60) {
            if (pioneer.secondMonthDelegate >= assessmentReturnCriteria[cityLevel][2] * 100) {
                uint256 rate = assessmentReturnRate[cityLevel][2];
                if (pioneer.firstMonthReturnEarnest) {
                    rate -= assessmentReturnRate[cityLevel][1];
                }
                earnestMoneyReturn = earnestMoney * rate / 100;
                pioneer.secondMonthReturnEarnest = true;
            }
        } else if (day == 30) {
            if (pioneer.firstMonthDelegate >= assessmentReturnCriteria[cityLevel][1] * 100) {
                uint256 rate = assessmentReturnRate[cityLevel][1];
                earnestMoneyReturn = earnestMoney * rate / 100;
                pioneer.firstMonthReturnEarnest = true;
            }
        }

        return earnestMoneyReturn;
    }

    // 每日奖励
    function reward(bytes32 cityId, address pioneerAddress_) private {
        Pioneer storage pioneer = pioneerInfo[pioneerAddress_];
        // 考核状态（按月考核）不正常，不发放奖励
        if (!pioneer.assessmentMonthStatus) {
            return;
        }
        // 任期结束，不发放奖励
        if (pioneer.lifeTime == LifeTime.moreThanSixMonth) {
            return;
        }

        // 福利包奖励
        uint bonus = 93333333333333333333;
        benefitPackageReward[pioneerAddress_] += bonus;

        // 社交基金5%奖励
        fundsReward[pioneerAddress_] += rewardsForSocialFunds[cityId] * 5 / 100;

        // 该城市新增质押量1%奖励，不累加
        uint256 day = getDay() - pioneer.day;
        delegateReward[pioneerAddress_] += pioneerDelegateInfo[pioneerAddress_][day] / 100;

        emit DailyRewardRecord(
            pioneerAddress_,
            bonus,
            fundsReward[pioneerAddress_],
            delegateReward[pioneerAddress_]
        );
    }

    // 用户提取福利包奖励
    function withdrawalBenefitPackageReward(uint256 amount_) public {
        // 判断是否是先锋
        require(pioneerInfo[msg.sender].day > 0, "you are not pioneer");
        // 检测余额
        require(amount_ <= benefitPackageReward[msg.sender], "insufficient balance");

        emit WithdrawalRewardRecord(
            msg.sender,
            amount_
        );
    }

    // 用户提取社交基金奖励
    function withdrawalFundsRewardReward(uint256 amount_) public {
        // 判断是否是先锋
        require(pioneerInfo[msg.sender].day > 0, "you are not pioneer");
        // 检测余额
        require(amount_ <= fundsReward[msg.sender], "insufficient balance");

        emit WithdrawalRewardRecord(
            msg.sender,
            amount_
        );
    }

    // 用户提取新增质押奖励
    function withdrawalDelegateRewardReward(uint256 amount_) public {
        // 判断是否是先锋
        require(pioneerInfo[msg.sender].day > 0, "you are not pioneer");
        // 检测余额
        require(amount_ <= delegateReward[msg.sender], "insufficient balance");

        emit WithdrawalRewardRecord(
            msg.sender,
            amount_
        );
    }

    // 用户提取退还的保证金
    function withdrawalEarnestMoneyReward(uint256 amount_) public {
        // 判断是否是先锋
        require(pioneerInfo[msg.sender].day > 0, "you are not pioneer");
        // 检测余额
        require(amount_ <= earnestMoneyReward[msg.sender], "insufficient balance");

        emit WithdrawalRewardRecord(
            msg.sender,
            amount_
        );
    }

    // 获取先锋所需缴纳的保证金
    function getEarnestMoney() public view returns (uint256){
        IntoCity city = IntoCity(cityAddress);
        // 获取先锋所在城市
        bytes32 cityId = city.pioneerCity(msg.sender);
        // 获取该城市所需要缴纳的保证金
        return city.earnestMoney(cityId);
    }

    // 判断一个地址是否是先锋
    function isPioneer(address user) public view returns (bool){
        return pioneerInfo[user].day > 0;
    }

}