package main

import (
	"city-node-server/api"
	"city-node-server/blockchain"
	"city-node-server/services"
	"city-node-server/tasks"
	"github.com/jasonlvhit/gocron"
	"time"
)

func main() {

	//services.InitTestNet()
	//services.InitCity()
	//task()

	//db.InitMysql()
	//blockchain.TriggerAllPioneerTask()

	//services.InitMainNet()
	//time.Sleep(time.Second * 120)
	//services.AdminSetRechargeAmountTask2500()
	//time.Sleep(time.Second * 60 * 7)
	//services.AdminSetRechargeAmountTask5000()
	//time.Sleep(time.Second * 60 * 50)

	//services.InitCityPioneer()
	//services.InitUserLocation()
	//

	//blockchain.AdminEditSurety("0x3edf72f7ab938a14e6aae3701a1e5acbb2512c840a019c6ead9d01415dbac864", 1, 100000) // 设置城市保证金、等级

	//blockchain.SetUserLocation("0xad02e2bafae66da7b495ff56ccc86f2906814d6f5ab13e6dcd0348f12dc8bf0b", "Vx4fMYGPb3M=") // 设置用户位置

	//
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
	//blockchain.SetCityMapping("0xc0cbf0aab88509841c40d83825fdd91fdaf6884ab253d63aa559979ecb633e95", "0xf0d851e51c9ebb39dcecd7685ab590d6f69ff2b3c32b225e96ac51928e57d957", "ihooqau20JSZxfjB0gZCEQ==")

	//services.AdminSetRechargeAmountTask()

	// 根据城市ID查询往日新增质押
	//blockchain.GetNewlyDelegateRecordByChengId("0x1d78790980b5a4917a2dfa1a4016ad10af0de8987f7a310acbc9baf85cad17f0", 19613) // 183430 000000000000000000
	//blockchain.GetNewlyDelegateRecordByChengId("0x64C970619EDC45A6484AAD8ACC9FEC4153DFC65ED6F061FE2458EB13755C1EE1", 19615) // 19616 000000000000000000

	// 根据区县ID和天查询新增充值权重
	//blockchain.GetRechargeDailyWeightRecordByChengId("0x936557454f464570813fe0c6e164d114cfc2d49eb5478f4e9867805eb4d3667b", 19616) // 19616 000000000000000000

	//fmt.Println(utils.ThreeDesEncrypt("0,126"))
	//fmt.Println(utils.ThreeDesDecrypt("j5BSvQJvDEQ="))
	//fmt.Println(utils.Keccak256("0,136"))
	//fmt.Println(utils.Keccak256("0,136,621122")) //0x5BC49BF2BB0F74F6B3321451F04FE277696611BC6A0A2D144F034B27C914F8F0
	//0x5BC49BF2BB0F74F6B3321451F04FE277696611BC6A0A2D144F034B27C914F8F0
	//0xF5C444D5602DE5D09258A569DF33128E716AD2AC84A5AE2B2CC90A4D06B7AD33
	//0xF5C444D5602DE5D09258A569DF33128E716AD2AC84A5AE2B2CC90A4D06B7AD33

	//tasks.GetPioneerRechargeWeight()
	//services.IntoBind()

	go func() {
		defer func() {
			recover()
		}()
		api.Start()
	}()
	tasks.Start()

	//clearTestNet()

	//blockchain.TriggerAllPioneerTask()

}

func clearTestNet() {
	services.InitTestNet()
	services.InitCity()
	task()
}
func task() {

	s := gocron.NewScheduler()
	s.ChangeLoc(time.UTC)
	//_ = s.Every(3).Seconds().From(gocron.NextTick()).Do(services.AdminSetDelegateTask)
	_ = s.Every(4).Seconds().From(gocron.NextTick()).Do(blockchain.TriggerAllPioneerTaskTestNet)
	//_ = s.Every(10).Seconds().From(gocron.NextTick()).Do(services.AdminSetRechargeAmountTask)
	<-s.Start() // Start all the pending jobs

}
