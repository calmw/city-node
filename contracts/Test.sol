// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;


contract Test {

    event Delegate(
        uint256 amount
    );

    uint256 public ctime;

    constructor(){
        ctime = getDay();
    }

    function getDay() public view returns (uint256){
        uint day = block.timestamp / 300;
        return uint256(day);
    }

    // 数据去重
    function getCityNum() public returns (uint256){
        uint256 day = getDay() - ctime;
        emit Delegate(day);
        return day;
    }
}
