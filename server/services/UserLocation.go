package services

import "city-node-server/blockchain"

func InitUserLocation() {

	// 管理员设置城市合约地址
	blockchain.AdminSetCityAddressUserLocation()
	// 管理员设置获取用户质押量合约地址
	blockchain.AdminSetPledgeStakeAddressUserLocation()
}
