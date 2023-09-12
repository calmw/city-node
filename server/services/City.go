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

	//blockchain.AdminSetChengShiLevelAndSurety("0x95fecb63031dfc09623ae52f647df2e749533e3acda9a0ea73f34c29cd07d667", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xf7d89a46c223f9704a010c7077eb586fd7208822821abefa1a5e0fc5a420cbe3", 2, 60000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xa3c9a043a733fc013604651b53591dec8cc6a09c7ebebb413fa826d0a3388cfe", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xa3c9a043a733fc013604651b53591dec8cc6a09c7ebebb413fa826d0a3388cfe", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xbb2d35b76dc3358825b55db25f3696363681ccd940f5d39cc6b1767b8e069404", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x0a23fb256dc6340f9287a2edf424eb5cb92f37f39a8f58cc6c1a9fa349d51b34", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xe1b0f7944ee94273aa896973319bda53389ef3f1168c81a50e1c6c05307659ad", 2, 60000)
	// --------

	// 管理员设置城市先锋 ------------
	//blockchain.AdminSetPioneer("0x0a23fb256dc6340f9287a2edf424eb5cb92f37f39a8f58cc6c1a9fa349d51b34", "0xC31Bd725D0C15c085a0860Be0c054F60B22c937F")
	//blockchain.AdminSetPioneer("0xbb2d35b76dc3358825b55db25f3696363681ccd940f5d39cc6b1767b8e069404", "0x9a04896807Aa61458cB3f996F811022086C61698")
	//blockchain.AdminRemovePioneer("0xaa3f3f93c743005d497d2a21cf5ac5132960e42496d9b45437722351f8524496", "0x6B35ba8b3b383714338686BcE4066B387Eab16C6")
	blockchain.AdminSetPioneer("0xe1b0f7944ee94273aa896973319bda53389ef3f1168c81a50e1c6c05307659ad", "0x2F443072324b1777a79c2530Db1eE36270F7D2Aa")
	//time.Sleep(time.Second * 3)

	// 给城市先锋合约、用户定位合约、设置质押量合约添加管理员权限
	//blockchain.AddCityAdmin()

}

func AdminSetDelegateTask() {
	randInt := rand.Int() % 2
	if randInt == 0 {
		randInt = 2
	}
	inc := randInt * 1000
	dec := randInt * 500
	accounts := []string{
		"0x0d15C5f863Ff5CdD96A4B21060eE25DDe667182b",
		"0x2F443072324b1777a79c2530Db1eE36270F7D2Aa",
	}

	for _, account := range accounts {
		time.Sleep(time.Second * 2)
		// 批量执行增加或减少质押量
		blockchain.AdminSetDelegate(account, int64(inc), 1)
		time.Sleep(time.Second * 2)
		blockchain.AdminSetDelegate(account, int64(dec), 2)
	}
}

func AdminSetRechargeAmountTask() {
	inc := 2000
	accounts := []string{
		"0x0d15C5f863Ff5CdD96A4B21060eE25DDe667182b",
		"0x2F443072324b1777a79c2530Db1eE36270F7D2Aa",
	}

	for _, account := range accounts {
		time.Sleep(time.Second * 2)
		// 批量执行增加或减少质押量
		blockchain.AdminSetRechargeAmount(account, int64(inc))
	}
}
