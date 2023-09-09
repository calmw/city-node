package services

import (
	"city-node-server/blockchain"
	"fmt"
	"math/rand"
	"time"
)

func InitCity() {
	// 管理员设置先锋计划合约地址
	//blockchain.AdminSetCityPioneerAddress()
	// 管理员设置用户定位合约地址
	//blockchain.AdminSetUserLocationAddress()
	// 管理员设置先锋计划，城市等级以及该等级城市所需缴纳的保证金数额
	//blockchain.AdminSetChengShiLevelAndSurety("0xaa3f3f93c743005d497d2a21cf5ac5132960e42496d9b45437722351f8524496", 1, 100000)
	// -----
	//blockchain.AdminSetChengShiLevelAndSurety("0x7651f967164a411bfd394cab35d575ddf5cada6e0d11bcf4c984c6867f6b73e2", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xaa3f3f93c743005d497d2a21cf5ac5132960e42496d9b45437722351f8524496", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xb3b348b5d62537a0d87eb58f09045cd2a2d2631fb72a5c58d912bf8f751f1ede", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x895e139e0a4c629694d8c590f4c0f05ada11904b7997cf30756d69022c64301c", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x890d4253b0eaef52eb154158ccbc49aafe7d53db49eb15215f8874024a9fbd91", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x91fa5e0e901e28b581707c10829194cd85d1d281e4d7dbf0dfd0d78463cd2e87", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x95fecb63031dfc09623ae52f647df2e749533e3acda9a0ea73f34c29cd07d667", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xa3c9a043a733fc013604651b53591dec8cc6a09c7ebebb413fa826d0a3388cfe", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x2cc69c7ca78ccbc134cbd3eef57f3c882263ecadf071a0f7b14fc8cc4cc1dc44", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x5d2e8bf23f54504a2f6afbe8abe5eafc331b7da7c166287da97c40935c2a49b2", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x59eaad55ef447a968257797d75e289157b780ef63027267b85c07b1309b6edee", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x80af9895d0010ac6650f3048ff3d21fcbf3e2583faf72a01c57d377981f91181", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xdf94a1319cc8f5fb67fdfe859ef61d7f883589f463cbd0ef14a95f4bcbf03159", 3, 40000)
	// --------

	//blockchain.AdminSetChengShiLevelAndSurety("0xf7d89a46c223f9704a010c7077eb586fd7208822821abefa1a5e0fc5a420cbe3", 2, 60000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x35cbb5a35a63817975c8f118ac59cf9e48a4f8109bb03276c58729c04540ebe7", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x3a88d7ca2d9ba961030f92a27f5b42aaf63d78d2b477fe9dfd140a1f7dbea576", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x6478eab9198cc2c0f53912730434acabfaa199b7471f6b5d5fd5b1413aa672f1", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xc8b0a5ed8a761c9e4d3fb44625cf80d29f38c1115a381154b9ab236058414f9c", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xb736dbc2078ef993011021edaffaab7286696e849d205ea3b2e43323a3466d24", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x4de9096348869d087c953a1b6df5267276b15929bf8c3eedd52332bb946dadda", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x0a23fb256dc6340f9287a2edf424eb5cb92f37f39a8f58cc6c1a9fa349d51b34", 3, 40000)
	////
	////// 管理员设置城市先锋 ------------
	//blockchain.AdminSetPioneer("0xaa3f3f93c743005d497d2a21cf5ac5132960e42496d9b45437722351f8524496", "0x5a1ECB5fb2e1EbAbF70f2a331F3738f94563e89a")
	blockchain.AdminRemovePioneer("0xaa3f3f93c743005d497d2a21cf5ac5132960e42496d9b45437722351f8524496", "0x5a1ECB5fb2e1EbAbF70f2a331F3738f94563e89a")
	time.Sleep(time.Second * 3)
	//blockchain.AdminSetPioneer("0xb3b348b5d62537a0d87eb58f09045cd2a2d2631fb72a5c58d912bf8f751f1ede", "0xda20b78bF48389DfBD66FDA4725c4b8BaA0d7b04")
	//blockchain.AdminRemovePioneer("0xb3b348b5d62537a0d87eb58f09045cd2a2d2631fb72a5c58d912bf8f751f1ede", "0xda20b78bF48389DfBD66FDA4725c4b8BaA0d7b04")
	//time.Sleep(time.Second * 3)
	//blockchain.AdminSetPioneer("0xb3b348b5d62537a0d87eb58f09045cd2a2d2631fb72a5c58d912bf8f751f1ede", "0x5a1ECB5fb2e1EbAbF70f2a331F3738f94563e89a")
	blockchain.AdminRemovePioneer("0xb3b348b5d62537a0d87eb58f09045cd2a2d2631fb72a5c58d912bf8f751f1ede", "0x5a1ECB5fb2e1EbAbF70f2a331F3738f94563e89a")
	time.Sleep(time.Second * 3)
	// ---- 开始
	blockchain.AdminSetPioneer("0xaa3f3f93c743005d497d2a21cf5ac5132960e42496d9b45437722351f8524496", "0x77bd41fdE654FE0054b771Ec6985dC9d5247BAfe")
	time.Sleep(time.Second * 3)
	blockchain.AdminSetPioneer("0xb3b348b5d62537a0d87eb58f09045cd2a2d2631fb72a5c58d912bf8f751f1ede", "0x17dC6411D638672A073f23267C4735ca877AA623")

	//blockchain.AdminSetPioneer("0xa3c9a043a733fc013604651b53591dec8cc6a09c7ebebb413fa826d0a3388cfe", "0xef3E99B0A86284e57c73d096f20c3B983cd4c18e")
	//blockchain.AdminSetPioneer("0x895e139e0a4c629694d8c590f4c0f05ada11904b7997cf30756d69022c64301c", "0xE137fF4FCdDA90C3665562F52491B511155e19FF")
	//blockchain.AdminSetPioneer("0x2cc69c7ca78ccbc134cbd3eef57f3c882263ecadf071a0f7b14fc8cc4cc1dc44", "0xF22fc6b430C265Fc6eC39d434897cE4E0bC2fD21")
	//blockchain.AdminSetPioneer("0x5d2e8bf23f54504a2f6afbe8abe5eafc331b7da7c166287da97c40935c2a49b2", "0x2E418BF2d36eE6464E7Dcf08fBb994C78971F50A")
	//blockchain.AdminSetPioneer("0x59eaad55ef447a968257797d75e289157b780ef63027267b85c07b1309b6edee", "0x7295f5a83337163aA06c70dcdd51903aCc5E27fa")
	//blockchain.AdminSetPioneer("0x80af9895d0010ac6650f3048ff3d21fcbf3e2583faf72a01c57d377981f91181", "0xFc542654708477B585a8C219397C34f921Ed0089")
	//blockchain.AdminSetPioneer("0xdf94a1319cc8f5fb67fdfe859ef61d7f883589f463cbd0ef14a95f4bcbf03159", "0x2Fde6c8b270dEeDb011F54eC8abDf08CA8cA5dfF")
	//blockchain.AdminSetPioneer("0x890d4253b0eaef52eb154158ccbc49aafe7d53db49eb15215f8874024a9fbd91", "0xcE492Ae90D030DB5C833296F1ABdE6c088013E50")
	////blockchain.AdminSetPioneer("", "0x8407DC14bFD44FDcEE0274766C08477f4C9748d2")
	//blockchain.AdminSetPioneer("0x91fa5e0e901e28b581707c10829194cd85d1d281e4d7dbf0dfd0d78463cd2e87", "0xF032A7e9556f3a9688278e995c1Ac9CC7A676eD5")
	//blockchain.AdminSetPioneer("0x95fecb63031dfc09623ae52f647df2e749533e3acda9a0ea73f34c29cd07d667", "0x7763aE1D565Dcc422b7Bc07D5567281338a92A25")
	// ------ 结束
	//blockchain.AdminSetPioneer("0xaa3f3f93c743005d497d2a21cf5ac5132960e42496d9b45437722351f8524496", "0x77bd41fdE654FE0054b771Ec6985dC9d5247BAfe")
	//blockchain.AdminSetPioneer("0x35cbb5a35a63817975c8f118ac59cf9e48a4f8109bb03276c58729c04540ebe7", "0xD5f92Fd92F8c7f9391513E3019D9441aAf5b2D9E")
	//blockchain.AdminSetPioneer("0x6478eab9198cc2c0f53912730434acabfaa199b7471f6b5d5fd5b1413aa672f1", "0x2D4EF8abEF8E90DE0cea3Ee302A68e25d0d055AC")
	//blockchain.AdminSetPioneer("0x3a88d7ca2d9ba961030f92a27f5b42aaf63d78d2b477fe9dfd140a1f7dbea576", "0x0d15C5f863Ff5CdD96A4B21060eE25DDe667182b")
	//blockchain.AdminSetPioneer("0xc8b0a5ed8a761c9e4d3fb44625cf80d29f38c1115a381154b9ab236058414f9c", "0xa10edEb8b3486C7433925081dC09Ec8A8B941677")
	//blockchain.AdminSetPioneer("0xb736dbc2078ef993011021edaffaab7286696e849d205ea3b2e43323a3466d24", "0xa75a076c5529b3813f53c9bd24ab1f7da37994fc")
	//blockchain.AdminSetPioneer("0x4de9096348869d087c953a1b6df5267276b15929bf8c3eedd52332bb946dadda", "0x5a1ECB5fb2e1EbAbF70f2a331F3738f94563e89a")
	//blockchain.AdminSetPioneer("0x7651f967164a411bfd394cab35d575ddf5cada6e0d11bcf4c984c6867f6b73e2", "0xda20b78bF48389DfBD66FDA4725c4b8BaA0d7b04")

	// 给城市先锋合约、用户定位合约、设置质押量合约添加管理员权限
	//blockchain.AddCityAdmin()
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
