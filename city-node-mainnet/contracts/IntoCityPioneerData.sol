// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./RoleAccess.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract IntoCityPioneerData is RoleAccess, Initializable {

    // 用户地址=>用户需要扣除的奖励
    mapping(address => uint256) public subReward;

    function initialize() public initializer {
        _addAdmin(msg.sender);
    }

//     管理员设置用户需要扣除的奖励
    function adminSetSubReward(address user_, uint256 reward_) public onlyAdmin {
        subReward[user_] = reward_;
    }

//     减去用户需要扣除的奖励
    function adminSubReward(address user_, uint256 reward_) public onlyAdmin {
        subReward[user_] -= reward_;
    }
}