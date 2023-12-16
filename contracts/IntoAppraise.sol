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

    // 考核天，正式86400秒，测试300秒
    //    uint public secondsPerDay;

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

    // 考核， 只针对第三批,返回是否需要退保证金
    function appraise(address pioneerAddress_, bytes32 chengShiId_) public {
        bool returnMargin;
        // 第一二批计算，沿用之前逻辑
        if (pioneerBatch[pioneerAddress_] <= 2) {
            checkPioneer(chengShiId_, pioneerAddress_, pioneer.getDay());
            return;
        }
        // 根据先锋交保证金时间获取现在第几个月
        Pioneer memory pioneerInfo = pioneer.pioneerInfo(pioneerAddress_);
        uint256 daysSinceCreate = pioneer.getDay() -
            (pioneerInfo.ctime / pioneer.secondsPerDay()); // 当前距离交完保证金的天数
        if (daysSinceCreate == 31) {} else if (
            daysSinceCreate == 61
        ) {} else if (daysSinceCreate == 91) {} else if (
            daysSinceCreate == 121
        ) {} else if (daysSinceCreate == 151) {} else if (
            daysSinceCreate == 181
        ) {}
        // 获取星级
        uint256 pioneerStar = star.ownerVip(pioneerAddress_);
        // 获取当月新增累加权重
    }

    // 检测考核与保证金退还,每日执行一次,考核失败的城市，可以参与城市节点竞选
    function checkPioneer(
        bytes32 chengShiId_,
        address pioneerAddress_,
        uint256 day
    ) private {
        Pioneer memory pioneerInfo = pioneer.pioneerInfo(pioneerAddress_);
        //        Pioneer storage pioneer = pioneerInfo[pioneerAddress_];
        //        IntoCity city = IntoCity(cityAddress);
        //        uint256 day = getDay() - pioneer.ctime / secondsPerDay;

        // 计算退还保证金额度,并更新退还状态
        //        calculateRefund(chengShiId_, pioneer, city, day);
        calculateRefund(pioneerAddress_, chengShiId_, day);
        // 考核
        if (pioneerInfo.assessmentStatus == true) {
            return;
        }
        // 前三个月考核
        //        assessmentPioneer(chengShiId_, pioneer, city, day);
        assessmentPioneer(pioneerInfo.pioneerAddress, chengShiId_, day);
    }

    // 前三个月考核
    function assessmentPioneer(
        address pioneerAddress_,
        bytes32 chengShiId_,
        //        Pioneer storage pioneer,
        //        IntoCity city,
        uint256 day
    ) private {
        Pioneer memory pioneerInfo = pioneer.pioneerInfo(pioneerAddress_);
        if (pioneerInfo.assessmentStatus) {
            // 已经通过终极考核
            return;
        }
        uint256 pioneerChengShiTotalRechargeWeight = city
            .getChengShiRechargeWeight(chengShiId_) / 1e18; // 先锋绑定的城市总的新增充值权重,这里的值是充值量除以100的权重
        uint256 assessmentCriteriaThreshold; // 考核标准金额
        bool execStatus;
        if (!pioneerInfo.assessmentMonthStatus) {
            // 如果按月考核失败，将不再考核
            return;
        }

        if (day == 90) {
            execStatus = true;
            assessmentCriteriaThreshold = pioneer.assessmentCriteria(
                pioneerInfo.cityLevel,
                3
            );
            //                                assessmentCriteria[ pioneerInfo.cityLevel ][3];
        } else if (day == 60) {
            // 检测是否满足直接考核通过
            if (
                pioneerChengShiTotalRechargeWeight >=
                pioneer.assessmentCriteria(pioneerInfo.cityLevel, 3)
                //                assessmentCriteria[pioneerInfo.cityLevel][3]
            ) {
                //直接达到m3考核标准，也就是直接通过终极考核
                //                pioneer.assessmentStatus = true;
                pioneer.saveAssessmentStatus(pioneerInfo.pioneerAddress, true);
                pioneer.saveAppraiseTime(
                    pioneerInfo.pioneerAddress,
                    block.timestamp,
                    1
                );
                //                successTime[pioneerInfo.pioneerAddress] = block.timestamp;
                return;
            }
            execStatus = true;
            // 没达到M3，考核M2
            assessmentCriteriaThreshold = pioneer.assessmentCriteria(
                pioneerInfo.cityLevel,
                2
            );
            //                                assessmentCriteria[pioneer.cityLevel][2];
        } else if (day == 30) {
            // 检测是否满足直接考核通过
            if (
                pioneerChengShiTotalRechargeWeight >=
                pioneer.assessmentCriteria(pioneerInfo.cityLevel, 3)
                //                assessmentCriteria[pioneerInfo.cityLevel][3]
            ) {
                //直接达到m3考核标准，也就是直接通过终极考核
                //                pioneer.assessmentStatus = true;
                pioneer.saveAssessmentStatus(pioneerInfo.pioneerAddress, true);
                //                successTime[pioneerInfo.pioneerAddress] = block.timestamp;
                pioneer.saveAppraiseTime(
                    pioneerInfo.pioneerAddress,
                    block.timestamp,
                    1
                );
                return;
            }
            execStatus = true;
            // 没达到M3，考核M1
            assessmentCriteriaThreshold = pioneer.assessmentCriteria(
                pioneerInfo.cityLevel,
                1
            );
            //            assessmentCriteria[pioneer.cityLevel][  1 ];
        }
        if (!execStatus) {
            return;
        }
        // 检测是否满足直接考核通过
        if (
            pioneerChengShiTotalRechargeWeight >=
            pioneer.assessmentCriteria(pioneerInfo.cityLevel, 3)
            //            assessmentCriteria[pioneer.cityLevel][3]
        ) {
            //直接达到m3考核标准，也就是直接通过终极考核
            //            pioneer.assessmentStatus = true;
            pioneer.saveAssessmentStatus(pioneerInfo.pioneerAddress, true);
            //            successTime[pioneerInfo.pioneerAddress] = block.timestamp;
            pioneer.saveAppraiseTime(
                pioneerInfo.pioneerAddress,
                block.timestamp,
                1
            );
            return;
        }
        // 检测其他标准考核
        if (pioneerChengShiTotalRechargeWeight < assessmentCriteriaThreshold) {
            //            pioneer.assessmentMonthStatus = false;
            pioneer.saveAssessmentStatus(pioneerInfo.pioneerAddress, false);
            if (pioneer.failedAt(pioneerInfo.pioneerAddress) <= 0) {
                //                failedAt[pioneerInfo.pioneerAddress] = block.timestamp;
                pioneer.saveAppraiseTime(
                    pioneerInfo.pioneerAddress,
                    block.timestamp,
                    2
                );
            }
            //            failedDelegate[chengShiId_] = pioneerChengShiTotalRechargeWeight * 1e18;
            pioneer.saveFailedWeight(
                chengShiId_,
                pioneerChengShiTotalRechargeWeight * 1e18
            );
            city.setChengShiPioneerAssessment(chengShiId_); // 将该城市设置为先锋计划洛选城市
        } else {
            //            successTime[pioneerInfo.pioneerAddress] = block.timestamp;
            pioneer.saveAppraiseTime(
                pioneerInfo.pioneerAddress,
                block.timestamp,
                1
            );
        }
        execStatus = false;
    }

    // 计算退还保证金额度,并更可退还金额
    function calculateRefund(
        address pioneerAddress_,
        bytes32 chengShiId,
        //        Pioneer storage pioneer,
        //        IntoCity city,
        uint256 day
    ) private {
        // 不退还保证金的用户，不再计算
        //        if (isPioneerReturnSurety[pioneer.pioneerAddress]) {
        if (pioneerBatch[pioneerAddress_] > 1) {
            // 只有第一批对保证金
            return;
        }
        uint256 chengLevel = city.chengShiLevel(chengShiId); // 城市等级
        uint256 surety = city.chengShiLevelSurety(chengLevel); // 城市保证金
        uint256 pioneerChengShiTotalRechargeWeight = city
            .getChengShiRechargeWeight(chengShiId) / 1e18; // 先锋绑定的城市总的新充值权重
        uint256 suretyReturn; // 退还保证金的金额
        Pioneer memory pioneerInfo = pioneer.pioneerInfo(pioneerAddress_);
        if (day == 31) {
            // 直接考核通过，退还100%，满足M3退还标准，直接退100%
            for (uint j = 3; j > 0; j--) {
                if (
                    pioneerChengShiTotalRechargeWeight >=
                    pioneer.assessmentReturnCriteria(chengLevel, j)
                ) {
                    //                    pioneer.returnSuretyRate += assessmentReturnRate[chengLevel][j];
                    pioneer.addSuretyRewardRate(
                        pioneerAddress_,
                        pioneer.assessmentReturnRate(chengLevel, j),
                        day / 30
                    );
                    //                    suretyReturn = (surety * pioneer.returnSuretyRate) / 100;
                    pioneerInfo = pioneer.pioneerInfo(pioneerAddress_);
                    suretyReturn =
                        (surety * pioneerInfo.returnSuretyRate) /
                        100;
                    //                    pioneer.returnSuretyStatus = true;
                    //                    pioneer.returnSuretyTime = block.timestamp;
                    //                    alreadyRewardRate[pioneer.pioneerAddress][1] = assessmentReturnRate[chengLevel][j]; // 第一个月退的比例
                    pioneer.addSuretyRewardRate(
                        pioneerAddress_,
                        pioneer.assessmentReturnRate(chengLevel, j),
                        1
                    ); // 设置先锋第一个月退的比例
                    //                    suretyMonthWeight[pioneer.pioneerAddress][1] = pioneerChengShiTotalRechargeWeight; // 第1个月结束的时候，权重值
                    pioneer.saveSuretyMonthWeight(
                        pioneerAddress_,
                        1,
                        pioneerChengShiTotalRechargeWeight
                    ); // 设置第1个月结束的时候，权重值

                    //                    suretyReward[pioneer.pioneerAddress] += suretyReturn; // 增加可退还保证金
                    //                    emit SuretyRecord(
                    //                        pioneer.pioneerAddress,
                    //                        suretyReturn,
                    //                        day / 30
                    //                    );
                    // 增加可退还保证金
                    pioneer.addSuretyReward(
                        pioneerAddress_,
                        suretyReturn,
                        day / 30
                    );
                    break;
                }
            }
        } else if (day == 60) {
            //            uint256 firstMonthRate = alreadyRewardRate[pioneer.pioneerAddress][1];
            uint256 firstMonthRate = pioneer.alreadyRewardRate(
                pioneerAddress_,
                1
            );
            for (uint i = 6; i > 3; i--) {
                if (
                    //                    pioneerChengShiTotalRechargeWeight >= assessmentReturnCriteria[chengLevel][i]
                    pioneerChengShiTotalRechargeWeight >=
                    pioneer.assessmentReturnRate(chengLevel, i)
                ) {
                    //                    if (assessmentReturnRate[chengLevel][i] <= firstMonthRate) {
                    if (
                        pioneer.assessmentReturnRate(chengLevel, i) <=
                        firstMonthRate
                    ) {
                        // 满足退还额度标准的情况下，需要第二个月的退还比例大于第一个月的
                        break;
                    }

                    pioneerInfo = pioneer.pioneerInfo(pioneerAddress_);

                    //                    pioneer.returnSuretyRate +=  assessmentReturnRate[chengLevel][i] - firstMonthRate;
                    //                    pioneer.returnSuretyRate += pioneer.assessmentReturnRate(chengLevel, i) - firstMonthRate;

                    //                    suretyReturn = (surety * (assessmentReturnRate[chengLevel][i] - firstMonthRate)) / 100;
                    suretyReturn =
                        (surety *
                            (pioneer.assessmentReturnRate(chengLevel, i) -
                                firstMonthRate)) /
                        100;
                    //                    pioneer.returnSuretyStatus = true;
                    //                    pioneer.returnSuretyTime = block.timestamp;
                    //                    alreadyRewardRate[pioneer.pioneerAddress][2] = assessmentReturnRate[chengLevel][i] - firstMonthRate; // 第2个月退的比例
                    pioneer.addSuretyRewardRate(
                        pioneerAddress_,
                        pioneer.assessmentReturnRate(chengLevel, i) -
                            firstMonthRate,
                        day / 30
                    ); // 设置先锋第2个月退的比例
                    /// 增加可退还保证金
                    //                    suretyReward[pioneer.pioneerAddress] += suretyReturn; // 增加可退还保证金
                    //                    emit SuretyRecord(
                    //                        pioneer.pioneerAddress,
                    //                        suretyReturn,
                    //                        day / 30
                    //                    );
                    pioneer.addSuretyReward(
                        pioneerAddress_,
                        suretyReturn,
                        day / 30
                    );
                    break;
                }
            }

            //            suretyMonthWeight[pioneer.pioneerAddress][2] = pioneerChengShiTotalRechargeWeight; // 第2个月结束的时候，权重值
            pioneer.saveSuretyMonthWeight(
                pioneerAddress_,
                2,
                pioneerChengShiTotalRechargeWeight
            ); // 第2个月结束的时候，权重值
        }
    }
}
