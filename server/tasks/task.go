package tasks

import (
	"city-node-server/blockchain"
	"github.com/jasonlvhit/gocron"
	"time"
)

func Task() {

	s := gocron.NewScheduler()
	s.ChangeLoc(time.UTC)
	_ = s.Every(30).Seconds().From(gocron.NextTick()).Do(blockchain.GetCityDelegateEvent)
	<-s.Start() // Start all the pending jobs

}
