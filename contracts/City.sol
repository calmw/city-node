// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./Events.sol";
import "./RoleAccess.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract City is RoleAccess, Initializable {

    bytes32 public cityIdEmpty;

    // 城市ID => 城市等级
    mapping(bytes32 => uint256) public cityLevel;

    // 城市ID => 城市先锋地址
    mapping(bytes32 => address) public cityPioneer;
    // 城市先锋地址 => 城市ID
    mapping(address => bytes32) public pioneerCity;

    // 城市ID => 城市先锋需要缴纳的保证金地址
    mapping(bytes32 => uint256) public earnestMoney;
    // 城市先锋地址 => 是否被设置过城市先锋
    mapping(address => bool) public hasSetPioneer;

    uint256[50] private __gap;

    function initialize() public initializer {
        _addAdmin(msg.sender);
    }

    function AdminSetPioneer(bytes32 cityId_, address pioneer_) public onlyAdmin {
        require(!hasSetPioneer[pioneer_], "can not set any more");
        cityPioneer[cityId_] = pioneer_;
        pioneerCity[pioneer_] = cityId_;
        hasSetPioneer[pioneer_] = true;
    }

    function AdminSetCityLevel(bytes32 cityId_, uint256 level_, uint256 earnestMoney_) public onlyAdmin {
        cityLevel[cityId_] = level_;
        earnestMoney[cityId_] = earnestMoney_;
    }

}