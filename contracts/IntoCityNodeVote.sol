// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./RoleAccess.sol";
import "./IntoUserLocation.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract IntoCityNodeVote is RoleAccess, Initializable {

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

    // 投票记录
    event VoteRecord(
        address user,
        address candidate,
        bytes32 cityId,
        uint256 toxNum
    );

    enum VoteStatus{
        none, // 未开始
        globalVoting,// 全球节点投票进行中
        globalVoteEnd, // 全球节点投票结束
        userVoting, // 用户投票进行中
        userVoteEnd, // 用户投票结束
        incumbency,// 任期
        endOfTerm// 任期结束
    }

    struct NodeInfo {
        address node;
        string avatar;
        string nodeName;
        string flag;
        uint256 status; //0 不存在 1未审核，2已审核，3审核失败
    }

    struct CandidateInfo {
        bytes32 cityId; // 城市ID
        address candidate; // 参选人地址
        uint256 rate; //分配比例
        uint256 votes; //当前累计投票数量
    }

    struct CityInfo {
        bytes32 cityId; // 城市ID
        uint256 userNumber; // 该城市定位的人数
        uint256 votes; //当前累计投票数量
        uint256 startTime; // 城市开始竞选时间
        uint256 cityLevel; // 城市等级
        uint256 status; // 0竞选未开始，1竞选已开始，2竞选已结束
    }

    address public TOX;
    // 用户定位合约地址
    address public userLocationAddress;
    uint256 public startTime; // 投票开始时间（按秒算的时间戳）
    uint256 public applyCityFee = 10000 * 1e18; // 城市节点报名费
    VoteStatus  public voteStatus; // 投票状态

    address[] public allCityNode; // 所有城市节点
    bytes32[] public allCityIds; // 所有城市
    mapping(address => NodeInfo) public nodeInfo; // 参选人信息（节点信息）
    // 城市ID=>CityInfo
    mapping(bytes32 => CityInfo) public cityInfo;
    // 城市ID=>cityLevel
    mapping(bytes32 => uint256) public cityLevel;
    // 城市ID=>城市节点开始投票时间
//    mapping(bytes32 => uint256) public cityStartTime;
    // 城市等级 => (类别 => 数值) ， 类别（1用户阀值，2质押量阀值）
    mapping(uint256 => mapping(uint256 => uint256))public  activityCriteria; // 城市激活标准
    // 城市ID => 该城市所有申请竞选的费用
    mapping(bytes32 => uint256) public cityApplyFee;
    // 城市ID => (投票人地址 => (候选人地址 => 投票数量)) 投票数量为投票结束，追加投票
    mapping(bytes32 => mapping(address => mapping(address => uint256))) public userVoteRecord;
    // 城市ID => (投票人地址 => (候选人地址 => 投票数量)) 投票数量为投票期间正常投票
    mapping(bytes32 => mapping(address => mapping(address => uint256))) public userVoteVoteEndRecord;
    // 城市ID => (候选人地址 => 候选人当前票数)，候选人累计票数,包含竞选结束的
    mapping(bytes32 => mapping(address => uint256)) public candidateCount;
    // 城市ID => (候选人地址 => 候选人当前票数)，候选人累计票数,只累加竞选期间的
    mapping(bytes32 => mapping(address => uint256)) public candidateVotingCount;
    // 候选人地址 => 候选人参选时间
    mapping(address => uint256) public candidateApplyTime;
    // 候选人地址 => 候选人参选的城市ID
