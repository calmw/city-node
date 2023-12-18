// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.18;

interface IAppraise {
    // 获取先锋批次
    function pioneerBatch(address pioneerAddress_) external returns (uint256);

    // 第三批考核
    function appraiseBeth3(
        address pioneerAddress_,
        bytes32 chengShiId_
    ) external returns (bool, bool, uint256);
}
