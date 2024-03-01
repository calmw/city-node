// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import "./RoleAccess.sol";
import "./IntoCityPioneer.sol";
import "./IntoUserLocation.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./interface/IWithdrawLimit.sol";
import "./interface/IAuth.sol";

interface Founds {
    function getFifteenDayAverage() external view returns (uint256);
}

contract IntoCity is RoleAccess, Initializable {
    event IncreaseCityDelegate(address user, bytes32 cityId, uint256 amount);

    event DecreaseCityDelegate(address user, bytes32 cityId, uint256 amount);

    event RechargeRecord(
        address user,
        bytes32 countyId,
        uint256 amount,
        uint256 time
    );

    event PioneerBlack(address pioneer);

    // 先锋计划合约地址
    address public cityPioneerAddress;
    // 用户定位合约地址
    address public userLocationAddress;
    // 用户定位过的区县ID集合
    bytes32[] public allCityIds;
    // 区县ID => 区县等级
    mapping(bytes32 => uint256) public cityLevel;
    // 区县ID => 区县累计质押量
    mapping(bytes32 => uint256) public cityDelegate;
    // 区县ID => 区县最高质押量2质押量，1质押时间（天）
    mapping(bytes32 => mapping(uint256 => uint256)) public cityMaxDelegate;
    // 区县ID/城市ID => (天=>质押量）
    mapping(bytes32 => mapping(uint256 => uint256)) public cityDelegateRecord; // 废弃
    // 区县ID => 区县先锋是否考核通过，如果考核不通过，后面进入区县节点竞选,ture(考核失败)
    mapping(bytes32 => bool) public cityPioneerAssessment;
    // 区县ID => 区县先锋地址
    mapping(bytes32 => address) public cityPioneer; // 废弃
    // 区县先锋地址 => 区县ID
    mapping(address => bytes32) public pioneerCity; // 废弃
    // 先锋城市ID集合
    bytes32[] public pioneerCityIds; // 废弃
    // 区县ID => 区县先锋需要缴纳的保证金
    mapping(bytes32 => uint256) public earnestMoney;
    // 先锋地址 => 是否被设置过先锋
    mapping(address => bool) public hasSetPioneer;
    // 每天定时任务执行状态
    mapping(uint256 => bool) public dailyTaskStatus;
    // 区县ID => (天=>社交基金量）
    mapping(bytes32 => mapping(uint256 => uint256)) public cityFoundsRecord;
    // 区县ID/城市ID => (天=>质押量）,新增质押量
    mapping(bytes32 => mapping(uint256 => uint256))
        public cityNewlyDelegateRecord;
    // 区县ID/城市ID => 质押 ,区县先锋所绑定区县新增质押量（只用于区县先锋）的累计值
    mapping(bytes32 => uint256) public cityDelegateTotal; // 上线删除
    // 所有区县先锋所绑定区县新增质押量（只用于区县先锋）的累计值
    uint256 public allCityDelegateTotal;
    // 区县ID => 社交基金 ,区县先锋所绑定区县新增社交基金（只用于区县先锋）的累计值
    mapping(bytes32 => uint256) public cityFoundsTotal;
    // 天 => 当天所有区县累计的社交基金值， 新增社交基金,所有区县某天的累计值
    mapping(uint256 => uint256) public allCityDailyFoundsTotal;
    // 天 => 当天所有累计的社交基金值， 新增社交基金,不区分区县某天的累计值
    mapping(uint256 => uint256) public allDailyFoundsTotal;
    // 新增社交基金,所有区县的累计值
    uint256 public allCityFoundsTotal;
    // 新增社交基金累计值，所有的加一起，不区分区县
    uint256 public allFoundsTotal;
    // 城市/区县ID => 城市/区县先锋需要缴纳的保证金,TOX
    mapping(bytes32 => uint256) public surety;
    // 城市ID => 城市/区域等级
    mapping(bytes32 => uint256) public chengShiLevel;
    // 先锋城市ID集合
    bytes32[] public pioneerChengShiIds;
    // 城市等级 => 该城市先锋需要缴纳的保证金
    mapping(uint256 => uint256) public chengShiLevelSurety;
    // 区县ID => 质押 ,区县新增充值权重（只用于区县先锋）的累计值
    mapping(bytes32 => uint256) public cityRechargeTotal;
    // 城市ID => 该城市的充值权重
    mapping(bytes32 => uint256) public chengShiRechargeWeight;
    // 城市ID => 城市先锋地址
    mapping(bytes32 => address) public chengShiPioneer;
    // 城市先锋地址 => 城市ID
    mapping(address => bytes32) public pioneerChengShi;
    // 考核天，正式86400秒，测试300秒
    uint public secondsPerDay;
    // 过去15天社交基金平均值的合约地址
    address public foundsAddress;
    //  城市ID=>(天=>当天累计充值)   充值权重
    mapping(bytes32 => mapping(uint256 => uint256))
        public rechargeDailyWeightRecord;
    //  天=>累计充值)   充值权重
    mapping(uint256 => uint256) public rechargeDailyWeight;
    //  全部累计充值权重
    uint256 public rechargeWeight;
    // 区县ID/城市ID => 先锋所绑定区县/城市新增充值权重（只用于先锋）的累计值
    mapping(bytes32 => uint256) public cityOrChengShiWeightTotal; // 废弃
    // 区县ID => (天=>质押量）,新增质押量，不算减去的
    mapping(bytes32 => mapping(uint256 => uint256))
        public countyNewlyPioneerDelegateRecord;
    //  城市ID=>(天=>当天到之前累计充值)   充值权重
    mapping(bytes32 => mapping(uint256 => uint256))
        public rechargeDailyTotalWeightRecord;
    //  城市先锋地址=>需要补加的充值权重
    mapping(address => uint256) public rechargeWeightAdditional;
    //  城市先锋地址=>状态(true 停止定时任务)
    mapping(address => bool) public pioneerStatus;
    IAuth public auth; // SBT认证合约
    IWithdrawLimit public withdrawLimit; // 是否在小黑屋合约
    address public authAddress; // SBT认证合约
    address public withdrawLimitAddress; // 是否在小黑屋合约
    // 区县ID => 质押 ,区县先锋所绑定区县新增充值权重（只用于区县先锋）的累计值
    mapping(bytes32 => uint256) public countyPioneerRechargeTotal;

    //        function initialize() public initializer {
    //            _addAdmin(msg.sender);
    //        }

    // 管理员设置先锋计划合约地址
    function adminSetCityPioneerAddress(
        address cityPioneerAddress_
    ) public onlyAdmin {
        cityPioneerAddress = cityPioneerAddress_;
    }

    // 管理员设置用户定位合约地址
    function adminSetUserLocationAddress(
        address userLocationAddress_
    ) public onlyAdmin {
        userLocationAddress = userLocationAddress_;
    }

    // 管理员设置获取过去15天社交基金平均值的合约地址
    function adminSetFoundsAddress(address foundsAddress_) public onlyAdmin {
        foundsAddress = foundsAddress_;
    }

    //设置Auth合约
    function adminSetAuthAddress(address authAddress_) public onlyAdmin {
        authAddress = authAddress_;
    }

    // 设置检测小黑屋合约
    function adminSetWithdrawLimitAddress(
        address withdrawLimitAddress_
    ) public onlyAdmin {
        withdrawLimitAddress = withdrawLimitAddress_;
    }

    // 管理员设置先锋是否可以正常分红、考核和退还保证金
    function adminSetPioneerStatus(
        address pioneer_,
        bool status_
    ) public onlyAdmin {
        pioneerStatus[pioneer_] = status_;
    }

    function adminSetPioneer(
        bytes32 chengShiId_,
        address pioneer_,
        uint256 pioneerBatch_,
        uint256 pioneerType_
    ) public onlyAdmin {
        require(!hasSetPioneer[pioneer_], "can not set any more");
        chengShiPioneer[chengShiId_] = pioneer_;
        pioneerChengShi[pioneer_] = chengShiId_;
        hasSetPioneer[pioneer_] = true;
        IntoCityPioneer intoCityPioneer = IntoCityPioneer(cityPioneerAddress);
        intoCityPioneer.setPioneerBefore(pioneer_, pioneerBatch_, pioneerType_); // 设置预添加先锋的信息

        if (!pioneerChengShiIdExits(chengShiId_)) {
            pioneerChengShiIds.push(chengShiId_);
        }
    }

    function changePioneerAddress(
        address newPioneerAddress_,
        address oldPioneerAddress_
    ) public {
        pioneerChengShi[newPioneerAddress_] = pioneerChengShi[
            oldPioneerAddress_
        ];
    }

    function pioneerChengShiIdExits(
        bytes32 chengShiId
    ) public view returns (bool) {
        for (uint256 i = 0; i < pioneerChengShiIds.length; i++) {
            if (chengShiId == pioneerChengShiIds[i]) {
                return true;
            }
        }
        return false;
    }

    function adminRemovePioneer(
        bytes32 chengShiId_,
        address pioneer_
    ) public onlyAdmin {
        chengShiPioneer[chengShiId_] = address(0);
        pioneerChengShi[pioneer_] = bytes32(0);
        hasSetPioneer[pioneer_] = false;
        IntoCityPioneer intoCityPioneer = IntoCityPioneer(cityPioneerAddress);
        intoCityPioneer.delPioneer(pioneer_);
        intoCityPioneer.removePioneer(pioneer_, chengShiId_);

        for (uint256 i = 0; i < pioneerChengShiIds.length; i++) {
            if (pioneerChengShiIds[i] == chengShiId_) {
                pioneerChengShiIds[i] = pioneerChengShiIds[
                    pioneerChengShiIds.length - 1
                ];
                pioneerChengShiIds.pop();
            }
        }
    }

    function adminChangePioneerCityId(
        bytes32 chengShiId_,
        address pioneer_
    ) public onlyAdmin {
        chengShiPioneer[chengShiId_] = pioneer_;
        pioneerChengShi[pioneer_] = chengShiId_;
    }

    function adminPopPioneerChengShiId() public onlyAdmin {
        pioneerChengShiIds.pop();
    }

    // 设置用户充值量
    function adminSetRechargeAmount(
        address user,
        uint256 amount
    ) public onlyAdmin {
        IntoCityPioneer intoCityPioneer = IntoCityPioneer(cityPioneerAddress);
        if (block.timestamp < intoCityPioneer.startTime()) {
            return;
        }
        IntoUserLocation intoUserLocation = IntoUserLocation(
            userLocationAddress
        );
        bytes32 countyId = intoUserLocation.getCountyId(user);
        if (countyId == bytes32(0)) {
            return;
        }
        // 判断用户是否在黑名单
        if (IWithdrawLimit(withdrawLimitAddress).isBlack(user)) {
            return;
        }
        // 判断用户是SBT
        if (!IAuth(authAddress).getAuthStatus(user)) {
            return;
        }
        uint256 today = getDay();
        amount = amount / 100;
        rechargeWeight += amount; // 全部累计充值权重
        rechargeDailyWeight[today] += amount; // 天=>累计充值)   充值权重
        rechargeDailyWeightRecord[countyId][today] += amount; // 区县ID=>(天=>当天累计充值)   充值权重
        // rechargeDailyTotalWeightRecord[countyId][today] += amount;// 区县ID=>(天=>当天累计充值)   充值权重

        cityRechargeTotal[countyId] += amount; // 区县ID=>累计充值权重   充值权重
        countyPioneerRechargeTotal[countyId] += amount; // 区县ID=>累计充值权重   区县先锋充值权重

        emit RechargeRecord(user, countyId, amount, block.timestamp);
    }

    // 管理员设置城市先锋需要额外增加的用户充值权重
    function adminSetRechargeWeightAdditional(
        address pioneer_,
        uint256 weight_
    ) public onlyAdmin {
        rechargeWeightAdditional[pioneer_] = weight_;
    }

    // 管理员设置先锋计划，城市等级以及该等级区县所需缴纳的保证金数额，第三期不需要更改或到期的话，可删除
    function adminSetChengShiLevelAndSurety(
        bytes32 chengShiId_,
        uint256 level_,
        uint256 surety_
    ) public onlyAdmin {
        chengShiLevel[chengShiId_] = level_;
        surety[chengShiId_] = surety_;
        chengShiLevelSurety[level_] = surety_;
    }

    // 设置城市/区域等级
    function adminSetAreaLevel(
        bytes32 AreaId_,
        uint256 level_
    ) public onlyAdmin {
        chengShiLevel[AreaId_] = level_; // 城市/区域等级
    }

    // 管理员修改先锋计划，区县等级以及该等级区县所需缴纳的保证金数额
    function adminEditSurety(
        bytes32 chengShiId_,
        uint256 level_,
        uint256 surety_
    ) public onlyAdmin {
        chengShiLevel[chengShiId_] = level_;
        surety[chengShiId_] = surety_;
        chengShiLevelSurety[level_] = surety_;
    }

    // 管理员设置每天秒数，用于测试
    function adminSetSecondsPerDay(uint56 secondsPerDay_) public onlyAdmin {
        secondsPerDay = secondsPerDay_;
    }

    // 先锋城市数量(这些先锋不一定都交保证金)
    function getPioneerCityNumber() public view returns (uint256) {
        return pioneerChengShiIds.length;
    }

    // 设置竞选失败的先锋城市
    function setChengShiPioneerAssessment(
        bytes32 chengShiId_
    ) public onlyAdmin {
        cityPioneerAssessment[chengShiId_] = true;
    }

    // 增加区县/城市质押量
    function incrCountyOrChengShiDelegate(
        address user_,
        bytes32 countyId_,
        uint256 amount_,
        uint256 today
    ) public onlyAdmin {
        // 判断用户是否在黑名单
        if (IWithdrawLimit(withdrawLimitAddress).isBlack(user_)) {
            return;
        }
        // 判断用户是SBT
        if (!IAuth(authAddress).getAuthStatus(user_)) {
            return;
        }

        cityDelegate[countyId_] += amount_;
        // 增加区县按天的质押记录
        cityNewlyDelegateRecord[countyId_][today] += amount_;
        countyNewlyPioneerDelegateRecord[countyId_][today] += amount_; // 增加先锋城市每天质押量
        emit IncreaseCityDelegate(user_, countyId_, amount_);
    }

    // 减少区县质押量
    function descCountyOrChengShiDelegate(
        address user_,
        bytes32 countyId_,
        uint256 amount_,
        uint256 today
    ) public onlyAdmin {
        if (cityDelegate[countyId_] >= amount_) {
            cityDelegate[countyId_] -= amount_;
        } else {
            cityDelegate[countyId_] = 0;
        }
        // 减少区县按天的质押记录
        if (cityNewlyDelegateRecord[countyId_][today] >= amount_) {
            cityNewlyDelegateRecord[countyId_][today] -= amount_;
        } else {
            cityNewlyDelegateRecord[countyId_][today] = 0;
        }

        emit DecreaseCityDelegate(user_, countyId_, amount_);
    }

    // 获取先锋城市某一天质押量
    function getDailyDelegateByChengShiID(
        bytes32 chengShiId_,
        uint256 day,
        uint256 pioneerType
    ) public view returns (uint256) {
        uint256 total;
        if (pioneerType == 1) {
            total += countyNewlyPioneerDelegateRecord[chengShiId_][day]; // 这里的chengShiId_就是区县ID
        } else {
            IntoUserLocation intoUserLocation = IntoUserLocation(
                userLocationAddress
            );
            bytes32[] memory countyIds_ = intoUserLocation
                .getCountyIdsByChengShiId(chengShiId_);
            for (uint256 i = 0; i < countyIds_.length; i++) {
                bytes32 countyId = countyIds_[i];
                total += countyNewlyPioneerDelegateRecord[countyId][day];
            }
        }

        return total;
    }

    function getDay() public view returns (uint256) {
        uint day = block.timestamp / secondsPerDay;
        return uint256(day);
    }

    // 获取先锋对应城市的充值权重
    function getChengShiRechargeWeight(
        bytes32 chengShiId
    ) public view returns (uint256) {
        uint256 weight;
        IntoCityPioneer intoCityPioneer = IntoCityPioneer(cityPioneerAddress);
        uint256 pioneerType_ = intoCityPioneer.getPioneerType(
            chengShiPioneer[chengShiId]
        );
        if (pioneerType_ == 1) {
            // 区域先锋
            weight += countyPioneerRechargeTotal[chengShiId];
        } else {
            IntoUserLocation intoUserLocation = IntoUserLocation(
                userLocationAddress
            );
            bytes32[] memory cityIds = intoUserLocation
                .getCountyIdsByChengShiId(chengShiId);
            for (uint256 i; i < cityIds.length; i++) {
                weight += cityRechargeTotal[cityIds[i]];
            }
        }

        // 额外增加的权重
        weight += rechargeWeightAdditional[chengShiPioneer[chengShiId]];
        return weight;
    }

    // 设置区县历史最大质押量，mapping(bytes32 => mapping(uint256 => uint256)) public cityMaxDelegate; //  区县最高质押量2质押量，1质押时间（天）
    function setCityMaxDelegate(
        bytes32 cityId_,
        uint256 amount_,
        uint256 day_
    ) public onlyAdmin {
        cityMaxDelegate[cityId_][1] = day_;
        cityMaxDelegate[cityId_][2] = amount_;
    }

    // 管理员设置用户每天质押量变更（新增1和减少2）
    function adminSetDelegate(
        address userAddress_,
        uint256 amount_,
        uint256 setType
    ) public onlyAdmin {
        amount_ *= 100;
        //         判断用户是否有对应的区县
        IntoUserLocation intoUserLocation = IntoUserLocation(
            userLocationAddress
        );
        bytes32 countyId = intoUserLocation.userCityId(userAddress_);
        if (countyId == bytes32(0)) {
            return;
        }
        uint256 today = getDay();
        if (setType == 1) {
            // 增加区县质押量
            incrCountyOrChengShiDelegate(
                userAddress_,
                countyId,
                amount_,
                today
            );
        } else if (setType == 2) {
            // 减少区县质押量
            descCountyOrChengShiDelegate(
                userAddress_,
                countyId,
                amount_,
                today
            );
        }
        // 更新区县历史某天最大质押值
        uint256 yesterdayDelegate = cityDelegateRecord[countyId][today - 1];
        uint256 maxDelegate = cityMaxDelegate[countyId][2];
        if (yesterdayDelegate > maxDelegate) {
            setCityMaxDelegate(countyId, yesterdayDelegate, today - 1);
        }
    }

    function pioneerDailyTask(address pioneerAddress) external {
        if (pioneerStatus[pioneerAddress]) {
            emit PioneerBlack(pioneerAddress);
            return;
        }
        IntoCityPioneer intoCityPioneer = IntoCityPioneer(cityPioneerAddress);
        bytes32 chengShiId = pioneerChengShi[pioneerAddress];
        if (chengShiId != bytes32(0)) {
            intoCityPioneer.pioneerTask(pioneerAddress, chengShiId); // 考核和保证金定时任务
        }
    }

    // 获取前15天社交基金平均值
    function getFifteenDayAverageFounds() public view returns (uint256) {
        Founds founds = Founds(foundsAddress);
        return founds.getFifteenDayAverage();
    }

    // 获取某一天社交基金量增量
    function getFounds(
        bytes32 cityId_,
        uint256 day
    ) public view returns (uint256) {
        return cityFoundsRecord[cityId_][day];
    }

    // 获取某一天质押量（包含增加和减少）
    function getDelegate(
        bytes32 cityId_,
        uint256 day
    ) public view returns (uint256) {
        return cityDelegateRecord[cityId_][day];
    }

    //初始化区县先锋绑定区县的累计质押量
    function initCityRechargeWeight(
        bytes32 chengShiId_,
        uint256 pioneerType_
    ) public onlyAdmin {
        if (pioneerType_ == 1) {
            countyPioneerRechargeTotal[chengShiId_] = 0; // 初始化区域先锋的充值权重
        } else {
            // 初始化城市先锋的充值权重
            IntoUserLocation intoUserLocation = IntoUserLocation(
                userLocationAddress
            );
            bytes32[] memory countyIds_ = intoUserLocation
                .getCountyIdsByChengShiId(chengShiId_);
            for (uint256 i = 0; i < countyIds_.length; i++) {
                cityRechargeTotal[countyIds_[i]] = 0;
            }
        }
    }

    // 获取先锋所需TOX保证金，根据先锋地址
    function getSuretyByPioneerAddress(address pioneer_) public view {
        bytes32 chengShiId = pioneerChengShi[pioneer_];
        uint256 level = chengShiLevel[chengShiId];
        chengShiLevelSurety[level];
    }
}
