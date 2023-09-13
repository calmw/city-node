// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./RoleAccess.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
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
        none,
        globalVoting,
        globalVoteEnd,
        userVoting,
        userVoteEnd,
        incumbency,// 任期
        endOfTerm// 任期结束
    }

    address public TOX;
    uint256 public startTime; // 投票开始时间（按秒算的时间戳）
    uint256 public applyCityFee = 10000 * 1e18; // 城市节点报名费
    VoteStatus  public voteStatus; // 投票状态

    address[] public allCityNode; // 所有城市节点
    bytes32[] public allCity; // 所有城市
    // 城市ID=>cityLevel
    mapping(bytes32 => uint256) public cityLevel;
    // 城市等级 => (类别 => 数值) ， 类别（1用户阀值，2质押量阀值）
    mapping(uint256 => mapping(uint256 => uint256))public  activityCriteria; // 城市激活标准
    // 城市ID => 该城市所有申请竞选的费用
    mapping(bytes32 => uint256) public cityApplyFee;
    // 城市ID => (投票人地址 => (候选人地址 => 投票数量))
    mapping(bytes32 => mapping(address => mapping(address => uint256))) public userVoteRecord;
    // 城市ID => (候选人地址 => 候选人当前票数)
    mapping(bytes32 => mapping(address => uint256)) public candidateCount;
    // 城市ID => []候选人地址
    mapping(bytes32 => address[]) public cityCandidate;
    // 城市ID => 数组，参与该城市竞选的全球节点地址
    mapping(bytes32 => address[]) public globalNodeCityCandidate;
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

    function initialize() public initializer {
        _addAdmin(msg.sender);
        startTime = block.timestamp;
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
    function adminAddCity(bytes32 cityId, uint256 cityLevel_) public onlyAdmin {
        if (!isBytes32Exits(allCity, cityId)) { // 添加城市
            allCity.push(cityId);
        }
        // 设置城市等级
        cityLevel[cityId] = cityLevel_;
    }

    // 管理员删除开放的节点城市
    function adminRemoveCity(bytes32 cityId) public onlyAdmin {
        for (uint256 i = 0; i < allCity.length; i++) {
            if (cityId == allCity[i]) {
                allCity[i] = allCity[allCity.length - 1];
                allCity.pop();
            }
        }
        cityLevel[cityId] = 0;
    }

    // 检测城市是否开放
    function isCityExits(bytes32 cityId) public view returns (bool){
        for (uint256 i = 0; i < allCity.length; i++) {
            if (cityId == allCity[i]) {
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

        if (!isAddressExits(cityCandidate[cityId], msg.sender)) {// 添加城市候选人
            cityCandidate[cityId].push(msg.sender);
        }
        cityNodeControlRate[msg.sender] = controlRate;// 设置候选人的支配比例
        cityApplyFee[cityId] += applyCityFee; // 该城市总的报名费
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
        if (isAddressExits(globalNodeCityCandidate[cityId], msg.sender)) {// 添加竞选该城市的全球节点地址
            globalNodeCityCandidate[cityId].push(msg.sender);
        }
        cityNodeControlRate[msg.sender] = controlRate; // 设置收益支配比例
        // 更新申请该城市节点票数最高者信息
        uint256 votes = globalNodeCandidateVotes[msg.sender];// 获取该节点当前票数
        if (votes > cityNodeGlobalNodeMaxVotes[cityId]) {// 更新票数最高者信息
            cityNodeGlobalNodeMaxVotes[cityId] = votes;
            cityNodeGlobalNode[cityId] = msg.sender;
        }
    }

    // 检测某地址是否存在
    function isAddressExits(address[] memory globals, address newGlobal) private pure returns (bool){
        for (uint256 i = 0; i < globals.length; i++) {
            if (newGlobal == globals[i]) {
                return true;
            }
        }
        return false;
    }

    // 检测某地址是否存在
    function isBytes32Exits(bytes32[] memory cityIds, bytes32 newCityId) private pure returns (bool){
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
        require(voteStatus == VoteStatus.userVoting, "voting is not in progress");
        require(isAddressExits(cityCandidate[cityId], candidate), "candidate and cityId are mismatching");
        uint256 userBalance = IERC20(TOX).balanceOf(msg.sender);
        require(userBalance >= toxNum, "Insufficient balance");

        IERC20(TOX).transferFrom(msg.sender, address(this), toxNum);
        userVoteRecord[cityId][msg.sender][candidate] += toxNum;
        candidateCount[cityId][candidate] += toxNum;// 更新当前候选人票数

        // 判断当前候选人，在得到该投票后，是否是当前最多票数者
        if (candidateCount[cityId][candidate] > cityNodeMaxVotes[cityId]) {
            cityNodeMaxVotes[cityId] = candidateCount[cityId][candidate];// 更新最大票数
            cityNodeWinner[cityId] = candidate; // 更新票数最多的候选人
        }

        emit VoteRecord(
            msg.sender,
            candidate,
            cityId,
            toxNum
        );
    }

    // 撤票，在竞选期间，投票可以撤票（全部撤回）
    function revokeVoting(address candidate, bytes32 cityId) public {
        require(!isCityExits(cityId), "cityId error");
        require(isAddressExits(cityCandidate[cityId], candidate), "candidate and cityId are mismatching");
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

    // 奖励 ，考虑怎么触发


}
