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
        incumbency,
        endOfTerm
    }

    address public TOX;
    uint256 public applyCityFee = 10000 * 1e18; // 城市节点报名费
    VoteStatus  public voteStatus; // 投票状态

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
    mapping(address => uint256) public globalNodeCandidateCount;

    // 城市节点 => control rate
    mapping(address => uint256) public cityNodeControlRate;

    address[] public allCityNode; // 所有城市节点
    mapping(bytes32 => address)  public cityNode; // 竞选成功的城市节点

    uint256[50] private __gap;

    function initialize() public initializer {
        _addAdmin(msg.sender);
    }

    constructor() {
        _addAdmin(msg.sender);
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

    // 管理员设置TOX代币地址
    function adminSetTOXAddress(address TOX_) public onlyAdmin {
        TOX = TOX_;
    }

    // 管理员添加全球节点
    function adminAddGlobalNode(address globalNode, uint256 voteNumber) public onlyAdmin {
        require(globalNodeCandidateCount[globalNode] == 0, "already exists");
        globalNodeCandidateCount[globalNode] = voteNumber;
    }

    // 申请城市节点
    function applyCityCandidate(bytes32 cityId, uint256 controlRate) external {
        uint256 userBalance = IERC20(TOX).balanceOf(msg.sender);
        require(userBalance >= applyCityFee, "Insufficient balance");
        IERC20(TOX).transferFrom(msg.sender, address(this), applyCityFee);

        cityCandidate[cityId].push(msg.sender);
        cityNodeControlRate[msg.sender] = controlRate;
        cityApplyFee[cityId] += applyCityFee;
        emit ApplyCandidateRecord(
            msg.sender,
            applyCityFee
        );
    }

    // 全球节点申请城市节点
    function globalNodeApplyCityNode(address global, bytes32 cityId, uint256 controlRate) external {
        require(globalNodeCandidateCount[global] > 0, "You are not global node");
        require(voteStatus == VoteStatus.globalVoting, "voting is not in progress");

        globalNodeCityCandidate[cityId].push(global);
        cityNodeControlRate[msg.sender] = controlRate;
    }

    // 普通用户给候选人投票
    function vote(address candidate, bytes32 cityId, uint256 toxNum) external {
        require(voteStatus == VoteStatus.userVoting, "voting is not in progress");
        uint256 userBalance = IERC20(TOX).balanceOf(msg.sender);
        require(userBalance >= toxNum, "Insufficient balance");

        IERC20(TOX).transferFrom(msg.sender, address(this), toxNum);
        userVoteRecord[cityId][msg.sender][candidate] += toxNum;

        emit VoteRecord(
            msg.sender,
            candidate,
            cityId,
            toxNum
        );

        candidateCount[cityId][candidate] += toxNum;
    }

    // 撤票，在竞选期间，投票可以撤票（全部撤回）
    function revokeVoting(address candidate, bytes32 cityId) public {
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

    // 计算第一阶段全球节点竞选城市节点结果
    function countGlobalNodeVoteResult(bytes32 cityId) public onlyAdmin returns (address){
        require(voteStatus == VoteStatus.globalVoteEnd, "voting is not in the end");
        address[] storage allCandidate = globalNodeCityCandidate[cityId];
        if (allCandidate.length == 0) {
            return address(0);
        } else if (allCandidate.length == 1) {
            return allCandidate[1];
        }
        address small = address(0);
        address big = address(0);
        for (uint i = 1; i < allCandidate.length; i++) {
            for (uint j; j < allCandidate.length - 1; j++) {
                if (globalNodeCandidateCount[allCandidate[j]] > globalNodeCandidateCount[allCandidate[j + 1]]) {
                    small = allCandidate[j + 1];
                    big = allCandidate[j];
                    allCandidate[j + 1] = big;
                    allCandidate[j] = small;
                }
            }
        }
        allCityNode.push(allCandidate[allCandidate.length - 1]);
        cityNode[cityId] = allCandidate[allCandidate.length - 1];

        emit VoteSuccessRecord(
            allCandidate[allCandidate.length - 1],
            cityId,
            globalNodeCandidateCount[allCandidate[allCandidate.length - 1]]
        );

        return allCandidate[allCandidate.length - 1];
    }

    // 计算第二阶段普通用户投票候选人结果
    function countVoteResult(bytes32 cityId) public onlyAdmin returns (address){
        require(voteStatus == VoteStatus.userVoteEnd, "voting is not in the end");
        address[] storage allCandidate = cityCandidate[cityId];
        if (allCandidate.length == 0) {
            return address(0);
        } else if (allCandidate.length == 1) {
            return allCandidate[1];
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
        allCityNode.push(allCandidate[allCandidate.length - 1]);
        cityNode[cityId] = allCandidate[allCandidate.length - 1];

        emit VoteSuccessRecord(
            allCandidate[allCandidate.length - 1],
            cityId,
            candidateCount[cityId][allCandidate[allCandidate.length - 1]]
        );

        return allCandidate[allCandidate.length - 1];
    }

}
