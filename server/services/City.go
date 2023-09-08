package services

import (
	"city-node-server/blockchain"
	"fmt"
	"math/rand"
)

func InitCity() {
	// 管理员设置先锋计划合约地址
	blockchain.AdminSetCityPioneerAddress()
	// 管理员设置用户定位合约地址
	blockchain.AdminSetUserLocationAddress()
	// 管理员设置先锋计划，城市等级以及该等级城市所需缴纳的保证金数额
	blockchain.AdminSetCityLevelAndSurety("0x2bd1a23d1755451bd63917d541b8b823c5ef95e119abfee92f2ca95550f7b1f2", 1, 100000)
	blockchain.AdminSetCityLevelAndSurety("0x6a35f6c500cbb91ed6c2ed6cfd80d5cdfe40f01ed5879e26a0050a26395ba781", 2, 60000)
	blockchain.AdminSetCityLevelAndSurety("0x35cbb5a35a63817975c8f118ac59cf9e48a4f8109bb03276c58729c04540ebe7", 3, 4000)
	blockchain.AdminSetCityLevelAndSurety("3b41ac54ca28591721f3a04355e07f6b23385bc91424754ce799625cd06bf494", 3, 4000)

	// 管理员设置城市先锋 ------------
	blockchain.AdminSetPioneer("0x2bd1a23d1755451bd63917d541b8b823c5ef95e119abfee92f2ca95550f7b1f2", "0x0DbE6e626B1085a2Cb4D5F77389bBb54Ec89Fa30")
	blockchain.AdminSetPioneer("0x6a35f6c500cbb91ed6c2ed6cfd80d5cdfe40f01ed5879e26a0050a26395ba781", "0x153F5bf2cCf235bDeF530d09387dd456f6389a81")
	blockchain.AdminSetPioneer("0x0a23fb256dc6340f9287a2edf424eb5cb92f37f39a8f58cc6c1a9fa349d51b34", "0xa75a076c5529b3813f53c9bd24ab1f7da37994fc")
	blockchain.AdminSetPioneer("3b41ac54ca28591721f3a04355e07f6b23385bc91424754ce799625cd06bf494", "0x26baf7283a7fe1FA182e57716100f97385dE7bDa")

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
