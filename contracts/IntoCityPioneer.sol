// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./IntoCity.sol";
import "./RoleAccess.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

interface IntoMining {
    function addBalanceWithTypes(address sender, uint256 amount, uint256 types) external; // 增加用户合约余额
}

interface IERC20 {
    function balanceOf(address account) external view returns (uint256);

    function transfer(address to, uint256 amount) external returns (bool);

    function transferFrom(address from, address to, uint256 amount) external returns (bool);
}

contract IntoCityPioneer is RoleAccess, Initializable {

    // 每日新增质押事件
    event DailyIncreaseDelegateRecord(
        address pioneerAddress, // 先锋地址
        bytes32 cityId, //城市ID
        uint256 amount, // 新增质押金额
        uint256 ctime // 新增质押的时间
    );

    // 保证金退还事件
    event SuretyRecord(
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
        uint256 reward, // 提现的奖励数量
        uint256 rewardType // 提现的奖励的类型. 1 福袋，2 社交基金， 3 质押奖励，4 保证金奖励
    );

    struct Pioneer {
        address pioneerAddress;
        uint256 ctime; // 成为城市节点的时间戳
        uint256 cityLevel; // 所在城市等级
        bool assessmentMonthStatus; // 按月考核状态
        bool assessmentStatus; // 最终考核状态
        bool returnSuretyStatus; // 保证金退还状态
        uint256 returnSuretyRate; // 保证金退还比例
        uint256 returnSuretyTime; // 保证金退还时间
    }

    address public TOXAddress; // TOX代币合约地址
    address public miningAddress; // 用户增加合约余额合约
    address public cityAddress; // 城市合约地址

    // 城市等级 => (月份 => 质押点数)
    mapping(uint256 => mapping(uint256 => uint256))public  assessmentCriteria; // 城市先锋考核标准
    // 城市等级 => (索引 => 退还标准)，索引1，2，3为第一个月的，索引4，5，6为第二个月的
    mapping(uint256 => mapping(uint256 => uint256))public  assessmentReturnCriteria; // 城市先锋保证金退还标准
    // 城市等级 => (索引 => 退还比例)，索引1，2，3为第一个月的，索引4，5，6为第二个月的
    mapping(uint256 => mapping(uint256 => uint256))public  assessmentReturnRate; // 城市先锋保证金退还比例
    // 先锋地址 => 先锋信息， 先锋信息
    mapping(address => Pioneer) public pioneerInfo;
    // 先锋地址 => 先锋福利包收益
    mapping(address => uint256) public benefitPackageReward;
    // 先锋地址 => 已领取的先锋福利包收益
    mapping(address => uint256) public benefitPackageRewardReceived;
    // 先锋地址 => 先锋社交基金收益
    mapping(address => uint256) public fundsReward;
    // 先锋地址 => 已领取的先锋社交基金收益
    mapping(address => uint256) public fundsRewardReceived;
    // 先锋地址 => 先锋新增质押收益
    mapping(address => uint256) public delegateReward;
    // 先锋地址 => 已领取的先锋新增质押收益
    mapping(address => uint256) public delegateRewardReceived;
    // 先锋地址 => 先锋可以退的保证金
    mapping(address => uint256) public suretyReward;
    // 先锋地址 => (月份=>退的比例)，月份为1和2，比例为整数
    mapping(address => mapping(uint256 => uint256)) public alreadyRewardRate;
    // 先锋地址 => 先锋已经退还的保证金
    mapping(address => uint256) public suretyRewardRecord;
    // 先锋城市ID => 先锋考核失败时候的累计新增质押量（绑定城市的累计质押量）
    mapping(bytes32 => uint256) public failedDelegate;
    // 每天定时任务执行状态
    mapping(uint256 => bool)public  dailyTaskStatus;
    mapping(address => bool)public benefitPackageRewardStatus; // 用户福袋奖励提取状态
    mapping(address => bool)public  fundsRewardStatus; // 用户社交基金奖励提取状态
    mapping(address => bool)public  delegateRewardStatus; // 用户新增质押奖励提取状态
    // 城市先锋地址=>(天数=>是否执行过)
    mapping(address => mapping(uint256 => bool))public  checkPioneerStatus; //每天一次考核
    // 城市先锋地址=>(天数=>是否执行过)
    mapping(address => mapping(uint256 => bool))public  checkRewardStatus; //每天一次奖励

    function initialize() public initializer {
        _addAdmin(msg.sender);
    }

    // 管理员设置TOX代币地址
    function adminSetTOXAddress(address TOXAddress_) public onlyAdmin {
        TOXAddress = TOXAddress_;
    }

    // 管理员设置城市合约地址
    function adminSetCityAddress(address cityAddress_) public onlyAdmin {
        cityAddress = cityAddress_;
    }

    // 管理员设置增加用户合约余额合约地址
    function adminSetMiningAddress(address miningAddress_) public onlyAdmin {
        miningAddress = miningAddress_;
    }

    // 管理员设置考核标准,月份为1，2，3
    function adminSetAssessmentCriteria(uint256 cityLevel_, uint256 month_, uint256 point_) public onlyAdmin {
        assessmentCriteria[cityLevel_][month_] = point_;
    }

    // 管理员设置先锋每天新增质押量
    function setPioneerDelegate(address pioneerAddress_, uint256 amount_, bytes32 cityId_) public onlyAdmin {
        Pioneer storage pioneer = pioneerInfo[pioneerAddress_];

        IntoCity city = IntoCity(cityAddress);
        // 增加先锋绑定的城市的累计质押量
        city.addCityDelegate(cityId_, amount_);
        // 更新奖励和考核
        uint256 today = getDay();
        if (!checkPioneerStatus[pioneerAddress_][today]) {
            // 考核
            checkPioneer(cityId_, pioneerAddress_);
            checkPioneerStatus[pioneerAddress_][today] = true;
        }
        if (!checkRewardStatus[pioneerAddress_][today]) {
            // 奖励
            reward(cityId_, pioneerAddress_);
            checkRewardStatus[pioneerAddress_][today] = true;
        }

        emit DailyIncreaseDelegateRecord(
            pioneer.pioneerAddress,
            cityId_,
            amount_,
            block.timestamp
        );
    }

    function getDay() public view returns (uint256){
        uint day = block.timestamp / 300;
        return uint256(day);
    }

    // 缴纳保证金
    function depositSurety() public {
        IntoCity city = IntoCity(cityAddress);
        // 查询调用者地址是否是城市先锋
        bytes32 cityId = city.pioneerCity(msg.sender);
        require(cityId != bytes32(0), "you are not pioneer"); // 不是城市先锋
        // 查询该城市先锋对应的城市，根据城市查询需要缴纳的保证金数量
        uint256 surety = city.surety(cityId);
        IERC20 TOXContract = IERC20(TOXAddress);
        uint256 userBalance = TOXContract.balanceOf(msg.sender);

        require(userBalance >= surety, "your balance is insufficient"); // 余额不足
        TOXContract.transferFrom(msg.sender, address(this), surety);
        pioneerInfo[msg.sender].pioneerAddress = msg.sender;
        pioneerInfo[msg.sender].assessmentMonthStatus = true;
        pioneerInfo[msg.sender].ctime = block.timestamp;
        pioneerInfo[msg.sender].cityLevel = city.cityLevel(cityId);
        city.initCityDelegate(cityId);// 将先锋绑定的城市的新增质押量变为0
    }

    // 检测考核与保证金退还,每日执行一次,考核失败的城市，可以参与城市节点竞选
    function checkPioneer(bytes32 cityId, address pioneerAddress_) private {
        Pioneer storage pioneer = pioneerInfo[pioneerAddress_];
        if (pioneer.assessmentStatus == true) {
            return;
        }
        IntoCity city = IntoCity(cityAddress);
        uint256 day = getDay() - pioneer.ctime / 300;

        // 前三个月考核
        assessmentPioneer(cityId, pioneer, city, day);

        // 计算退还保证金额度,并更新退还状态
        calculateRefund(cityId, pioneer, city, day);

    }

    // 前三个月考核
    function assessmentPioneer(bytes32 cityId, Pioneer storage pioneer, IntoCity city, uint256 day) private {
        if (pioneer.assessmentStatus) { // 已经通过终极考核
            return;
        }
        uint256 assessmentCriteriaThreshold; // 考核标准金额

        uint256 pioneerCityTotalNewlyDelegate = city.cityDelegateTotal(cityId); // 先锋绑定的城市总的新增质押量

        if (day == 90) {
            assessmentCriteriaThreshold = assessmentCriteria[pioneer.cityLevel][3] * 100;
            if (pioneerCityTotalNewlyDelegate < assessmentCriteriaThreshold) {
                pioneer.assessmentMonthStatus = false;
                failedDelegate[cityId] = pioneerCityTotalNewlyDelegate;
                city.setCityPioneerAssessment(cityId); // 将该城市设置为先锋计划洛选城市
            }
        } else if (day == 60) {
            assessmentCriteriaThreshold = assessmentCriteria[pioneer.cityLevel][2] * 100;
            if (pioneerCityTotalNewlyDelegate < assessmentCriteriaThreshold) {
                pioneer.assessmentMonthStatus = false;
                failedDelegate[cityId] = pioneerCityTotalNewlyDelegate;
                city.setCityPioneerAssessment(cityId); // 将该城市设置为先锋计划洛选城市
            }
        } else if (day == 30) {
            // 检测是否满足直接考核通过
            if (pioneerCityTotalNewlyDelegate >= assessmentCriteria[pioneer.cityLevel][3] * 100) { //直接达到m3考核标准，也就是直接通过终极考核
                pioneer.assessmentStatus = true;
                return;
            }
            // 没达到M3，考核M1
            assessmentCriteriaThreshold = assessmentCriteria[pioneer.cityLevel][1] * 100;
            if (pioneerCityTotalNewlyDelegate < assessmentCriteriaThreshold) {
                pioneer.assessmentMonthStatus = false;
                failedDelegate[cityId] = pioneerCityTotalNewlyDelegate;
                city.setCityPioneerAssessment(cityId); // 将该城市设置为先锋计划洛选城市
            }
        }
    }

    // 计算退还保证金额度,并更可退还金额
    function calculateRefund(bytes32 cityId, Pioneer storage pioneer, IntoCity city, uint256 day) private {
        if (pioneer.returnSuretyStatus = true) {// 已经退还
            return;
        }
        uint256 cityLevel = city.cityLevel(cityId); // 城市等级
        uint256 surety = city.surety(cityId); // 城市保证金

        uint256 pioneerCityTotalNewlyDelegate = city.cityDelegateTotal(cityId); // 先锋绑定的城市总的新增质押量
        uint256 suretyReturn; // 退还保证金的金额
        if (day == 30) {
            // 直接考核通过，退还100%，满足M3退还标准，直接退100%
            if (pioneerCityTotalNewlyDelegate >= assessmentReturnCriteria[cityLevel][3] * 100) {
                pioneer.returnSuretyRate = 100;
                suretyReturn = surety; // 100% 退还
                pioneer.returnSuretyStatus = true;
                pioneer.returnSuretyTime = block.timestamp;
                alreadyRewardRate[pioneer.pioneerAddress][1] = 100; // 第一个月退了100%
            } else {
                // 第一个月满足第一档
                if (pioneerCityTotalNewlyDelegate >= assessmentReturnCriteria[cityLevel][1] * 100) {
                    pioneer.returnSuretyRate = assessmentReturnRate[cityLevel][1];
                    suretyReturn = surety * pioneer.returnSuretyRate / 100;
                    pioneer.returnSuretyStatus = true;
                    pioneer.returnSuretyTime = block.timestamp;
                    alreadyRewardRate[pioneer.pioneerAddress][1] = assessmentReturnRate[cityLevel][1]; // 第一个月退的比例
                } else if (pioneerCityTotalNewlyDelegate >= assessmentReturnCriteria[cityLevel][2] * 100) {// 第一个月满足第2档
                    pioneer.returnSuretyRate = assessmentReturnRate[cityLevel][2];
                    suretyReturn = surety * pioneer.returnSuretyRate / 100;
                    pioneer.returnSuretyStatus = true;
                    pioneer.returnSuretyTime = block.timestamp;
                    alreadyRewardRate[pioneer.pioneerAddress][1] = assessmentReturnRate[cityLevel][2]; // 第一个月退的比例
                }
            }

        } else if (day == 60) {
            uint256 firstMonthRate = alreadyRewardRate[pioneer.pioneerAddress][1];
            // 第2个月满足第一档
            if (pioneerCityTotalNewlyDelegate >= assessmentReturnCriteria[cityLevel][4] * 100) {
                if (assessmentReturnRate[cityLevel][4] <= firstMonthRate) { // 满足退还额度标准的情况下，需要第二个月的退还比例大于第一个月的
                    return;
                }
                pioneer.returnSuretyRate = assessmentReturnRate[cityLevel][4];
                suretyReturn = surety * pioneer.returnSuretyRate / 100;
                pioneer.returnSuretyStatus = true;
                pioneer.returnSuretyTime = block.timestamp;
                alreadyRewardRate[pioneer.pioneerAddress][2] = assessmentReturnRate[cityLevel][4]; // 第2个月退的比例
            } else if (pioneerCityTotalNewlyDelegate >= assessmentReturnCriteria[cityLevel][5] * 100) {// 第2个月满足第2档
                if (assessmentReturnRate[cityLevel][5] <= firstMonthRate) {
                    return;
                }
                pioneer.returnSuretyRate = assessmentReturnRate[cityLevel][5];
                suretyReturn = surety * pioneer.returnSuretyRate / 100;
                pioneer.returnSuretyStatus = true;
                pioneer.returnSuretyTime = block.timestamp;
                alreadyRewardRate[pioneer.pioneerAddress][2] = assessmentReturnRate[cityLevel][5]; // 第2个月退的比例
            } else if (pioneerCityTotalNewlyDelegate >= assessmentReturnCriteria[cityLevel][6] * 100) {// 第2个月满足第3档
                if (assessmentReturnRate[cityLevel][5] <= firstMonthRate) {
                    return;
                }
                pioneer.returnSuretyRate = assessmentReturnRate[cityLevel][6];
                suretyReturn = surety * pioneer.returnSuretyRate / 100;
                pioneer.returnSuretyStatus = true;
                pioneer.returnSuretyTime = block.timestamp;
                alreadyRewardRate[pioneer.pioneerAddress][2] = assessmentReturnRate[cityLevel][6]; // 第2个月退的比例
            }
        }
        suretyReward[pioneer.pioneerAddress] += suretyReturn;

    }

    // 每日奖励发放
    function reward(bytes32 cityId, address pioneerAddress_) private {
        Pioneer storage pioneer = pioneerInfo[pioneerAddress_];
        // 考核状态（按月考核）不正常，不发放奖励
        if (!pioneer.assessmentMonthStatus) {
            return;
        }
        uint256 today = getDay();
        // 任期结束，不发放奖励
        if (today - pioneer.ctime / 300 >= 1800) {
            return;
        }
        // 昨日，天
        uint256 yesterday = getDay() - 1;

        // 福利包奖励
        uint bonus = 93333333333333333333;
        benefitPackageReward[pioneerAddress_] += bonus;

        // 社交基金5%奖励
        IntoCity city = IntoCity(cityAddress);
        uint256 yesterdayFounds = city.getFounds(cityId, yesterday);// 昨日先锋绑定城市新增社交基金
        uint256 allDailyFoundsTotal = city.allDailyFoundsTotal(yesterday);// 全网昨日所有城市新增社交基金
        fundsReward[pioneerAddress_] += yesterdayFounds * 5 / 100 / allDailyFoundsTotal;

        // 该城市新增质押量1%奖励，不累加
        uint256 yesterdayDelegate = city.getNewlyDelegate(cityId, yesterday);// 昨日新增质押权重
        delegateReward[pioneerAddress_] += yesterdayDelegate / 100;

        emit DailyRewardRecord(
            pioneerAddress_,
            bonus,
            fundsReward[pioneerAddress_],
            delegateReward[pioneerAddress_]
        );
    }

    // 用户提取福利包奖励
    function withdrawalBenefitPackageReward() public {
        // 判断是否是先锋
        require(pioneerInfo[msg.sender].ctime > 0, "you are not pioneer");

        // 将奖励转账到用户合约余额
        setUserBalance(msg.sender, benefitPackageReward[msg.sender], 19);
        // 更新领取状态(全部领完才算已领取)
        benefitPackageRewardStatus[msg.sender] = true;
        benefitPackageRewardReceived[msg.sender] += benefitPackageReward[msg.sender]; // 更新已领取福利包奖励
        emit WithdrawalRewardRecord(
            msg.sender,
            benefitPackageReward[msg.sender],
            1
        );
    }

    // 用户提取社交基金奖励
    function withdrawalFundsReward() public {
        // 判断是否是先锋
        require(pioneerInfo[msg.sender].ctime > 0, "you are not pioneer");
        // 将奖励转账到用户合约余额
        setUserBalance(msg.sender, fundsReward[msg.sender], 18);
        // 更新领取状态(全部领完才算已领取)
        fundsRewardStatus[msg.sender] = true;
        fundsRewardReceived[msg.sender] += fundsReward[msg.sender]; // 更新已领取社交基金奖励
        emit WithdrawalRewardRecord(
            msg.sender,
            fundsReward[msg.sender],
            2
        );
    }

    // 用户提取新增质押奖励
    function withdrawalDelegateReward() public {
        // 判断是否是先锋
        require(pioneerInfo[msg.sender].ctime > 0, "you are not pioneer");
        // 将奖励转账到用户合约余额
        setUserBalance(msg.sender, delegateReward[msg.sender], 17);
        delegateRewardStatus[msg.sender] = true;
        delegateRewardReceived[msg.sender] += delegateReward[msg.sender]; // 更新已领取新增质押奖励
        emit WithdrawalRewardRecord(
            msg.sender,
            delegateReward[msg.sender],
            3
        );
    }

    // 用户提取退还的保证金
    function withdrawalSuretyReward() public {
        // 判断是否是先锋
        require(pioneerInfo[msg.sender].ctime > 0, "you are not pioneer");
        // 退还保证金到钱包账户
        IERC20 TOX = IERC20(TOXAddress);
        TOX.transfer(msg.sender, suretyReward[msg.sender]);
        // 更新已退还记录
        suretyRewardRecord[msg.sender] += suretyReward[msg.sender];
        emit WithdrawalRewardRecord(
            msg.sender,
            suretyReward[msg.sender],
            4
        );
    }

    // 将奖励转账到用户合约余额 17 节点 18 基金 19 福利
    function setUserBalance(address user_, uint256 amount_, uint256 types_) private {
        IntoMining intoMining = IntoMining(miningAddress);
        intoMining.addBalanceWithTypes(user_, amount_, types_);
    }

    // 获取先锋所需缴纳的保证金
    function getSurety() public view returns (uint256){
        IntoCity city = IntoCity(cityAddress);
        // 获取先锋所在城市
        bytes32 cityId = city.pioneerCity(msg.sender);
        // 获取该城市所需要缴纳的保证金
        return city.surety(cityId);
    }

    // 判断一个地址是否是先锋
    function isPioneer(address user) public view returns (bool){
        return pioneerInfo[user].ctime > 0;
    }

    // 缴纳保证金
    function depositSuretyTest(address pioneerTest_) public {
        IntoCity city = IntoCity(cityAddress);
        // 查询调用者地址是否是城市先锋
        bytes32 cityId = city.pioneerCity(pioneerTest_);
        require(cityId != bytes32(0), "you are not pioneer"); // 不是城市先锋

        // 查询该城市先锋对应的城市，根据城市查询需要缴纳的保证金数量
//        uint256 surety = city.surety(cityId);
//        IERC20 TOXContract = IERC20(TOXAddress);
//        uint256 userBalance = TOXContract.balanceOf(msg.sender);
//        require(userBalance >= surety, "your balance is insufficient"); // 余额不足
//        TOXContract.transfer(address(this), surety);
        pioneerInfo[pioneerTest_].pioneerAddress = pioneerTest_;
        pioneerInfo[pioneerTest_].assessmentMonthStatus = true;
        pioneerInfo[pioneerTest_].ctime = block.timestamp;
        pioneerInfo[pioneerTest_].cityLevel = city.cityLevel(cityId);
        city.initCityDelegate(cityId);// 将先锋绑定的城市的新增质押量变为0
    }

}