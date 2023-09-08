package services

import (
	"city-node-server/blockchain"
	"fmt"
	"math/rand"
	"time"
)

func InitCity() {
	// 管理员设置先锋计划合约地址
	blockchain.AdminSetCityPioneerAddress()
	// 管理员设置用户定位合约地址
	blockchain.AdminSetUserLocationAddress()
	time.Sleep(time.Second * 5)
	// 管理员设置先锋计划，城市等级以及该等级城市所需缴纳的保证金数额
	//blockchain.AdminSetChengShiLevelAndSurety("0x0a23fb256dc6340f9287a2edf424eb5cb92f37f39a8f58cc6c1a9fa349d51b34", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xf7d89a46c223f9704a010c7077eb586fd7208822821abefa1a5e0fc5a420cbe3", 2, 60000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x35cbb5a35a63817975c8f118ac59cf9e48a4f8109bb03276c58729c04540ebe7", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x3a88d7ca2d9ba961030f92a27f5b42aaf63d78d2b477fe9dfd140a1f7dbea576", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x6478eab9198cc2c0f53912730434acabfaa199b7471f6b5d5fd5b1413aa672f1", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xc8b0a5ed8a761c9e4d3fb44625cf80d29f38c1115a381154b9ab236058414f9c", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xb736dbc2078ef993011021edaffaab7286696e849d205ea3b2e43323a3466d24", 3, 40000)
	blockchain.AdminSetChengShiLevelAndSurety("0x5b4d60ccba6e56a3bb2ac1b1e8eaa4f55320987e453b2c5947d4e6964bd25e39", 3, 40000)
	////
	////// 管理员设置城市先锋 ------------
	//blockchain.AdminSetPioneer("0x0a23fb256dc6340f9287a2edf424eb5cb92f37f39a8f58cc6c1a9fa349d51b34", "0x0DbE6e626B1085a2Cb4D5F77389bBb54Ec89Fa30")
	//blockchain.AdminSetPioneer("0xf7d89a46c223f9704a010c7077eb586fd7208822821abefa1a5e0fc5a420cbe3", "0x56a2066B4cD0954f0B49Df6835a670Aed946B265")
	//blockchain.AdminSetPioneer("0x35cbb5a35a63817975c8f118ac59cf9e48a4f8109bb03276c58729c04540ebe7", "0xD5f92Fd92F8c7f9391513E3019D9441aAf5b2D9E")
	//blockchain.AdminSetPioneer("0x6478eab9198cc2c0f53912730434acabfaa199b7471f6b5d5fd5b1413aa672f1", "0x2D4EF8abEF8E90DE0cea3Ee302A68e25d0d055AC")
	//blockchain.AdminSetPioneer("0x3a88d7ca2d9ba961030f92a27f5b42aaf63d78d2b477fe9dfd140a1f7dbea576", "0x0d15C5f863Ff5CdD96A4B21060eE25DDe667182b")
	//blockchain.AdminSetPioneer("0xc8b0a5ed8a761c9e4d3fb44625cf80d29f38c1115a381154b9ab236058414f9c", "0xa10edEb8b3486C7433925081dC09Ec8A8B941677")
	//blockchain.AdminSetPioneer("0xb736dbc2078ef993011021edaffaab7286696e849d205ea3b2e43323a3466d24", "0xa75a076c5529b3813f53c9bd24ab1f7da37994fc")
	blockchain.AdminSetPioneer("0x5b4d60ccba6e56a3bb2ac1b1e8eaa4f55320987e453b2c5947d4e6964bd25e39", "0xB37f8cfFE78698DBe04CE36D47b8ff1E5D140b0a")

	// 给城市先锋合约、用户定位合约、设置质押量合约添加管理员权限
	blockchain.AddCityAdmin()
}

func AdminSetDelegate() {
	randInt := rand.Int() % 2
	if randInt == 0 {
		randInt = 2
	}
	inc := randInt * 1000
	dec := randInt * 500
	fmt.Println("增加", inc)
	fmt.Println("减少", dec)
	accounts := []string{
		"0x0b6d66f125b165fd41bcaa7b4e7aa721cda47f71",
		"0xbcf477b0f8add3249acfbc76b0e1e3134ec9b6c1",
		"0x2d4ef8abef8e90de0cea3ee302a68e25d0d055ac",
		"0x6B35ba8b3b383714338686BcE4066B387Eab16C6",
		"0x153F5bf2cCf235bDeF530d09387dd456f6389a81",
		"0x0DbE6e626B1085a2Cb4D5F77389bBb54Ec89Fa30",
	}

	for _, account := range accounts {
		// 批量执行增加或减少质押量
		fmt.Println(account)
		blockchain.AdminSetDelegate(account, int64(inc), 1)
		blockchain.AdminSetDelegate(account, int64(dec), 2)
	}
	//blockchain.AdminSetDelegate("0xD5f92Fd92F8c7f9391513E3019D9441aAf5b2D9E", int64(inc), 1)
	//blockchain.AdminSetDelegate("0xD5f92Fd92F8c7f9391513E3019D9441aAf5b2D9E", int64(dec), 2)

}
