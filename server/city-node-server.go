package main

import (
	"city-node-server/blockchain"
	"city-node-server/services"
)

func main() {
	//db.InitMysql()
	//tasks.Task()

	services.InitCity()
	services.InitCityPioneer()
	services.InitUserLocation()

	// DepositSuretyTest 交保证金成为先锋
	blockchain.DepositSuretyTest("0xD5f92Fd92F8c7f9391513E3019D9441aAf5b2D9E")

	blockchain.DepositSurety()
}
