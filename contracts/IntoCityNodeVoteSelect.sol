// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./RoleAccess.sol";
import "./IntoCityNodeVote.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract IntoCityNodeVoteSelect is RoleAccess, Initializable {


    address public TOX;
    address public IntoCityVoteAddress;


    function initialize() public initializer {
        _addAdmin(msg.sender);
    }

    // 管理员设置投票合约地址
    function adminSetIntoCityVoteAddress(address IntoCityVoteAddress_) public onlyAdmin {
        IntoCityVoteAddress = IntoCityVoteAddress_;
    }


}
