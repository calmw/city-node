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

    // 获取区县先锋的充值权重
    function countyPioneerRechargeTotal(
        bytes32 countyId_
    ) external view returns (uint256);

    // 获取城市需要缴纳的保证金,城市等级 => 该城市先锋需要缴纳的保证金
    function chengShiLevelSurety(
        uint256 chengShiLevel_
    ) external view returns (uint256);

    // 获取先锋所在城市,先锋地址 => 城市ID
    function pioneerChengShi(
        address pioneerAddress_
    ) external view returns (bytes32);

    // 城市/区县ID => 城市/区县先锋需要缴纳的保证金,TOX
    function surety(bytes32 areaId) external view returns (uint);
}
