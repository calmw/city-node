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
    mapping(uint256 => uint256) public weightByCityLevel; // 先锋考核标准，三期用户
    mapping(address => mapping(uint256 => uint256))
        public pioneerMonthWeightProcess; // 先锋按月的每月实时新增权重, 废弃
    mapping(address => uint256) public pioneerPreMonthWeight; // 先锋累积到上个月的权重
    mapping(address => uint256) public filedMonth; // 先锋考核失败的月份
    mapping(address => uint256) public pioneerPrePreMonthWeight; // 先锋累积到上上个月的权重
    mapping(address => uint256) public pioneerType; // 先锋类型，0 城市节点，1 区县节点
    address[] public pioneerCounty; // 区县节点先锋地址
    uint public pioneerCountyNo; // 区县节点先锋数量
    // 期数=>（城市/区域节点 =>（等级=>(月份=>业绩值)））
    mapping(uint256 => mapping(uint256 => mapping(uint256 => mapping(uint256 => uint256))))
        public weightByAreaLevel; // 先锋考核标准,从四期开始。

    //    function initialize() public initializer {
    //        _addAdmin(msg.sender);
    //    }

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

    // 管理员设置先锋考核标准
    function adminSetWeightByAreaLevel(
        uint256 pioneerBatch_,
        uint256 isArea_, // 0 城市 1 区域
        uint256 areaLevel_,
        uint256 month_,
        uint256 weight_
    ) public onlyAdmin {
        // weightByAreaLevel // 期数=>（城市/区域节点 =>（等级=>(月份=>业绩值)））
        weightByAreaLevel[pioneerBatch_][isArea_][areaLevel_][month_] = weight_;
    }

    // 管理员设置先锋批次
    function adminSetPioneerBatch(
        address pioneerAddress_,
        uint256 batch_,
        uint256 pioneerType_
    ) public onlyAdmin {
        pioneerBatch[pioneerAddress_] = batch_;
        pioneerType[pioneerAddress_] = pioneerType_;
    }

    function adminSetPioneerBatchAndPioneerType(
        address pioneerAddress_,
        uint256 batch_,
        uint256 pioneerType_
    ) public onlyAdmin {
        pioneerBatch[pioneerAddress_] = batch_;
        pioneerType[pioneerAddress_] = pioneerType_;

        bool exist;
        for (uint i = 0; i < pioneerCounty.length; i++) {
            if (pioneerCounty[i] == pioneerAddress_) {
                exist = true;
            }
        }
        if (!exist) {
            pioneerCounty.push(pioneerAddress_);
            pioneerCountyNo++;
        }
    }

    function adminAreaPioneerNo(address pioneerAddress_) public onlyAdmin {
        bool exist;
        for (uint i = 0; i < pioneerCounty.length; i++) {
            if (pioneerCounty[i] == pioneerAddress_) {
                exist = true;
            }
        }
        if (!exist) {
            pioneerCounty.push(pioneerAddress_);
            pioneerCountyNo++;
        }
    }

    // 第三四批考核,返回（是否考核[可能有数据变更]，考核是否成功，考核月份，考核失败时候的总充值权重）
    function appraiseBeth(
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
        // 获取城市/区域等级
        uint256 areaLevel = city.chengShiLevel(chengShiId_);

        // 获取用户星级
        uint256 userStar = star.ownerVip(pioneerAddress_);
        if (
            daysSinceCreate == 30 ||
            daysSinceCreate == 60 ||
            daysSinceCreate == 90 ||
            daysSinceCreate == 120 ||
            daysSinceCreate == 150 ||
            daysSinceCreate == 180
        ) {
            // 按月更新充值权重
            savePioneerMonthWeight(
                pioneerAddress_,
                chengShiRechargeWeight,
                daysSinceCreate / 30
            );
        } else {
            // 不在考核时间点
            return (false, false, 0, 0);
        }

        // 第6个月都免考核
        if (daysSinceCreate == 180) {
            // 无需考核
            return (false, false, 0, 0);
        }

        // 第三批城市先锋，根据星级可以免考核，区域先锋只有第一个月免考核
        if (pioneerBatch[pioneerAddress_] == 3) {
            if (
                (daysSinceCreate == 30 && userStar >= 3) ||
                (daysSinceCreate == 60 && userStar >= 4) ||
                (daysSinceCreate == 90 && userStar >= 5)
            ) {
                // 无需考核
                return (false, false, 0, 0);
            }
        } else if (pioneerBatch[pioneerAddress_] == 4) {
            // 第四期用户第一个月都免考核
            if (daysSinceCreate == 30) {
                // 无需考核
                return (false, false, 0, 0);
            }
        }

        address pioneerAddr_ = pioneerAddress_;

        return (
            true,
            appraisePioneer(
                pioneerAddr_,
                pioneerBatch[pioneerAddr_],
                pioneerType[pioneerAddr_],
                areaLevel,
                daysSinceCreate / 30
            ), //考核
            daysSinceCreate / 30,
            chengShiRechargeWeight
        );
    }

    // 根据先锋城市/区域等级考核
    function appraisePioneer(
        address pioneerAddress_,
        uint256 pioneerBatch_,
        uint256 isChengShi_,
        uint256 areaLevel_,
        uint256 month_
    ) private returns (bool) {
        if (pioneerBatch_ == 3) {
            /// 三期考核
            if (
                pioneerMonthWeight[pioneerAddress_][month_] >=
                weightByCityLevel[areaLevel_]
            ) {
                return true;
            } else {
                filedMonth[pioneerAddress_] = month_;
                return false;
            }
            /// 三期考核
        } else {
            // 四期
            if (
                pioneerMonthWeight[pioneerAddress_][month_] >=
                weightByAreaLevel[pioneerBatch_][isChengShi_][areaLevel_][
                    month_
                ]
            ) {
                return true;
            } else {
                filedMonth[pioneerAddress_] = month_;
                return false;
            }
        }
    }

    // 按月更新充值权重
    function savePioneerMonthWeight(
        address pioneerAddress_,
        uint256 chengShiRechargeWeight_,
        uint256 month_
    ) private {
        pioneerMonthWeight[pioneerAddress_][month_] =
            chengShiRechargeWeight_ -
            pioneerPreMonthWeight[pioneerAddress_]; // 每月充值权重
        pioneerPreMonthWeight[pioneerAddress_] = chengShiRechargeWeight_; // 保存截止当前月份的权重值
        pioneerPrePreMonthWeight[pioneerAddress_] =
            chengShiRechargeWeight_ -
            pioneerMonthWeight[pioneerAddress_][month_]; // 保存截止上个月份的权重值
    }

    // 获取先锋当前的进度
    function pioneerProcess(
        address pioneerAddress_
    ) public view returns (uint256, uint256) {
        uint256 weightTotal;
        uint256 weightTarget;
        bytes32 chengShiId = city.pioneerChengShi(pioneerAddress_);
        uint256 cityLevel = city.chengShiLevel(chengShiId);
        uint256 failedWeight = pioneer.failedDelegate(chengShiId);
        if (pioneerBatch[pioneerAddress_] == 3) {
            weightTarget = weightByCityLevel[cityLevel];
        } else if (pioneerBatch[pioneerAddress_] == 4) {
            // 期数=>（城市/区域节点 =>（等级=>(月份=>业绩值)））
            weightTarget = weightByAreaLevel[4][pioneerType[pioneerAddress_]][
                cityLevel
            ][1];
        }
        if (pioneerType[pioneerAddress_] == 1) {
            weightTotal = city.countyPioneerRechargeTotal(chengShiId);
        } else {
            weightTotal = city.getChengShiRechargeWeight(chengShiId);
        }

        uint256 proc1;
        if (filedMonth[pioneerAddress_] > 0) {
            if (failedWeight > pioneerPrePreMonthWeight[pioneerAddress_]) {
                proc1 =
                    failedWeight -
                    pioneerPrePreMonthWeight[pioneerAddress_];
            }
        } else {
            if (weightTotal > pioneerPreMonthWeight[pioneerAddress_]) {
                proc1 = weightTotal - pioneerPreMonthWeight[pioneerAddress_];
            }
        }

        return (proc1, weightTarget);
    }

    // 获取先锋当前的进度，测试用
    function pioneerProcess2(
        address pioneerAddress_
    ) public view returns (uint256, uint256, uint256) {
        uint256 weightTotal;
        uint256 weightTarget;
        bytes32 chengShiId = city.pioneerChengShi(pioneerAddress_);
        uint256 cityLevel = city.chengShiLevel(chengShiId);
        uint256 failedWeight = pioneer.failedDelegate(chengShiId);
        if (pioneerBatch[pioneerAddress_] == 3) {
            weightTarget = weightByCityLevel[cityLevel];
        } else if (pioneerBatch[pioneerAddress_] == 4) {
            // 期数=>（城市/区域节点 =>（等级=>(月份=>业绩值)））
            weightTarget = weightByAreaLevel[4][pioneerType[pioneerAddress_]][
                cityLevel
            ][1];
        }
        if (pioneerType[pioneerAddress_] == 1) {
            weightTotal = city.countyPioneerRechargeTotal(chengShiId);
        } else {
            weightTotal = city.getChengShiRechargeWeight(chengShiId);
        }

        //        if (pioneerBatch[pioneerAddress_] > 3) {
        uint256 proc1;
        if (filedMonth[pioneerAddress_] > 0) {
            if (failedWeight > pioneerPrePreMonthWeight[pioneerAddress_]) {
                proc1 =
                    failedWeight -
                    pioneerPrePreMonthWeight[pioneerAddress_];
            }
        } else {
            if (weightTotal > pioneerPreMonthWeight[pioneerAddress_]) {
                proc1 = weightTotal - pioneerPreMonthWeight[pioneerAddress_];
            }
        }
        //        } else { // 一二三期
        //            return (weightTotal, weightByCityLevel[cityLevel]);
        //        }
        return (proc1, weightTarget, 0);
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

    // 获取先锋当前的星级
    function getStar(address pioneerAddress_) public view returns (uint256) {
        return star.ownerVip(pioneerAddress_);
    }

    // 设置先锋批次和类型
    function setPioneerBatchAndType(
        address pioneerAddress_,
        uint256 batch_,
        uint256 pioneerType_
    ) public onlyAdmin {
        pioneerBatch[pioneerAddress_] = batch_;
        pioneerType[pioneerAddress_] = pioneerType_;
    }
}
