const {ethers, upgrades} = require("hardhat")
require('@openzeppelin/hardhat-upgrades')

// const ProxyAddress = "0x9BCafDc5e97f25d586aBfb4ac5466fA8827b2e5E"
const ProxyAddress = "0x9FEf79ebfcEa35323A7A1eb0C52B2998D36B975d"

// 可升级合约
async function main() {

    console.log("ImplementationAddress is", await upgrades.erc1967.getImplementationAddress(ProxyAddress));
    const MyContractV2 = await ethers.getContractFactory("IntoCityPioneerData");
    // 升级
    const contract = await upgrades.upgradeProxy(ProxyAddress, MyContractV2);

    await contract.deployed();

    console.log("proxy address is ", contract.address)
    console.log("ImplementationAddress is", await upgrades.erc1967.getImplementationAddress(contract.address));
    console.log("ProxyAdmin is", await upgrades.erc1967.getAdminAddress(contract.address));

    // console.log("ImplementationAddress is", await upgrades.erc1967.getImplementationAddress(ProxyAddress));
    // console.log("ProxyAdmin is", await upgrades.erc1967.getAdminAddress(ProxyAddress));
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });