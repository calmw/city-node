const {ethers, upgrades} = require("hardhat")
require('@openzeppelin/hardhat-upgrades')

// const ProxyAddr = "0x9c189E3989A2c0B6C3b73eA6A77F947cB71b50B1" // 测试网
const ProxyAddr = "0x9c189E3989A2c0B6C3b73eA6A77F947cB71b50B1" // 主网

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
// proxyAdminAddress="0xb6e35b62eC8077538Cb0209Ce744d05E8ea686d9"
// logicAddress="0x6eC189557273d7AFF1A12Ab1Bd3b935Cb5eFbf11"

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });