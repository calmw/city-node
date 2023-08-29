// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./events.sol";
import "./RoleAccess.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract City is RoleAccess, Initializable {

    // 城市ID => 城市等级
    mapping(bytes32 => string) public cityLevel;

    // 城市ID => 城市先锋地址
    mapping(bytes32 => address) public pioneerInfo;

    uint256[50] private __gap;

    function initialize() public initializer {
        _addAdmin(msg.sender);
    }

    function AdminSetPioneer(bytes32 cityId_, address pioneer_) private pure onlyAdmin {
        pioneerInfo[cityId_] = pioneer_;
    }

    function AdminSetCityLevel(bytes32 cityId_, uint256 level_) private pure onlyAdmin {
        cityLevel[cityId_] = level_;
    }

}