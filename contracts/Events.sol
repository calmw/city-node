// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract Events {

    enum VoteStatus{
        none,
        globalVoting,
        globalVoteEnd,
        userVoting,
        userVoteEnd,
        incumbency,
        endOfTerm
    }

    // 投票记录
    event VoteRecord(
        address user,
        address candidate,
        bytes32 cityId,
        uint256 toxNum
    );

///////////////////////////////////////// 投票 ////////////////////////////////////////////

    // 投票成功
    event VoteSuccessRecord(
        address node,
        bytes32 cityId,
        uint256 toxNum
    );

    // 投票状态改变
    event VoteStatusChangeRecord(
        VoteStatus voteStatus
    );

    // 申请城市候选人记录
    event ApplyCandidateRecord(
        address user,
        uint256 applyFee
    );

///////////////////////////////////////// 投票收益 ////////////////////////////////////////////

    // 昨日城市社交基金值记录
    event UpdateSocialFundsRecord (
        bytes32 cityId,
        uint256 rewardsForSocialFunds,
        uint256 timeStemp
    );

    // 昨日城市新增质押记录
    event UpdateLatestDelegateRecord (
        bytes32 cityId,
        uint256 latestDelegate,
        uint256 timeStemp
    );

///////////////////////////////////////// 城市先锋 ////////////////////////////////////////////

    // 每日新增质押事件
    event DailyIncreaseDelegateRecord(
        address pioneerAddress, // 先锋地址
        bytes32 cityId, //城市ID
        uint256 amount, // 新增质押金额
        uint256 ctime // 新增质押的时间
    );

    // 保证金退还事件
    event EarnestMoneyRecord(
        address pioneerAddress, // 先锋地址
        uint256 amount // 退还保证金的额度
    );

    // 城市先锋奖励事件
    event DailyRewardRecord(
        address pioneerAddress, // 先锋地址
        uint256 toxReward, // 赠送质押包奖励的额度
        uint256 foundsReward, // 社交基金奖励的额度
        uint256 delegateReward // 新增质押奖励的额度
    );
}