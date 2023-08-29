// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./RoleAccess.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract Withdrawal is RoleAccess {

    address public TOX; // TOX 代币地址
    address public withdrawalAddress; // TOX提币地址

    constructor() {
        _addAdmin(msg.sender);
    }

    // 管理员设置投票状态
    function adminWithdrawal() public onlyAdmin {
        IERC20 TOXContract = IERC20(TOX);
        uint256 balance = TOXContract.balanceOf(address(this));
        TOXContract.transfer(withdrawalAddress, balance);
    }

    // 管理员设置TOX代币地址
    function adminSetTOXAddress(address TOX_) public onlyAdmin {
        TOX = TOX_;
    }

    // 管理员设置提款地址
    function adminAddWithdrawalAddress(address withdrawalAddress_) public onlyAdmin {
        withdrawalAddress = withdrawalAddress_;
    }
}
