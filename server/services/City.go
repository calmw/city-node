package services

import (
	"city-node-server/blockchain"
	"math/rand"
	"time"
)

func InitCity() {
	// 管理员设置先锋计划合约地址
	//blockchain.AdminSetCityPioneerAddress()
	// 管理员设置用户定位合约地址
	//blockchain.AdminSetUserLocationAddress()
	// 管理员设置获取过去15天社交基金平均值的合约地址
	//blockchain.AdminSetFoundsAddress()

	// 管理员设置先锋计划，城市等级以及该等级城市所需缴纳的保证金数额

	//blockchain.AdminSetChengShiLevelAndSurety("0xe48baef0767f2198d4a783075148f0d7650294f840652960e9ea74f56c9171a6", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x95fecb63031dfc09623ae52f647df2e749533e3acda9a0ea73f34c29cd07d667", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xf7d89a46c223f9704a010c7077eb586fd7208822821abefa1a5e0fc5a420cbe3", 2, 60000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xf7d89a46c223f9704a010c7077eb586fd7208822821abefa1a5e0fc5a420cbe3", 2, 60000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xa3c9a043a733fc013604651b53591dec8cc6a09c7ebebb413fa826d0a3388cfe", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xb6fc3268053782f854d25c8835b5b14310b43055e85822cb7a429b38e509672f", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xf0d851e51c9ebb39dcecd7685ab590d6f69ff2b3c32b225e96ac51928e57d957", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x491f0fa71e9db46360a7ac880539f923cda458800f22fb344e10e98491f1f39a", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x71672b46f47097e2f0ebbd289c1c88e4e8558f02b70e3efa669142c1c9e9b6a1", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x8791cf7226e6b3c30f0557f8f0aaa201d606061d431d70458876f22dca6bb210", 2, 60000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x491f0fa71e9db46360a7ac880539f923cda458800f22fb344e10e98491f1f39a", 2, 60000)
	//time.Sleep(time.Second * 5)
	// --------
	//time.Sleep(time.Second * 3)
	// 管理员设置城市先锋 ------------ 0x95023738d7165bdc0da5544a021c1e9cfc1bc3f5f3e2351b8edd6abff1c02aab
	//blockchain.AdminSetPioneer("0x8791cf7226e6b3c30f0557f8f0aaa201d606061d431d70458876f22dca6bb210", "0xeAddEef3df4687d26bF6E160515AbC424c55D4F4")
	//blockchain.AdminRemovePioneer("0x8549d0e3957c74e498e26c5207d83edbfcf4f2df01e4105964f334c752e1a51a", "0xCe07c0Df73EE82A1F77b0bD3c6967a7f6a29342B")
	//time.Sleep(time.Second * 5)
	//blockchain.AdminSetPioneer("0x8549d0e3957c74e498e26c5207d83edbfcf4f2df01e4105964f334c752e1a51a", "0xCe07c0Df73EE82A1F77b0bD3c6967a7f6a29342B")
	//time.Sleep(time.Second * 5)
	//blockchain.AdminRemovePioneer("0x71672b46f47097e2f0ebbd289c1c88e4e8558f02b70e3efa669142c1c9e9b6a1", "0xFAa3e964bB91fc62270DCf3DBfa7Bf9e8239E697")
	//time.Sleep(time.Second * 5)
	//blockchain.AdminSetPioneer("0x71672b46f47097e2f0ebbd289c1c88e4e8558f02b70e3efa669142c1c9e9b6a1", "0xFAa3e964bB91fc62270DCf3DBfa7Bf9e8239E697")
	//blockchain.AdminSetPioneer("0xb6fc3268053782f854d25c8835b5b14310b43055e85822cb7a429b38e509672f", "0xe7FD766A7F0f1A8566Ead2bdEfFe88b80b3fA3A9")
	//blockchain.AdminRemovePioneer("0xe48baef0767f2198d4a783075148f0d7650294f840652960e9ea74f56c9171a6", "0x0151D7707B3D20d8Fcaa9A6448bc42663cF46736")
	//blockchain.AdminSetPioneer("0xbb2d35b76dc3358825b55db25f3696363681ccd940f5d39cc6b1767b8e069404", "0x9a04896807Aa61458cB3f996F811022086C61698")
	////time.Sleep(time.Second * 3)

	// 给城市先锋合约、用户定位合约、设置质押量合约添加管理员权限
	blockchain.AddCityAdmin()

	//blockchain.AdminSetChengShiLevelAndSurety("0xb6fc3268053782f854d25c8835b5b14310b43055e85822cb7a429b38e509672f", 1, 100000)
	//blockchain.SetCityMapping("0xb73482b8711de7a9ea6d605da1662617f1f75d62597e3a5fe7b56dafc029bc2a", "0xb6fc3268053782f854d25c8835b5b14310b43055e85822cb7a429b38e509672f", "R9TjhZrWJi8=")
	// 管理员设置城市先锋 ------------ 0xD5f92Fd92F8c7f9391513E3019D9441aAf5b2D9E
	//blockchain.AdminSetPioneer("0xb6fc3268053782f854d25c8835b5b14310b43055e85822cb7a429b38e509672f", "0xe7FD766A7F0f1A8566Ead2bdEfFe88b80b3fA3A9")

}

