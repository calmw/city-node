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

interface ICityPioneer {
    function pioneerInfo(
        address pioneerAddress_
    ) external view returns (Pioneer memory);

    function getDay() external view returns (uint256);

    // 增加先锋可退还保证金
    function addSuretyReward(
        address pioneerAddress_,
        uint256 suretyReturn_,
        uint256 month_
    ) external;

    // 增加先锋可退还保证金比例,先锋地址 => (月份=>退的比例)，月份为1和2，比例为整数
    // 并设置每月退还的比例
    function addSuretyRewardRate(
        address pioneerAddress_,
        uint256 suretyReturnRate_,
        uint256 month_
    ) external;

    // 城市先锋考核标准：城市等级 => (月份 => 质押点数);考核标准
    function assessmentCriteria(
        uint256 cityLevel_,
        uint256 month_
    ) external view returns (uint256);

    // 更新考核通过时间
    function saveAppraiseTime(
        address pioneerAddress_,
        uint256 timeStamp_,
        uint256 saveType_ // 1 考试通过，2 考核失败
    ) external;

    // 获取考核失败时间
    function failedAt(address pioneerAddress_) external view returns (uint256);

    // 更新考核失败时候先锋绑定城市的充值权重
    function saveFailedWeight(bytes32 chengShiId_, uint256 weight_) external;

    // 更新可领取保证金时刻的充值权重
    function saveSuretyMonthWeight(
        address pioneerAddress_,
        uint256 month_,
        uint256 weight_
    ) external;

    // 获取城市先锋保证金退还标准,城市等级 => (索引 => 退还标准)，索引1，2，3为第一个月的，索引4，5，6为第二个月的；退保证金标准（点数）
    function assessmentReturnCriteria(
        uint256 cityLevel_,
        uint256 month_
    ) external view returns (uint256);

    // 获取城市先锋保证金退还比例, 城市等级 => (索引 => 退还比例)，索引1，2，3为第一个月的，索引4，5，6为第二个月的；退保证金标准（比例）
    function assessmentReturnRate(
        uint256 cityLevel_,
        uint256 month_
    ) external view returns (uint256);

    // 获取先锋已退保证金的比例，先锋地址 => (月份=>退的比例)，月份为1和2，比例为整数
    function alreadyRewardRate(
        address pioneerAddress_,
        uint256 month_
    ) external view returns (uint256);

    // 获取每天秒数
    function secondsPerDay() external view returns (uint256);

    // 更改先锋直接通过终极考核状态
    function saveAssessmentStatus(
        address pioneerAddress_,
        bool status_
    ) external;
}
