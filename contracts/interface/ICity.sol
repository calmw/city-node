// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.18;

interface ICity {
    // 设置竞选失败的先锋城市
    function setChengShiPioneerAssessment(bytes32 chengShiId_) external;

    // 获取先锋对应城市的充值权重
    function getChengShiRechargeWeight(
        bytes32 chengShiId
    ) external view returns (uint256);

    // 获取先锋对应城市的等级
    function chengShiLevel(bytes32 chengShiId_) external view returns (uint256);

    // 获取城市需要缴纳的保证金,城市等级 => 该城市先锋需要缴纳的保证金
    function chengShiLevelSurety(
        uint256 chengShiLevel_
    ) external view returns (uint256);
}
