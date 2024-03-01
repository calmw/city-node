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
    event DepositSurety(address pioneer, uint256 isUsdt, uint256 amount);

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
    mapping(address => bool) public isGlobalNode; // 是否是全球节点。节点地址=>是否是全球节点
    address[] public failedPioneers; // --------------  上线删除
    address[] public failedAreaPioneers; // 考核失败的先锋
    address[] public failedCityPioneers; // 考核失败的先锋

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

    // 管理员设置全球节点，全球节点用户不能参与先锋计划
    function adminSetIsGlobalNode(
        address globalNodeAddress_,
        bool isGlobalNode_
    ) public onlyAdmin {
        isGlobalNode[globalNodeAddress_] = isGlobalNode_;
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
        // 没有交过TOX的话，交TOX
        if (suretyDepositTOX[msg.sender] < _suretyTOX) {
            // 没交过TOX
            uint256 userBalance = tox.balanceOf(msg.sender);
            require(
                userBalance >= _suretyTOX,
                "your balance of TOX is insufficient"
            ); // 余额不足
            tox.transferFrom(msg.sender, address(this), _suretyTOX);
            suretyDepositTOX[msg.sender] += _suretyTOX;

            emit DepositSurety(msg.sender, 0, _suretyTOX);
        }

        // 如果已经交过USDT,并且不是先锋就设置成为先锋
        if (
            (suretyDepositUSDT[msg.sender] >= _suretyUSDT) &&
            !cityPioneerContract.isPioneer(msg.sender)
        ) {
            bytes32 chengShiId = cityContract.pioneerChengShi(msg.sender);
            uint cityLevel = cityContract.chengShiLevel(chengShiId);
            cityPioneerContract.initPioneer(msg.sender, chengShiId, cityLevel);
        }
    }

    function depositSuretyUSDT() public {
        (uint256 _suretyTOX, uint256 _suretyUSDT) = getSurety(msg.sender);
        // 没有交过USDT的话，交USDT
        if (suretyDepositUSDT[msg.sender] < _suretyUSDT) {
            // 没交过USDT
            uint256 userBalance = usdt.balanceOf(msg.sender);
            require(
                userBalance >= _suretyUSDT,
                "your balance of USDT is insufficient"
            ); // 余额不足
            usdt.transferFrom(msg.sender, address(this), _suretyUSDT);
            suretyDepositUSDT[msg.sender] += _suretyUSDT;

            emit DepositSurety(msg.sender, 1, _suretyUSDT);
        }

        // 如果已经交过TOX,并且不是先锋就设置成为先锋
        if (
            (suretyDepositTOX[msg.sender] >= _suretyTOX) &&
            !cityPioneerContract.isPioneer(msg.sender)
        ) {
            bytes32 chengShiId = cityContract.pioneerChengShi(msg.sender);
            uint cityLevel = cityContract.chengShiLevel(chengShiId);
            cityPioneerContract.initPioneer(msg.sender, chengShiId, cityLevel);
        }
    }

    // 重置先锋账户已经交的保证金
    function clearPioneerSurety(address pioneerAddress_) public {
        suretyDepositTOX[pioneerAddress_] = 0;
        suretyDepositUSDT[pioneerAddress_] = 0;
    }

    function addFailedPioneer(
        address pioneerAddress_,
        uint256 pioneerType_
    ) public onlyAdmin {
        bool exist;
        if (pioneerType_ == 1) {
            for (uint i = 0; i < failedAreaPioneers.length; i++) {
                if (failedAreaPioneers[i] == pioneerAddress_) {
                    exist = true;
                }
            }
            if (!exist) {
                failedAreaPioneers.push(pioneerAddress_);
            }
        } else {
            for (uint i = 0; i < failedCityPioneers.length; i++) {
                if (failedCityPioneers[i] == pioneerAddress_) {
                    exist = true;
                }
            }
            if (!exist) {
                failedCityPioneers.push(pioneerAddress_);
            }
        }
    }
}
