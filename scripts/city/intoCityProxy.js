const {ethers, upgrades} = require("hardhat")
require('@openzeppelin/hardhat-upgrades')

const ProxyAddr = "0xEe4e83f11c197B8C717E4F32902132eE08957c04"

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

// proxyAdminAddress="0x3c00D5b53b529ca9992E9c925FB87e831A70F535"
// logicAddress="0x7a7b85774F6F6e40cd964DBC7B48fc053352bc3D"

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });