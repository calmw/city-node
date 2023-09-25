// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

interface IUserLocation {

    function getChengShiUserNumber(bytes32 chengShiId) external returns (uint256);

}