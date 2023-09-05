// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;


import "./IntoCity.sol";
import "./RoleAccess.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

interface IPledgeStake {
    function ownerWeight(address _addr) external view returns (uint256 count);
}

contract IntoUserLocation is RoleAccess, Initializable {

    event UserLocationRecord(
        address user,
        bytes32 cityId,
        string location
    );
    // 新增事件开始

    // 新增事件结束

    // 质押合约地址 ,测试合约地址：0x575493F35AA4926decF877004056380538C8eB52 ,正式合约地址：0x909448fBb09880AF2812d263f7E5C77dcfC2AB53
    address public pledgeStakeAddress;
    // 城市合约地址
    address public intoCityAddress;

    // 城市ID集合
    bytes32[] public cityIds; // 废弃
    // 参与定位的用户数量
    uint256 public cityIdNum;
    // 用户是否设置过定位信息
    mapping(address => bool) private userHaveSetLocation;
    // 城市ID => 位置信息（加密）
    mapping(bytes32 => string) public cityInfo;
    // 城市ID => 城市用户数量
    mapping(bytes32 => uint256) public userNumberOfCity;
    // 用户地址 => 位置信息（加密）
    mapping(address => string) public userLocationInfo;
    // 用户地址 => 城市ID
    mapping(address => bytes32) public userCityId;
    // --
    // 城市ID => 该城市ID是否存在
    mapping(bytes32 => bool) public cityIdExists; // 废弃
    // 城市ID集合
    bytes32[] public cityIdsNoRepeat;
    // 城市ID => 该城市ID是否存在
    mapping(bytes32 => bool) public cityIdExist;

    function initialize() public initializer {
        _addAdmin(msg.sender);
    }

    // 管理员设置城市合约地址
    function adminSetCityAddress(address intoCityAddress_) public onlyAdmin {
        intoCityAddress = intoCityAddress_;
    }

    // 管理员设置城市合约地址
    function adminSetPledgeStakeAddress(address pledgeStakeAddress_) public onlyAdmin {
        pledgeStakeAddress = pledgeStakeAddress_;
    }

    function compareStr(string calldata _str, string memory str) private pure returns (bool) {
        return keccak256(abi.encodePacked(_str)) == keccak256(abi.encodePacked(str));
    }

    function setUserLocation(bytes32 cityId_, string calldata location_) public {
        require(!userHaveSetLocation[msg.sender], "cant not set any more");
        require(!compareStr(location_, ""), "location is empty");
        userNumberOfCity[cityId_] += 1;
        cityIds.push(cityId_); // 废弃
        cityInfo[cityId_] = location_;
        userLocationInfo[msg.sender] = location_;
        userCityId[msg.sender] = cityId_;
        userHaveSetLocation[msg.sender] = true;
        cityIdNum++;
        if (!cityIdExist[cityId_]) {
            cityIdsNoRepeat.push(cityId_);
            cityIdExist[cityId_] = true;
        }

        // 给用户所在城市增加质押量
        setUserDelegate(cityId_, msg.sender);

        emit UserLocationRecord(
            msg.sender,
            cityId_,
            location_
        );
    }

    // 获取当前用户质押量，给对应城市增加累计质押量
    function setUserDelegate(bytes32 cityId_, address user) private {
        IntoCity intoCity = IntoCity(intoCityAddress);
        IPledgeStake pledgeStake = IPledgeStake(pledgeStakeAddress);
        uint256 delegate = pledgeStake.ownerWeight(user);
        intoCity.incrCityDelegate(cityId_, delegate * 100);
    }

    // 定位过的用户数量
    function getUserNumber() public view returns (uint256) {
        return cityIds.length;
    }

    // 定位过的城市数量
    function getCityNumber() public view returns (uint256) {
        return cityIdsNoRepeat.length;
    }

    // 数据去重,执行到6547
    function noRepeatCityIds(uint256 start, uint256 end) public onlyAdmin {
        for (uint i = start; i < end; i++) {
            if (!cityIdExist[cityIds[i]]) {
                cityIdsNoRepeat.push(cityIds[i]);
                cityIdExist[cityIds[i]] = true;
            }
        }
    }

}
