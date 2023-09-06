// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./RoleAccess.sol";
import "./IntoCityPioneer.sol";
import "./IntoUserLocation.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
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

    // 用户定位过的城市ID集合
    bytes32[] public allCityIds;
    // 城市ID => 城市等级
    mapping(bytes32 => uint256) public cityLevel;
    // 城市ID => 城市累计质押量
    mapping(bytes32 => uint256) public cityDelegate;
    // 城市ID => 城市最高质押量2质押量，1质押时间（天）
    mapping(bytes32 => mapping(uint256 => uint256)) public cityMaxDelegate;
    // 城市ID => (天=>质押量）
    mapping(bytes32 => mapping(uint256 => uint256)) public cityDelegateRecord;
    // 城市ID => 城市先锋是否考核通过，如果考核不通过，后面进入城市节点竞选,ture(考核失败)
    mapping(bytes32 => bool) public cityPioneerAssessment;
    // 城市ID => 城市先锋地址
    mapping(bytes32 => address) public cityPioneer;
    // 城市先锋地址 => 城市ID
    mapping(address => bytes32) public pioneerCity;
    // 先锋城市ID集合
    bytes32[] public pioneerCityIds;
    // 城市ID => 城市先锋需要缴纳的保证金
    mapping(bytes32 => uint256) public surety;
    // 城市等级 => 该城市等级所需缴纳的保证金数额
    mapping(uint256 => uint256) public cityLevelSurety;
    // 城市先锋地址 => 是否被设置过城市先锋
    mapping(address => bool) public hasSetPioneer;
    // 每天定时任务执行状态
    mapping(uint256 => bool)public  dailyTaskStatus;

    // 新增变量 -------------------------------------
    // 城市ID => (天=>社交基金量）
    mapping(bytes32 => mapping(uint256 => uint256)) public cityFoundsRecord;
    // 城市ID => (天=>质押量）,新增质押量，不算减去的
    mapping(bytes32 => mapping(uint256 => uint256)) public cityNewlyDelegateRecord;
    // 城市ID => 质押 ,城市先锋所绑定城市新增质押量（只用于城市先锋）的累计值
    mapping(bytes32 => uint256) public cityDelegateTotal;
    // 所有城市先锋所绑定城市新增质押量（只用于城市先锋）的累计值
    uint256 public allCityDelegateTotal;
    // 城市ID => 社交基金 ,城市先锋所绑定城市新增社交基金（只用于城市先锋）的累计值
    mapping(bytes32 => uint256) public cityFoundsTotal;
    // 天 => 当天所有城市累计的社交基金值， 新增社交基金,所有城市某天的累计值
    mapping(uint256 => uint256) public allCityDailyFoundsTotal;
    // 新增社交基金,所有城市的累计值
    uint256 public allCityFoundsTotal;


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

    function adminSetPioneer(bytes32 cityId_, address pioneer_) public onlyAdmin {
        require(!hasSetPioneer[pioneer_], "can not set any more");
        cityPioneer[cityId_] = pioneer_;
        pioneerCity[pioneer_] = cityId_;
        hasSetPioneer[pioneer_] = true;
        if (cityPioneer[cityId_] != address(0)) {
            pioneerCityIds.push(cityId_);
        }
    }

    // 管理员设置先锋计划，城市等级以及该等级城市所需缴纳的保证金数额
    function adminSetCityLevelAndSurety(bytes32 cityId_, uint256 level_, uint256 surety_) public onlyAdmin {
        cityLevel[cityId_] = level_;
        surety[cityId_] = surety_;
        cityLevelSurety[level_] = surety_;
    }

    // 先锋城市数量
    function getPioneerCityNumber() public view returns (uint256) {
        return pioneerCityIds.length;
    }

    // 设置竞选失败的先锋城市
    function setCityPioneerAssessment(bytes32 cityId_) public onlyAdmin {
        cityPioneerAssessment[cityId_] = true;
    }

    // 增加城市质押量
    function incrCityDelegate(bytes32 cityId_, uint256 amount_) public onlyAdmin {
        cityDelegate[cityId_] += amount_;
        emit IncreaseCityDelegate(
            cityId_,
            amount_
        );
    }

    // 减少城市质押量
    function descCityDelegate(bytes32 cityId_, uint256 amount_) public onlyAdmin {
        if (cityDelegate[cityId_] >= amount_) {
            cityDelegate[cityId_] -= amount_;
        } else {
            cityDelegate[cityId_] = 0;
        }
//        cityDelegate[cityId_] -= amount_;

        emit DecreaseCityDelegate(
            cityId_,
            amount_
        );
    }

    function getDay() public view returns (uint256){
        uint day = block.timestamp / 86400;
        return uint256(day);
    }

    function dailyTask(uint256 start, uint256 end) public onlyAdmin returns (bool){
        uint256 day = getDay();
        require(!dailyTaskStatus[day], "can not execute any more");
        IntoUserLocation intoUserLocation = IntoUserLocation(userLocationAddress);
        // 更新城市每天累计质押最大值
        uint256 today = getDay();
        for (uint256 i = start; i < end; i++) {
            bytes32 cityId = intoUserLocation.cityIdsNoRepeat(i);
            uint256 yesterdayDelegate = cityDelegateRecord[cityId][1];
            uint256 maxDelegate = cityMaxDelegate[cityId][2];
            if (yesterdayDelegate > maxDelegate) {
                setCityMaxDelegate(intoUserLocation.cityIds(i), yesterdayDelegate, today - 1);
            }
        }
//        dailyTaskStatus[day] = true;
        return true;
    }

    // 设置城市历史最大质押量，mapping(bytes32 => mapping(uint256 => uint256)) public cityMaxDelegate; //  城市最高质押量2质押量，1质押时间（天）
    function setCityMaxDelegate(bytes32 cityId_, uint256 amount_, uint256 day_) public onlyAdmin {
        cityMaxDelegate[cityId_][1] = day_;
        cityMaxDelegate[cityId_][2] = amount_;
    }

    // 管理员设置用户每天质押量变更（新增1和减少2）
    function adminSetDelegate(address userAddress_, uint256 amount_, uint256 setType) public onlyAdmin {
        amount_ *= 100;
        // 判断用户是否有对应的城市
        IntoUserLocation intoUserLocation = IntoUserLocation(userLocationAddress);
        bytes32 cityId = intoUserLocation.userCityId(userAddress_);
        if (cityId == bytes32(0)) {
            return;
        }
        uint256 today = getDay();
        if (setType == 1) {// 增加
            IntoCityPioneer intoCityPioneer = IntoCityPioneer(cityPioneerAddress);
//             判断是否是先锋,先锋累计新增质押，不统计减少的
            if (intoCityPioneer.isPioneer(userAddress_)) { // 是先锋
                intoCityPioneer.setPioneerDelegate(userAddress_, amount_, cityId);
                cityNewlyDelegateRecord[cityId][today] += amount_;
            }
            // 增加城市质押量
            incrCityDelegate(cityId, amount_);
            // 增加城市质押记录
            cityDelegateRecord[cityId][today] += amount_;
        } else if (setType == 2) {// 减少
            // 减少城市质押量
            descCityDelegate(cityId, amount_);
            // 减少城市质押记录
            if (cityDelegateRecord[cityId][today] >= amount_) {
                cityDelegateRecord[cityId][today] -= amount_;
            } else {
                cityDelegateRecord[cityId][today] = 0;
            }
//            cityDelegateRecord[cityId][today] -= amount_;
        }

    }

    // 管理员设置用户每天社交基金变更（只有增加）
    function adminSetFounds(address userAddress_, uint256 amount_) public onlyAdmin {
        // 判断用户是否有对应的城市
        IntoUserLocation intoUserLocation = IntoUserLocation(userLocationAddress);
        bytes32 cityId = intoUserLocation.userCityId(userAddress_);
        if (cityId == bytes32(0)) {
            return;
        }
        uint256 today = getDay();
        cityFoundsRecord[cityId][today] += amount_; // 增加城市每天社交基金值
        allCityDailyFoundsTotal[today] += amount_; // 新增社交基金,所有城市某天的累计值
        cityFoundsTotal[cityId] += amount_; // 增加城市累计社交基金值
        allCityFoundsTotal += amount_; // 增加所有城市累计社交基金值
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

    // 增加城市先锋绑定城市的累计质押量,增加总的质押量
    function addCityDelegate(bytes32 cityId_, uint256 amount_) public onlyAdmin {
        cityDelegateTotal[cityId_] += amount_;
        allCityDelegateTotal += amount_;
    }

    //初始化城市先锋绑定城市的累计质押量
    function initCityDelegate(bytes32 cityId_) public onlyAdmin {
        cityDelegateTotal[cityId_] = 0;
    }

}