const {ethers, upgrades} = require("hardhat")
require('@openzeppelin/hardhat-upgrades')

// const ProxyAddr = "0xE8380395ad6f8a7744e4e6965228595Df5a8f60B" // 测试网
const ProxyAddr = "0xebD06631510A66968f0379A4deB896d3eE7DD6ED" // 主网

// 可升级合约
async function main() {

    const MyContractV2 = await ethers.getContractFactory("IntoCity");
    // 升级
    const contract = await upgrades.upgradeProxy(ProxyAddr, MyContractV2);

    await contract.deployed();

    console.log("proxy address is ", contract.address)
    console.log("ImplementationAddress is", await upgrades.erc1967.getImplementationAddress(contract.address));
    console.log("ProxyAdmin is", await upgrades.erc1967.getAdminAddress(contract.address));
}

// 测试网
// proxyAdminAddress="0xFafD5293E5a936da119508f69e1B5C5F5a1B03bb"
// logicAddress="0x601AC4dc0Eb861537B4e3E7c217fEc01883eb44E"

// 主网
// proxyAdminAddress="0x64A9fa6Fd332d85552E365b6877055D7C92Fa45A"
// logicAddress="0x4e97f449488CccD31c0EB995D25440AA64a8ab9e"

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });