// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;


contract Test {

    event TestEvent(
        address user,
        uint256 amount
    );

    function test(uint256 amount_) public {
        emit TestEvent(
            msg.sender,
            amount_
        );
    }
}
