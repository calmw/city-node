/** @type import('hardhat/config').HardhatUserConfig */
require('@openzeppelin/hardhat-upgrades')
require("@nomiclabs/hardhat-waffle");
const account = require("./secretkey.json")

module.exports = {
    solidity: "0.8.18",
    settings: {
        optimizer: {
            enabled: true,
            runs: 2000
        },
    },
    networks: {
        match: {
            url: "http://18.179.50.41:8545",
            accounts: [account.key]
        },
        match_test: {
            url: "https://lisbon-testnet-rpc.matchtest.co",
            accounts: [account.key]
        }
    }

};
