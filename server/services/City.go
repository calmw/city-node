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
	//blockchain.AdminSetChengShiLevelAndSurety("0b3094b785e1c1d3479aa8955e9e6b898927828d5ac90406db5f078a48542ad3", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0b3094b785e1c1d3479aa8955e9e6b898927828d5ac90406db5f078a48542ad3", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x158b43ff3f04bfd1487d9112b3589528a7eb62dd38916bba315b536e69f97fd4", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xbd3f3948984277e24c4d15a7c885bf744fc9529d59cbac338e6ef13b10216e3d", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x9159fc828b5ea826b73279d2bee6e5a18cbcabad7f90b12264505d4d61f5b064", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x21897f13a80f729953b217858c04dc07dcc6c909a4fe19dedd85745de71698d5", 2, 60000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xb8f09d587defd54aa24dbb770b3773f29a658b080f95d6cf12aa0c8f6e97cef8", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xee72b2ea43e56148d63a0adf23653790ab0e0ab31bc0f8a4326ab9f377ee931a", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xd6b40febaa0ddd0e1e2674e83d08653b02ee51bef854cb11fe76ee91c1e9cfee", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xb6fc3268053782f854d25c8835b5b14310b43055e85822cb7a429b38e509672f", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x1d78790980b5a4917a2dfa1a4016ad10af0de8987f7a310acbc9baf85cad17f0", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xad1e2a3d586cb0bb665d7259e7521371730784148ba9a2e73fbac1ee79fdf5bb", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x936557454f464570813fe0c6e164d114cfc2d49eb5478f4e9867805eb4d3667b", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xc1e225d0bd8dc9302ee5642c2cecde889298c6e06fb527c43327f054efec8c38", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xa9c60bef6066612174e2e8ef885466d6cb8e46664cacadf073006b52bba8a109", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x1981a72a1320024da2570a1fa9c91299a0c7e3b999ebe7990b2611a251cb94a3", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xe5987567e81c5727f23926455949c1e7ef563497a7bb14e3cefe56ebbc79c6f6", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x44ab852f6b40d7cf25299e016cdfc5425c06b381ebcfe60be3575fb995f56694", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x64c970619edc45a6484aad8acc9fec4153dfc65ed6f061fe2458eb13755c1ee1", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x0bbc240bf18d0daa56f87a833ef6e4da5829169785b529a84599cf77ac04d917", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xdc07541f043227d960504e3aea526cd9e221ce7847ab0271bc988b862d4cbd57", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xbd3f3948984277e24c4d15a7c885bf744fc9529d59cbac338e6ef13b10216e3d", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x46edd53d46025bb997dfb63590e8b1a05006076bbe6cbea9503ea9dba07b13d3", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xbb2d35b76dc3358825b55db25f3696363681ccd940f5d39cc6b1767b8e069404", 3, 40000)
	// --------
	//time.Sleep(time.Second * 3)
	// 管理员设置城市先锋 ------------ 0xD5f92Fd92F8c7f9391513E3019D9441aAf5b2D9E
	//blockchain.AdminRemovePioneer("0xee72b2ea43e56148d63a0adf23653790ab0e0ab31bc0f8a4326ab9f377ee931a", "0x2F443072324b1777a79c2530Db1eE36270F7D2Aa")
	//blockchain.AdminSetPioneer("0xee72b2ea43e56148d63a0adf23653790ab0e0ab31bc0f8a4326ab9f377ee931a", "0x77bd41fdE654FE0054b771Ec6985dC9d5247BAfe")
	//blockchain.AdminSetPioneer("0xd6b40febaa0ddd0e1e2674e83d08653b02ee51bef854cb11fe76ee91c1e9cfee", "0x17dC6411D638672A073f23267C4735ca877AA623")
	//blockchain.AdminSetPioneer("0xb6fc3268053782f854d25c8835b5b14310b43055e85822cb7a429b38e509672f", "0xef3E99B0A86284e57c73d096f20c3B983cd4c18e")
	//blockchain.AdminSetPioneer("0x1d78790980b5a4917a2dfa1a4016ad10af0de8987f7a310acbc9baf85cad17f0", "0xE137fF4FCdDA90C3665562F52491B511155e19FF")
	//blockchain.AdminSetPioneer("0xad1e2a3d586cb0bb665d7259e7521371730784148ba9a2e73fbac1ee79fdf5bb", "0xF22fc6b430C265Fc6eC39d434897cE4E0bC2fD21")
	//blockchain.AdminSetPioneer("0x936557454f464570813fe0c6e164d114cfc2d49eb5478f4e9867805eb4d3667b", "0x2E418BF2d36eE6464E7Dcf08fBb994C78971F50A")
	//blockchain.AdminSetPioneer("0xc1e225d0bd8dc9302ee5642c2cecde889298c6e06fb527c43327f054efec8c38", "0x7295f5a83337163aA06c70dcdd51903aCc5E27fa")
	//blockchain.AdminSetPioneer("0xa9c60bef6066612174e2e8ef885466d6cb8e46664cacadf073006b52bba8a109", "0xFc542654708477B585a8C219397C34f921Ed0089")
	//blockchain.AdminSetPioneer("0x1981a72a1320024da2570a1fa9c91299a0c7e3b999ebe7990b2611a251cb94a3", "0x2Fde6c8b270dEeDb011F54eC8abDf08CA8cA5dfF")
	//blockchain.AdminSetPioneer("0xe5987567e81c5727f23926455949c1e7ef563497a7bb14e3cefe56ebbc79c6f6", "0xcE492Ae90D030DB5C833296F1ABdE6c088013E50")
	//blockchain.AdminSetPioneer("0x44ab852f6b40d7cf25299e016cdfc5425c06b381ebcfe60be3575fb995f56694", "0x8407DC14bFD44FDcEE0274766C08477f4C9748d2") // -----
	//blockchain.AdminSetPioneer("0x64c970619edc45a6484aad8acc9fec4153dfc65ed6f061fe2458eb13755c1ee1", "0xF032A7e9556f3a9688278e995c1Ac9CC7A676eD5")
	//blockchain.AdminSetPioneer("0x0bbc240bf18d0daa56f87a833ef6e4da5829169785b529a84599cf77ac04d917", "0x7763aE1D565Dcc422b7Bc07D5567281338a92A25")
	//blockchain.AdminSetPioneer("0xdc07541f043227d960504e3aea526cd9e221ce7847ab0271bc988b862d4cbd57", "0x06EfCb450059517f5b34a7cC4eFD80298825124E")
	//blockchain.AdminSetPioneer("0xbd3f3948984277e24c4d15a7c885bf744fc9529d59cbac338e6ef13b10216e3d", "0x1763271960D827D1DE4654D9E754021A0f5f94A5")
	//blockchain.AdminSetPioneer("0x46edd53d46025bb997dfb63590e8b1a05006076bbe6cbea9503ea9dba07b13d3", "0xD01BB85d6eB4fbd0EdF9A9a4885B3defC7e77ae8")
	//blockchain.AdminSetPioneer("0xbb2d35b76dc3358825b55db25f3696363681ccd940f5d39cc6b1767b8e069404", "0x9a04896807Aa61458cB3f996F811022086C61698")
	////time.Sleep(time.Second * 3)

	// 给城市先锋合约、用户定位合约、设置质押量合约添加管理员权限
	//blockchain.AddCityAdmin()

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
	dec := randInt * 500
	accounts := []string{
		"0xDB257161aa49a2353a8D7E5eee0f25D55a20c720",
		//"0x2F443072324b1777a79c2530Db1eE36270F7D2Aa",
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
		//"0x0d15C5f863Ff5CdD96A4B21060eE25DDe667182b",
		//"0x2F443072324b1777a79c2530Db1eE36270F7D2Aa",
		"0xe7FD766A7F0f1A8566Ead2bdEfFe88b80b3fA3A9",
	}

	for _, account := range accounts {
		time.Sleep(time.Second * 2)
		// 批量执行增加或减少质押量
		blockchain.AdminSetRechargeAmount(account, int64(inc))
	}
}
