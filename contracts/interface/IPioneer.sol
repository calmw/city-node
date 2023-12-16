// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.18;

struct Pioneer {
    address pioneerAddress;
    uint256 ctime; // 开始考核的时间戳
    uint256 cityLevel; // 所在城市等级,这里是城市ID
    bool assessmentMonthStatus; // 按月考核状态
    bool assessmentStatus; // 最终考核状态
    bool returnSuretyStatus; // 保证金退还状态
    uint256 returnSuretyRate; // 保证金退还比例
    uint256 returnSuretyTime; // 保证金退还时间
    uint256 suretyTime; // 交保证金，成为城市节点的时间戳
}
