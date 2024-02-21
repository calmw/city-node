// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./RoleAccess.sol";
import "./interface/ICity.sol";
import "./interface/IERC20.sol";
import "./interface/ICityPioneer.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

interface City {} // 废弃
interface CityPioneer {} // 废弃

contract IntoCityPioneerData is RoleAccess, Initializable {
    // 用户地址=>用户需要扣除的奖励
    mapping(address => uint256) public subReward; // 废弃
    // 用户地址=>第几期用户
    mapping(address => uint256) public userPeriod;

    CityPioneer public cityPioneer; // 先锋合约,废弃
    City public city; // 城市合约，废弃

    // 城市/区县ID => 城市/区县先锋需要缴纳的保证金,USDT
    mapping(bytes32 => uint256) public suretyUSDT;
    IERC20 public tox; // TOX
    IERC20 public usdt; // USDT
    ICity public cityContract; // 城市合约
    ICityPioneer public cityPioneerContract; // 先锋合约
    mapping(address => uint256) public suretyDepositUSDT; // 先锋已经缴纳的USDT保证金数量
    mapping(address => uint256) public suretyDepositTOX; // 先锋已经缴纳的TOX保证金数量

    function initialize() public initializer {
        _addAdmin(msg.sender);
    }

    // 管理员设置用户需要扣除的奖励
    function adminSetSubReward(
        address user_,
        uint256 reward_
    ) public onlyAdmin {
        subReward[user_] = reward_;
    }

    // 减去用户需要扣除的奖励
    function adminSubReward(address user_, uint256 reward_) public onlyAdmin {
        subReward[user_] -= reward_;
    }

    // 设置先锋合约
    function adminSetCityPioneerAddress(
        address cityPioneerAddress_
    ) public onlyAdmin {
        cityPioneerContract = ICityPioneer(cityPioneerAddress_);
    }

    // 设置城市合约
    function adminSetCityAddress(address cityAddress_) public onlyAdmin {
        cityContract = ICity(cityAddress_);
    }

    // 设置TOX
    function adminSetTOX(address toxAddress_) public onlyAdmin {
        tox = IERC20(toxAddress_);
    }

    // 设置USDT
    function adminSetUSDT(address usdtAddress_) public onlyAdmin {
        usdt = IERC20(usdtAddress_);
    }

    // 管理员设置先锋计划，城市或者区县先锋所需缴纳的USDT保证金数额,chengShiId_可为城市ID或者区县ID
    function adminSetChengShiSuretyUSDT(
        bytes32 chengShiId_,
        uint256 suretyUSDT_
    ) public onlyAdmin {
        suretyUSDT[chengShiId_] = suretyUSDT_;
    }

    function getSurety(address user) public view returns (uint, uint) {
        // 查询调用者地址是否是城市先锋
        bytes32 chengShiId = cityContract.pioneerChengShi(user);
        require(chengShiId != bytes32(0), "you are not pioneer"); // 不是城市先锋
        // 查询该城市先锋对应的城市，根据城市查询需要缴纳的保证金数量
        uint256 suretyTOX = cityContract.surety(chengShiId);

        return (suretyTOX, suretyUSDT[chengShiId]);
    }

    function depositSuretyTOX() public {
        (uint256 _suretyTOX, uint256 _suretyUSDT) = getSurety(msg.sender);
        if (suretyDepositTOX[msg.sender] < _suretyTOX) {
            // 没交过TOX
            uint256 userBalance = tox.balanceOf(msg.sender);
            require(
                userBalance >= _suretyTOX,
                "your balance of TOX is insufficient"
            ); // 余额不足
            tox.transferFrom(msg.sender, address(this), _suretyTOX);
        }
        // 如果交过USDT,设置成为先锋
        if (suretyDepositUSDT[msg.sender] >= _suretyUSDT) {
            bytes32 chengShiId = cityContract.pioneerChengShi(msg.sender);
            uint cityLevel = cityContract.chengShiLevel(chengShiId);
            cityPioneerContract.initPioneer(msg.sender, chengShiId, cityLevel);
        }
    }

    function depositSuretyUSDT() public {
        (uint256 _suretyTOX, uint256 _suretyUSDT) = getSurety(msg.sender);
        if (suretyDepositUSDT[msg.sender] < _suretyUSDT) {
            // 没交过USDT
            uint256 userBalance = usdt.balanceOf(msg.sender);
            require(
                userBalance >= _suretyUSDT,
                "your balance of USDT is insufficient"
            ); // 余额不足
            tox.transferFrom(msg.sender, address(this), _suretyUSDT);
        }
        // 如果交过USDT,设置成为先锋
        if (suretyDepositTOX[msg.sender] >= _suretyTOX) {
            bytes32 chengShiId = cityContract.pioneerChengShi(msg.sender);
            uint cityLevel = cityContract.chengShiLevel(chengShiId);
            cityPioneerContract.initPioneer(msg.sender, chengShiId, cityLevel);
        }
    }
}
