const {ethers, upgrades} = require("hardhat")
require('@openzeppelin/hardhat-upgrades')

// const ProxyAddr = "0x9c189E3989A2c0B6C3b73eA6A77F947cB71b50B1" // 测试网
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
// proxyAdminAddress="0xb6e35b62eC8077538Cb0209Ce744d05E8ea686d9"
// logicAddress="0x6eC189557273d7AFF1A12Ab1Bd3b935Cb5eFbf11"

// 主网
// proxyAdminAddress="0x64A9fa6Fd332d85552E365b6877055D7C92Fa45A"
// logicAddress="0x21BFE73CB4f3c136a108435855e1420d6844d983"

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });