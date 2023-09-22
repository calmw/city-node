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

    // 区县ID集合
    bytes32[] public cityIds; // 废弃
    // 参与定位的用户数量
    uint256 public cityIdNum;
    // 用户是否设置过定位信息
    mapping(address => bool) private userHaveSetLocation;
    // 区县ID => 位置信息（加密）
    mapping(bytes32 => string) public cityInfo;
    // 区县ID => 城市用户数量
    mapping(bytes32 => uint256) public userNumberOfCity;
    // 用户地址 => 位置信息（加密）
    mapping(address => string) public userLocationInfo;
    // 用户地址 => 区县ID
    mapping(address => bytes32) public userCityId;

    // 新增变量 -------------------------------------
    // 区县ID => 该区县ID是否存在
    mapping(bytes32 => bool) public cityIdExists; // 废弃
    // 区县ID集合
    bytes32[] public cityIdsNoRepeat;
    // 区县ID => 该区县ID是否存在
    mapping(bytes32 => bool) public cityIdExist;

    // 区县ID=>城市ID
    mapping(bytes32 => bytes32) public cityIdChengShiID;
    // 城市ID=>区县ID[]
    mapping(bytes32 => bytes32[]) public chengShiIDCityIdSet;
    // 防止重复添加，区县ID=>(城市ID=>是否存在)
    mapping(bytes32 => mapping(bytes32 => bool)) public cityIdToChengShiIDExits;
    // 考核天，正式86400秒，测试300秒
    uint public secondsPerDay;
    // 关闭定位，定位黑名单,后期花钱定位（花钱设置）
    mapping(address => bool) public userDisable;
    uint256 public enableToxNumber; // 用户开启定位所需TOX数量
    address public TOXAddress; // TOX代币合约地址
    address[] public disableLbsUsers;

    function initialize() public initializer {
        _addAdmin(msg.sender);
    }

    // 管理员设置TOX代币地址
    function adminSetTOXAddress(address TOXAddress_) public onlyAdmin {
        TOXAddress = TOXAddress_;
    }

    // 管理员设置城市合约地址
    function adminSetCityAddress(address intoCityAddress_) public onlyAdmin {
        intoCityAddress = intoCityAddress_;
    }

    // 管理员设置城市合约地址
    function adminSetPledgeStakeAddress(address pledgeStakeAddress_) public onlyAdmin {
        pledgeStakeAddress = pledgeStakeAddress_;
    }

    // 管理员设置每天秒数，用于测试
    function adminSetSecondsPerDay(uint56 secondsPerDay_) public onlyAdmin {
        secondsPerDay = secondsPerDay_;
    }

    function compareStr(string calldata _str, string memory str) private pure returns (bool) {
        return keccak256(abi.encodePacked(_str)) == keccak256(abi.encodePacked(str));
    }

    function getDay() public view returns (uint256){
        uint day = block.timestamp / secondsPerDay;
        return uint256(day);
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

    // 本次上线放开
    function setUserLocationV2(bytes32 countyId_, bytes32 chengShiId_, string calldata location_) public {
        require(!userDisable[msg.sender], "LBS disable");
        require(!userHaveSetLocation[msg.sender], "cant not set any more");
        require(!compareStr(location_, ""), "location is empty");
        userNumberOfCity[countyId_] += 1;
        cityIds.push(countyId_); // 废弃
        cityInfo[countyId_] = location_;
        userLocationInfo[msg.sender] = location_;
        userCityId[msg.sender] = countyId_;
        userHaveSetLocation[msg.sender] = true;
        cityIdNum++;
        if (!cityIdExist[countyId_]) {
            cityIdsNoRepeat.push(countyId_);
            cityIdExist[countyId_] = true;
        }

        // 设置城市ID，本次上线放开
        SetCityChengShi(countyId_, chengShiId_, location_);

        // 给用户所在城市增加质押量
        setUserDelegate(chengShiId_, msg.sender);

        emit UserLocationRecord(
            msg.sender,
            countyId_,
            location_
        );
    }

    // 获取当前用户质押量，给对应城市增加累计质押量
    function setUserDelegate(bytes32 cityId_, address user) private {
        IntoCity intoCity = IntoCity(intoCityAddress);
        IPledgeStake pledgeStake = IPledgeStake(pledgeStakeAddress);
        uint256 delegate = pledgeStake.ownerWeight(user);
        uint256 today = getDay();
        intoCity.incrCountyOrChengShiDelegate(user, cityId_, delegate * 100, today);
    }

    // 定位过的用户数量
    function getUserNumber() public view returns (uint256) {
        return cityIds.length;
    }

    // 获取城市ID对应的所有区县ID
    function getCountyIdsByChengShiId(bytes32 chengShiId) public view returns (bytes32[] memory) {
        return chengShiIDCityIdSet[chengShiId];
    }

    // 定位过的城市数量
    function getCityNumber() public view returns (uint256) {
        return cityIdsNoRepeat.length;
    }

    // 设置城市与区县的映射，上线后过段时间关闭
    function SetCityMapping(bytes32 countyId, bytes32 chengShiId, string calldata location_) public returns (bool){
        cityIdChengShiID[countyId] = chengShiId;
        cityInfo[chengShiId] = location_;
        if (!countyIdInChengShiId(countyId, chengShiId)) {
            chengShiIDCityIdSet[chengShiId].push(countyId);
        }
        return true;
    }

    function countyIdInChengShiId(bytes32 countyId, bytes32 chengShiId) public view returns (bool){
        bytes32[] memory countyIds = chengShiIDCityIdSet[chengShiId];
        for (uint256 i = 0; i < countyIds.length; i++) {
            if (countyId == countyIds[i]) {
                return true;
            }
        }
        return false;
    }

    // 测试用
    function SetCityMappingDel(bytes32 chengShiId, uint256 index) public onlyAdmin {
        chengShiIDCityIdSet[chengShiId][index] = chengShiIDCityIdSet[chengShiId][chengShiIDCityIdSet[chengShiId].length - 1];
        chengShiIDCityIdSet[chengShiId].pop();
    }

    // 设置城市与区县的映射，上线后过段时间关闭
    function SetCityMappingBatch(bytes32[] calldata countyIds, bytes32[] calldata chengShiIds, string[] calldata locations_) public {
        for (uint i; i < countyIds.length; i++) {
            SetCityMapping(countyIds[i], chengShiIds[i], locations_[i]);
        }

    }

    // 设置城市与区县的映射，本次上线放开
    function SetCityChengShi(bytes32 countyId, bytes32 chengShiId, string calldata location_) private {
        cityIdChengShiID[countyId] = chengShiId;
        cityInfo[chengShiId] = location_;
        if (!countyIdInChengShiId(countyId, chengShiId)) {
            chengShiIDCityIdSet[chengShiId].push(countyId);
        }

    }

    // 根据用户获取区县ID
    function getCountyId(address user) public view returns (bytes32){
        return userCityId[user];
    }

    // 根据用户获取城市ID
    function getChengShiIdByAddress(address user) public view returns (bytes32){
        return cityIdChengShiID[userCityId[user]];
    }

    // 根据区县ID获取城市ID
    function getChengShiIdByCountyId(bytes32 countyId) public view returns (bytes32){
        return cityIdChengShiID[countyId];
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

    // 用户关闭定位
    function disableLbs() public {
        userDisable[msg.sender] = true;
        if (!isDisableLbsUsersExits(msg.sender)) {
            disableLbsUsers.push(msg.sender);
        }
    }

    // 判断用户是否开启定位
    function isDisableLbsUsersExits(address user) public view returns (bool){
        for (uint256 i = 0; i < disableLbsUsers.length; i++) {
            if (user == disableLbsUsers[i]) {
                return true;
            }
        }
        return false;
    }

    // 获取不开启定位的数量
    function getDisableLbsUsersNumber() public view returns (uint256){
        return disableLbsUsers.length;
    }

    // 获取城市的总人数
    function getChengShiUserNumber(bytes32 chengShiId) public view returns (uint256){
        bytes32[] memory cityIds_ = getCountyIdsByChengShiId(chengShiId);
        uint256 userNumber;
        for (uint256 i = 0; i < cityIds_.length; i++) {
            userNumber += userNumberOfCity[cityIds_[i]];
        }
        return userNumber;
    }

    // 用户开启定位,暂不开放
//    function enableLbs() public {
//        // 交钱
//        IERC20 TOXContract = IERC20(TOXAddress);
//        uint256 userBalance = TOXContract.balanceOf(msg.sender);
//
//        require(userBalance >= 1e25, "your balance is insufficient"); // 余额不足
//        TOXContract.transferFrom(msg.sender, address(this), surety);
//        // 更新状态
//        userDisable[msg.sender] = false;
//    }
    // 开启用户定位,测试
    function enableLbs(address user) public onlyAdmin {
        // 更新状态
        userDisable[user] = false;
    }

//    function setUserLocationTest(bytes32 cityId_, address user, string calldata location_) public onlyAdmin {
//        userNumberOfCity[cityId_] += 1;
//        cityIds.push(cityId_); // 废弃
//        cityInfo[cityId_] = location_;
//        userLocationInfo[user] = location_;
//        userCityId[user] = cityId_;
//        userHaveSetLocation[user] = true;
//        cityIdNum++;
//        if (!cityIdExist[cityId_]) {
//            cityIdsNoRepeat.push(cityId_);
//            cityIdExist[cityId_] = true;
//        }
//
//        // 给用户所在城市增加质押量
//        setUserDelegate(cityId_, msg.sender);
//
//        emit UserLocationRecord(
//            msg.sender,
//            cityId_,
//            location_
//        );
//    }

}
