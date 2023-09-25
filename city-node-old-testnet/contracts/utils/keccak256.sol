// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract Keccak256 {

    function getKeccak256(string calldata str_) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(str_));
    }

}
