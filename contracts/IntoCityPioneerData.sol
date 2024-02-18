// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./RoleAccess.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

interface CityPioneer {}

interface City {
    function getChengShiRechargeWeight(bytes32 chengShiId) external;
}

contract IntoCityPioneerData is RoleAccess, Initializable {
    // 用户地址=>用户需要扣除的奖励
    mapping(address => uint256) public subReward; // 废弃
    // 用户地址=>第几期用户
    mapping(address => uint256) public userPeriod;

    CityPioneer public cityPioneer; // 先锋合约
    City public city; // 城市合约

    function initialize() public initializer {
        _addAdmin(msg.sender);
    }

//    // 管理员设置用户需要扣除的奖励
//    function adminSetSubReward(
//        address user_,
//        uint256 reward_
//    ) public onlyAdmin {
//        subReward[user_] = reward_;
//    }
//
//    // 减去用户需要扣除的奖励
//    function adminSubReward(address user_, uint256 reward_) public onlyAdmin {
//        subReward[user_] -= reward_;
//    }

    // 设置先锋合约
    function adminSetCityPioneerAddress(
        address cityPioneerAddress_
    ) public onlyAdmin {
        cityPioneer = CityPioneer(cityPioneerAddress_);
    }

    // 设置城市合约
    function adminSetCityAddress(address cityAddress_) public onlyAdmin {
        city = City(cityAddress_);
    }
}
