// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./RoleAccess.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract IntoCity is RoleAccess, Initializable {

    bytes32 public cityIdEmpty;

    // 城市ID => 城市等级
    mapping(bytes32 => uint256) public cityLevel;
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

    uint256[50] private __gap;
    mapping(uint256 => uint256) public Aaa;

    function initialize() public initializer {
        _addAdmin(msg.sender);
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

    function getCityNumber() public view returns (uint256) {
        return pioneerCityIds.length;
    }

    // 设置竞选失败的先锋城市
    function setCityPioneerAssessment(bytes32 cityId_) public onlyAdmin {
        cityPioneerAssessment[cityId_] = true;
    }

}