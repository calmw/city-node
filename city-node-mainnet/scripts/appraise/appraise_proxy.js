const {ethers, upgrades} = require("hardhat")
require('@openzeppelin/hardhat-upgrades')

const ProxyAddress = "0x164ffB92521BfbBE2732465788b5208AE2817E97"

// 可升级合约
async function main() {


    console.log("ProxyAdmin is", await upgrades.erc1967.getAdminAddress(ProxyAddress));
    const MyContractV2 = await ethers.getContractFactory("IntoAppraise");
    // 升级
    const contract = await upgrades.upgradeProxy(ProxyAddress, MyContractV2);

    await contract.deployed();

    console.log("proxy address is ", contract.address)
    console.log("ImplementationAddress is", await upgrades.erc1967.getImplementationAddress(contract.address));
    console.log("ProxyAdmin is", await upgrades.erc1967.getAdminAddress(contract.address));

    // console.log("ImplementationAddress is", await upgrades.erc1967.getImplementationAddress(ProxyAddress));
    // console.log("ProxyAdmin is", await upgrades.erc1967.getAdminAddress(ProxyAddress));
}


// proxyAdminAddress="0x3c00D5b53b529ca9992E9c925FB87e831A70F535"
// logicAddress="0xCb4c9f3E2646Dbfe40b5B4F11fEE6CF3445D349C"

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });