// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract UserLocation is Initializable {

    event UserLocationRecord(
        address user,
        bytes32 cityId,
        string location
    );

    // 城市ID => 位置信息（加密）
    mapping(bytes32 => string) public cityInfo;

    // 城市ID集合
    bytes32[] public cityIds;

    // 城市ID数量
    uint256 public cityIdNum;

    // 城市ID => 城市用户数量
    mapping(bytes32 => uint256) public userNumberOfCity;

    // 用户设置信息
    mapping(address => bool) private userHaveSetLocation;

    uint256[50] private __gap;

    function initialize() public initializer {}

    function compareStr(string calldata _str, string memory str) private pure returns (bool) {
        return keccak256(abi.encodePacked(_str)) == keccak256(abi.encodePacked(str));
    }

    function setUserLocation(bytes32 cityId_, string calldata location_) public {
        require(!userHaveSetLocation[msg.sender], "cant not set any more");
        require(!compareStr(location_, ""), "location is empty");
        userNumberOfCity[cityId_] += 1;
        cityIds.push(cityId_);
        cityInfo[cityId_] = location_;
        userHaveSetLocation[msg.sender] = true;
        cityIdNum++;

        emit UserLocationRecord(
            msg.sender,
            cityId_,
            location_
        );
    }
}