func AdminSetDelegateTask() {
	randInt := rand.Int() % 2
	if randInt == 0 {
		randInt = 2
	}
	inc := randInt * 1000
	//dec := randInt * 500
	accounts := []string{
		"0xCe07c0Df73EE82A1F77b0bD3c6967a7f6a29342B",
		"0xFAa3e964bB91fc62270DCf3DBfa7Bf9e8239E697",
	}

	for _, account := range accounts {
		//time.Sleep(time.Second * 2)
		// 批量执行增加或减少质押量
		blockchain.AdminSetDelegate(account, int64(inc), 1)
		time.Sleep(time.Second * 1)
		//blockchain.AdminSetDelegate(account, int64(dec), 2)
	}
}

func AdminSetRechargeAmountTask() {
	//inc := 5000
	inc := 50000
	//inc := 1000
	//inc := 1000000000
	accounts := []string{
		"0xCe07c0Df73EE82A1F77b0bD3c6967a7f6a29342B",
	}

	for _, account := range accounts {
		time.Sleep(time.Second * 2)
		// 批量执行增加或减少质押量
		blockchain.AdminSetRechargeAmount(account, int64(inc))
	}
}

func AdminSetRechargeAmountTask2500() {
	//inc := 5000
	inc := 250000
	//inc := 1000
	//inc := 1000000000
	accounts := []string{
		"0xCe07c0Df73EE82A1F77b0bD3c6967a7f6a29342B",
	}

	for _, account := range accounts {
		time.Sleep(time.Second * 2)
		// 批量执行增加或减少质押量
		blockchain.AdminSetRechargeAmount(account, int64(inc))
	}
}

func AdminSetRechargeAmountTask5000() {
	//inc := 5000
	inc := 500000
	//inc := 1000
	//inc := 1000000000
	accounts := []string{
		"0xCe07c0Df73EE82A1F77b0bD3c6967a7f6a29342B",
	}

	for _, account := range accounts {
		time.Sleep(time.Second * 2)
		// 批量执行增加或减少质押量
		blockchain.AdminSetRechargeAmount(account, int64(inc))
	}
}

func AdminSetRechargeAmountTask7500() {
	//inc := 5000
	inc := 750000
	//inc := 1000
	//inc := 1000000000
	accounts := []string{
		"0xCe07c0Df73EE82A1F77b0bD3c6967a7f6a29342B",
	}

	for _, account := range accounts {
		time.Sleep(time.Second * 2)
		// 批量执行增加或减少质押量
		blockchain.AdminSetRechargeAmount(account, int64(inc))
	}
}
