// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;


interface IPledgeStake {
    function ownerWeight(address _addr) external view returns (uint256 count);
}

contract Test {

    event Delegate(
        uint256 amount
    );

    // 质押合约地址 ,测试合约地址：0x575493F35AA4926decF877004056380538C8eB52
    //正式合约地址：0x909448fBb09880AF2812d263f7E5C77dcfC2AB53
    address public pledgeStakeAddress;
    address public userAddress;

    // 管理员设置城市合约地址
    function adminSetPledgeStakeAddress(address pledgeStakeAddress_) public {
        pledgeStakeAddress = pledgeStakeAddress_;
    }

    function adminSetUserAddress(address userAddress_) public {
        userAddress = userAddress_;
    }

    // 设置用户当前质押量
    function setUserDelegate() public {
        IPledgeStake pledgeStake = IPledgeStake(pledgeStakeAddress);
        uint256 delegate = pledgeStake.ownerWeight(msg.sender);

        emit Delegate(delegate);
    }

}
