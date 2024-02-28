// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface CityPioneerData {
    function subReward(address user_) external returns (uint256);

    function adminSetSubReward(address user_, uint256 reward_) external;

    function adminSubReward(address user_, uint256 reward_) external;

    function suretyUSDT(bytes32 chengShiID) external returns (uint256);

    function isGlobalNode(address user_) external returns (bool);

    // 清除先锋
    function clearPioneerSurety(address pioneerAddress_) external;
}
