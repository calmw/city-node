package tasks

import (
	"city-node-server/blockchain"
	"city-node-server/log"
	"context"
	"fmt"
)

const PoolStep = 999

func PollBlockTask() {
	Cli := blockchain.Client(blockchain.CityNodeConfig)
	startBlock := blockchain.GetStartBlock() // 2306974
	number, err := Cli.BlockNumber(context.Background())
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(number, startBlock)
	endBlock := startBlock + PoolStep
	if endBlock > number {
		return
	}
	// 用户定位事件处理
	err = blockchain.GetUserLocationRecordEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	//
	blockchain.SetSTartBlock(int64(startBlock + 1000))
}
