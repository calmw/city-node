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
            url: "https://rpc-sen.matchscan.io",
            accounts: [account.key]
        },
        match_test: {
            url: "https://testnet-rpc.d2ao.com/",
            accounts: [account.key]
        }
    }

};
