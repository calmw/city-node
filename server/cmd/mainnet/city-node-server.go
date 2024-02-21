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
	services.InitCity(86400)
	//services.InitAppraise()

	services.AddPioneerBeth3()

	//services.ReadExcel("./副本城市节点报名表11.xlsx")
	//services.ReadExcel("./assets/城市节点报名表合肥.xlsx")
	//services.ReadExcel5("./assets/副本INTO工作室申请统计表(审核12月31日)发给技术.xlsx") // 查询用户所在城市的网体业绩
	//services.ReadExcel5("./assets/INTO工作室补贴业绩查询1.31.xlsx") // 查询用户所在城市的网体业绩
	//services.ReadCityFIle("./assets/HaNoi.txt")
	//services.ReadCityFIle("./assets/Bangkok.txt")
	//services.CheckPioneer("./assets/城市先锋-用户信息.xlsx") // 确认用户是否交保证金
	//services.CheckPioneer4("./assets/城市先锋-用户信息.xlsx") // 确认用户是否交保证金，批次，绑定城市用户数量
	//services.CheckPioneer3("./assets/城市先锋-用户信息.xlsx") // 设置先锋批次
	//services.CheckPioneer2("./assets/城市先锋-用户信息.xlsx", "./assets/副本城市节点汇总11.26.2.xlsx") // 确认用户是否交保证金
	//services.CheckLocation("./assets/越南2.xlsx") // 查看位置是否存在,把县城映射到省（越南等国家）

}
