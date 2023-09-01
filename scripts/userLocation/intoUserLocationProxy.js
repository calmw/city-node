const {ethers, upgrades} = require("hardhat")
require('@openzeppelin/hardhat-upgrades')

// const ProxyAddress = "0x694D23cDA706eC8816C9598A523005cb4410Bf21" // 测试网
const ProxyAddress = "0x694D23cDA706eC8816C9598A523005cb4410Bf21"

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
// proxyAdminAddress="0xb6e35b62eC8077538Cb0209Ce744d05E8ea686d9"
// logicAddress="0x17435E20dDd57A6Ff5374B6e85CCAEfcE2422f27"

// 主网
// proxyAdminAddress="0xb6e35b62eC8077538Cb0209Ce744d05E8ea686d9"
// logicAddress="0x17435E20dDd57A6Ff5374B6e85CCAEfcE2422f27"

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });