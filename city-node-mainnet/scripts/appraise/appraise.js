const {ethers, upgrades} = require("hardhat")

// 可升级合约
async function main() {

    const contract = await ethers.getContractFactory("IntoAppraise")
    console.log("Deploying .........")


//   这里先需要提交MetaData，获取到代理地址，将代理地址写入
    const City = await upgrades.deployProxy(contract, [], {initializer: "initialize"});
    console.log("Proxy address is ", City.address)
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });