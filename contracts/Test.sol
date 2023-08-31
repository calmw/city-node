// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;


contract Test {

    event Delegate(
        uint256 amount
    );

    mapping(address => bytes32) public pioneerCity;

    function test(address user) public {
        bytes32 cityId = pioneerCity[user];

        if (cityId == bytes32(0)) {
            emit Delegate(3);
        } else {
            emit Delegate(4);
        }
        if (cityId != bytes32(0)) {
            emit Delegate(5);
        } else {
            emit Delegate(6);
        }

    }

}
