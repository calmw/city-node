// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract UserLocation {

    enum LocationType{
        none,
        global,
        city,
        county
    }

    event UserLocationRecord(
        address user,
        string country,
        string city,
        string county,
        bytes32 cityId,
        LocationType locationType
    );

    event LogBytes32(
        bytes32 data
    );

    struct Location {
        string country;
        string city;
        string county;
        bytes32 cityId;
        LocationType locationType;
    }

    // user address => location info
    mapping(address => Location) public userCityInfo;
    mapping(address => Location) public userCountyInfo;

    // user address => userLocationHash
    mapping(address => bytes32) public userLocationHash;

    // cityId => userNumber
    mapping(bytes32 => uint256) public userNumberOfCity;

    function compareStr(string calldata _str, string memory str) private pure returns (bool) {
        return keccak256(abi.encodePacked(_str)) == keccak256(abi.encodePacked(str));
    }

    function getKeccak256(string calldata _str) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(_str));
    }

    function getKeccak2562() external pure returns (bytes32) {
        bytes32 bytes320;
        return keccak256(abi.encodePacked(bytes320));
    }

    function setUserLocation(bytes32 cityId_, string calldata country_, string calldata city_, string calldata county_) public {

        require(!compareStr(country_, "") && !compareStr(city_, ""), "country or city is empty");
        LocationType locationType;
        string memory county = "";
        if (!compareStr(county_, "")) {
            require(userCountyInfo[msg.sender].locationType == LocationType.none, "can not set any more");
            county = county_;
            userCountyInfo[msg.sender].county = county;
            locationType = LocationType.county;
            userCountyInfo[msg.sender].country = country_;
            userCountyInfo[msg.sender].city = city_;
            userCountyInfo[msg.sender].cityId = cityId_;
            userCountyInfo[msg.sender].locationType = locationType;
        } else {
            require(userCityInfo[msg.sender].locationType == LocationType.none, "can not set any more");
            locationType = LocationType.city;
            userCityInfo[msg.sender].country = country_;
            userCityInfo[msg.sender].city = city_;
            userCityInfo[msg.sender].cityId = cityId_;
            userCityInfo[msg.sender].locationType = locationType;
        }

        userNumberOfCity[cityId_] += 1;

        bytes32 cityId = cityId_;
        emit UserLocationRecord(
            msg.sender,
            country_,
            city_,
            county,
            cityId,
            locationType
        );
    }
}
