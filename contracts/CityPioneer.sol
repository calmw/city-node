// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./City.sol";
import "./Events.sol";
import "./RoleAccess.sol";
import "./CityNodeVote.sol";
import "./Withdrawal.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract CityPioneer is RoleAccess, Events, Initializable {

    address public TOX;
    address public cityAddress; // 城市合约地址
    uint256 public termOfOffice; // 任期，单位秒

    enum Month {
        none,
        firstMonth,
        secondMonth,
        thirdMonth
    }

    event newDelegate(
        address pioneerAddress, // 先锋地址
        uint256 amount, // 新增质押金额
        uint256 ctime // 新增质押的时间
    );

    struct Pioneer {
        address pioneerAddress;
        uint256 ctime; // 成为城市节点的时间
        uint256 firstMonth; // 第1个月累计质押量
        uint256 secondMonth; // 第2个月累计质押量
        uint256 thirdMonth; // 第3个月累计质押量
        bool testStatus; // 最终考核状态
    }

    struct Delegate {
        address pioneer; // 先锋地址
        uint256 amount; // 质押金额
        uint256 ctime; // 质押时间（每天新增质押）
    }

    // 先锋地址 => 先锋信息， 先锋信息
    mapping(address => Pioneer) public pioneerInfo;

    mapping(address => mapping(Month => uint256[])) public pioneerDelegateInfo;

    uint256[50] private __gap;

    function initialize() public initializer {
        _addAdmin(msg.sender);
    }

    // 管理员设置TOX代币地址
    function adminSetTOXAddress(address TOX_) public onlyAdmin {
        TOX = TOX_;
    }

    // 缴纳保证金
    function depositEarnestMoney() public {
        City city = City(cityAddress);
        // 查询调用者地址是否是城市先锋
        bytes32 cityId = city.pioneerCity(msg.sender);
        require(cityId != city.cityIdEmpty(), "you are not pioneer"); // 不是城市先锋

        // 查询该城市先锋对应的城市，根据城市查询需要缴纳的保证金数量
        uint256 earnestMoney = city.earnestMoney(cityId);
        IERC20 TOXContract = IERC20(TOX);
        uint256 userBalance = TOXContract.balanceOf(msg.sender);
        require(userBalance >= earnestMoney, "your balance is insufficient"); // 余额不足
        TOXContract.transfer(address(this), earnestMoney);
        pioneerInfo[msg.sender].ctime = block.timestamp;
        pioneerInfo[msg.sender].pioneerAddress = msg.sender;
        pioneerInfo[msg.sender].testStatus = true;
    }
}