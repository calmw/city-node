const {ethers, upgrades} = require("hardhat")
require('@openzeppelin/hardhat-upgrades')

const ProxyAddress = "0x60C541388077d524178521A7ceD95d0c7a016B72"

// 可升级合约
async function main() {
    // console.log("ImplementationAddress is", await upgrades.erc1967.getImplementationAddress(ProxyAddress));
    console.log("ProxyAdmin is", await upgrades.erc1967.getAdminAddress(ProxyAddress));
    const MyContractV2 = await ethers.getContractFactory("IntoCityPioneer");
    // 升级
    const contract = await upgrades.upgradeProxy(ProxyAddress, MyContractV2);

    await contract.deployed();

    console.log("proxy address is ", contract.address)
    console.log("ImplementationAddress is", await upgrades.erc1967.getImplementationAddress(contract.address));
    console.log("ProxyAdmin is", await upgrades.erc1967.getAdminAddress(contract.address));
}


// proxyAdminAddress="0xb6e35b62eC8077538Cb0209Ce744d05E8ea686d9"
// logicAddress="0x60C541388077d524178521A7ceD95d0c7a016B72"
// logicAddress="0x60C541388077d524178521A7ceD95d0c7a016B72" // 上一版

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });