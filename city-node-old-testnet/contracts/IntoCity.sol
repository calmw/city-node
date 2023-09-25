// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./RoleAccess.sol";
import "./IntoCityPioneer.sol";
import "./IntoUserLocation.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

interface Founds {
    function getFifteenDayAverage() external view returns (uint256);
}

contract IntoCity is RoleAccess, Initializable {

    event IncreaseCityDelegate(
        address user,
        bytes32 cityId,
        uint256 amount
    );

    event DecreaseCityDelegate(
        address user,
        bytes32 cityId,
        uint256 amount
    );

    event RechargeRecord(
        address user,
        bytes32 countyId,
        uint256 amount,
        uint256 time
    );

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
    // 区县先锋地址 => 是否被设置过区县先锋
    mapping(address => bool) public hasSetPioneer;
    // 每天定时任务执行状态
    mapping(uint256 => bool)public  dailyTaskStatus;

    // 新增变量 -------------------------------------
    // 区县ID => (天=>社交基金量）
    mapping(bytes32 => mapping(uint256 => uint256)) public cityFoundsRecord;
    // 区县ID/城市ID => (天=>质押量）,新增质押量
    mapping(bytes32 => mapping(uint256 => uint256)) public cityNewlyDelegateRecord;
    // 区县ID/城市ID => 质押 ,区县先锋所绑定区县新增质押量（只用于区县先锋）的累计值
    mapping(bytes32 => uint256) public cityDelegateTotal;  // 上线删除
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
    // 城市ID => 区县先锋需要缴纳的保证金
    mapping(bytes32 => uint256) public surety;
    // 城市等级 => 该城市先锋需要缴纳的保证金
//    mapping(uint256 => uint256) public cityLevelSurety; // 上线删除
    // 城市ID => 该城市的权重
//    mapping(bytes32 => uint256) public chengShiWeight; // 上线删除
    // 城市ID => 城市等级
    mapping(bytes32 => uint256) public chengShiLevel;
    // 先锋城市ID集合
    bytes32[] public pioneerChengShiIds;
    // 城市等级 => 该城市先锋需要缴纳的保证金
    mapping(uint256 => uint256) public chengShiLevelSurety;
    // 区县ID => 质押 ,区县先锋所绑定区县新增充值权重（只用于区县先锋）的累计值
    mapping(bytes32 => uint256) public cityRechargeTotal;
    // 城市ID => 该城市的充值权重
//    mapping(bytes32 => uint256) public chengShiRecharge; // 上线删除
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
    mapping(bytes32 => mapping(uint256 => uint256)) public rechargeDailyWeightRecord;
    //  天=>累计充值)   充值权重
    mapping(uint256 => uint256) public rechargeDailyWeight;
    //  全部累计充值权重
    uint256  public rechargeWeight;
    // 区县ID/城市ID => 先锋所绑定区县/城市新增充值权重（只用于先锋）的累计值
    mapping(bytes32 => uint256) public cityOrChengShiWeightTotal; // 废弃
    // 区县ID => (天=>质押量）,新增质押量，不算减去的
    mapping(bytes32 => mapping(uint256 => uint256)) public countyNewlyPioneerDelegateRecord;
    //  城市ID=>(天=>当天到之前累计充值)   充值权重
    mapping(bytes32 => mapping(uint256 => uint256)) public rechargeDailyTotalWeightRecord;


    function initialize() public initializer {
        _addAdmin(msg.sender);
    }

    // 管理员设置先锋计划合约地址
    function adminSetCityPioneerAddress(address cityPioneerAddress_) public onlyAdmin {
        cityPioneerAddress = cityPioneerAddress_;
    }

    // 管理员设置用户定位合约地址
    function adminSetUserLocationAddress(address userLocationAddress_) public onlyAdmin {
        userLocationAddress = userLocationAddress_;
    }

    // 管理员设置获取过去15天社交基金平均值的合约地址
    function adminSetFoundsAddress(address foundsAddress_) public onlyAdmin {
        foundsAddress = foundsAddress_;
    }

    function adminSetPioneer(bytes32 chengShiId_, address pioneer_) public onlyAdmin {
        require(!hasSetPioneer[pioneer_], "can not set any more");
        chengShiPioneer[chengShiId_] = pioneer_;
        pioneerChengShi[pioneer_] = chengShiId_;
        hasSetPioneer[pioneer_] = true;
        if (!pioneerChengShiIdExits(chengShiId_)) {
            pioneerChengShiIds.push(chengShiId_);
        }
    }

    function adminRemoveAllPioneer() public onlyAdmin {
        for (uint256 i = 0; i < pioneerChengShiIds.length; i++) {
            pioneerChengShiIds.pop();
        }
    }


    function pioneerChengShiIdExits(bytes32 chengShiId) public view returns (bool){
        for (uint256 i = 0; i < pioneerChengShiIds.length; i++) {
            if (chengShiId == pioneerChengShiIds[i]) {
                return true;
            }
        }
        return false;
    }

    function adminRemovePioneer(bytes32 chengShiId_, address pioneer_) public onlyAdmin {
        chengShiPioneer[chengShiId_] = address(0);
        pioneerChengShi[pioneer_] = bytes32(0);
        hasSetPioneer[pioneer_] = false;
        IntoCityPioneer intoCityPioneer = IntoCityPioneer(cityPioneerAddress);
        intoCityPioneer.initPioneer(pioneer_);
        intoCityPioneer.removePioneer(pioneer_);

        for (uint256 i = 0; i < pioneerChengShiIds.length; i++) {
            if (pioneerChengShiIds[i] == chengShiId_) {
                pioneerChengShiIds[i] = pioneerChengShiIds[pioneerChengShiIds.length - 1];
                pioneerChengShiIds.pop();
            }
        }
    }

    function adminPopPioneerChengShiId() public onlyAdmin {
        pioneerChengShiIds.pop();
    }

    // 设置用户充值量
    function adminSetRechargeAmount(address user, uint256 amount) public onlyAdmin {
        IntoCityPioneer intoCityPioneer = IntoCityPioneer(cityPioneerAddress);
        if (block.timestamp < intoCityPioneer.startTime()) {
            return;
        }
        IntoUserLocation intoUserLocation = IntoUserLocation(userLocationAddress);
        bytes32 countyId = intoUserLocation.getCountyId(user);
        if (countyId == bytes32(0)) {
            return;
        }
        uint256 today = getDay();
        amount = amount / 100;
        rechargeWeight += amount;// 全部累计充值权重
        rechargeDailyWeight[today] += amount;// 天=>累计充值)   充值权重
        rechargeDailyWeightRecord[countyId][today] += amount;// 区县ID=>(天=>当天累计充值)   充值权重
//        rechargeDailyTotalWeightRecord[countyId][today] += amount;// 区县ID=>(天=>当天累计充值)   充值权重

        cityRechargeTotal[countyId] += amount;// 区县ID=>累计充值权重   充值权重

        emit RechargeRecord(
            user,
            countyId,
            amount,
            block.timestamp
        );
    }

    // 管理员设置先锋计划，城市等级以及该等级区县所需缴纳的保证金数额
    function adminSetChengShiLevelAndSurety(bytes32 chengShiId_, uint256 level_, uint256 surety_) public onlyAdmin {
        chengShiLevel[chengShiId_] = level_;
        surety[chengShiId_] = surety_;
        chengShiLevelSurety[level_] = surety_;
    }

    // 管理员修改先锋计划，区县等级以及该等级区县所需缴纳的保证金数额
    function adminEditSurety(bytes32 chengShiId_, uint256 level_, uint256 surety_) public onlyAdmin {
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
    function setChengShiPioneerAssessment(bytes32 chengShiId_) public onlyAdmin {
        cityPioneerAssessment[chengShiId_] = true;
    }

    // 增加区县/城市质押量
    function incrCountyOrChengShiDelegate(address user_, bytes32 countyId_, uint256 amount_, uint256 today) public onlyAdmin {
        cityDelegate[countyId_] += amount_;
        // 增加区县按天的质押记录
        cityNewlyDelegateRecord[countyId_][today] += amount_;
        countyNewlyPioneerDelegateRecord[countyId_][today] += amount_; // 增加先锋城市每天质押量
        emit IncreaseCityDelegate(
            user_,
            countyId_,
            amount_
        );
    }

    // 减少区县质押量
    function descCountyOrChengShiDelegate(address user_, bytes32 countyId_, uint256 amount_, uint256 today) public onlyAdmin {
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

        emit DecreaseCityDelegate(
            user_,
            countyId_,
            amount_
        );
    }

    // 获取先锋城市某一天质押量
    function getDailyDelegateByChengShiID(bytes32 chengShiId_, uint256 day) public view returns (uint256){
        uint256 total;
        IntoUserLocation intoUserLocation = IntoUserLocation(userLocationAddress);
        bytes32[] memory countyIds_ = intoUserLocation.getCountyIdsByChengShiId(chengShiId_);
        for (uint256 i = 0; i < countyIds_.length; i++) {
            bytes32 countyId = countyIds_[i];
            total += countyNewlyPioneerDelegateRecord[countyId][day];
        }
        return total;

    }

    function getDay() public view returns (uint256){
        uint day = block.timestamp / secondsPerDay;
        return uint256(day);
    }

    // 获取先锋对应城市的充值权重
    function getChengShiRechargeWeight(bytes32 chengShiId) public view returns (uint256){
        IntoUserLocation intoUserLocation = IntoUserLocation(userLocationAddress);
        bytes32[] memory cityIds = intoUserLocation.getCountyIdsByChengShiId(chengShiId);
        uint256 weight;
        for (uint256 i; i < cityIds.length; i++) {
            weight += cityRechargeTotal[cityIds[i]];
        }
        return weight;
    }

// 设置区县历史最大质押量，mapping(bytes32 => mapping(uint256 => uint256)) public cityMaxDelegate; //  区县最高质押量2质押量，1质押时间（天）
    function setCityMaxDelegate(bytes32 cityId_, uint256 amount_, uint256 day_) public onlyAdmin {
        cityMaxDelegate[cityId_][1] = day_;
        cityMaxDelegate[cityId_][2] = amount_;
    }

// 管理员设置用户每天质押量变更（新增1和减少2）
    function adminSetDelegate(address userAddress_, uint256 amount_, uint256 setType) public onlyAdmin {
        amount_ *= 100;
//         判断用户是否有对应的区县
        IntoUserLocation intoUserLocation = IntoUserLocation(userLocationAddress);
        bytes32 countyId = intoUserLocation.userCityId(userAddress_);
        if (countyId == bytes32(0)) {
            return;
        }
        uint256 today = getDay();
        if (setType == 1) {// 增加
            IntoCityPioneer intoCityPioneer = IntoCityPioneer(cityPioneerAddress);
//            // 判断是否是先锋,先锋累计新增质押，不统计减少的
            if (intoCityPioneer.isPioneer(userAddress_)) { // 是先锋
                bytes32 chengShiId = intoUserLocation.getChengShiIdByCountyId(countyId);
                if (chengShiId != bytes32(0)) {
                    intoCityPioneer.pioneerTask(userAddress_, chengShiId); // 考核和保证金检测
                }
            }
//              增加区县质押量
            incrCountyOrChengShiDelegate(userAddress_, countyId, amount_, today);// 同时增加区县质押量，和城市质押量可以各取各的
        } else if (setType == 2) {// 减少
            // 减少区县质押量
            descCountyOrChengShiDelegate(userAddress_, countyId, amount_, today);
        }
//         更新区县历史某天最大质押值
        uint256 yesterdayDelegate = cityDelegateRecord[countyId][today - 1];
        uint256 maxDelegate = cityMaxDelegate[countyId][2];
        if (yesterdayDelegate > maxDelegate) {
            setCityMaxDelegate(countyId, yesterdayDelegate, today - 1);
        }
    }

    function triggerPioneerTask(address user) public {
        IntoUserLocation intoUserLocation = IntoUserLocation(userLocationAddress);
        bytes32 countyId = intoUserLocation.getCountyId(user);
        bytes32 chengShiId = intoUserLocation.cityIdChengShiID(countyId);
        if (chengShiId == bytes32(0)) {
            return;
        }
        if (!isPioneerChengShi(chengShiId)) {
            return;
        }
        address pioneerAddress = chengShiPioneer[chengShiId];
        if (chengShiPioneer[chengShiId] != address(0)) {
            IntoCityPioneer intoCityPioneer = IntoCityPioneer(cityPioneerAddress);
            intoCityPioneer.pioneerTask(pioneerAddress, chengShiId); // 考核和保证金检测
        }
    }

    function cityPioneerTask(address pioneerAddress) public {
        IntoCityPioneer intoCityPioneer = IntoCityPioneer(cityPioneerAddress);
        bytes32  chengShiId = pioneerChengShi[pioneerAddress];
        intoCityPioneer.pioneerTask(pioneerAddress, chengShiId); // 考核和保证金检测
    }

    function isPioneerChengShi(bytes32 chengShiId) public view returns (bool){
        for (uint256 i; i < pioneerChengShiIds.length; i++) {
            if (chengShiId == pioneerChengShiIds[i]) {
                return true;
            }
        }
        return false;
    }

// 获取前15天社交基金平均值
    function getFifteenDayAverageFounds() public view returns (uint256) {
        Founds founds = Founds(foundsAddress);
        return founds.getFifteenDayAverage();
    }

// 获取某一天社交基金量增量
    function getFounds(bytes32 cityId_, uint256 day) public view returns (uint256){
        return cityFoundsRecord[cityId_][day];
    }

// 获取某一天质押量（包含增加和减少）
    function getDelegate(bytes32 cityId_, uint256 day) public view returns (uint256){
        return cityDelegateRecord[cityId_][day];
    }

//初始化区县先锋绑定区县的累计质押量
    function initCityRechargeWeight(bytes32 chengShiId_) public onlyAdmin {
        IntoUserLocation intoUserLocation = IntoUserLocation(userLocationAddress);
        bytes32[] memory countyIds_ = intoUserLocation.getCountyIdsByChengShiId(chengShiId_);
        for (uint256 i = 0; i < countyIds_.length; i++) {
            cityRechargeTotal[countyIds_[i]] = 0;
        }
    }

// 获取先锋所需保证金，根据先锋地址
    function getSuretyByPioneerAddress(address pioneer_) public view {
        bytes32 chengShiId = pioneerChengShi[pioneer_];
        uint256 level = chengShiLevel[chengShiId];
        chengShiLevelSurety[level];
    }

// 获取先锋所需保证金，根据先锋地址
    function getPioneerAndCityNodeNumber() public view returns (uint256){
        IntoCityPioneer intoCityPioneer = IntoCityPioneer(cityPioneerAddress);
        uint256 pioneerNumber = intoCityPioneer.getPioneerNumber(); // 先锋数量
        uint256 cityNodeNumber = 0; // 城市节点数量
        return pioneerNumber + cityNodeNumber;
    }

}