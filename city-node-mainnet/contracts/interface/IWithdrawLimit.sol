// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.18;

interface IWithdrawLimit {
    function isBlack(address sender) external view returns (bool);
}
