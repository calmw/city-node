// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./events.sol";
import "./RoleAccess.sol";
import "./CityNodeVote.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract CityPioneer is RoleAccess, Events, Initializable {

}