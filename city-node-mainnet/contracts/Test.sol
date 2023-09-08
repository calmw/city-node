// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;


contract Test {

    event Delegate(
        uint256 amount
    );


    // 城市ID集合
    bytes32[] public cityIds; // 废弃
    // 城市ID => 该城市ID是否存在
    mapping(bytes32 => bool) public cityIdExists;
    // 城市ID集合
    bytes32[] public cityIdsNoRepeat;

    function getDay() public view returns (uint256){
        uint day = block.timestamp / 300;
        return uint256(day);
    }

//    function test() public {
//        uint256 today = getDay();
//        bytes32 aa = keccak256(abi.encodePacked("sdfsdf"));
//        uint256 yesterdayDelegate = cityDelegateRecord[aa][today-1];
//        emit Delegate(yesterdayDelegate);
//    }

    // 数据删除
    function delNoRepeatCityIds() public {
        for (uint256 i = 0; i < cityIds.length; i++) {
            cityIdsNoRepeat.pop();
        }
    }

    // 数据去重
    function noRepeatCityIds() public {
        for (uint256 i = 0; i < cityIds.length; i++) {
            if (!cityIdExists[cityIds[i]]) {
                cityIdsNoRepeat.push(cityIds[i]);
                cityIdExists[cityIds[i]] = true;
            }
        }
    }

    // 数据去重
    function getCityNum() public view returns(uint256){
        return cityIds.length;
    }

    // 数据去重
    function getCityNum(uint256 data) public {
       bytes32 cityId =  keccak256(abi.encodePacked(data));
        return cityIds.push(cityId);
    }

}
