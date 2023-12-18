// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.18;

import "./RoleAccess.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./interface/IStar.sol";
import "./interface/ICity.sol";

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

contract IntoAppraise is RoleAccess, Initializable {
    IStar public star; // 获取星级合约
    ICityPioneer public pioneer; // 获取先锋合约
    ICity public city; // 获取城市合约
    // 先锋地址=>（月份=>当月累加新增权重）
    mapping(address => mapping(uint256 => uint256)) public pioneerMonthWeight; // 先锋按月的每月新增权重
    mapping(address => uint256) public pioneerBatch; // 先锋批次
    // 城市等级=>充值权重
    mapping(uint256 => uint256) public weightByCityLevel; // 先锋考核标准

    function initialize() public initializer {
        _addAdmin(msg.sender);
    }

    // 管理员设置Star合约地址
    function adminSetStar(address starAddress_) public onlyAdmin {
        star = IStar(starAddress_);
    }

    // 管理员设置城市合约
    function adminSetCity(address cityAddress_) public onlyAdmin {
        city = ICity(cityAddress_);
    }

    // 管理员设置先锋合约
    function adminSetPioneer(address cityPioneerAddress_) public onlyAdmin {
        pioneer = ICityPioneer(cityPioneerAddress_);
    }

    // 管理员设置先锋批次
    function adminSetPioneerBatch(address pioneerAddress_, uint256 batch_) public onlyAdmin {
        pioneerBatch[pioneerAddress_] = batch_;
    }

    // 第三批考核,返回（是否考核[可能有数据变更]，考核是否成功，考核月份）
    function appraiseBeth3(
        address pioneerAddress_,
        bytes32 chengShiId_
    ) public onlyAdmin returns (bool, bool, uint256) {
        // 第一二批计算，沿用之前逻辑,跳过
        if (pioneerBatch[pioneerAddress_] <= 2) {
            return (false, false, 0);
        }
        // 根据先锋交保证金时间获取现在第几个月
        Pioneer memory pioneerInfo = pioneer.pioneerInfo(pioneerAddress_);
        // 当前距离交完保证金的天数
        uint256 daysSinceCreate = pioneer.getDay() -
            (pioneerInfo.ctime / pioneer.secondsPerDay());
        // 获取当前充值权重
        uint256 chengShiRechargeWeight = city.getChengShiRechargeWeight(
            chengShiId_
        );
        // 获取城市等级
        uint256 cityLevel = city.chengShiLevel(chengShiId_);
        // 获取用户星级
        uint256 userStar = star.ownerVip(pioneerAddress_);
        if (daysSinceCreate == 31) {
            // 按月更新充值权重
            savePioneerMonthWeight(
                pioneerAddress_,
                chengShiRechargeWeight,
                daysSinceCreate / 30
            );
            if (userStar >= 3) {
                return (false, false, 0);
            }
            // 考核
            return (
                true,
                appraiseByStar(
                pioneerAddress_,
                cityLevel,
                daysSinceCreate / 30
            ),
                daysSinceCreate / 30
            );
        } else if (daysSinceCreate == 61) {
            // 按月更新充值权重
            savePioneerMonthWeight(
                pioneerAddress_,
                chengShiRechargeWeight,
                daysSinceCreate / 30
            );
            if (userStar >= 4) {
                return (false, false, 0);
            }
            // 考核
            return (
                true,
                appraiseByStar(
                pioneerAddress_,
                cityLevel,
                daysSinceCreate / 30
            ),
                daysSinceCreate / 30
            );
        } else if (daysSinceCreate == 91) {
            // 按月更新充值权重
            savePioneerMonthWeight(
                pioneerAddress_,
                chengShiRechargeWeight,
                daysSinceCreate / 30
            );
            if (userStar >= 5) {
                return (false, false, 0);
            }
            // 考核
            return (
                true,
                appraiseByStar(
                pioneerAddress_,
                cityLevel,
                daysSinceCreate / 30
            ),
                daysSinceCreate / 30
            );
        } else if (daysSinceCreate == 121) {
            // 按月更新充值权重
            savePioneerMonthWeight(
                pioneerAddress_,
                chengShiRechargeWeight,
                daysSinceCreate / 30
            );
            // 考核
            return (
                true,
                appraiseByStar(
                pioneerAddress_,
                cityLevel,
                daysSinceCreate / 30
            ),
                daysSinceCreate / 30
            );
        } else if (daysSinceCreate == 151) {
            // 按月更新充值权重
            savePioneerMonthWeight(
                pioneerAddress_,
                chengShiRechargeWeight,
                daysSinceCreate / 30
            );
            // 考核
            return (
                true,
                appraiseByStar(
                pioneerAddress_,
                cityLevel,
                daysSinceCreate / 30
            ),
                daysSinceCreate / 30
            );
        } else if (daysSinceCreate == 181) {
            // 按月更新充值权重
            savePioneerMonthWeight(
                pioneerAddress_,
                chengShiRechargeWeight,
                daysSinceCreate / 30
            );
            // 第六个月无需考核
        }
        return (false, false, 0);
    }

    // 按星级考核
    function appraiseByStar(
        address pioneerAddress_,
        uint256 cityLevel_,
        uint256 month_
    ) private view returns (bool) {
        if (
            pioneerMonthWeight[pioneerAddress_][month_] >=
            weightByCityLevel[cityLevel_]
        ) {
            return true;
        } else {
            return false;
        }
    }

    // 按月更新充值权重
    function savePioneerMonthWeight(
        address pioneerAddress_,
        uint256 chengShiRechargeWeight_,
        uint256 month_
    ) private {
        if (month_ == 1) {
            pioneerMonthWeight[pioneerAddress_][
            month_
            ] = chengShiRechargeWeight_;
        } else {
            pioneerMonthWeight[pioneerAddress_][month_] =
                chengShiRechargeWeight_ -
                pioneerMonthWeight[pioneerAddress_][month_ - 1];
        }
    }
}
