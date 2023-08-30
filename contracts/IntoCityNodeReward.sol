// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./RoleAccess.sol";
import "./IntoCityNodeVote.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract IntoCityNodeReward is RoleAccess, Initializable {

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

    address public voteAddress;
    mapping(bytes32 => uint256)public  rewardsForSocialFunds; // 城市每日社交基金
    mapping(bytes32 => uint256)public  latestDelegate; // 城市每日新增质押

    function initialize() public initializer {
        _addAdmin(msg.sender);
    }

    constructor()  {
        _addAdmin(msg.sender);
    }

    // 管理员设置投票合约地址
    function adminSetVoteAddress(address voteAddress_) public onlyAdmin {
        voteAddress = voteAddress_;
    }

    // 管理员设置昨日最新城市社交基金值
    // ？ 此时需要把钱打到这个合约？？
    function adminSetRewardsForSocialFunds(bytes32 cityId, uint256 founds_) public onlyAdmin {
        rewardsForSocialFunds[cityId] = founds_;
        emit UpdateSocialFundsRecord(
            cityId,
            founds_,
            block.timestamp
        );
    }

    // 管理员设置昨日城市新增质押量
    function adminSetLatestDelegate_(bytes32 cityId, uint256 latestDelegate_) public onlyAdmin {
        latestDelegate[cityId] = latestDelegate_;
        emit UpdateLatestDelegateRecord(
            cityId,
            latestDelegate_,
            block.timestamp
        );
    }

    // 竞选成功，每天奖励
    function rewardNode(bytes32 cityId) public onlyAdmin {
        IntoCityNodeVote cityNodeVote = IntoCityNodeVote(voteAddress);
        address cityNode = cityNodeVote.cityNode(cityId); // 当选者
        // 奖励收益支配比例
        uint256 rate = cityNodeVote.cityNodeControlRate(cityNode);
        // 该节点总的投票数
        cityNodeVote.candidateCount(cityId, cityNode);
        // 总收益=该城市社交基金奖励+该城市昨日新增质押
        uint256 totalRewards = rewardsForSocialFunds[cityId] + latestDelegate[cityId];
        // 当选者今日收益发放
        uint256 userElectedRewards = (100 - rate) * totalRewards / 100;// 当选者今日收益
        address TOX = cityNodeVote.TOX();
        IERC20(TOX).transfer(cityNode, userElectedRewards);
        // 投票者今日收益发放
        uint256 voterRewards = rate * totalRewards / 100;// 投票者今日总收益


         voterRewards = rate * totalRewards / 100;// 该城市总投票数
//        address[] memory allVoter = cityNodeVote.cityCandidate(cityId);
//        for (uint256 i = 0; i < allVoter.length; i++) {
//            uint256 memory voterNumber = cityNodeVote.userVoteRecord(cityId, msg.sender, candidate); // 当前用户给该城市节点（竞选成功者）投票的数量
//        }

    }

    // 竞选成功，一次性奖励,竞选结束调用
    function campaignSuccessRewards(bytes32 cityId) public onlyAdmin {
        IntoCityNodeVote cityNodeVote = IntoCityNodeVote(voteAddress);
        address cityNode = cityNodeVote.cityNode(cityId); // 当选者
        uint256 cityApplyFee = cityNodeVote.cityApplyFee(cityId); // 该城市竞选所有的报名费
        // 获得参与该城市节点竞选人所有的报名费
        address TOX = cityNodeVote.TOX();
        IERC20(TOX).transfer(cityNode, cityApplyFee);

    }

    // 锁仓期截止，释放TOX
    function unlock() public {

    }

}
