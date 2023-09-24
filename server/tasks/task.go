package tasks

//
//package main
//
//import (
//	"city-node-server/db"
//	"city-node-server/services"
//	"city-node-server/tasks/service"
//	"github.com/jasonlvhit/gocron"
//	"time"
//)
//
//func main() {
//	db.InitMysql()
//
//	services.InitMainNet()
//
//	s := gocron.NewScheduler()
//	s.ChangeLoc(time.UTC)
//	_ = s.Every(10).Seconds().From(gocron.NextTick()).Do(service.PollBlockTask)
//	//_ = s.Every(5).Minutes().From(gocron.NextTick()).Do(service.PollBlockTask)
//	//_ = s.Every(1).Day().From(gocron.NextTick()).Do(service.GetPioneerRechargeWeight)
//	<-s.Start()
//
//}
