package tasks

import (
	"city-node-server/services"
	"github.com/jasonlvhit/gocron"
	"time"
)

func Task() {

	s := gocron.NewScheduler()
	s.ChangeLoc(time.UTC)
	//_ = s.Every(30).Seconds().From(gocron.NextTick()).Do(blockchain.GetCityDelegateEvent)
	_ = s.Every(50).Seconds().From(gocron.NextTick()).Do(services.AdminSetDelegate)
	<-s.Start() // Start all the pending jobs

}
