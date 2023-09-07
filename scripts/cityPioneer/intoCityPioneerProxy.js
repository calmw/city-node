const {ethers, upgrades} = require("hardhat")
require('@openzeppelin/hardhat-upgrades')

const ProxyAddress = "0xB1AE37BBD0ABF855727741b4c2Aa9104Db9931B7" // 测试网
// const ProxyAddress = "0xDf8CFD04CC66b8a775bF77303C4F57F0DfB91b0a" // 主网

// 可升级合约
async function main() {

    const MyContractV2 = await ethers.getContractFactory("IntoCityPioneer");
    // 升级
    const contract = await upgrades.upgradeProxy(ProxyAddress, MyContractV2);

    await contract.deployed();

    console.log("proxy address is ", contract.address)
    console.log("ImplementationAddress is", await upgrades.erc1967.getImplementationAddress(contract.address));
    console.log("ProxyAdmin is", await upgrades.erc1967.getAdminAddress(contract.address));
}

// 测试网
// proxyAdminAddress="0xb6e35b62eC8077538Cb0209Ce744d05E8ea686d9"
// logicAddress="0x600955EbFfE0e422891d5BA03284f1F58e8e6c4C"

// 主网
// proxyAdminAddress="0xb6e35b62eC8077538Cb0209Ce744d05E8ea686d9"
// logicAddress="0x600955EbFfE0e422891d5BA03284f1F58e8e6c4C"

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });