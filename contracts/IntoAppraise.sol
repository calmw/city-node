// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.18;

import "./RoleAccess.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./interface/IStar.sol";
import "./interface/ICity.sol";
import "./interface/ICityPioneer.sol";

contract IntoAppraise is RoleAccess, Initializable {
    IStar public star; // 获取星级合约
    ICityPioneer public pioneer; // 获取先锋合约
    ICity public city; // 获取城市合约
    // 先锋地址=>（月份=>当月累加新增权重）
    mapping(address => mapping(uint256 => uint256)) public pioneerMonthWeight; // 先锋截止考核时间，按月的每月新增权重
    mapping(address => uint256) public pioneerBatch; // 先锋批次
    // 城市等级=>充值权重
    mapping(uint256 => uint256) public weightByCityLevel; // 先锋考核标准
    mapping(address => mapping(uint256 => uint256))
        public pioneerMonthWeightProcess; // 先锋按月的每月实时新增权重, 废弃
    mapping(address => uint256) public pioneerPreMonthWeight; // 先锋累积到上个月的权重
    mapping(address => uint256) public filedMonth; // 先锋考核失败的月份

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

    // 管理员设置先锋考核标准
    function adminSetWeightByCityLevel(
        uint256 cityLevel_,
        uint256 weight_
    ) public onlyAdmin {
        weightByCityLevel[cityLevel_] = weight_;
    }

    // 管理员设置先锋批次
    function adminSetPioneerBatch(
        address pioneerAddress_,
        uint256 batch_
    ) public onlyAdmin {
        pioneerBatch[pioneerAddress_] = batch_;
    }

    // 第三批考核,返回（是否考核[可能有数据变更]，考核是否成功，考核月份，考核失败时候的总充值权重）
    function appraiseBeth3(
        address pioneerAddress_,
        bytes32 chengShiId_
    ) public onlyAdmin returns (bool, bool, uint256, uint256) {
        // 第一二批计算，沿用之前逻辑,跳过
        if (pioneerBatch[pioneerAddress_] <= 2) {
            return (false, false, 0, 0);
        }
        // 根据先锋交保证金时间获取现在第几个月
        Pioneer memory pioneerInfo = pioneer.pioneerInfo(pioneerAddress_);

        // 已经考核失败的不需要再考核,跳过
        if (!pioneerInfo.assessmentMonthStatus) {
            return (false, false, 0, 0);
        }
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
                return (false, false, 0, 0);
            }
            // 考核
            return (
                true,
                appraiseByStar(
                    pioneerAddress_,
                    cityLevel,
                    daysSinceCreate / 30
                ),
                daysSinceCreate / 30,
                chengShiRechargeWeight
            );
        } else if (daysSinceCreate == 61) {
            // 按月更新充值权重
            savePioneerMonthWeight(
                pioneerAddress_,
                chengShiRechargeWeight,
                daysSinceCreate / 30
            );
            if (userStar >= 4) {
                return (false, false, 0, 0);
            }
            // 考核
            return (
                true,
                appraiseByStar(
                    pioneerAddress_,
                    cityLevel,
                    daysSinceCreate / 30
                ),
                daysSinceCreate / 30,
                chengShiRechargeWeight
            );
        } else if (daysSinceCreate == 91) {
            // 按月更新充值权重
            savePioneerMonthWeight(
                pioneerAddress_,
                chengShiRechargeWeight,
                daysSinceCreate / 30
            );
            if (userStar >= 5) {
                return (false, false, 0, 0);
            }
            // 考核
            return (
                true,
                appraiseByStar(
                    pioneerAddress_,
                    cityLevel,
                    daysSinceCreate / 30
                ),
                daysSinceCreate / 30,
                chengShiRechargeWeight
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
                daysSinceCreate / 30,
                chengShiRechargeWeight
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
                daysSinceCreate / 30,
                chengShiRechargeWeight
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
        return (false, false, 0, 0);
    }

    // 按星级考核
    function appraiseByStar(
        address pioneerAddress_,
        uint256 cityLevel_,
        uint256 month_
    ) private returns (bool) {
        if (
            pioneerMonthWeight[pioneerAddress_][month_] >=
            weightByCityLevel[cityLevel_]
        ) {
            return true;
        } else {
            filedMonth[pioneerAddress_] = month_;
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
            pioneerPreMonthWeight[pioneerAddress_] = 0;
        } else if (month_ == 2) {
            pioneerPreMonthWeight[pioneerAddress_] = pioneerMonthWeight[
                pioneerAddress_
            ][1];
        } else if (month_ == 3) {
            pioneerPreMonthWeight[pioneerAddress_] =
                pioneerMonthWeight[pioneerAddress_][1] +
                pioneerMonthWeight[pioneerAddress_][2];
        } else if (month_ == 4) {
            pioneerPreMonthWeight[pioneerAddress_] =
                pioneerMonthWeight[pioneerAddress_][1] +
                pioneerMonthWeight[pioneerAddress_][2] +
                pioneerMonthWeight[pioneerAddress_][3];
        } else if (month_ == 5) {
            pioneerPreMonthWeight[pioneerAddress_] =
                pioneerMonthWeight[pioneerAddress_][1] +
                pioneerMonthWeight[pioneerAddress_][2] +
                pioneerMonthWeight[pioneerAddress_][3] +
                pioneerMonthWeight[pioneerAddress_][4];
        } else if (month_ == 6) {
            pioneerPreMonthWeight[pioneerAddress_] =
                pioneerMonthWeight[pioneerAddress_][1] +
                pioneerMonthWeight[pioneerAddress_][2] +
                pioneerMonthWeight[pioneerAddress_][3] +
                pioneerMonthWeight[pioneerAddress_][4] +
                pioneerMonthWeight[pioneerAddress_][5];
        }
        pioneerMonthWeight[pioneerAddress_][month_] =
            chengShiRechargeWeight_ -
            pioneerPreMonthWeight[pioneerAddress_];
    }

    // 获取先锋当前的进度
    function pioneerProcess(
        address pioneerAddress_
    ) public view returns (uint256, uint256) {
        bytes32 chengShiId = city.pioneerChengShi(pioneerAddress_);
        uint256 cityLevel = city.chengShiLevel(chengShiId);
        uint256 weightTotal = city.getChengShiRechargeWeight(chengShiId);
        uint256 failedWeight = pioneer.failedDelegate(chengShiId);

        if (filedMonth[pioneerAddress_] > 0) {
            if (filedMonth[pioneerAddress_] == 1) {
                return (failedWeight, weightByCityLevel[cityLevel]);
            } else {
                return (
                    failedWeight - pioneerPreMonthWeight[pioneerAddress_],
                    weightByCityLevel[cityLevel]
                );
            }
        } else {
            if (filedMonth[pioneerAddress_] == 1) {
                return (weightTotal, weightByCityLevel[cityLevel]);
            } else {
                return (
                    weightTotal - pioneerPreMonthWeight[pioneerAddress_],
                    weightByCityLevel[cityLevel]
                );
            }
        }
    }

    // 获取先锋当前的进度
    function pioneerProcess2(
        address pioneerAddress_
    ) public view returns (uint256, uint256) {
        bytes32 chengShiId = city.pioneerChengShi(pioneerAddress_);
        //        uint256 cityLevel = city.chengShiLevel(chengShiId);
        uint256 weightTotal = city.getChengShiRechargeWeight(chengShiId);
        uint256 failedWeight = pioneer.failedDelegate(chengShiId);

        return (failedWeight, weightTotal);
    }

    // 删除先锋当前的进度
    function delPioneerPreMonthWeight(
        address pioneerAddress_
    ) public onlyAdmin {
        delete pioneerPreMonthWeight[pioneerAddress_];
        delete pioneerMonthWeight[pioneerAddress_][1];
        delete pioneerMonthWeight[pioneerAddress_][2];
        delete pioneerMonthWeight[pioneerAddress_][3];
        delete pioneerMonthWeight[pioneerAddress_][4];
        delete pioneerMonthWeight[pioneerAddress_][5];
        delete pioneerMonthWeight[pioneerAddress_][6];
        delete filedMonth[pioneerAddress_];
    }

    // 删除先锋当前的进度
    function getStar(address pioneerAddress_) public view returns (uint256) {
        return star.ownerVip(pioneerAddress_);
    }

    // 设置先锋批次
    function setPioneerBatch(
        address pioneerAddress_,
        uint256 batch_
    ) public onlyAdmin {
        pioneerBatch[pioneerAddress_] = batch_;
    }
}
