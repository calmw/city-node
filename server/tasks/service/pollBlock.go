package service

import (
	"city-node-server/blockchain"
	"city-node-server/log"
	"context"
)

//const PoolStep = 999

const PoolStep = 1000

func PollBlockTaskGetUserLocationRecordEvent() {
	Cli := blockchain.Client(blockchain.CityNodeConfig)
	err, startBlock := blockchain.GetStartBlock(blockchain.LocationEvent) // 2306974
	if err != nil {
		return
	}
	number, err := Cli.BlockNumber(context.Background())
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	endBlock := startBlock + PoolStep
	if endBlock > number {
		log.Logger.Sugar().Error(err)
		return
	}
	// 用户定位事件处理
	blockchain.GetUserLocationRecordEvent(Cli, int64(startBlock), int64(endBlock))
	//err = blockchain.GetUserLocationRecordEvent(Cli, int64(startBlock), int64(endBlock))
	//if err != nil {
	//	return
	//}
	// 更新区块高度
	blockchain.SetSTartBlock(int64(startBlock+1000), blockchain.LocationEvent)
}

func PollBlockTaskGetDailyRewardRecordEvent() {
	Cli := blockchain.Client(blockchain.CityNodeConfig)
	err, startBlock := blockchain.GetStartBlock(blockchain.RewardEvent) // 2306974
	if err != nil {
		return
	}
	number, err := Cli.BlockNumber(context.Background())
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	endBlock := startBlock + PoolStep
	if endBlock > number {
		log.Logger.Sugar().Error(err)
		return
	}
	err = blockchain.GetDailyRewardRecordEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	// 更新区块高度
	blockchain.SetSTartBlock(int64(startBlock+1000), blockchain.RewardEvent)
}

func PollBlockTaskGetRechargeRecordEvent() {
	Cli := blockchain.Client(blockchain.CityNodeConfig)
	err, startBlock := blockchain.GetStartBlock(blockchain.RechargeEvent) // 2306974
	if err != nil {
		return
	}
	number, err := Cli.BlockNumber(context.Background())
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	endBlock := startBlock + PoolStep
	if endBlock > number {
		log.Logger.Sugar().Error(err)
		return
	}
	// 用户定位事件处理
	err = blockchain.GetRechargeRecordEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	// 更新区块高度
	blockchain.SetSTartBlock(int64(startBlock+1000), blockchain.RechargeEvent)
}

func PollBlockTaskGetIncreaseCityDelegateEvent() {
	Cli := blockchain.Client(blockchain.CityNodeConfig)
	err, startBlock := blockchain.GetStartBlock(blockchain.DelegateEvent) // 2306974
	if err != nil {
		return
	}
	number, err := Cli.BlockNumber(context.Background())
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	endBlock := startBlock + PoolStep
	if endBlock > number {
		log.Logger.Sugar().Error(err)
		return
	}
	// 用户定位事件处理
	err = blockchain.GetIncreaseCityDelegateEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	// 更新区块高度
	blockchain.SetSTartBlock(int64(startBlock+1000), blockchain.DelegateEvent)
}

func PollBlockTaskGetWithdrawalRewardRecordEvent() {
	Cli := blockchain.Client(blockchain.CityNodeConfig)
	err, startBlock := blockchain.GetStartBlock(blockchain.WithdrawEvent) // 2306974
	if err != nil {
		return
	}
	number, err := Cli.BlockNumber(context.Background())
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	endBlock := startBlock + PoolStep
	if endBlock > number {
		log.Logger.Sugar().Error(err)
		return
	}
	// 用户定位事件处理
	err = blockchain.GetWithdrawalRewardRecordEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	// 更新区块高度
	blockchain.SetSTartBlock(int64(startBlock+1000), blockchain.WithdrawEvent)
}
