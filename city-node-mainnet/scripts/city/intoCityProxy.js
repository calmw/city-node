const {ethers, upgrades} = require("hardhat")
require('@openzeppelin/hardhat-upgrades')

const ProxyAddr = "0xebD06631510A66968f0379A4deB896d3eE7DD6ED" // 主网

// 可升级合约
async function main() {
    console.log("ImplementationAddress is", await upgrades.erc1967.getImplementationAddress(ProxyAddr));
    console.log("ProxyAdmin is", await upgrades.erc1967.getAdminAddress(ProxyAddr));
    const MyContractV2 = await ethers.getContractFactory("IntoCity");
    // 升级
    // const overrides = {
    //     nonce: 505
    // };
    const contract = await upgrades.upgradeProxy(ProxyAddr, MyContractV2);

    await contract.deployed();

    console.log("proxy address is ", contract.address)
    console.log("ImplementationAddress is", await upgrades.erc1967.getImplementationAddress(contract.address));
    console.log("ProxyAdmin is", await upgrades.erc1967.getAdminAddress(contract.address));
}

// proxyAdminAddress="0x64A9fa6Fd332d85552E365b6877055D7C92Fa45A"
// logicAddress="0x83d2BBA1eeF850574e1dbe20336545513FD1C2F8"
// logicAddress="0xcd507929f9f8B79f02192837eaD33B30c89752Ce" // 上一版

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });