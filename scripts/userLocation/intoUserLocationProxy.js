const {ethers, upgrades} = require("hardhat")
require('@openzeppelin/hardhat-upgrades')

const ProxyAddress = "0x556814d8C9B85B32181F2cCcB9E02b7FE2E69277"

// 可升级合约
async function main() {

    const MyContractV2 = await ethers.getContractFactory("IntoUserLocation");
    // 升级
    const contract = await upgrades.upgradeProxy(ProxyAddress, MyContractV2);

    await contract.deployed();

    console.log("proxy address is ", contract.address)
    console.log("ImplementationAddress is", await upgrades.erc1967.getImplementationAddress(contract.address));
    console.log("ProxyAdmin is", await upgrades.erc1967.getAdminAddress(contract.address));
}

// 测试网
// proxyAdminAddress="0xc9DBD06547FAE83C309A9cC80EEcc207CD71d200"
// logicAddress="0xA025C2Dd6FDF956A11941aB3D3C2Fc296F66D5Cb"

// 主网
// proxyAdminAddress="0x64A9fa6Fd332d85552E365b6877055D7C92Fa45A"
// logicAddress="0x073Eab9354E9A9Cb9C02fC0a7dD36ABA5Dc732B6"

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });