const {ethers, upgrades} = require("hardhat")
require('@openzeppelin/hardhat-upgrades')

const ProxyAddress = "0x1B535f616B0465891Bc0bb71307A8781A8cCB8f2"

// 可升级合约
async function main() {

    const MyContractV2 = await ethers.getContractFactory("IntoUserLocation");
    // 升级
    const contract = await upgrades.upgradeProxy(ProxyAddress, MyContractV2);

    await contract.deployed();

    console.log("proxy address is ", contract.address)
    console.log("ImplementationAddress is", await upgrades.erc1967.getImplementationAddress(contract.address));
    console.log("ProxyAdmin is", await upgrades.erc1967.getAdminAddress(contract.address));
    // console.log("ImplementationAddress is", await upgrades.erc1967.getImplementationAddress(ProxyAddress));
    // console.log("ProxyAdmin is", await upgrades.erc1967.getAdminAddress(ProxyAddress));
}

// proxyAdminAddress="0x64A9fa6Fd332d85552E365b6877055D7C92Fa45A"
// logicAddress="0xF7fCC8100CBe0BA7889f8623Baab4cd3868d3c53"
// logicAddress="0x073Eab9354E9A9Cb9C02fC0a7dD36ABA5Dc732B6" // 上一版

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });