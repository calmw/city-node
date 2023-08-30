const {ethers, upgrades} = require("hardhat")
require('@openzeppelin/hardhat-upgrades')

const ProxyAddr = "0x713e54A842dc81FC2d1E86594ca5526C6dD7C303"

// 可升级合约
async function main() {

    const MyContractV2 = await ethers.getContractFactory("IntoUserLocation.sol");
    // 升级
    const contract = await upgrades.upgradeProxy(ProxyAddr, MyContractV2);

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