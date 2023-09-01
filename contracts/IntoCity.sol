// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./RoleAccess.sol";
import "./IntoCityPioneer.sol";
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

    // 城市ID => 城市先锋需要缴纳的保证金地址
    mapping(bytes32 => uint256) public earnestMoney;
    // 城市先锋地址 => 是否被设置过城市先锋
    mapping(address => bool) public hasSetPioneer;

    function initialize() public initializer {
        _addAdmin(msg.sender);
    }

    // 管理员设置先锋计划合约地址
    function adminSetCityPioneerAddress(address cityPioneerAddress_) public onlyAdmin {
        cityPioneerAddress = cityPioneerAddress_;
    }

    function AdminSetPioneer(bytes32 cityId_, address pioneer_) public onlyAdmin {
        require(!hasSetPioneer[pioneer_], "can not set any more");
        cityPioneer[cityId_] = pioneer_;
        pioneerCity[pioneer_] = cityId_;
        hasSetPioneer[pioneer_] = true;
        if (cityPioneer[cityId_] != address(0)) {
            pioneerCityIds.push(cityId_);
        }
    }

    function AdminSetCityLevel(bytes32 cityId_, uint256 level_, uint256 earnestMoney_) public onlyAdmin {
        cityLevel[cityId_] = level_;
        earnestMoney[cityId_] = earnestMoney_;
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
        cityDelegate[cityId_] -= amount_;
        emit DecreaseCityDelegate(
            cityId_,
            amount_
        );
    }

    function getDay() public view returns (uint256){
        uint day = block.timestamp / 86400;
        return uint256(day);
    }

    // 管理员设置用户每天质押量变更（新增和减少）
    function adminSetDelegate(address userAddress_, uint256 amount_, uint256 setType) public onlyAdmin {
        amount_ *= 100;
        bytes32 cityId = pioneerCity[userAddress_];
        if (cityId == bytes32(0)) {
            return;
        }
        uint256 today = getDay();
        if (setType == 1) {// 增加
//            IntoCityPioneer intoCityPioneer = IntoCityPioneer(cityPioneerAddress);
            // 判断是否是先锋,先锋累计新增质押，不统计减少的
//            if (intoCityPioneer.isPioneer(userAddress_)) { // 是先锋
//                intoCityPioneer.setPioneerDelegate(userAddress_, amount_);
//            }
            // 增加城市质押量
            if (cityId != bytes32(0)) {
                incrCityDelegate(cityId, amount_);
            }
            cityDelegateRecord[cityId][today] += amount_;
        } else {// 减少
            // 减少城市质押量
            if (cityId != bytes32(0)) {
                descCityDelegate(cityId, amount_);
            }
            cityDelegateRecord[cityId][today] -= amount_;
        }
        // 更新历史每天累计最大值
        if (cityDelegateRecord[cityId][today - 1] > cityDelegateRecord[cityId][today - 2]) {
            cityMaxDelegate[cityId][1] = today - 1;
            cityMaxDelegate[cityId][2] = cityDelegateRecord[cityId][today - 1];
        } else {
            cityMaxDelegate[cityId][1] = today - 2;
            cityMaxDelegate[cityId][2] = cityDelegateRecord[cityId][today - 2];
        }
    }
}