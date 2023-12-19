package main

import (
	"city-node-server/pkg/blockchain"
	"city-node-server/pkg/db"
	"city-node-server/services"
	"github.com/jasonlvhit/gocron"
	"time"
)

func main() {

	db.InitMysql()
	services.InitMainNet()
	//services.InitUserLocation()
	services.InitCityPioneer(86400)
	services.InitCity(86400)
	services.InitAppraise()

	//services.InitTestNet()

	//blockchain.TriggerAllPioneerTask()

	//time.Sleep(time.Second * 120)
	//services.AdminSetRechargeAmountTask2500()
	//time.Sleep(time.Second * 60 * 7)
	//services.AdminSetRechargeAmountTask5000()
	//time.Sleep(time.Second * 60 * 50)

	//blockchain.AdminEditSurety("0x3edf72f7ab938a14e6aae3701a1e5acbb2512c840a019c6ead9d01415dbac864", 1, 100000) // 设置城市保证金、等级

	//blockchain.SetUserLocation("0xad02e2bafae66da7b495ff56ccc86f2906814d6f5ab13e6dcd0348f12dc8bf0b", "Vx4fMYGPb3M=") // 设置用户位置

	//blockchain.ApproveToxCityPioneer() // 向城市先锋合约approve
	//blockchain.ApproveTox("0x35b821D00e9733Eb0F51195b173EA0AF2ac81736")
	//blockchain.CityPioneerBalance() // 城市先锋合约余额
	//blockchain.DepositSurety() // 交保证金，慎用，扣钱
	//blockchain.DepositSuretyTest("0x08a01BE67fF47Ba2652b7dCE2005B47D81bAaC13") // DepositSuretyTest 交保证金成为先锋

	// 管理员设置开始考核时间,先交保证金，后考核
	//blockchain.AdminSetStartTime(time.Now().Add(time.Hour * 1000000).Unix())
	//blockchain.AdminSetStartTime(time.Now().Unix())
	// 批量执行增加或减少质押量
	//task()

	// 获取前15天社交基金平均值
	//blockchain.GetFifteenDayAverageFounds()

	//blockchain.GetCityId(22)

	// 设置城市ID位置,设置映射关系
	//blockchain.SetCityMapping("0xc56f998ce0b465aec56076b1e61edc5f275aa96c2f6d7190d9d2e597ece7454d", "0xca98b991031d3b83e27132a5373c2a8b7d68c8e6477cb6ab43baf2268b3e9639", "3ALUNlzFlwETLRJDC/7mqg==")

	//services.AdminSetRechargeAmountTask()

	// 根据城市ID查询往日新增质押
	//blockchain.GetNewlyDelegateRecordByChengId("0x1d78790980b5a4917a2dfa1a4016ad10af0de8987f7a310acbc9baf85cad17f0", 19613) // 183430 000000000000000000

	// 根据区县ID和天查询新增充值权重
	//blockchain.GetRechargeDailyWeightRecordByChengId("0x936557454f464570813fe0c6e164d114cfc2d49eb5478f4e9867805eb4d3667b", 19616) // 19616 000000000000000000

	//fmt.Println(utils.ThreeDesEncrypt("0,126"))
	//fmt.Println(utils.ThreeDesDecrypt("7gZ4qrL761s="))
	//fmt.Println(utils.Keccak256("0,276"))
	//fmt.Println(strings.ToLower("0x17dC6411D638672A073f23267C4735ca877AA623"))
	//fmt.Println(utils.Keccak256("0,136,621122")) //0x5BC49BF2BB0F74F6B3321451F04FE277696611BC6A0A2D144F034B27C914F8F0

	//tasks.GetPioneerRechargeWeight()

	// init cache
	//go utils.InitCache()
	//err := services2.GetUserSons("0x74E552c1591A9Bc767057B270090Fd88b82Ea0a1")
	//fmt.Println(123456, err)
	//services2.GetToxTxBridgeBsc()
	//for true {
	//
	//}
	//services2.GetToxTxBridgeBsc()
	//api.Start()
	//tasks.Start()
	//blockchain.GetAllPioneer()
	//blockchain.TriggerAllPioneerTask()
	//ticker := time.NewTicker(30 * time.Minute)
	//for {
	//	<-ticker.C
	//	blockchain.TriggerAllPioneerTask()
	//}
	//services.ReadExcel("./副本城市节点报名表11.xlsx")
	//services.ReadExcel("./城市节点报名表合肥.xlsx")
	//services.CheckPioneer("./城市先锋-用户信息.xlsx") // 确认用户是否交保证金
	//services.CheckLocation("./泰国.xlsx") // 查看位置是否存在
	//重新获取并存储用户位置信息，从链上查询，不是从事件获取
	//blockchain.ReSaveUserLocation()
	//err := blockchain.RestoreUserLocation("0xd7922692c157ee415facfe700e7a3e616f7b12c8")
	//fmt.Println(err)

	// 更新先锋信息
	//service.UpdatePioneer()

	//services.UserCityChange()
	//clearTestNet()
	//services.CheckData()
	// 获取先锋合约TOX的交易
	//services.GetToxTx()
}

func clearTestNet() {
	services.InitTestNet()
	services.InitCity(86400)
	//taskTest()
}
func taskTest() {

	s := gocron.NewScheduler()
	s.ChangeLoc(time.UTC)
	//_ = s.Every(3).Seconds().From(gocron.NextTick()).Do(services.AdminSetDelegateTask)
	_ = s.Every(6).Seconds().From(gocron.NextTick()).Do(blockchain.TriggerAllPioneerTaskTestNet)
	//_ = s.Every(10).Seconds().From(gocron.NextTick()).Do(services.AdminSetRechargeAmountTask)
	<-s.Start() // Start all the pending jobs

}
