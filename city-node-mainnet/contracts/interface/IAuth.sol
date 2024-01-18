// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.18;

interface IAuth {
    function getAuthStatus(address account) external view returns (bool);
}