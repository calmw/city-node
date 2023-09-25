//package main

package tasks

import (
	"city-node-server/blockchain"
	"city-node-server/db"
	"city-node-server/services"
	"city-node-server/tasks/service"
	"github.com/jasonlvhit/gocron"
	"time"
)

func Start() {
	db.InitMysql()

	services.InitMainNet()

	s := gocron.NewScheduler()
	s.ChangeLoc(time.UTC)
	//_ = s.Every(30).Seconds().From(gocron.NextTick()).Do(service.PollBlockTask)
	_ = s.Every(5).Minutes().From(gocron.NextTick()).Do(service.PollBlockTask)
	_ = s.Every(1).Day().From(gocron.NextTick()).From(gocron.NextTick()).Do(service.GetPioneerRechargeWeight)
	_ = s.Every(4).Hours().From(gocron.NextTick()).From(gocron.NextTick()).Do(blockchain.TriggerAllPioneerTask)
	<-s.Start()

}
