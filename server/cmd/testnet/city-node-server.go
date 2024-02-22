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
	services.InitTestNet()
	//services.InitUserLocation()
	//services.InitAppraise()
	//services.InitCityPioneer(20)
	services.InitCity(20)
	//services.InitCityPioneerData()

	//db.InitMysql()

	//taskTest()
	/// 四期设置
	//err := services.AddPioneer("0x97033fc90d09fa7e17279c173eb0a62ab7d4dc04031195bfcf87e37b58e3b407", "0xeb09570B7841b2280F6340728a092Fde216dBbca", 1,
	//	5000,
	//	1000,
	//	4,
	//	false)
	//fmt.Println(err, 123456)

}

func taskTest() {
	services.InitTestNet()

	s := gocron.NewScheduler()
	s.ChangeLoc(time.UTC)
	//_ = s.Every(3).Seconds().From(gocron.NextTick()).Do(services.AdminSetDelegateTask)
	_ = s.Every(25).Seconds().From(gocron.NextTick()).Do(blockchain.TriggerAllPioneerTaskTestNet)
	//_ = s.Every(10).Seconds().From(gocron.NextTick()).Do(services.AdminSetRechargeAmountTask)
	<-s.Start() // Start all the pending jobs
}
