package tasks

import (
	"github.com/jasonlvhit/gocron"
	"time"
)

func Task() {

	s := gocron.NewScheduler()
	s.ChangeLoc(time.UTC)
	_ = s.Every(30).Seconds().From(gocron.NextTick()).Do(PollBlockTask)
	_ = s.Every(1).Day().From(gocron.NextTick()).Do(GetPioneerRechargeWeight)
	//_ = s.Every(20).Seconds().From(gocron.NextTick()).Do(services.AdminSetDelegateTask)
	//_ = s.Every(10).Seconds().From(gocron.NextTick()).Do(services.AdminSetRechargeAmountTask)
	<-s.Start() // Start all the pending jobs

}
