package main

import (
	"city-node-server/pkg/blockchain"
	"city-node-server/pkg/db"
	"city-node-server/services"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"time"
)

func main() {

	db.InitMysql()
	services.InitTestNet()
	//services.InitUserLocation()
	//services.InitAppraise()
	//services.InitCityPioneer(20)
	//services.InitCity(20)
	//services.InitCityPioneerData()

	//db.InitMysql()

	//taskTest()
	// 四期设置
	//services.AddPioneerBeth4()
	err := services.AddPioneerBatch4(
		"0x3edf72f7ab938a14e6aae3701a1e5acbb2512c840a019c6ead9d01415dbac864",
		"0x365Cae736D93Ad3e388919d0E4d3EE6Ed364b060",
		1,
		10, //5000 6
		5,  // 1000 3
		4,
		1)
	fmt.Println(err, 123456)
	//0x0a23fb256dc6340f9287a2edf424eb5cb92f37f39a8f58cc6c1a9fa349d51b34
	//services.RemovePioneer(
	//	"0x3edf72f7ab938a14e6aae3701a1e5acbb2512c840a019c6ead9d01415dbac864",
	//	"0x365Cae736D93Ad3e388919d0E4d3EE6Ed364b060",
	//)

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
