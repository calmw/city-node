package services

import (
	"city-node-server/pkg/blockchain"
	"city-node-server/pkg/db"
	"city-node-server/pkg/log"
	"city-node-server/pkg/models"
	"os"
)

func InitTestNet() {
	key := os.Getenv("META_ACCOUNT")
	//key := os.Getenv("HUZHI")
	var config models.Config
	err := db.Mysql.Model(models.Config{}).Where("id=1").First(&config).Error
	if err != nil {
		log.Logger.Sugar().Fatal(err)
	}
	blockchain.CityNodeConfig = blockchain.CityNodeConfigs{
		ChainId: 698,
		//RPC:                   "https://lisbon-testnet-rpc.matchtest.co",
		RPC:                    config.RpcTestnet,
		AuthAddress:            "0xD712221FE8b655C981Ac47385aEd078E4E497d3A", // SBT认证合约
		WithdrawLimitAddress:   "0x0a359A2C1214435a433C9686B4B3064345211671", // 是否在小黑屋合约
		AppraiseAddress:        "0x96A45d1966B0bd08B5F3f6460f1C240527E69F72", // 考核合约地址
		StarAddress:            "0xe8739b502df3A3dC5C0f14c0F27288c06A5ad887", // 获取用户星级合约
		CityAddress:            "0xDfFA9bfB4D6376DB617fD2Fc56682cC7646cCb9C", // 城市合约地址
		CityPioneerAddress:     "0x11Ef2544A6f62a2cF589390325ab0A56df1eeF15", // 成市先锋合约地址
		CityPioneerDataAddress: "0x7D05883B19cD14c216AD82222c48bCD4eE98914c", // 成市先锋数据合约地址
		UserLocationAddress:    "0x12CEC1A760E1e25Ce5b143cDB9D115D83603fAba", // 用户定位合约地址
		MiningAddress:          "0xD8c1d40a6FF4E53577389C8008343081949C373D", // 需要杨森设置IntoMining合约的管理员权限，权限给到cityPioneer
		SetDelegateAddress:     "0x73A8f49C231ffBF9D190C623361c332bEc59F95A", // 该合约为增加或减少质押量的合约，艳杰，在city给管理员权限
		PledgeStakeAddress:     "0x744c8E9FaABB263a8c003Fc7CfbD5136dC670c1d", // 获取质押量合约。不需要权限
		RechargeWeightAddress:  "0xD8c1d40a6FF4E53577389C8008343081949C373D", // 需要杨森给合约地址，在cityPioneer给管理员权限
		GetFoundsAddress:       "0x6Fd8375249cC5FAb32099697A301901378E77a12", // 获取社交基金值的合约
		ToxAddress:             "0x3eE243ff68074502b1D9D65443fa97b99f634570", // TOX代币合约地址
		USDTAddress:            "0x05171e5C88b43ef35D223f64E1304D3D5210701D", // USDT代币合约地址
		PrivateKey:             key,
	}

	//blockchain.AdminSetSecondsPerDayCityUserLocation(300) // 86400 , 300
	//blockchain.AdminSetSecondsPerDayCityPioneer(300)      // 86400 , 300
	//blockchain.AdminSetSecondsPerDayCity(300)             // 86400 , 300
	//blockchain.AdminSetPresidencyTime(300 * 180)          // 86400*180 , 1800
}
