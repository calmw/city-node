// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./utils/Datetime.sol";

contract Test {

    event DateLog(
        uint day
    );

    function getDay() public returns (uint256){
        uint day = block.timestamp/86400;
        emit DateLog(day);
        return uint256(day);
    }
}