package main

import (
	"city-node-server/pkg/blockchain"
	"city-node-server/pkg/db"
	"city-node-server/services"
	"github.com/jasonlvhit/gocron"
	"time"
)

func main() {

	services.InitTestNet()
	//services.InitCityPioneer()
	//services.InitUserLocation()
	//services.InitCity()

	db.InitMysql()
}

func clearTestNet() {
	services.InitTestNet()
	services.InitCity()
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
