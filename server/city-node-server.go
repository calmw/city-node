package main

import (
	"city-node-server/services"
)

func main() {
	//db.InitMysql()
	//tasks.Task()

	//services.InitTestNet()
	services.InitMainNet()

	//services.InitCity()
	//services.InitCityPioneer()
	//services.InitUserLocation()
	//

	//blockchain.AdminEditSurety("0x3edf72f7ab938a14e6aae3701a1e5acbb2512c840a019c6ead9d01415dbac864", 1, 100000) // 设置城市保证金、等级

	//blockchain.SetUserLocation("0xad02e2bafae66da7b495ff56ccc86f2906814d6f5ab13e6dcd0348f12dc8bf0b", "Vx4fMYGPb3M=") // 设置用户位置

	//
	//blockchain.ApproveToxCityPioneer() // 向城市先锋合约approve
	//blockchain.CityPioneerBalance() // 城市先锋合约余额
	//blockchain.DepositSurety() // 交保证金，慎用，扣钱
	//blockchain.DepositSuretyTest("0x08a01BE67fF47Ba2652b7dCE2005B47D81bAaC13") // DepositSuretyTest 交保证金成为先锋

	// 管理员设置开始考核时间,先交保证金，后考核
	//blockchain.AdminSetStartTime(time.Now().Add(time.Hour * 1000000).Unix())
	//blockchain.AdminSetStartTime(time.Now().Unix())
	// 批量执行增加或减少质押量
	//tasks.Task()

	// 获取前15天社交基金平均值
	//blockchain.GetFifteenDayAverageFounds()

	//blockchain.GetCityId(22)

	// 设置城市ID位置,设置映射关系
	//blockchain.SetCityMapping("0x818c1d2b3417786f35375c0a532bbd70a98cfcba5040c9dae6a64ac90de11d37", "0xb8f09d587defd54aa24dbb770b3773f29a658b080f95d6cf12aa0c8f6e97cef8", "E48M4KQJ3To=")

	//services.AdminSetRechargeAmountTask()

}
