// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./IntoCity.sol";
import "./RoleAccess.sol";
import "./interface/IERC20.sol";
import "./IntoUserLocation.sol";
import "./interface/IAppraise.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

interface IntoMining {
    function addBalanceWithTypes(
        address sender,
        uint256 amount,
        uint256 types
    ) external; // 增加用户合约余额
}

interface CityPioneerData {
    function subReward(address user_) external returns (uint256);

    function adminSetSubReward(address user_, uint256 reward_) external;

    function adminSubReward(address user_, uint256 reward_) external;
}

contract IntoCityPioneer is RoleAccess, Initializable {
    // 每日新增质押事件
    event DailyIncreaseDelegateRecord(
        address pioneerAddress, // 先锋地址
        bytes32 cityId, //城市ID
        uint256 amount, // 新增质押金额
        uint256 ctime // 新增质押的时间
    );

    // 保证金退还事件
    event SuretyRecord(
        address pioneerAddress, // 先锋地址
        uint256 amount, // 退还保证金的额度
        uint256 month // 第几个月（1，2）
    );

    // 城市先锋奖励事件
    event DailyRewardRecord(
        address pioneerAddress, // 先锋地址
        uint256 toxReward, // 赠送质押包奖励的额度
        uint256 foundsReward, // 社交基金奖励的额度
        uint256 delegateReward // 新增质押奖励的额度
    );
    event DailyRewardRecordV2(
        address pioneerAddress, // 先锋地址
        uint256 toxReward, // 赠送质押包奖励的额度
        uint256 foundsReward, // 社交基金奖励的额度
        uint256 delegateReward // 新增质押奖励的额度
    );

    // 城市先锋提取奖励事件
    event WithdrawalRewardRecord(
        address pioneerAddress, // 先锋地址
        uint256 reward, // 提现的奖励数量
        uint256 rewardType // 提现的奖励的类型. 1 福袋，2 社交基金， 3 质押奖励，4 保证金奖励
    );

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

    address public TOXAddress; // TOX代币合约地址
    address public miningAddress; // 用户增加合约余额合约
    address public cityAddress; // 城市合约地址

    // 城市等级 => (月份 => 质押点数);考核标准
    mapping(uint256 => mapping(uint256 => uint256)) public assessmentCriteria; // 城市先锋考核标准
    // 城市等级 => (索引 => 退还标准)，索引1，2，3为第一个月的，索引4，5，6为第二个月的；退保证金标准（点数）
    mapping(uint256 => mapping(uint256 => uint256))
        public assessmentReturnCriteria; // 城市先锋保证金退还标准
    // 城市等级 => (索引 => 退还比例)，索引1，2，3为第一个月的，索引4，5，6为第二个月的；退保证金标准（比例）
    mapping(uint256 => mapping(uint256 => uint256)) public assessmentReturnRate; // 城市先锋保证金退还比例
    // 先锋地址 => 先锋信息， 先锋信息
    mapping(address => Pioneer) public pioneerInfo;
    // 先锋地址 => 先锋福利包收益
    mapping(address => uint256) public benefitPackageReward;
    // 先锋地址 => 已领取的先锋福利包收益
    mapping(address => uint256) public benefitPackageRewardReceived;
    // 先锋地址 => 先锋社交基金收益
    mapping(address => uint256) public fundsReward;
    // 先锋地址 => 已领取的先锋社交基金收益
    mapping(address => uint256) public fundsRewardReceived;
    // 先锋地址 => 先锋新增质押收益
    mapping(address => uint256) public delegateReward;
    // 先锋地址 => 已领取的先锋新增质押收益
    mapping(address => uint256) public delegateRewardReceived;
    // 先锋地址 => 先锋可以退的保证金
    mapping(address => uint256) public suretyReward;
    // 先锋地址 => (月份=>退的比例)，月份为1和2，比例为整数
    mapping(address => mapping(uint256 => uint256)) public alreadyRewardRate;
    // 先锋地址 => 先锋已经退还的保证金
    mapping(address => uint256) public suretyRewardRecord;
    // 先锋城市ID => 先锋考核失败时候的累计充值权重（绑定城市的累计充值权重）
    mapping(bytes32 => uint256) public failedDelegate;
    // 每天定时任务执行状态
    mapping(uint256 => bool) public dailyTaskStatus;
    mapping(address => bool) public benefitPackageRewardStatus; // 用户福袋奖励提取状态
    mapping(address => bool) public fundsRewardStatus; // 用户社交基金奖励提取状态
    mapping(address => bool) public delegateRewardStatus; // 用户新增质押奖励提取状态
    // 城市先锋地址=>(天数=>是否执行过)
    mapping(address => mapping(uint256 => bool)) public checkPioneerStatus; //每天一次考核，废弃
    // 城市先锋地址=>(天数=>是否执行过)
    mapping(address => mapping(uint256 => bool)) public checkRewardStatus; //每天一次奖励
    uint256 public startTime; //开始考核时间
    address[] public pioneers; // 所有先锋的集合
    // 用户定位合约地址
    address public userLocationAddress;
    // 考核天，正式86400秒，测试300秒
    uint public secondsPerDay;
    // 考核天，正式86400*180秒，测试10*180秒
    uint public presidencyTime;
    // 先锋地址 => 已经退的比例，领取保证金时，累加
    mapping(address => uint256) public alreadyRewardRateTotal;
    // 先锋地址 => 考核失败时间
    mapping(address => uint256) public failedAt;
    // 先锋地址 => (类别=>考核成功时间)，类别（1第一个月，2第二个月，3第三个月，4第一个月直接达到最高标准通过，后面不在考核）
    mapping(address => mapping(uint256 => uint256)) public successAt; // 废弃
    // 先锋地址 => 考核成功时间
    mapping(address => uint256) public successTime;
    // 先锋地址 => （月份=>权重值）；权重值，计算可领取保证金时刻的充值权重
    mapping(address => mapping(uint => uint)) public suretyMonthWeight;
    mapping(address => mapping(uint => uint)) public testValue; //用于测试
    // 天数=>(城市先锋地址=>是否执行过)
    mapping(uint256 => mapping(address => bool)) public checkPioneerDailyStatus; //每天一次考核
    // 城市先锋地址=>需要补交保证金数量
    mapping(address => uint256) public pioneerPaySurety;
    // 城市先锋地址=>是否退还保证金（需求更新，后面增加的先锋不退还保证金），true的话不退还保证金
    mapping(address => bool) public isPioneerReturnSurety;
    CityPioneerData public cityPioneerData; // 先锋数据合约
    IAppraise public appraise; // 先锋考核合约

    //            function initialize() public initializer {
    //                _addAdmin(msg.sender);
    //            }

    // 管理员设置TOX代币地址
    //    function adminSetTOXAddress(address TOXAddress_) public onlyAdmin {
    //        TOXAddress = TOXAddress_;
    //    }

    // 管理员设置数据合约地址
    //    function adminSetCityPioneerDataAddress(
    //        address cityPioneerData_
    //    ) public onlyAdmin {
    //        cityPioneerData = CityPioneerData(cityPioneerData_);
    //    }

    // 管理员设置城市合约地址
    //    function adminSetCityAddress(address cityAddress_) public onlyAdmin {
    //        cityAddress = cityAddress_;
    //    }

    // 管理员设置考核合约
    //    function adminSetAppraise(address appraiseAddress_) public onlyAdmin {
    //        appraise = IAppraise(appraiseAddress_);
    //    }

    //    // 管理员设置用户定位合约地址
    //    function adminSetUserLocationAddress(address userLocationAddress_) public onlyAdmin {
    //        userLocationAddress = userLocationAddress_;
    //    }

    // 管理员设置增加用户合约余额合约地址
    //    function adminSetMiningAddress(address miningAddress_) public onlyAdmin {
    //        miningAddress = miningAddress_;
    //    }
    //
    //    // 管理员设置考核标准,月份为1，2，3
    //    function adminSetAssessmentCriteria(uint256 cityLevel_, uint256 month_, uint256 point_) public onlyAdmin {
    //        assessmentCriteria[cityLevel_][month_] = point_;
    //    }
    //
    // 管理员设置先锋定时任务是否执行过
    //    function adminSetCheckPioneerDailyStatus(address pioneer_, uint256 day_, bool execution_) public onlyAdmin {
    //        checkPioneerDailyStatus[day_][pioneer_] = execution_;
    //    }
    //
    //    // 管理员设置城市先锋保证金退还标准；点数
    //    function adminSetAssessmentReturnCriteria(uint256 chengShiLevel_, uint256 month_, uint256 point_) public onlyAdmin {
    //        assessmentReturnCriteria[chengShiLevel_][month_] = point_;
    //    }
    //   // 管理员设置城市先锋保证金退还标准，比例
    //    function adminSetAssessmentReturnRate(uint256 chengShiLevel_, uint256 month_, uint256 point_) public onlyAdmin {
    //        assessmentReturnRate[chengShiLevel_][month_] = point_;
    //    }
    // 管理员设置城市先锋保证金退还标准，比例
    //    function adminEditSuretyReward(address pioneerAddress_, uint256 amount_) public onlyAdmin {
    //        suretyReward[pioneerAddress_] -= amount_;
    //    }

    // 管理员设置开始考核时间
    //    function adminSetStartTime(uint256 startTime_) public onlyAdmin {
    //        startTime = startTime_;
    //        // 将已经交保证金的先锋，重置开始考核时间
    //        for (uint256 i; i < pioneers.length; i++) {
    //            Pioneer storage pioneer = pioneerInfo[pioneers[i]];
    //            pioneer.ctime = startTime;
    //        }
    //    }

    // 管理员设置先锋是否退还保证金
    //    function adminSetPioneerReturnSurety(address pioneer_, bool pay_) public onlyAdmin {
    //        isPioneerReturnSurety[pioneer_] = pay_;
    //    }

    // 管理员设置每天秒数，用于测试
    //    function adminSetSecondsPerDay(uint56 secondsPerDay_) public onlyAdmin {
    //        secondsPerDay = secondsPerDay_;
    //        presidencyTime = secondsPerDay_ * 180;
    //    }

    // 管理员设置任期
    //    function adminSetPresidencyTime(uint56 presidencyTime_) public onlyAdmin {
    //        presidencyTime = presidencyTime_;
    //    }

    // 考核和奖励任务
    function pioneerTask(
        address pioneerAddress_,
        bytes32 chengShiId_
    ) public onlyAdmin {
        // 更新奖励和考核
        uint256 today = getDay();
        if (!checkPioneerDailyStatus[today][pioneerAddress_]) {
            // 考核
            checkPioneer(chengShiId_, pioneerAddress_);
            // 奖励
            reward(chengShiId_, pioneerAddress_);
            checkPioneerDailyStatus[today][pioneerAddress_] = true;
        }
    }

    function getDay() public view returns (uint256) {
        uint day = block.timestamp / secondsPerDay;
        return uint256(day);
    }

    // 缴纳保证金
    function depositSurety() public {
        IntoCity city = IntoCity(cityAddress);
        // 查询调用者地址是否是城市先锋
        bytes32 chengShiId = city.pioneerChengShi(msg.sender);
        require(chengShiId != bytes32(0), "you are not pioneer"); // 不是城市先锋
        // 查询该城市先锋对应的城市，根据城市查询需要缴纳的保证金数量
        uint256 surety = city.surety(chengShiId);
        IERC20 TOXContract = IERC20(TOXAddress);
        uint256 userBalance = TOXContract.balanceOf(msg.sender);

        require(userBalance >= surety, "your balance is insufficient"); // 余额不足
        TOXContract.transferFrom(msg.sender, address(this), surety);
        pioneerInfo[msg.sender].pioneerAddress = msg.sender;
        pioneerInfo[msg.sender].assessmentMonthStatus = true;
        pioneerInfo[msg.sender].ctime = block.timestamp;
        pioneerInfo[msg.sender].suretyTime = block.timestamp;
        pioneerInfo[msg.sender].cityLevel = city.chengShiLevel(chengShiId);
        failedDelegate[chengShiId] = 0;
        if (chengShiId != bytes32(0)) {
            city.initCityRechargeWeight(chengShiId); // 将先锋绑定的城市的新增质押量变为0
        }
        // 更新已交保证金先锋数量
        pioneers.push(msg.sender);
    }

    // 修改先锋信息
    //    function editPioneerInfo(address pioneerAddress_) public onlyAdmin {
    //        isPioneerReturnSurety[pioneerAddress_] = false;
    //    }

    // 管理员设置先锋需要补交的保证金
    //    function adminSetPioneerPaySurety(address pioneer_, uint256 amount_) public onlyAdmin {
    //        pioneerPaySurety[pioneer_] = amount_;
    //    }

    // 检查先锋是否需要补交保证金
    function checkSurety(address pioneer_) public view returns (uint256) {
        return pioneerPaySurety[pioneer_];
    }

    // 先锋补交保证金
    function paySurety() public {
        require(pioneerPaySurety[msg.sender] > 0, "you have already paid");
        IERC20 TOXContract = IERC20(TOXAddress);
        TOXContract.transferFrom(
            msg.sender,
            address(this),
            pioneerPaySurety[msg.sender]
        );
        pioneerPaySurety[msg.sender] = 0;
    }

    function setIsPioneerReturnSurety(address pioneer_) public onlyAdmin {
        isPioneerReturnSurety[pioneer_] = true;
        // 设置先锋批次
        appraise.setPioneerBatch(pioneer_, 3);
    }

    function removePioneer(address pioneer_) public onlyAdmin {
        //        for (uint256 i = 0; i < pioneers.length; i++) {
        //            if (pioneer_ == pioneers[i]) {
        //                pioneers[i] = pioneers[pioneers.length - 1];
        //                pioneers.pop();
        //            }
        //        }
        //        pioneerInfo[pioneer_].assessmentStatus = false;
        //        alreadyRewardRate[pioneer_][1] = 0;
        //        alreadyRewardRate[pioneer_][2] = 0;
        //        alreadyRewardRateTotal[pioneer_] = 0;
        //
        //        // 先锋地址 => 先锋福利包收益
        //        benefitPackageReward[pioneer_] = 0;
        //        // 先锋地址 => 已领取的先锋福利包收益
        //        benefitPackageRewardReceived[pioneer_] = 0;
        //        // 先锋地址 => 先锋社交基金收益
        //        fundsReward[pioneer_] = 0;
        //        // 先锋地址 => 已领取的先锋社交基金收益
        //        fundsRewardReceived[pioneer_] = 0;
        //        // 先锋地址 => 先锋新增质押收益
        //        delegateReward[pioneer_] = 0;
        //        // 先锋地址 => 已领取的先锋新增质押收益
        //        delegateRewardReceived[pioneer_] = 0;
        //        // 先锋地址 => 先锋可以退的保证金
        //        suretyReward[pioneer_] = 0;
        //        // 先锋地址 => (月份=>退的比例)，月份为1和2，比例为整数
        //        alreadyRewardRate[pioneer_][1] = 0;
        //        alreadyRewardRate[pioneer_][2] = 0;
        //        // 先锋地址 => 先锋已经退还的保证金
        //        suretyRewardRecord[pioneer_] = 0;
        //        benefitPackageRewardStatus[pioneer_] = false; // 用户福袋奖励提取状态
        //        fundsRewardStatus[pioneer_] = false; // 用户社交基金奖励提取状态
        //        delegateRewardStatus[pioneer_] = false; // 用户新增质押奖励提取状态
        //        failedAt[pioneer_] = 0; // 清楚失败时间
        //        IntoCity city = IntoCity(cityAddress);
        //        failedDelegate[city.pioneerChengShi(pioneer_)] = 0; // 先锋考核失败时候的累计充值权重
        //        successTime[pioneer_] = 0; // 清楚成功时间
        //        suretyMonthWeight[pioneer_][1] = 0; // 清楚成功时间
        //        suretyMonthWeight[pioneer_][2] = 0; // 清楚成功时间
        //        appraise.delPioneerPreMonthWeight(pioneer_);
    }

    // 检测考核与保证金退还,每日执行一次,考核失败的城市，可以参与城市节点竞选
    function checkPioneer(
        bytes32 chengShiId_,
        address pioneerAddress_
    ) private {
        Pioneer storage pioneer = pioneerInfo[pioneerAddress_];
        IntoCity city = IntoCity(cityAddress);
        uint256 day = getDay() - pioneer.ctime / secondsPerDay;

        // 考核
        if (appraise.pioneerBatch(pioneerAddress_) == 3) {
            // 第三批次考核
            bool dataChange;
            bool appraiseStatus;
            uint256 month;
            uint256 chengShiRechargeWeight;
            (
                dataChange,
                appraiseStatus,
                month,
                chengShiRechargeWeight
            ) = appraise.appraiseBeth3(pioneer.pioneerAddress, chengShiId_);
            if (dataChange) {
                if (!appraiseStatus) {
                    // 考核失败
                    // 更新考核失败状态
                    pioneer.assessmentMonthStatus = false;
                    // 更新考核失败时间
                    if (failedAt[pioneer.pioneerAddress] <= 0) {
                        failedAt[pioneer.pioneerAddress] = block.timestamp;
                    }
                    // 先锋当前充值权重
                    failedDelegate[chengShiId_] = chengShiRechargeWeight;
                    // 将该城市设置为先锋计划洛选城市
                    city.setChengShiPioneerAssessment(chengShiId_);
                } else {
                    successTime[pioneer.pioneerAddress] = block.timestamp;
                }
            }
        } else {
            // 计算退还保证金额度,并更新退还状态
            calculateRefund(chengShiId_, pioneer, city, day);
            /// 前三个月考核，第一二批次
            if (pioneer.assessmentStatus == true) {
                return;
            }
            // 前三个月考核
            assessmentPioneer(chengShiId_, pioneer, city, day);
        }
    }

    // 前三个月考核
    function assessmentPioneer(
        bytes32 chengShiId_,
        Pioneer storage pioneer,
        IntoCity city,
        uint256 day
    ) private {
        if (pioneer.assessmentStatus) {
            // 已经通过终极考核
            return;
        }
        uint256 pioneerChengShiTotalRechargeWeight = city
            .getChengShiRechargeWeight(chengShiId_) / 1e18; // 先锋绑定的城市总的新增充值权重,这里的值是充值量除以100的权重
        uint256 assessmentCriteriaThreshold; // 考核标准金额
        bool execStatus;
        if (!pioneer.assessmentMonthStatus) {
            // 如果按月考核失败，将不再考核
            return;
        }

        if (day == 90) {
            execStatus = true;
            assessmentCriteriaThreshold = assessmentCriteria[pioneer.cityLevel][
                3
            ];
        } else if (day == 60) {
            // 检测是否满足直接考核通过
            if (
                pioneerChengShiTotalRechargeWeight >=
                assessmentCriteria[pioneer.cityLevel][3]
            ) {
                //直接达到m3考核标准，也就是直接通过终极考核
                pioneer.assessmentStatus = true;
                successTime[pioneer.pioneerAddress] = block.timestamp;
                return;
            }
            execStatus = true;
            // 没达到M3，考核M2
            assessmentCriteriaThreshold = assessmentCriteria[pioneer.cityLevel][
                2
            ];
        } else if (day == 30) {
            // 检测是否满足直接考核通过
            if (
                pioneerChengShiTotalRechargeWeight >=
                assessmentCriteria[pioneer.cityLevel][3]
            ) {
                //直接达到m3考核标准，也就是直接通过终极考核
                pioneer.assessmentStatus = true;
                successTime[pioneer.pioneerAddress] = block.timestamp;
                return;
            }
            execStatus = true;
            // 没达到M3，考核M1
            assessmentCriteriaThreshold = assessmentCriteria[pioneer.cityLevel][
                1
            ];
        }
        if (!execStatus) {
            return;
        }
        // 检测是否满足直接考核通过
        if (
            pioneerChengShiTotalRechargeWeight >=
            assessmentCriteria[pioneer.cityLevel][3]
        ) {
            //直接达到m3考核标准，也就是直接通过终极考核
            pioneer.assessmentStatus = true;
            successTime[pioneer.pioneerAddress] = block.timestamp;
            return;
        }
        // 检测其他标准考核
        if (pioneerChengShiTotalRechargeWeight < assessmentCriteriaThreshold) {
            pioneer.assessmentMonthStatus = false;
            if (failedAt[pioneer.pioneerAddress] <= 0) {
                failedAt[pioneer.pioneerAddress] = block.timestamp;
            }
            failedDelegate[chengShiId_] =
                pioneerChengShiTotalRechargeWeight *
                1e18;
            city.setChengShiPioneerAssessment(chengShiId_); // 将该城市设置为先锋计划洛选城市
        } else {
            successTime[pioneer.pioneerAddress] = block.timestamp;
        }
        execStatus = false;
    }

    // 计算退还保证金额度,并更可退还金额
    function calculateRefund(
        bytes32 chengShiId,
        Pioneer storage pioneer,
        IntoCity city,
        uint256 day
    ) private {
        if (isPioneerReturnSurety[pioneer.pioneerAddress]) {
            // 不退还保证金的用户，不再计算
            return;
        }
        uint256 chengLevel = city.chengShiLevel(chengShiId); // 城市等级
        uint256 surety = city.chengShiLevelSurety(chengLevel); // 城市保证金
        uint256 pioneerChengShiTotalRechargeWeight = city
            .getChengShiRechargeWeight(chengShiId) / 1e18; // 先锋绑定的城市总的新充值权重
        uint256 suretyReturn; // 退还保证金的金额
        if (day == 30) {
            // 直接考核通过，退还100%，满足M3退还标准，直接退100%
            for (uint j = 3; j > 0; j--) {
                if (
                    pioneerChengShiTotalRechargeWeight >=
                    assessmentReturnCriteria[chengLevel][j]
                ) {
                    pioneer.returnSuretyRate += assessmentReturnRate[
                        chengLevel
                    ][j];
                    suretyReturn = (surety * pioneer.returnSuretyRate) / 100;
                    pioneer.returnSuretyStatus = true;
                    pioneer.returnSuretyTime = block.timestamp;
                    alreadyRewardRate[pioneer.pioneerAddress][
                        1
                    ] = assessmentReturnRate[chengLevel][j]; // 第一个月退的比例
                    suretyMonthWeight[pioneer.pioneerAddress][
                        1
                    ] = pioneerChengShiTotalRechargeWeight; // 第1个月结束的时候，权重值
                    suretyReward[pioneer.pioneerAddress] += suretyReturn; // 增加可退还保证金
                    emit SuretyRecord(
                        pioneer.pioneerAddress,
                        suretyReturn,
                        day / 30
                    );
                    break;
                }
            }
        } else if (day == 60) {
            uint256 firstMonthRate = alreadyRewardRate[pioneer.pioneerAddress][
                1
            ];
            for (uint i = 6; i > 3; i--) {
                if (
                    pioneerChengShiTotalRechargeWeight >=
                    assessmentReturnCriteria[chengLevel][i]
                ) {
                    if (assessmentReturnRate[chengLevel][i] <= firstMonthRate) {
                        // 满足退还额度标准的情况下，需要第二个月的退还比例大于第一个月的
                        break;
                    }
                    pioneer.returnSuretyRate +=
                        assessmentReturnRate[chengLevel][i] -
                        firstMonthRate;
                    suretyReturn =
                        (surety *
                            (assessmentReturnRate[chengLevel][i] -
                                firstMonthRate)) /
                        100;
                    pioneer.returnSuretyStatus = true;
                    pioneer.returnSuretyTime = block.timestamp;
                    alreadyRewardRate[pioneer.pioneerAddress][2] =
                        assessmentReturnRate[chengLevel][i] -
                        firstMonthRate; // 第2个月退的比例
                    suretyReward[pioneer.pioneerAddress] += suretyReturn; // 增加可退还保证金
                    emit SuretyRecord(
                        pioneer.pioneerAddress,
                        suretyReturn,
                        day / 30
                    );
                    break;
                }
            }
            suretyMonthWeight[pioneer.pioneerAddress][
                2
            ] = pioneerChengShiTotalRechargeWeight; // 第2个月结束的时候，权重值
        }
    }

    // 每日奖励发放
    function reward(bytes32 ChengShiId_, address pioneerAddress_) private {
        Pioneer storage pioneer = pioneerInfo[pioneerAddress_];
        // 考核状态（按月考核）不正常，不发放奖励
        if (!pioneer.assessmentMonthStatus) {
            return;
        }
        uint256 today = getDay();
        // 任期结束，不发放奖励
        if (
            today - (pioneer.ctime / secondsPerDay) >=
            (presidencyTime / secondsPerDay)
        ) {
            return;
        }
        // 昨日，天
        uint256 yesterday = getDay() - 1;
        /// 浙江丽水的先锋发放收益地址由 0x12Cc2278b37c2751D11B5A64712Ff439b57F6E6a更改为0x5db860601869Dad7Eb2961341056b389C3149e5f
        address rewardAddress = pioneerAddress_;
        if (pioneerAddress_ == 0x12Cc2278b37c2751D11B5A64712Ff439b57F6E6a) {
            rewardAddress = 0x5db860601869Dad7Eb2961341056b389C3149e5f;
        }
        ///

        // 福利包奖励
        uint bonus = 93333333333333333333;
        benefitPackageReward[rewardAddress] += bonus;

        // 社交基金5%奖励
        IntoCity city = IntoCity(cityAddress);
        uint pioneerAndCityNodeNumber = city.getPioneerAndCityNodeNumber(); // 城市节点上线后，需要加上城市节点的数量
        uint256 allDailyFoundsTotal = city.getFifteenDayAverageFounds(); // 全网昨日所有城市新增社交基金
        uint256 f;
        if (pioneerAndCityNodeNumber == 0) {
            fundsReward[rewardAddress] += 0;
        } else {
            fundsReward[rewardAddress] +=
                (allDailyFoundsTotal * 5) /
                100 /
                pioneerAndCityNodeNumber;
            f = (allDailyFoundsTotal * 5) / 100 / pioneerAndCityNodeNumber;
        }

        // 该城市新增质押量1%奖励，不累加
        uint256 yesterdayDelegate = city.getDailyDelegateByChengShiID(
            ChengShiId_,
            yesterday
        ); // 昨日该城市新增质押权重
        delegateReward[rewardAddress] += yesterdayDelegate / 100;

        address pioneerAddress = pioneerAddress_;
        emit DailyRewardRecord(
            pioneerAddress,
            bonus,
            fundsReward[pioneerAddress],
            delegateReward[pioneerAddress]
        );
        emit DailyRewardRecordV2(
            pioneerAddress,
            bonus,
            f,
            yesterdayDelegate / 100
        );
    }

    //     用户提取福利包奖励
    function withdrawalBenefitPackageReward() public {
        //        require(false, "withdraw service paused");
        // 判断是否是先锋
        //        require(pioneerInfo[msg.sender].ctime > 0, "you are not pioneer");
        require(benefitPackageReward[msg.sender] > 0, "balance insufficient");
        require(pioneerPaySurety[msg.sender] >= 0, "you need pay surety");
        uint256 balance = benefitPackageReward[msg.sender];

        balance = checkBalance(msg.sender, balance);

        benefitPackageReward[msg.sender] = 0; // 更新可领取福利包奖励
        // 将奖励转账到用户合约余额
        setUserBalance(msg.sender, balance, 19);
        // 更新领取状态(全部领完才算已领取)
        benefitPackageRewardStatus[msg.sender] = true;
        benefitPackageRewardReceived[msg.sender] += balance; // 更新已领取福利包奖励
        emit WithdrawalRewardRecord(msg.sender, balance, 1);
    }

    // 用户提取社交基金奖励
    function withdrawalFundsReward() public {
        //        require(false, "withdraw service paused");
        // 判断是否是先锋
        //        require(pioneerInfo[msg.sender].ctime > 0, "you are not pioneer");
        require(fundsReward[msg.sender] > 0, "balance insufficient");
        require(pioneerPaySurety[msg.sender] >= 0, "you need pay surety");
        uint256 balance = fundsReward[msg.sender];

        balance = checkBalance(msg.sender, balance);

        fundsReward[msg.sender] = 0; // 更新可领取社交基金奖励
        // 将奖励转账到用户合约余额
        setUserBalance(msg.sender, balance, 18);
        // 更新领取状态(全部领完才算已领取)
        fundsRewardStatus[msg.sender] = true;
        fundsRewardReceived[msg.sender] += balance; // 更新已领取社交基金奖励
        emit WithdrawalRewardRecord(msg.sender, balance, 2);
    }

    // 用户提取新增质押奖励
    function withdrawalDelegateReward() public {
        //        require(false, "withdraw service paused");
        // 判断是否是先锋
        //        require(pioneerInfo[msg.sender].ctime > 0, "you are not pioneer");
        require(delegateReward[msg.sender] > 0, "balance insufficient");
        require(pioneerPaySurety[msg.sender] >= 0, "you need pay surety");
        uint256 balance = delegateReward[msg.sender];

        balance = checkBalance(msg.sender, balance);

        delegateReward[msg.sender] = 0; // 更新可领取新增质押奖励
        // 将奖励转账到用户合约余额
        setUserBalance(msg.sender, balance, 17);
        delegateRewardStatus[msg.sender] = true;
        delegateRewardReceived[msg.sender] += balance; // 更新已领取新增质押奖励
        emit WithdrawalRewardRecord(msg.sender, balance, 3);
    }

    // 用户提取退还的保证金
    function withdrawalSuretyReward() public {
        //        require(false, "withdraw service paused");
        // 判断是否是先锋
        //        require(pioneerInfo[msg.sender].ctime > 0, "you are not pioneer");
        require(suretyReward[msg.sender] > 0, "balance insufficient");
        require(pioneerPaySurety[msg.sender] >= 0, "you need pay surety");
        uint256 balance = suretyReward[msg.sender];

        balance = checkBalance(msg.sender, balance);
        // 更新可退还保证金数额
        suretyReward[msg.sender] = 0;
        // 退还保证金到钱包账户
        IERC20 TOX = IERC20(TOXAddress);
        TOX.transfer(msg.sender, balance);
        // 更新已领取的数量
        suretyRewardRecord[msg.sender] += balance;
        // 更新已领取的比例
        alreadyRewardRateTotal[msg.sender] = pioneerInfo[msg.sender]
            .returnSuretyRate;
        // alreadyRewardRate

        emit WithdrawalRewardRecord(msg.sender, balance, 4);
    }

    function checkBalance(
        address user,
        uint256 balance
    ) private returns (uint256) {
        uint256 subReward = cityPioneerData.subReward(user);
        uint256 subRes;
        if (balance >= subReward) {
            balance = balance - subReward;
            cityPioneerData.adminSubReward(user, subReward);
            subRes = subReward;
        } else {
            cityPioneerData.adminSubReward(user, balance);
            subRes = balance;
            balance = 0;
        }
        if (subRes > 0) {
            emit WithdrawalRewardRecord(msg.sender, balance, 1000);
        }
        return balance;
    }

    // 将奖励转账到用户合约余额 17 节点 18 基金 19 福利
    function setUserBalance(
        address user_,
        uint256 amount_,
        uint256 types_
    ) private {
        IntoMining intoMining = IntoMining(miningAddress);
        intoMining.addBalanceWithTypes(user_, amount_, types_);
    }

    // 获取先锋所需缴纳的保证金
    function getSurety() public view returns (uint256) {
        IntoCity city = IntoCity(cityAddress);
        // 获取先锋所在城市
        bytes32 cityId = city.pioneerCity(msg.sender);
        // 获取该城市所需要缴纳的保证金
        return city.surety(cityId);
    }

    // 判断一个地址是否是先锋
    function isPioneer(address user) public view returns (bool) {
        return pioneerInfo[user].ctime > 0;
    }

    // 获取先锋所需保证金，根据先锋地址
    function getPioneerNumber() public view returns (uint256) {
        return pioneers.length;
    }

    // 初始化先锋用户
    function initPioneer(address pioneer) public {
        delete pioneerInfo[pioneer];
    }
}
