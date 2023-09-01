package tasks

import (
	"city-node-server/blockchain"
	"city-node-server/blockchain/event"
	"github.com/jasonlvhit/gocron"
	"time"
)

func Task() {

	s := gocron.NewScheduler()
	s.ChangeLoc(time.UTC)
	_ = s.Every(5).Seconds().From(gocron.NextTick()).Do(blockchain.GetLogs, event.SaleEvent)
	<-s.Start() // Start all the pending jobs

}