//    mapping(address => bytes32) public candidateApplyCityId;
    // 城市ID => (候选人地址 => 候选人当前票数)，候选人截止当选成功累计票数
    mapping(bytes32 => mapping(address => uint256)) public candidateVoteEndCount;
    // 参选人地址=>参选人城市; 一个人只能参选一个城市，下一次可以更换城市
    mapping(address => bytes32) public candidateApplyCityId;
    // 城市ID => []候选人地址
    mapping(bytes32 => address[]) public cityCandidate;
    // 城市ID => 投票数量(该城市累计投票数量),包含竞选结束的
    mapping(bytes32 => uint256) public cityVotes;
    // 城市ID => 数组，参与该城市竞选的全球节点地址
    mapping(bytes32 => address[]) public globalNodeCityCandidate;
    // 城市ID => 数组，参与该城市投票的用户地址
    mapping(bytes32 => address[]) public cityVoteUsers;
    // 全球节点地址 => 该全球节点的质押量（票数）
    mapping(address => uint256) public globalNodeCandidateVotes;
    // 城市节点 => control rate（收益支配比例）
    mapping(address => uint256) public cityNodeControlRate;
    // 城市ID=>全球节点竞选者票数最高者地址（当选者地址，只要不为空，就说明第一阶段有当选成功者）
    mapping(bytes32 => address)  public cityNodeGlobalNode;
    // 城市ID=>全球节点竞选者票数最高值（最大票数）
    mapping(bytes32 => uint256)  public cityNodeGlobalNodeMaxVotes;
    // 城市节点 => 票数最多的候选人(没值的话就是没有候选人)
    mapping(bytes32 => address) public cityNodeWinner;
    // 城市节点 => 候选人最大票数
    mapping(bytes32 => uint256) public cityNodeMaxVotes;
    // 城市ID=>(参选人地址=>追加票数)
    mapping(bytes32 => mapping(address => uint256)) public additionalVoting;
    // 城市ID=>(候选人地址=>投票结束排名)
    mapping(bytes32 => mapping(address => uint256)) public votingRanking;
    // 天，正式86400秒，测试300秒
    uint public secondsPerDay;

    function initialize() public initializer {
        _addAdmin(msg.sender);
        startTime = block.timestamp;
    }

    // 管理员设置用户定位合约地址
    function adminSetUserLocationAddress(address userLocationAddress_) public onlyAdmin {
        userLocationAddress = userLocationAddress_;
    }

    // 管理员设置投票状态
    function adminSetVoteStatus(VoteStatus voteStatus_) public onlyAdmin {
        voteStatus = voteStatus_;
        emit VoteStatusChangeRecord(voteStatus_);
    }

    // 管理员设置申请城市节点所需费用
    function adminSetApplyCityFee(uint256 applyCityFee_) public onlyAdmin {
        applyCityFee = applyCityFee_;
    }
    // 管理员设置投票开始时间
    function adminSetStartTime(uint256 startTime_) public onlyAdmin {
        startTime = startTime_;
    }

    // 管理员设置TOX代币地址
    function adminSetTOXAddress(address TOX_) public onlyAdmin {
        TOX = TOX_;
    }
    // 管理员设置城市激活标准,// 城市等级 => (类别 => 数值) ， 类别（1用户阀值，2质押量阀值）
    function adminSetActivityCriteria(uint256 cityLevel_, uint256 type_, uint256 val_) public onlyAdmin {
        activityCriteria[cityLevel_][type_] = val_;
    }
    // 管理员添加全球节点
    function adminAddGlobalNode(address globalNode, uint256 voteNumber) public onlyAdmin {
        require(globalNodeCandidateVotes[globalNode] == 0, "already exists");
        globalNodeCandidateVotes[globalNode] = voteNumber;
    }

    // 服务端设置开放的节点城市
    function adminAddCity(bytes32 cityId, uint256 cityLevel_, uint256 startTime_) public onlyAdmin {
        if (!inBytes32List(allCityIds, cityId)) { // 添加城市
            allCityIds.push(cityId);
        }
        // 设置城市等级
        cityLevel[cityId] = cityLevel_;
        // 设置城市信息
        cityInfo[cityId].cityId = cityId;
        cityInfo[cityId].startTime = startTime_;
        cityInfo[cityId].cityLevel = cityLevel_;
    }

    // 管理员审核参选人信息状态,//0 不存在 1已审核,2未审核，3审核失败
    function adminSetNodeInfoStatus(address node_, uint256 status_) public onlyAdmin {
        require(status_ == 1 || status_ == 2, "status_ error, only 1 or 2");
        NodeInfo storage node = nodeInfo[node_];
        node.status = status_;
    }

    // 用户修改节点信息,//0 不存在 1已审核,2待审核，【3审核失败，可选】
    function userEditNodeInfo(string calldata avatar_, string calldata flag_, string calldata nodeName_) public {
        NodeInfo storage node = nodeInfo[msg.sender];
        require(node.status != 1, "can not edit any more");
        node.node = msg.sender;
        node.avatar = avatar_;
        node.flag = flag_;
        node.nodeName = nodeName_;
        node.status = 2;
    }

    // 管理员删除开放的节点城市
    function adminRemoveCity(bytes32 cityId) public onlyAdmin {
        for (uint256 i = 0; i < allCityIds.length; i++) {
            if (cityId == allCityIds[i]) {
                allCityIds[i] = allCityIds[allCityIds.length - 1];
                allCityIds.pop();
            }
        }
        cityLevel[cityId] = 0;
    }

    // 检测城市是否开放
    function isCityExits(bytes32 cityId) public view returns (bool){
        for (uint256 i = 0; i < allCityIds.length; i++) {
            if (cityId == allCityIds[i]) {
                return true;
            }
        }
        return false;
    }

    // 申请城市节点
    function applyCityCandidate(bytes32 cityId, uint256 controlRate) external {
        require(!isCityExits(cityId), "cityId error");
        uint256 userBalance = IERC20(TOX).balanceOf(msg.sender);
        require(userBalance >= applyCityFee, "Insufficient balance");
        IERC20(TOX).transferFrom(msg.sender, address(this), applyCityFee);

        if (!inAddressList(cityCandidate[cityId], msg.sender)) {// 添加城市候选人
            cityCandidate[cityId].push(msg.sender);
        }
        cityNodeControlRate[msg.sender] = controlRate;// 设置候选人的支配比例
        cityApplyFee[cityId] += applyCityFee; // 该城市总的报名费
        candidateApplyCityId[msg.sender] = cityId;//参选人城市
        emit ApplyCandidateRecord(
            msg.sender,
            applyCityFee
        );
    }

    // 全球节点申请城市节点
    function globalNodeApplyCityNode(bytes32 cityId, uint256 controlRate) external {
        require(!isCityExits(cityId), "cityId error");
        require(globalNodeCandidateVotes[msg.sender] > 0, "You are not global node");
        require(voteStatus == VoteStatus.globalVoting, "voting is not in progress");
        if (inAddressList(globalNodeCityCandidate[cityId], msg.sender)) {// 添加竞选该城市的全球节点地址
            globalNodeCityCandidate[cityId].push(msg.sender);
        }
        cityNodeControlRate[msg.sender] = controlRate; // 设置收益支配比例
        // 更新申请该城市节点票数最高者信息
        uint256 votes = globalNodeCandidateVotes[msg.sender];// 获取该全球节点当前票数
        if (votes > cityNodeGlobalNodeMaxVotes[cityId]) {// 更新票数最高者信息
            cityNodeGlobalNodeMaxVotes[cityId] = votes;
            cityNodeGlobalNode[cityId] = msg.sender;
        }
        // 增加该城市投票累加量
        cityVotes[cityId] += votes;
    }

    // 检测某地址是否存在
    function inAddressList(address[] memory addressList, address newAddress) private pure returns (bool){
        for (uint256 i = 0; i < addressList.length; i++) {
            if (newAddress == addressList[i]) {
                return true;
            }
        }
        return false;
    }

    // 检测某地址是否存在
    function inBytes32List(bytes32[] memory cityIds, bytes32 newCityId) private pure returns (bool){
        for (uint256 i = 0; i < cityIds.length; i++) {
            if (newCityId == cityIds[i]) {
                return true;
            }
        }
        return false;
    }

    // 普通用户给候选人投票
    function vote(address candidate, bytes32 cityId, uint256 toxNum) external {
        require(!isCityExits(cityId), "cityId error");
        require(voteStatus == VoteStatus.userVoting || voteStatus == VoteStatus.incumbency, "voting is not in progress");
        require(inAddressList(cityCandidate[cityId], candidate), "candidate and cityId are mismatching");
        uint256 userBalance = IERC20(TOX).balanceOf(msg.sender);
        require(userBalance >= toxNum, "Insufficient balance");

        // additionalVoting
        if (voteStatus == VoteStatus.userVoting) {
            candidateVoteEndCount[cityId][candidate] += toxNum;// 更新当前候选人票数，只统计投票期间
            userVoteVoteEndRecord[cityId][msg.sender][candidate] += toxNum; // 用户投票，指统计投票期间
            candidateVotingCount[cityId][candidate] += toxNum;// 更新当前候选人票数，不包含追加投票
        }
        IERC20(TOX).transferFrom(msg.sender, address(this), toxNum);
        userVoteRecord[cityId][msg.sender][candidate] += toxNum; // 包含追加投票
        candidateCount[cityId][candidate] += toxNum;// 更新当前候选人票数，包含追加投票

        // 判断当前候选人，在得到该投票后，是否是当前最多票数者
        if (candidateCount[cityId][candidate] > cityNodeMaxVotes[cityId]) {
            cityNodeMaxVotes[cityId] = candidateCount[cityId][candidate];// 更新最大票数
            cityNodeWinner[cityId] = candidate; // 更新票数最多的候选人
            // 更新当前城市竞选期间票数排名
            sortVotingRanking(cityId);
        } else {
            // 更新当前城市竞选结束后票数排名
            sortRanking(cityId);
        }

        // 添加该城市的投票人
        if (!inAddressList(cityVoteUsers[cityId], msg.sender)) {// 添加给该城市投票的人
            cityVoteUsers[cityId].push(msg.sender);
        }

        // 增加该城市投票累加量
        cityVotes[cityId] += toxNum;

        // 更新城市排序
        sortAllCity();

        emit VoteRecord(
            msg.sender,
            candidate,
            cityId,
            toxNum
        );
    }

    // 更新当前城市竞选结束后票数排名
    function sortRanking(bytes32 cityId) private {
        address[] storage allCandidate = cityCandidate[cityId];
        if (allCandidate.length == 0 || allCandidate.length == 1) {
            return;
        }
        address small = address(0);
        address big = address(0);
        for (uint i = 1; i < allCandidate.length; i++) {
            for (uint j; j < allCandidate.length - 1; j++) {
                if (candidateCount[cityId][allCandidate[j]] > candidateCount[cityId][allCandidate[j + 1]]) {
                    small = allCandidate[j + 1];
                    big = allCandidate[j];
                    allCandidate[j + 1] = big;
                    allCandidate[j] = small;
                }
            }
        }
    }

    // 更新当前城市竞选期间票数排名
    function sortVotingRanking(bytes32 cityId) private {
        address[] storage allCandidate = cityCandidate[cityId];
        if (allCandidate.length == 0 || allCandidate.length == 1) {
            return;
        }
        address small = address(0);
        address big = address(0);
        for (uint i = 1; i < allCandidate.length; i++) {
            for (uint j; j < allCandidate.length - 1; j++) {
                if (candidateVotingCount[cityId][allCandidate[j]] > candidateVotingCount[cityId][allCandidate[j + 1]]) {
                    small = allCandidate[j + 1];
                    big = allCandidate[j];
                    allCandidate[j + 1] = big;
                    allCandidate[j] = small;
                }
            }
        }
    }

    // 撤票，在竞选期间，投票可以撤票（全部撤回）
    function revokeVoting(address candidate, bytes32 cityId) public {
        require(!isCityExits(cityId), "cityId error");
        require(inAddressList(cityCandidate[cityId], candidate), "candidate and cityId are mismatching");
        // 检查投票是否在进行中
        require(voteStatus == VoteStatus.userVoting, "voting is not in progress");
        // 查询当前用户投票数量
        uint256 toxNum = userVoteRecord[cityId][msg.sender][candidate];
        require(toxNum > 0, "You have not vote");
        // 减少用户投票数量
        userVoteRecord[cityId][msg.sender][candidate] = 0;
        IERC20(TOX).transfer(msg.sender, toxNum);
        // 减少候选人票数
        candidateCount[cityId][candidate] -= toxNum;
    }

    // 服务端任务，更新投票状态，默认0
    function setVoteStatus(uint256 status) public {
        if (status == uint256(VoteStatus.globalVoting)) { // 1 全球节点投票中
            voteStatus = VoteStatus.globalVoting;
        } else if (status == uint256(VoteStatus.globalVoteEnd)) { // 2 全球节点投票结束
            voteStatus = VoteStatus.globalVoteEnd;
        } else if (status == uint256(VoteStatus.userVoting)) { // 3 用户投票中
            voteStatus = VoteStatus.userVoting;
        } else if (status == uint256(VoteStatus.userVoteEnd)) { // 4 用户投票结束
            voteStatus = VoteStatus.userVoteEnd;
        } else if (status == uint256(VoteStatus.incumbency)) { // 5 任期
            voteStatus = VoteStatus.incumbency;
        } else if (status == uint256(VoteStatus.endOfTerm)) { // 6 任期结束
            voteStatus = VoteStatus.endOfTerm;
        }
    }

    // 第二阶段投票结束结算，在投票结束后调用，根据城市执行，服务端先获取城市，再循环调用该方法
    function checkVote(bytes32 cityId) public onlyAdmin {
        require(voteStatus == VoteStatus.userVoteEnd, "Voting is not in an end state");

        address winner = cityNodeWinner[cityId];
        if (winner == address(0)) {// 该城市没有竞选成功者，没人参与竞选
            return;
        }
        // 获得参与该城市节点竞选人所有的报名费
        IERC20(TOX).transfer(winner, cityApplyFee[cityId]);
    }

    // 若所投票的候選人末當選，24小时后解锁為其投票的TOX，需自行领回，服务端辅助调用
    function unlockTOX() public {

    }

    /// API  ------------------------------------------------------------------------------------开始------------------------------------------------------------------------------------
    // 查看所有城市（开放可参选的城市）
    function getAllCity() public returns (CityInfo[] memory){
        CityInfo[] memory cityInfo_ = new CityInfo[](allCityIds.length);
        IntoUserLocation intoUserLocation = IntoUserLocation(userLocationAddress);
        for (uint256 i = 0; i < allCityIds.length; i++) {
            // 更新城市状态
            updateCityStatus(allCityIds[i]);

            cityInfo_[i] = CityInfo(
                allCityIds[i],// 城市ID
                intoUserLocation.getChengShiUserNumber(allCityIds[i]),// 该城市定位的人数
                cityVotes[allCityIds[i]],//当前累计投票数量
                cityInfo[allCityIds[i]].startTime,
                cityInfo[allCityIds[i]].cityLevel,
                cityInfo[allCityIds[i]].status
            );
        }
        return cityInfo_;
    }

    function updateCityStatus(bytes32 cityId) private {
        if (block.timestamp - cityInfo[cityId].startTime > secondsPerDay * 10) {
            cityInfo[cityId].status = 2;
        } else if (block.timestamp - cityInfo[cityId].startTime > 0) {
            cityInfo[cityId].status = 1;
        }
    }

    // 更新城市排序
    function sortAllCity() private {
        bytes32 small = bytes32(0);
        bytes32 big = bytes32(0);
        for (uint i = 1; i < allCityIds.length; i++) {
            for (uint j; j < allCityIds.length - 1; j++) {
                if (cityVotes[allCityIds[j]] > cityVotes[allCityIds[j + 1]]) {
                    small = allCityIds[j + 1];
                    big = allCityIds[j];
                    allCityIds[j + 1] = big;
                    allCityIds[j] = small;
                }
            }
        }
    }
