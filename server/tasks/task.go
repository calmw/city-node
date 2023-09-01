package tasks

import (
	"city-node-server/blockchain"
	"github.com/jasonlvhit/gocron"
	"time"
)

func Task() {

	s := gocron.NewScheduler()
	s.ChangeLoc(time.UTC)
	_ = s.Every(1).Day().From(gocron.NextTick()).Do(blockchain.CityPioneerDailyTask)
	<-s.Start() // Start all the pending jobs

}
