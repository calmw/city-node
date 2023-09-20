package service

import (
	"city-node-server/blockchain"
	"city-node-server/log"
	"context"
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
	endBlock := startBlock + PoolStep
	if endBlock > number {
		return
	}
	// 用户定位事件处理
	//_ = blockchain.GetUserLocationRecordEvent(Cli, int64(startBlock), int64(endBlock))

	// 城市先锋奖励事件
	_ = blockchain.GetDailyRewardRecordEvent(Cli, int64(startBlock), int64(endBlock))

	// 更新区块高度
	blockchain.SetSTartBlock(int64(startBlock + 1000))
}
