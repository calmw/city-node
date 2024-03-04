package services

import (
	"city-node-server/pkg/blockchain"
	"city-node-server/pkg/db"
	"city-node-server/pkg/log"
	"city-node-server/pkg/models"
)

func InitMainNet() {
	//key := os.Getenv("META_ACCOUNT")
	//key := os.Getenv("HUZHI")
	key := "a12dc8efdc993a8a7e67700c471f4ef85ddd7d8dceb781c9104637ec194b7ed2"
	var config models.Config
	err := db.Mysql.Model(models.Config{}).Where("id=1").First(&config).Error
	if err != nil {
		log.Logger.Sugar().Fatal(err)
	}
	blockchain.CityNodeConfig = blockchain.CityNodeConfigs{
		ChainId:                9001,
		RPC:                    config.RpcMainnet,
		AuthAddress:            "0x424CaA5beA9bDfeF9F49796D7C7005632fAbE2E8", // SBT认证合约
		WithdrawLimitAddress:   "0xa3FF6A43b990A6AF220d1B376E9e97E2621bcaD3", // 是否在小黑屋合约
		AppraiseAddress:        "0x164ffB92521BfbBE2732465788b5208AE2817E97",
		StarAddress:            "0xBc4E6563af151d199726DF794903b0dcb4D1CAf3", //
		CityAddress:            "0xebD06631510A66968f0379A4deB896d3eE7DD6ED",
		CityPioneerAddress:     "0x60C541388077d524178521A7ceD95d0c7a016B72",
		CityPioneerDataAddress: "0x2548Ec31B6EBe2D0dD6c316658193789601830B1", // 成市先锋数据合约地址
		UserLocationAddress:    "0x1B535f616B0465891Bc0bb71307A8781A8cCB8f2",
		MiningAddress:          "0x0007B44b6Ca810EBff3ED4560cD7d997b08BA104", // 需要杨森设置IntoMining合约的管理员权限，权限给到cityPioneer
		SetDelegateAddress:     "0x1Ec4585a74aF7746C79263c2bc6C286B631554f1", // 该合约为增加或减少质押量的合约，艳杰，在city给管理员权限
		PledgeStakeAddress:     "0xFc207eC02eE9A675865db77869df049130C32a4A", // 获取质押量合约。不需要权限
		RechargeWeightAddress:  "0x0007B44b6Ca810EBff3ED4560cD7d997b08BA104", // 需要杨森给合约地址，在cityPioneer给管理员权限
		GetFoundsAddress:       "0x20142400c97AD2c0A74d4C0400e482a973087033", // 获取社交基金值的合约
		ToxAddress:             "0x96397347Ea2beE29713Bc48749eB277D6A36A407",
		USDTAddress:            "0x67Dc36C19835Fa65Bf4B100FAC9a80A9E9280fB9", // USDT代币合约地址
		PrivateKey:             key,
	}
}