//    // 查看参选人信息，可判断用户是否参选，作为候选人
//    function getCandidateByAddress() public returns (address[]memory ){
//        // 查询用户所在城市
//
//        // 查询该城市所有参选者，已审核的
//
//
//        return address[];
//    }
    // 获取某个城市参选人列表
    function getCandidateByCityId(bytes32 cityId) public view returns (CandidateInfo[] memory){
        address[] memory candidates = cityCandidate[cityId];
        uint256 candidatesLength = candidates.length;
        CandidateInfo[] memory candidateInfo = new CandidateInfo[](candidatesLength);

        // 查询该城市所有参选者，已审核的
        for (uint256 i = 0; i < candidates.length; i++) {
            candidateInfo[i] = CandidateInfo(
                cityId,
                candidates[i],
                cityNodeControlRate[candidates[i]],
                candidateCount[cityId][candidates[i]]
            );
        }

        return candidateInfo;
    }
    // 参选结果
    function getCandidateResult() public pure returns (uint256, bool){
        // 查询用户参选的城市
        // 查询当前名次

        // 查询竞选结果

        return (1, true);
    }

    // 参选详情,参选人自己查询
    function getCandidateInfo() public returns (uint256, uint256, uint256, uint256, uint256){
        // 查询用户参选的城市
        bytes32 cityId = candidateApplyCityId[msg.sender];
        // 参选城市节点报名费
//        applyCityFee;
        // 分红比例：候选人设置的支配比例
        uint256 controlRate = cityNodeControlRate[msg.sender];
        // 参与时间
        uint256 applyTime = candidateApplyTime[msg.sender];
        // 任期截止时间
        uint256 endTime = startTime += secondsPerDay * (10 + 180);
        // 获得投票数量
        uint256 voteNumber = candidateCount[cityId][msg.sender];
        return (applyCityFee, controlRate, applyTime, endTime, voteNumber);
    }
    /// API  ------------------------------------------------------------------------------------结束------------------------------------------------------------------------------------


}
