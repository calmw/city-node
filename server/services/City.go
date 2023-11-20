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

	//blockchain.AdminSetChengShiLevelAndSurety("0xc8164ecc2f74f797d66c6b5ca5ffbd5a415de35c2a5a3a45e8e0b008ad758625", 1, 100000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xe5987567e81c5727f23926455949c1e7ef563497a7bb14e3cefe56ebbc79c6f6", 2, 60000)
	//blockchain.AdminSetChengShiLevelAndSurety("0xa3c9a043a733fc013604651b53591dec8cc6a09c7ebebb413fa826d0a3388cfe", 3, 40000)

	blockchain.AdminSetChengShiLevelAndSurety("0x31a3bec1b8b480b83b011ad214d417397953990dc395d94f379a7dd54e699c70", 1, 100000)
	//time.Sleep(time.Second * 5)
	// --------
	//time.Sleep(time.Second * 3)
	// 管理员设置城市先锋 ------------ 0x95023738d7165bdc0da5544a021c1e9cfc1bc3f5f3e2351b8edd6abff1c02aab
	blockchain.AdminSetPioneer("0x31a3bec1b8b480b83b011ad214d417397953990dc395d94f379a7dd54e699c70", "0x5b1003eA0f799BD05b6E9750d0E79c570Ab52463")
	//blockchain.AdminRemovePioneer("0xe5987567e81c5727f23926455949c1e7ef563497a7bb14e3cefe56ebbc79c6f6", "0xD6439Dcb4CFacb56d0bc8DADFEAaf2bd7C8dEEd8")
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
		"0xa10edEb8b3486C7433925081dC09Ec8A8B941677",
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
