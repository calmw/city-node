package main

import (
	"city-node-server/services"
)

func main() {
	//db.InitMysql()

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
	//blockchain.SetCityMapping("0x180c563a51cbf2bde90bdd4fd506a5cb5b0b474c14164d37d5bb96dfe43f0024", "0x8549d0e3957c74e498e26c5207d83edbfcf4f2df01e4105964f334c752e1a51a", "gkJ6lFWahdQ=")

	//services.AdminSetRechargeAmountTask()

	// 根据城市ID查询往日新增质押
	//blockchain.GetNewlyDelegateRecordByChengId("0x1d78790980b5a4917a2dfa1a4016ad10af0de8987f7a310acbc9baf85cad17f0", 19613) // 183430 000000000000000000
	//blockchain.GetNewlyDelegateRecordByChengId("0xd6b40febaa0ddd0e1e2674e83d08653b02ee51bef854cb11fe76ee91c1e9cfee", 19613) // 655770 000000000000000000

}
