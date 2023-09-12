package main

import (
	"city-node-server/blockchain"
	"city-node-server/services"
)

func main() {
	//db.InitMysql()
	//tasks.Task()

	services.InitTestNet()
	//services.InitMainNet()

	//services.InitCity()
	//services.InitCityPioneer()
	//services.InitUserLocation()
	//

	//blockchain.AdminEditSurety("0x3edf72f7ab938a14e6aae3701a1e5acbb2512c840a019c6ead9d01415dbac864", 1, 100000) // 设置城市保证金、等级

	//blockchain.SetUserLocation("0xad02e2bafae66da7b495ff56ccc86f2906814d6f5ab13e6dcd0348f12dc8bf0b", "Vx4fMYGPb3M=") // 设置用户位置

	//
	//blockchain.ApproveToxCityPioneer() // 向城市先锋合约approve
	//blockchain.CityPioneerBalance() // 城市先锋合约余额
	//blockchain.DepositSurety()         // 交保证金，慎用，扣钱
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
	blockchain.SetCityMapping("0x0a23fb256dc6340f9287a2edf424eb5cb92f37f39a8f58cc6c1a9fa349d51b34", "0x85202e610a59bdbc7559cabd41c28f4ecfe18675575572ac5822ca697df5103d", "6WPQrWhP2DQ=")

	//services.AdminSetRechargeAmountTask()

}
