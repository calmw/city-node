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
	// 城市先锋
	//_ = s.Every(10).Minutes().From(gocron.NextTick()).Do(service.PollBlockTaskGetUserLocationRecordEvent)
	//_ = s.Every(3).Minutes().From(gocron.NextTick()).Do(service.PollBlockTaskGetDailyRewardRecordEvent)
	//_ = s.Every(5).Minutes().From(gocron.NextTick()).Do(service.PollBlockTaskGetRechargeRecordEvent)
	//_ = s.Every(5).Minutes().From(gocron.NextTick()).Do(service.PollBlockTaskGetIncreaseCityDelegateEvent)
	//_ = s.Every(2).Minutes().From(gocron.NextTick()).Do(service.PollBlockTaskGetWithdrawalRewardRecordEvent)
	_ = s.Every(6).Minutes().From(gocron.NextTick()).Do(service.PollBlockGetSuretyRecordEvent)
	//_ = s.Every(30).Minutes().From(gocron.NextTick()).Do(service.GetPioneerRechargeWeight)
	_ = s.Every(2).Hours().From(gocron.NextTick()).Do(blockchain.TriggerAllPioneerTask)

	// 新手任务
	//_ = s.Every(2).Minutes().From(gocron.NextTick()).Do(service.PollBlockTaskGetUserLocationRecordEvent)
	<-s.Start()

}
