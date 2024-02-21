// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.18;

interface IAppraise {
    // 获取先锋批次
    function pioneerBatch(address pioneerAddress_) external returns (uint256);

    // 第三批考核
    function appraiseBeth3(
        address pioneerAddress_,
        bytes32 chengShiId_
    ) external returns (bool, bool, uint256, uint256);

    //
    function delPioneerPreMonthWeight(address pioneerAddress_) external;

    // 设置先锋批次
    function setPioneerBatch(address pioneerAddress_, uint256 batch_) external;

    // 获取区域节点数量
    function pioneerCountyNo() external returns (uint256);

    // 获取先锋类型
    function pioneerType(address pioneerAddress_) external returns (uint256);
}
