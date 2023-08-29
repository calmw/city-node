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

    struct Pioneer {
       uint256 ctime; // 成为城市节点的时间
    }
//mapping(address => )  ;

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
    }
}