// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./City.sol";
import "./Events.sol";
import "./RoleAccess.sol";
import "./CityNodeVote.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract CityPioneer is RoleAccess, Events, Initializable {
    enum LifeTime {
        none,
        firstMonth,
        secondMonth,
        thirdMonth,
        moreThanThreeMonth,
        moreThanSixMonth
    }

    event DailyIncreaseDelegate(
        address pioneerAddress, // 先锋地址
        uint256 amount, // 新增质押金额
        uint256 ctime // 新增质押的时间
    );

    struct Pioneer {
        address pioneerAddress;
        uint256 ctime; // 成为城市节点的时间
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

    struct Delegate {
        address pioneer; // 先锋地址
        uint256 amount; // 质押金额
        uint256 ctime; // 质押时间（每天新增质押）
    }

    address public TOX;
    address public cityAddress; // 城市合约地址
    uint256 public termOfOffice; // 任期，单位秒
    uint256 public assessmentPassed = 10000; // 城市先锋第一个月直接通过考核的条件（点数）

    // 城市等级 => (月份 => 质押点数)
    mapping(uint256 => mapping(uint256 => uint256))public  assessmentCriteria; // 城市先锋考核标准
    // 城市等级 => (月份 => 退还标准)
    mapping(uint256 => mapping(uint256 => uint256))public  assessmentReturnCriteria; // 城市先锋保证金退还标准
    // 城市等级 => (月份 => 退还比例)
    mapping(uint256 => mapping(uint256 => uint256))public  assessmentReturnRate; // 城市先锋保证金退还比例
    // 先锋地址 => 先锋信息， 先锋信息
    mapping(address => Pioneer) public pioneerInfo;

    mapping(address => uint256[]) public pioneerDelegateInfo;

    mapping(bytes32 => uint256)public  rewardsForSocialFunds; // 城市每日社交基金
    mapping(bytes32 => uint256)public  latestDayDelegate; // 城市每日新增质押

    uint256[50] private __gap;

    function initialize() public initializer {
        _addAdmin(msg.sender);
    }

    // 管理员设置TOX代币地址
    function adminSetTOXAddress(address TOX_) public onlyAdmin {
        TOX = TOX_;
    }

    // 管理员设置考核标准,月份为1，2，3
    function adminSetAssessmentCriteria(uint256 cityLevel_, uint256 month_, uint256 point_) public onlyAdmin {
        assessmentCriteria[cityLevel_][month_] = point_;
    }

    // 管理员设置先锋每天新增质押量
    function adminSetPioneerDelegate(address pioneer_, uint256 amount_) public onlyAdmin {
        Pioneer storage pioneer = pioneerInfo[pioneer_];
        uint256 day = getDay() - pioneer.day; // 第几天质押
        pioneerDelegateInfo[pioneer_][day] = amount_;
        pioneer.day = day;
        // 更新城市先锋信息
        if (pioneer.day > 180) {
            pioneer.lifeTime = LifeTime.moreThanSixMonth;
        } else if (pioneer.day > 90) {
            pioneer.lifeTime = LifeTime.moreThanThreeMonth;
        } else if (pioneer.day > 60) {
            pioneer.lifeTime = LifeTime.thirdMonth;
            pioneer.thirdMonthDelegate += amount_;
        } else if (pioneer.day > 30) {
            pioneer.lifeTime = LifeTime.secondMonth;
            pioneer.secondMonthDelegate += amount_;
        } else {
            pioneer.lifeTime = LifeTime.firstMonth;
            pioneer.firstMonthDelegate += amount_;
        }

        emit DailyIncreaseDelegate(
            pioneer.pioneerAddress,
            amount_,
            block.timestamp
        );
    }

    function getDay() public returns (uint256){
        uint day = block.timestamp / 86400;
        return uint256(day);
    }

    // 缴纳保证金
    function depositEarnestMoney() public {
        City city = City(cityAddress);
        // 查询调用者地址是否是城市先锋
        bytes32 cityId = city.pioneerCity(msg.sender);
        require(cityId != city.cityIdEmpty(), "you are not pioneer"); // 不是城市先锋

        // 查询该城市先锋对应的城市，根据城市查询需要缴纳的保证金数量
        uint256 earnestMoney = city.earnestMoney(cityId);
        IERC20 TOXContract = IERC20(TOX);
        uint256 userBalance = TOXContract.balanceOf(msg.sender);
        require(userBalance >= earnestMoney, "your balance is insufficient"); // 余额不足
        TOXContract.transfer(address(this), earnestMoney);
        pioneerInfo[msg.sender].ctime = block.timestamp;
        pioneerInfo[msg.sender].pioneerAddress = msg.sender;
        pioneerInfo[msg.sender].assessmentMonthStatus = true;
        pioneerInfo[msg.sender].day = getDay();
        pioneerInfo[msg.sender].cityLevel = city.cityLevel(cityId);
        pioneerInfo[msg.sender].lifeTime = LifeTime.firstMonth;
    }

    // 每天执行一次，计算奖励和考核
    function dailyTask() public onlyAdmin {
        City city = City(cityAddress);
//        bytes32[] memory cityIds = city.pioneerCityIds;
        for (uint256 i = 0; i < city.getCityNumber(); i++) {
            // 考核
            assessUser(city.pioneerCityIds(i), city.cityPioneer(city.pioneerCityIds(i)));
            // 奖励
            reward(city.pioneerCityIds(i), city.cityPioneer(city.pioneerCityIds(i)));
        }
    }

    // 考核,每日执行一次,考核失败的城市，可以参与城市节点竞选
    function assessUser(bytes32 cityId, address pioneerAddress_) private {
        Pioneer storage pioneer = pioneerInfo[pioneerAddress_];
        if (pioneer.assessmentStatus == true) {
            return;
        }
        City city = City(cityAddress);
        uint256 cityLevel = city.cityLevel(cityId); // 城市等级
        uint256 earnestMoney = city.earnestMoney(cityId); // 城市保证金
        uint256 assessmentCriteriaThreshold; // 考核标准金额
        pioneer.day = getDay() - pioneer.day;
        if (pioneer.day == 90) {
            // 考核
            assessmentCriteriaThreshold = assessmentCriteria[pioneer.cityLevel][3] * 100;
            if (pioneer.thirdMonthDelegate < assessmentCriteriaThreshold) {
                pioneer.assessmentMonthStatus = false;
                city.setCityPioneerAssessment(cityId); // 将该城市设置为先锋计划洛选城市
            }
            // 退还保证金
            if (pioneer.secondMonthDelegate >= assessmentReturnCriteria[cityLevel][3] * 100) {
                uint256 rate = assessmentReturnRate[cityLevel][3];
                if (pioneer.firstMonthReturnEarnest) {
                    rate -= assessmentReturnRate[cityLevel][1];
                }
                if (pioneer.secondMonthReturnEarnest) {
                    rate -= assessmentReturnRate[cityLevel][2];
                }
                IERC20(TOX).transfer(pioneer.pioneerAddress, earnestMoney * rate / 100);
            }
        } else if (pioneer.day == 60) {
            // 考核
            assessmentCriteriaThreshold = assessmentCriteria[pioneer.cityLevel][2] * 100;
            if (pioneer.secondMonthDelegate < assessmentCriteriaThreshold) {
                pioneer.assessmentMonthStatus = false;
                city.setCityPioneerAssessment(cityId); // 将该城市设置为先锋计划洛选城市
            }
            // 退还保证金
            if (pioneer.secondMonthDelegate >= assessmentReturnCriteria[cityLevel][2] * 100) {
                uint256 rate = assessmentReturnRate[cityLevel][2];
                if (pioneer.firstMonthReturnEarnest) {
                    rate -= assessmentReturnRate[cityLevel][1];
                }
                IERC20(TOX).transfer(pioneer.pioneerAddress, earnestMoney * rate / 100);
            }
        } else if (pioneer.day == 30) {
            // 考核
            assessmentCriteriaThreshold = assessmentCriteria[pioneer.cityLevel][1] * 100;
            if (pioneer.firstMonthDelegate < assessmentCriteriaThreshold) {
                pioneer.assessmentMonthStatus = false;
                city.setCityPioneerAssessment(cityId); // 将该城市设置为先锋计划洛选城市
            }
            // 检测是否满足直接考核通过
            if (pioneer.firstMonthDelegate >= assessmentPassed) {
                pioneer.assessmentStatus = true;
            }
            // 退还保证金
            if (pioneer.firstMonthDelegate >= assessmentReturnCriteria[cityLevel][1] * 100) {
                uint256 rate = assessmentReturnRate[cityLevel][1];
                IERC20(TOX).transfer(pioneer.pioneerAddress, earnestMoney * rate / 100);
                pioneer.firstMonthReturnEarnest = true;
            }
        }
    }

    // 每日奖励
    function reward(bytes32 cityId, address pioneer) private {
        Pioneer storage pioneer = pioneerInfo[pioneer];
        if (!pioneer.assessmentStatus) {
            return;
        }
        pioneer.day = getDay() - pioneer.day;
        // 第一个月，每天分质押包的收益[獎勵一：TOX福利包]
        if (pioneer.day <= 30) {
            // 93 = 2800 / 30;
            IERC20(TOX).transfer(pioneer.pioneerAddress, 93);
        }
    }

}