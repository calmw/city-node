// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./RoleAccess.sol";
import "./IntoCityPioneer.sol";
import "./IntoUserLocation.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract IntoCity is RoleAccess, Initializable {

    event IncreaseCityDelegate(
        bytes32 cityId,
        uint256 amount
    );

    event DecreaseCityDelegate(
        bytes32 cityId,
        uint256 amount
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
    // 区县ID => (天=>质押量）
    mapping(bytes32 => mapping(uint256 => uint256)) public cityDelegateRecord;
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
    // 区县ID => (天=>质押量）,新增质押量，不算减去的
    mapping(bytes32 => mapping(uint256 => uint256)) public cityNewlyDelegateRecord;
    // 区县ID => 质押 ,区县先锋所绑定区县新增质押量（只用于区县先锋）的累计值
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
    // 区县ID => 质押 ,区县先锋所绑定区县新增质押量（只用于区县先锋）的累计值
    mapping(bytes32 => uint256) public cityRechargeTotal;
    // 城市ID => 该城市的充值权重
//    mapping(bytes32 => uint256) public chengShiRecharge; // 上线删除
    // 城市ID => 该城市的充值权重
    mapping(bytes32 => uint256) public chengShiRechargeWeight;
    // 城市ID => 城市先锋地址
    mapping(bytes32 => address) public chengShiPioneer;
    // 城市先锋地址 => 城市ID
    mapping(address => bytes32) public pioneerChengShi;


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

    function adminSetPioneer(bytes32 chengShiId_, address pioneer_) public onlyAdmin {
        require(!hasSetPioneer[pioneer_], "can not set any more");
        chengShiPioneer[chengShiId_] = pioneer_;
        pioneerChengShi[pioneer_] = chengShiId_;
        hasSetPioneer[pioneer_] = true;
        if (chengShiPioneer[chengShiId_] != address(0)) {
            pioneerChengShiIds.push(chengShiId_);
        }
    }

    function adminRemovePioneer(bytes32 chengShiId_, address pioneer_) public onlyAdmin {
        chengShiPioneer[chengShiId_] = address(0);
        pioneerChengShi[pioneer_] = bytes32(0);
        hasSetPioneer[pioneer_] = false;

        for (uint256 i = 0; i < pioneerChengShiIds.length; i++) {
            if (pioneerChengShiIds[i] == chengShiId_) {
                pioneerChengShiIds[i] = pioneerChengShiIds[pioneerChengShiIds.length - 1];
                pioneerChengShiIds.pop();
            }
        }
    }

    // 设置用户充值量
    function adminSetRechargeAmount(address user, uint256 amount) public onlyAdmin {
        IntoUserLocation intoUserLocation = IntoUserLocation(userLocationAddress);
        bytes32 countyId = intoUserLocation.getCountyId(user);

        if (countyId == bytes32(0)) {
            return;
        }

        // 增加充值权重
        addCityDelegate(countyId, amount);

        IntoCityPioneer intoCityPioneer = IntoCityPioneer(cityPioneerAddress);
        // 判断是否是先锋,先锋累计新增质押，不统计减少的
        if (intoCityPioneer.isPioneer(user)) { // 是先锋
            intoCityPioneer.setPioneerDelegate(user, amount);
            cityNewlyDelegateRecord[countyId][getDay()] += amount;
        }
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

    // 先锋城市数量
    function getPioneerCityNumber() public view returns (uint256) {
        return pioneerChengShiIds.length;
    }

    // 设置竞选失败的先锋城市
    function setChengShiPioneerAssessment(bytes32 chengShiId_) public onlyAdmin {
        cityPioneerAssessment[chengShiId_] = true;
    }

    // 增加区县质押量
    function incrCityDelegate(bytes32 cityId_, uint256 amount_, uint256 today) public onlyAdmin {
        cityDelegate[cityId_] += amount_;
        // 增加区县按天的质押记录
        cityDelegateRecord[cityId_][today] += amount_;
        emit IncreaseCityDelegate(
            cityId_,
            amount_
        );
    }

    // 减少区县质押量
    function descCityDelegate(bytes32 cityId_, uint256 amount_, uint256 today) public onlyAdmin {
        if (cityDelegate[cityId_] >= amount_) {
            cityDelegate[cityId_] -= amount_;
        } else {
            cityDelegate[cityId_] = 0;
        }
        // 减少区县按天的质押记录
        if (cityDelegateRecord[cityId_][today] >= amount_) {
            cityDelegateRecord[cityId_][today] -= amount_;
        } else {
            cityDelegateRecord[cityId_][today] = 0;
        }

        emit DecreaseCityDelegate(
            cityId_,
            amount_
        );
    }

    function getDay() public view returns (uint256){
        uint day = block.timestamp / 300;
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
                intoCityPioneer.pioneerTask(userAddress_, chengShiId);
                cityNewlyDelegateRecord[countyId][today] += amount_;
            }
//              增加区县质押量
            incrCityDelegate(countyId, amount_, today);
        } else if (setType == 2) {// 减少
            // 减少区县质押量
            descCityDelegate(countyId, amount_, today);
        }
//         更新区县历史某天最大质押值
        uint256 yesterdayDelegate = cityDelegateRecord[countyId][today - 1];
        uint256 maxDelegate = cityMaxDelegate[countyId][2];
        if (yesterdayDelegate > maxDelegate) {
            setCityMaxDelegate(countyId, yesterdayDelegate, today - 1);
        }
        // 考核和保证金检测

    }

    // 管理员设置用户每天社交基金变更（只有增加）
    function adminSetFounds(address userAddress_, uint256 amount_) public onlyAdmin {
        uint256 today = getDay();
        allFoundsTotal += amount_; // 全网不区分区县，所有累计新增社交基金值
        allDailyFoundsTotal[today] += amount_; // 全网不区分区县，某天所有累计新增社交基金值
        // 判断用户是否有对应的区县
        IntoUserLocation intoUserLocation = IntoUserLocation(userLocationAddress);
        bytes32 cityId = intoUserLocation.userCityId(userAddress_);
        if (cityId == bytes32(0)) {
            return;
        }
        cityFoundsRecord[cityId][today] += amount_; // 增加区县每天社交基金值
        allCityDailyFoundsTotal[today] += amount_; // 新增社交基金,所有区县某天的累计值
        cityFoundsTotal[cityId] += amount_; // 增加区县累计社交基金值
        allCityFoundsTotal += amount_; // 增加所有区县累计社交基金值
    }

    // 获取某一天社交基金量增量
    function getFounds(bytes32 cityId_, uint256 day) public view returns (uint256){
        return cityFoundsRecord[cityId_][day];
    }

    // 获取某一天质押量（包含增加和减少）
    function getDelegate(bytes32 cityId_, uint256 day) public view returns (uint256){
        return cityDelegateRecord[cityId_][day];
    }

    // 获取某一天新增质押（只算增加的）
    function getNewlyDelegate(bytes32 cityId_, uint256 day) public view returns (uint256){
        return cityNewlyDelegateRecord[cityId_][day];
    }

    // 增加区县先锋绑定区县的累计质押量,增加总的质押量
    function addCityDelegate(bytes32 cityId_, uint256 amount_) public onlyAdmin {
        cityRechargeTotal[cityId_] += amount_;
        allCityDelegateTotal += amount_;
    }

    //初始化区县先锋绑定区县的累计质押量
    function initCityDelegate(bytes32 cityId_) public onlyAdmin {
        cityRechargeTotal[cityId_] = 0;
    }

    // 获取先锋所需保证金，根据先锋地址
    function getSuretyByPioneerAddress(address pioneer_) public view {
        bytes32 chengShiId = pioneerChengShi[pioneer_];
        uint256 level = chengShiLevel[chengShiId];
        chengShiLevelSurety[level];
    }

}