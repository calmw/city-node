package main

import (
	"city-node-server/services"
)

func main() {
	//db.InitMysql()
	//tasks.Task()

	services.InitCity()
	services.InitCityPioneer()
	services.InitUserLocation()

	//blockchain.ApproveToxCityPioneer() // 向城市先锋合约approve
	//blockchain.CityPioneerBalance()    // 城市先锋合约余额
	//blockchain.DepositSurety() // 交保证金，慎用，扣钱
	//blockchain.DepositSuretyTest("0xD5f92Fd92F8c7f9391513E3019D9441aAf5b2D9E") // DepositSuretyTest 交保证金成为先锋
}
