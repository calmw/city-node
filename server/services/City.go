package services

import "city-node-server/blockchain"

func InitCity() {
	// 管理员设置先锋计划合约地址
	blockchain.AdminSetCityPioneerAddress()
	// 管理员设置用户定位合约地址
	blockchain.AdminSetUserLocationAddress()
	// 管理员设置先锋计划，城市等级以及该等级城市所需缴纳的保证金数额
	blockchain.AdminSetCityLevelAndSurety("0x367feef2d7c842ee11d3d5efe6ca7a5df6b9d3995a6d743c0a8bc11ac53f46c5", 1, 1000)
	blockchain.AdminSetCityLevelAndSurety("0xbffccd1d296a01011034d29d15df098fc72d9910750930042d973b671c2986bd", 2, 200)
	blockchain.AdminSetCityLevelAndSurety("0xfbbb3fd59e09ffc17b917954bdeb126312ae56c57c4ff78c086a94b55e364dfb", 3, 50)
	// 管理员设置城市先锋
	blockchain.AdminSetPioneer("0x367feef2d7c842ee11d3d5efe6ca7a5df6b9d3995a6d743c0a8bc11ac53f46c5", "0x6B35ba8b3b383714338686BcE4066B387Eab16C6")
	blockchain.AdminSetPioneer("0xbffccd1d296a01011034d29d15df098fc72d9910750930042d973b671c2986bd", "0xa75a076c5529b3813f53c9bd24ab1f7da37994fc")
	blockchain.AdminSetPioneer("0xfbbb3fd59e09ffc17b917954bdeb126312ae56c57c4ff78c086a94b55e364dfb", "0x8c69C5F4DbF59648682cAfe35557F94da4De1c28")
}
