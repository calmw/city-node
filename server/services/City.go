package services

import (
	"city-node-server/pkg/blockchain"
	"math/rand"
	"time"
)

func InitCity(secondsPerDay int64) {
	//appraise := blockchain.NewAppraise()
	// 管理员设置先锋计划合约地址
	//blockchain.AdminSetCityPioneerAddress()
	// 管理员设置用户定位合约地址
	//blockchain.AdminSetUserLocationAddress()
	// 管理员设置获取过去15天社交基金平均值的合约地址
	//blockchain.AdminSetFoundsAddress()
	// 设置每天秒数
	//blockchain.AdminSetSecondsPerDay(secondsPerDay)
	//blockchain.AdminSetAuthAddress()
	//blockchain.AdminSetWithdrawLimitAddress()

	// 管理员设置先锋计划，城市等级以及该等级城市所需缴纳的保证金数额

	//blockchain.AdminSetChengShiLevelAndSurety("0xc8164ecc2f74f797d66c6b5ca5ffbd5a415de35c2a5a3a45e8e0b008ad758625", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xe5987567e81c5727f23926455949c1e7ef563497a7bb14e3cefe56ebbc79c6f6", 2, 60000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xa3c9a043a733fc013604651b53591dec8cc6a09c7ebebb413fa826d0a3388cfe", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x9159fc828b5ea826b73279d2bee6e5a18cbcabad7f90b12264505d4d61f5b064", 2, 60000)

	//blockchain.AdminSetChengShiLevelAndSurety("0x45773751d6f8d5ecd20c3b2d3b0bc72abd5e605a1c239fc1c18bd302103dbc68", 2, 60000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x395bd804ecf082fce7050f6b3264d4a2f4f802dfb6067cef2d051353163a9eb2", 2, 60000)

	//blockchain.AdminSetChengShiLevelAndSurety("0xc50db5c60dc66068637ffc825954152fbce8b73a290285f252697c31b005a382", 3, 40000)
	//blockchain.AdminSetChengShiLevelAndSurety("0x223132fd87c1a316c5481cb862236f68a1f046598487b8348118d69d8d10ea22", 3, 40000)

	//time.Sleep(time.Second * 5)
	// --------
	//time.Sleep(time.Second * 3)
	// 管理员设置城市先锋

	//blockchain.PioneerChengShi("0xa10edEb8b3486C7433925081dC09Ec8A8B941677")
	//blockchain.AdminRemovePioneer("0x223132fd87c1a316c5481cb862236f68a1f046598487b8348118d69d8d10ea22", "0xc250bfc206ba2d9b761eb51537ef8d665e0a54d7")
	//blockchain.AdminSetPioneer("0x223132fd87c1a316c5481cb862236f68a1f046598487b8348118d69d8d10ea22", "0xc250bfc206ba2d9b761eb51537ef8d665e0a54d7")

	// 焦作市
	//blockchain.AdminRemovePioneer("0x85202e610a59bdbc7559cabd41c28f4ecfe18675575572ac5822ca697df5103d", "0x81F61e0b02d899956d2d96ACc2c6F9Cb43D6b99d")
	//blockchain.AdminSetPioneer("0x85202e610a59bdbc7559cabd41c28f4ecfe18675575572ac5822ca697df5103d", "0x81F61e0b02d899956d2d96ACc2c6F9Cb43D6b99d")

	// 新乡
	//blockchain.AdminRemovePioneer("0x9882a6038f5037928f8ceabccbbbe73652f94cda789484427957542073dbb2f5", "0x360C815e8C5F130913113801D0c57611Ee95723A")
	blockchain.AdminSetPioneer("0x9882a6038f5037928f8ceabccbbbe73652f94cda789484427957542073dbb2f5", "0x360C815e8C5F130913113801D0c57611Ee95723A", 3)

	// 杭州
	//blockchain.AdminRemovePioneer("0xe48baef0767f2198d4a783075148f0d7650294f840652960e9ea74f56c9171a6", "0x28ddE14d31Bd7025b3Db1FA8eC7C5707E4FFE1e8")
	//blockchain.AdminSetPioneer("0xe48baef0767f2198d4a783075148f0d7650294f840652960e9ea74f56c9171a6", "0x28ddE14d31Bd7025b3Db1FA8eC7C5707E4FFE1e8")

	//appraise := blockchain.NewAppraise()
	//appraise.AdminSetPioneerBatch("0x81F61e0b02d899956d2d96ACc2c6F9Cb43D6b99d", 1) // 焦作市
	//appraise.AdminSetPioneerBatch("0x360C815e8C5F130913113801D0c57611Ee95723A", 2) // 新乡
	//appraise.AdminSetPioneerBatch("0x28ddE14d31Bd7025b3Db1FA8eC7C5707E4FFE1e8", 2) // 杭州

	// 郑州
	//blockchain.AdminRemovePioneer("0xe5987567e81c5727f23926455949c1e7ef563497a7bb14e3cefe56ebbc79c6f6", "0x54129df2ed332ef039dc70a3fffe5ee8700b574c")
	//blockchain.AdminSetPioneer("0xe5987567e81c5727f23926455949c1e7ef563497a7bb14e3cefe56ebbc79c6f6", "0x54129df2ed332ef039dc70a3fffe5ee8700b574c")

	// 海口
	//blockchain.AdminRemovePioneer("0xb9cd82388e3ee54bbff9b4a2ed627865b7aa43466fccea379fef8264450a2a31", "0xeE25f9793dBAB14FA2b7143CDfc04f9F9eFb0c62")
	//blockchain.AdminSetPioneer("0xb9cd82388e3ee54bbff9b4a2ed627865b7aa43466fccea379fef8264450a2a31", "0xeE25f9793dBAB14FA2b7143CDfc04f9F9eFb0c62")

	// 佛山
	//blockchain.AdminRemovePioneer("0xc50db5c60dc66068637ffc825954152fbce8b73a290285f252697c31b005a382", "0xe4c5E5cEDd9b509E1Baf3f35b28FfA99Be37f4F3")
	//blockchain.AdminSetPioneer("0xc50db5c60dc66068637ffc825954152fbce8b73a290285f252697c31b005a382", "0xe4c5E5cEDd9b509E1Baf3f35b28FfA99Be37f4F3")

	// 成都
	//blockchain.AdminRemovePioneer("0xee72b2ea43e56148d63a0adf23653790ab0e0ab31bc0f8a4326ab9f377ee931a", "0xA3655DB8DA1466388c8C70D1200c22bb83823f28")
	//blockchain.AdminSetPioneer("0xee72b2ea43e56148d63a0adf23653790ab0e0ab31bc0f8a4326ab9f377ee931a", "0xA3655DB8DA1466388c8C70D1200c22bb83823f28")
	//appraise := blockchain.NewAppraise()
	//appraise.AdminSetPioneerBatch("0xA3655DB8DA1466388c8C70D1200c22bb83823f28", 1)
	//
	//blockchain.AdminRemovePioneer("0xee72b2ea43e56148d63a0adf23653790ab0e0ab31bc0f8a4326ab9f377ee931a", "0xD01BB85d6eB4fbd0EdF9A9a4885B3defC7e77ae8")
	//blockchain.AdminSetPioneer("0x969b96ab33464703a9fafb3d1bd8235389b83d2f16b107faa389f6a497fb0699", "0xD01BB85d6eB4fbd0EdF9A9a4885B3defC7e77ae8")

	//appraise := blockchain.NewAppraise()
	//appraise.AdminSetPioneerBatch("0xe4c5E5cEDd9b509E1Baf3f35b28FfA99Be37f4F3", 3)

	//blockchain.AdminSetPioneer("0xca98b991031d3b83e27132a5373c2a8b7d68c8e6477cb6ab43baf2268b3e9639", "0xd7922692C157Ee415FaCfe700e7A3e616f7B12C8")
	//blockchain.AdminRemovePioneer("0x3a92a49f884dba5f56758df8f440b9be669b3ea2501c6094ed4f1fabe3914014", "0x81F61e0b02d899956d2d96ACc2c6F9Cb43D6b99d")
	//blockchain.AdminSetPioneer("0xe5987567e81c5727f23926455949c1e7ef563497a7bb14e3cefe56ebbc79c6f6", "0xD6439Dcb4CFacb56d0bc8DADFEAaf2bd7C8dEEd8")
	//time.Sleep(time.Second * 5)
	//blockchain.AdminSetPioneer("0x44ab852f6b40d7cf25299e016cdfc5425c06b381ebcfe60be3575fb995f56694", "0x97f3A085bC4f88915F95D181e5bd9500B452D412")
	//blockchain.AdminRemovePioneer("0x291ab56b0fa225d0c7e26139f33fddc04343453345855ad5f0b207fe57932841", "0xe91E938F4A2a5f9e2388Fb5a6881115C74b3C8E2")
	//blockchain.AdminRemovePioneer("0xe5987567e81c5727f23926455949c1e7ef563497a7bb14e3cefe56ebbc79c6f6", "0xa75a076c5529b3813f53c9bd24ab1f7da37994fc")
	//time.Sleep(time.Second * 5)
	//blockchain.AdminRemovePioneer("0xe5987567e81c5727f23926455949c1e7ef563497a7bb14e3cefe56ebbc79c6f6", "0xa75a076c5529b3813f53c9bd24ab1f7da37994fc")
	//time.Sleep(time.Second * 5)
	//blockchain.AdminSetPioneer("0xe5987567e81c5727f23926455949c1e7ef563497a7bb14e3cefe56ebbc79c6f6", "0xa75a076c5529b3813f53c9bd24ab1f7da37994fc")

	//blockchain.AdminRemovePioneer("0x049ebcbf643d2a2eb061ae48960b133d5242e0b502b50e450ac59fabd62eae33", "0x7dd71454F60EA232b12BFE61C251F0b76D2D8E50")
	//time.Sleep(time.Second * 5)
	//blockchain.AdminSetPioneer("0x049ebcbf643d2a2eb061ae48960b133d5242e0b502b50e450ac59fabd62eae33", "0x7D574395807fd442F199Baa05166eC267b7776f8")
	//blockchain.AdminRemovePioneer("0x28cec2fd3e4a8670956e93ca9a9d9166302f0d9dd8086bdd0408f898873daadc", "0x8c69C5F4DbF59648682cAfe35557F94da4De1c28")
	//time.Sleep(time.Second * 5)
	//blockchain.AdminSetPioneer("0x28cec2fd3e4a8670956e93ca9a9d9166302f0d9dd8086bdd0408f898873daadc", "0x8c69C5F4DbF59648682cAfe35557F94da4De1c28")
	//time.Sleep(time.Second * 5)

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
	inc := 5000
	//inc := 50000
	//inc := 500000
	//inc := 1000
	//inc := 1000000000
	accounts := []string{
		//"0xa75a076c5529b3813f53c9bd24ab1f7da37994fc",
		"0xFf2B12085c7F7B8133eEf5006703A6c055a3ed7d",
		"0x0151D7707B3D20d8Fcaa9A6448bc42663cF46736",
		"0x81F61e0b02d899956d2d96ACc2c6F9Cb43D6b99d",
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
