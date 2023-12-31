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
	blockchain.AdminSetCityLevelAndSurety("0x367feef2d7c842ee11d3d5efe6ca7a5df6b9d3995a6d743c0a8bc11ac53f46c5", 1, 1000)
	blockchain.AdminSetCityLevelAndSurety("0xbffccd1d296a01011034d29d15df098fc72d9910750930042d973b671c2986bd", 2, 200)
	blockchain.AdminSetCityLevelAndSurety("0xfbbb3fd59e09ffc17b917954bdeb126312ae56c57c4ff78c086a94b55e364dfb", 3, 50)
	blockchain.AdminSetCityLevelAndSurety("0xad02e2bafae66da7b495ff56ccc86f2906814d6f5ab13e6dcd0348f12dc8bf0b", 3, 50)
	// 管理员设置城市先锋
	blockchain.AdminSetPioneer("0x367feef2d7c842ee11d3d5efe6ca7a5df6b9d3995a6d743c0a8bc11ac53f46c5", "0x6B35ba8b3b383714338686BcE4066B387Eab16C6")
	blockchain.AdminSetPioneer("0xbffccd1d296a01011034d29d15df098fc72d9910750930042d973b671c2986bd", "0xa75a076c5529b3813f53c9bd24ab1f7da37994fc")
	blockchain.AdminSetPioneer("0xfbbb3fd59e09ffc17b917954bdeb126312ae56c57c4ff78c086a94b55e364dfb", "0x8c69C5F4DbF59648682cAfe35557F94da4De1c28")
	blockchain.AdminSetPioneer("0xad02e2bafae66da7b495ff56ccc86f2906814d6f5ab13e6dcd0348f12dc8bf0b", "0xD5f92Fd92F8c7f9391513E3019D9441aAf5b2D9E")
	blockchain.AdminSetPioneer("0x1b25ba0ac0336a11af717696e4ef339e2ed36dbabfd33f3d623f3516fd38cf0a", "0x08a01BE67fF47Ba2652b7dCE2005B47D81bAaC13")
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

	// 批量执行增加或减少质押量
	blockchain.AdminSetDelegate("0xD5f92Fd92F8c7f9391513E3019D9441aAf5b2D9E", int64(inc), 1)
	blockchain.AdminSetDelegate("0xD5f92Fd92F8c7f9391513E3019D9441aAf5b2D9E", int64(dec), 2)

}
