package main

import (
	"city-node-server/pkg/db"
	"city-node-server/services"
)

func main() {

	db.InitMysql()
	services.InitMainNet()
	//services.InitUserLocation()
	//services.InitCityPioneer(86400)
	//services.InitCity(86400)
	//services.InitAppraise()
	//services.InitCityPioneerData()

	//services.AddPioneerBeth3() // 四期上线前可用，需要更新ABI
	//services.AddPioneerBeth4()

	// 获取一二期用户最近三个月重置权重详情
	//pioneers := []string{
	//	"0xE137fF4FCdDA90C3665562F52491B511155e19FF",
	//	"0x7295f5a83337163aA06c70dcdd51903aCc5E27fa",
	//	"0x90E9aAD6EE78F79058A9d7aA1655205465bef695",
	//	"0x342AcB8027b227fB4104ff35fE9cc373d103eC92",
	//	"0x21DCe990d025e909049E7966e42850491CB42afd",
	//	"0x7D574395807fd442F199Baa05166eC267b7776f8",
	//}
	//blockchain.RechargeWeightRecordByChengId(pioneers)
	//services.ReadExcel("./副本城市节点报名表11.xlsx")
	//services.ReadExcel("./assets/城市节点报名表合肥.xlsx")
	//services.ReadExcel("./assets/区县节点报名表2.xlsx")

	/// 添加先锋
	//blockchain.GetAllPioneer()
	//services.SyncAddPioneerInfoFromExcel("./assets/区县节点报名表2.xlsx")
	//services.AddPioneerBeth4FromDb()
	/// 添加先锋
	services.AddPioneerBeth4FromDb2()

	//services.ReadExcel5("./assets/副本INTO工作室申请统计表(审核12月31日)发给技术.xlsx") // 查询用户所在城市的网体业绩
	//services.ReadExcel5("./assets/INTO工作室补贴业绩查询1.31.xlsx") // 查询用户所在城市的网体业绩
	//services.ReadCityFIle("./assets/HaNoi.txt")
	//services.ReadCityFIle("./assets/Bangkok.txt")
	//services.CheckPioneer("./assets/城市先锋-用户信息.xlsx") // 确认用户是否交保证金
	//services.CheckPioneer4("./assets/城市先锋-用户信息.xlsx") // 确认用户是否交保证金，批次，绑定城市用户数量
	//services.CheckPioneer3("./assets/城市先锋-用户信息.xlsx") // 设置先锋批次
	//services.CheckPioneer2("./assets/城市先锋-用户信息.xlsx", "./assets/副本城市节点汇总11.26.2.xlsx") // 确认用户是否交保证金
	//services.CheckLocation("./assets/越南2.xlsx") // 查看位置是否存在,把县城映射到省（越南等国家）

	//err := services.RemovePioneer(
	//	"0xf5dfc352bad4aad6e9a31a8cc0cd00ff6bde8a4cf8c3302dbf715cdb2e728e40",
	//	"0x04AC71d61Ee8A74e522D35085CA117C9c9F4d933",
	//)
	//fmt.Println(err)

	//err := services.AddPioneerBatch4(
	//	"0x4de9096348869d087c953a1b6df5267276b15929bf8c3eedd52332bb946dadda",
	//	"0x62C6490dE592F439092eac4567f74e1622c9C9C9",
	//	2,
	//	3000, //5000 6
	//	600,  // 1000 3
	//	4,
	//	1)
	//fmt.Println(err)

	//db.InitFdb()
	//db.InitLevelDb()
	//go services.InitSyncTask()
	////services.SyncUserLocation()
	//services.SyncStatus("assets/区县节点报名表2.xlsx")
	//services.GetRaceNodeWeight("assets/区县节点报名表2.xlsx")

}
