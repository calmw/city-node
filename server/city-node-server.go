package main

import (
	"city-node-server/blockchain"
)

func main() {
	//db.InitMysql()
	//tasks.Task()
	//services.InitCity()
	//services.InitCityPioneer()

	// DepositSuretyTest 交保证金成为先锋
	blockchain.DepositSuretyTest("0x8c69C5F4DbF59648682cAfe35557F94da4De1c28")
}
