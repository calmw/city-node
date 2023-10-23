package services

import (
	"city-node-server/blockchain"
	"os"
)

func InitMainNet() {
	//key := os.Getenv("META_ACCOUNT")
	key := os.Getenv("HUZHI")
	blockchain.CityNodeConfig = blockchain.CityNodeConfigs{
		ChainId:               9001,
		RPC:                   "https://rpc-7.matchscan.io/",
		CityAddress:           "0xebD06631510A66968f0379A4deB896d3eE7DD6ED",
		CityPioneerAddress:    "0x60C541388077d524178521A7ceD95d0c7a016B72",
		UserLocationAddress:   "0x1B535f616B0465891Bc0bb71307A8781A8cCB8f2",
		MiningAddress:         "0x0007B44b6Ca810EBff3ED4560cD7d997b08BA104", // 需要杨森设置IntoMining合约的管理员权限，权限给到cityPioneer
		SetDelegateAddress:    "0x1Ec4585a74aF7746C79263c2bc6C286B631554f1", // 该合约为增加或减少质押量的合约，艳杰，在city给管理员权限
		PledgeStakeAddress:    "0xFc207eC02eE9A675865db77869df049130C32a4A", // 获取质押量合约。不需要权限
		RechargeWeightAddress: "0x0007B44b6Ca810EBff3ED4560cD7d997b08BA104", // 需要杨森给合约地址，在cityPioneer给管理员权限
		GetFoundsAddress:      "0x20142400c97AD2c0A74d4C0400e482a973087033", // 获取社交基金值的合约
		ToxAddress:            "0x96397347Ea2beE29713Bc48749eB277D6A36A407",
		PrivateKey:            key,
	}

	//blockchain.AdminSetSecondsPerDayCityUserLocation(86400) // 86400 , 300
	//blockchain.AdminSetSecondsPerDayCityPioneer(86400)      // 86400 , 300
	//blockchain.AdminSetSecondsPerDayCity(86400)             // 86400 , 300
	//blockchain.AdminSetPresidencyTime(86400 * 180)          // 86400*180 , 1800
}
