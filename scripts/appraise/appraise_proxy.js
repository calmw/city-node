const {ethers, upgrades} = require("hardhat")
require('@openzeppelin/hardhat-upgrades')

const ProxyAddress = "0x96A45d1966B0bd08B5F3f6460f1C240527E69F72"

// 可升级合约
async function main() {


    console.log("ProxyAdmin is", await upgrades.erc1967.getAdminAddress(ProxyAddress));
    const MyContractV2 = await ethers.getContractFactory("IntoAppraise");
    const contract = await upgrades.upgradeProxy(ProxyAddress, MyContractV2);

    await contract.deployed();

    console.log("proxy address is ", contract.address)
    console.log("ImplementationAddress is", await upgrades.erc1967.getImplementationAddress(contract.address));
    console.log("ProxyAdmin is", await upgrades.erc1967.getAdminAddress(contract.address));
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });