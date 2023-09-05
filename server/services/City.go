package services

import "city-node-server/blockchain"

func InitCity() {
	// 管理员设置先锋计划合约地址
	blockchain.AdminSetCityPioneerAddress()
	// 管理员设置用户定位合约地址
	blockchain.AdminSetUserLocationAddress()
	// 管理员设置城市先锋
	blockchain.AdminSetPioneer("0x367feef2d7c842ee11d3d5efe6ca7a5df6b9d3995a6d743c0a8bc11ac53f46c5", "0x6B35ba8b3b383714338686BcE4066B387Eab16C6")
	// 管理员设置城市等级和保证金
	blockchain.AdminSetCityLevel("0x367feef2d7c842ee11d3d5efe6ca7a5df6b9d3995a6d743c0a8bc11ac53f46c5", 1, 1000)
}
