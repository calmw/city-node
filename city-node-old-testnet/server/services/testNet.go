package services

import "city-node-server/blockchain"

func InitTestNet() {

	blockchain.CityNodeConfig = blockchain.CityNodeConfigs{
		ChainId:               9001,
		RPC:                   "https://testnet-rpc.d2ao.com/",
		CityAddress:           "0xDfFA9bfB4D6376DB617fD2Fc56682cC7646cCb9C", // 城市合约地址
		CityPioneerAddress:    "0x11Ef2544A6f62a2cF589390325ab0A56df1eeF15", // 成市先锋合约地址
		UserLocationAddress:   "0x12CEC1A760E1e25Ce5b143cDB9D115D83603fAba", // 用户定位合约地址
		MiningAddress:         "0xD8c1d40a6FF4E53577389C8008343081949C373D", // 需要杨森设置IntoMining合约的管理员权限，权限给到cityPioneer
		SetDelegateAddress:    "0x575493F35AA4926decF877004056380538C8eB52", // 该合约为增加或减少质押量的合约，艳杰，在city给管理员权限
		PledgeStakeAddress:    "0x575493F35AA4926decF877004056380538C8eB52", // 获取质押量合约。不需要权限
		RechargeWeightAddress: "0xD8c1d40a6FF4E53577389C8008343081949C373D", // 需要杨森给合约地址，在cityPioneer给管理员权限
		GetFoundsAddress:      "0x6Fd8375249cC5FAb32099697A301901378E77a12", // 获取社交基金值的合约
		ToxAddress:            "0x3eE243ff68074502b1D9D65443fa97b99f634570", // TOX代币合约地址
		PrivateKey:            "a12dc8efdc993a8a7e67700c471f4ef85ddd7d8dceb781c9104637ec194b7ed2",
	}

	//blockchain.AdminSetSecondsPerDayCityUserLocation(300) // 86400 , 300
	//blockchain.AdminSetSecondsPerDayCityPioneer(300)      // 86400 , 300
	//blockchain.AdminSetSecondsPerDayCity(300)             // 86400 , 300
	//blockchain.AdminSetPresidencyTime(300 * 180)          // 86400*180 , 1800
}
