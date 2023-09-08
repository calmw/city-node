package main

import (
	"city-node-server/services"
	"city-node-server/tasks"
)

func main() {
	//db.InitMysql()
	//tasks.Task()

	services.InitCity()
	services.InitCityPioneer()
	services.InitUserLocation()
	//

	//blockchain.AdminEditSurety("0x3edf72f7ab938a14e6aae3701a1e5acbb2512c840a019c6ead9d01415dbac864", 1, 100000)
	//blockchain.AdminEditSurety("0xa398d3dc089919305bd4f65d4b3caf5950702860341fa9cd9a59b53b75574c24", 2, 60000)
	//blockchain.AdminEditSurety("0x35cbb5a35a63817975c8f118ac59cf9e48a4f8109bb03276c58729c04540ebe7", 3, 4000)
	//blockchain.AdminEditSurety("0xe5987567e81c5727f23926455949c1e7ef563497a7bb14e3cefe56ebbc79c6f6", 1, 100000)

	//blockchain.SetUserLocation("0xad02e2bafae66da7b495ff56ccc86f2906814d6f5ab13e6dcd0348f12dc8bf0b", "Vx4fMYGPb3M=") // 设置用户位置
	//
	//blockchain.ApproveToxCityPioneer() // 向城市先锋合约approve
	//blockchain.CityPioneerBalance() // 城市先锋合约余额
	//blockchain.DepositSurety()         // 交保证金，慎用，扣钱
	//blockchain.DepositSuretyTest("0x08a01BE67fF47Ba2652b7dCE2005B47D81bAaC13") // DepositSuretyTest 交保证金成为先锋

	// 批量执行增加或减少质押量
	tasks.Task()

	//blockchain.GetCityId(22)

}
