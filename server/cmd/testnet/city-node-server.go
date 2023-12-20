package main

import (
	"city-node-server/pkg/blockchain"
	"city-node-server/services"
	"github.com/jasonlvhit/gocron"
	"time"
)

func main() {

	services.InitTestNet()
	//services.InitUserLocation()
	//services.InitAppraise()
	//services.InitCityPioneer(20)
	services.InitCity(20)

	//db.InitMysql()

	//clearTestNet()
}

func clearTestNet() {
	services.InitTestNet()
	taskTest()
}
func taskTest() {

	s := gocron.NewScheduler()
	s.ChangeLoc(time.UTC)
	//_ = s.Every(3).Seconds().From(gocron.NextTick()).Do(services.AdminSetDelegateTask)
	_ = s.Every(9).Seconds().From(gocron.NextTick()).Do(blockchain.TriggerAllPioneerTaskTestNet)
	//_ = s.Every(10).Seconds().From(gocron.NextTick()).Do(services.AdminSetRechargeAmountTask)
	<-s.Start() // Start all the pending jobs

}
