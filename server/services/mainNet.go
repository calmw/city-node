package services

import (
	"city-node-server/blockchain"
	"fmt"
	"os"
)

func InitMainNet() {
	key := os.Getenv("META_ACCOUNT")
	fmt.Println(key, 567567)
	blockchain.CityNodeConfig = blockchain.CityNodeConfigs{
		ChainId:               9001,
		RPC:                   "https://rpc-7.matchscan.io/",
		CityAddress:           "0xebD06631510A66968f0379A4deB896d3eE7DD6ED",
		CityPioneerAddress:    "0x60C541388077d524178521A7ceD95d0c7a016B72",
		UserLocationAddress:   "0x1B535f616B0465891Bc0bb71307A8781A8cCB8f2",
		MiningAddress:         "0x0007B44b6Ca810EBff3ED4560cD7d997b08BA104",
		SetDelegateAddress:    "0x909448fBb09880AF2812d263f7E5C77dcfC2AB53",
		PledgeStakeAddress:    "0x909448fBb09880AF2812d263f7E5C77dcfC2AB53",
		RechargeWeightAddress: "0x0007B44b6Ca810EBff3ED4560cD7d997b08BA104",
		GetFoundsAddress:      "0x20142400c97AD2c0A74d4C0400e482a973087033", // 获取社交基金值的合约
		ToxAddress:            "0x96397347Ea2beE29713Bc48749eB277D6A36A407",
		PrivateKey:            key,
	}

	//blockchain.AdminSetSecondsPerDayCityUserLocation(86400) // 86400 , 300
	//blockchain.AdminSetSecondsPerDayCityPioneer(86400)      // 86400 , 300
	//blockchain.AdminSetSecondsPerDayCity(86400)             // 86400 , 300
	//blockchain.AdminSetPresidencyTime(86400 * 180)          // 86400*180 , 1800
}
