package service

import (
	"city-node-server/blockchain"
	"city-node-server/log"
	"context"
)

//const PoolStep = 999

const PoolStep = 500

func PollBlockTaskGetUserLocationRecordEvent() {
	Cli := blockchain.Client(blockchain.CityNodeConfig)
	err, startBlock := blockchain.GetStartBlock(blockchain.LocationEvent) // 2306974
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	number, err := Cli.BlockNumber(context.Background())
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	endBlock := startBlock + PoolStep
	if endBlock > number {
		endBlock = number
	}
	// 用户定位事件处理
	err = blockchain.GetUserLocationRecordEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	if endBlock >= number {
		log.Logger.Sugar().Error("扫描区块高度>=最新高度", number)
		return
	}
	// 更新区块高度
	blockchain.SetSTartBlock(int64(endBlock), blockchain.LocationEvent)
}

func PollBlockTaskGetUserLocationRecordEvent2() {
	Cli := blockchain.Client(blockchain.CityNodeConfig)
	err, startBlock := blockchain.GetStartBlock(blockchain.LocationEvent2) // 2306974
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	number, err := Cli.BlockNumber(context.Background())
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	endBlock := startBlock + PoolStep
	if endBlock > number {
		endBlock = number
	}
	// 用户定位事件处理
	err = blockchain.GetUserLocationRecordEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	if endBlock >= number {
		log.Logger.Sugar().Error("扫描区块高度>=最新高度", number)
		return
	}
	// 更新区块高度
	blockchain.SetSTartBlock(int64(startBlock+1000), blockchain.LocationEvent2)
}

func PollBlockTaskGetDailyRewardRecordEvent() {
	Cli := blockchain.Client(blockchain.CityNodeConfig)
	err, startBlock := blockchain.GetStartBlock(blockchain.RewardEvent) // 2430927
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	number, err := Cli.BlockNumber(context.Background())
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	endBlock := startBlock + PoolStep
	if endBlock > number {
		endBlock = number
	}
	err = blockchain.GetDailyRewardRecordEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	if endBlock >= number {
		log.Logger.Sugar().Error("扫描区块高度>=最新高度", number)
		return
	}
	// 更新区块高度
	blockchain.SetSTartBlock(int64(endBlock), blockchain.RewardEvent)
}

func PollBlockTaskGetRechargeRecordEvent() {
	Cli := blockchain.Client(blockchain.CityNodeConfig)
	err, startBlock := blockchain.GetStartBlock(blockchain.RechargeEvent) // 2306974
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	number, err := Cli.BlockNumber(context.Background())
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	endBlock := startBlock + PoolStep
	if endBlock > number {
		endBlock = number
	}

	err = blockchain.GetRechargeRecordEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	if endBlock >= number {
		log.Logger.Sugar().Error("扫描区块高度>=最新高度", number)
		return
	}
	// 更新区块高度
	blockchain.SetSTartBlock(int64(endBlock), blockchain.RechargeEvent)
}

func PollBlockTaskGetIncreaseCityDelegateEvent() {
	Cli := blockchain.Client(blockchain.CityNodeConfig)
	err, startBlock := blockchain.GetStartBlock(blockchain.DelegateEvent) // 2306974
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	number, err := Cli.BlockNumber(context.Background())
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	endBlock := startBlock + PoolStep
	if endBlock > number {
		endBlock = number
	}

	err = blockchain.GetIncreaseCityDelegateEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	if endBlock >= number {
		log.Logger.Sugar().Error("扫描区块高度>=最新高度", number)
		return
	}
	// 更新区块高度
	blockchain.SetSTartBlock(int64(endBlock), blockchain.DelegateEvent)
}

func PollBlockTaskGetWithdrawalRewardRecordEvent() {
	Cli := blockchain.Client(blockchain.CityNodeConfig)
	err, startBlock := blockchain.GetStartBlock(blockchain.WithdrawEvent) // 2306974
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	number, err := Cli.BlockNumber(context.Background())
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	endBlock := startBlock + PoolStep
	if endBlock > number {
		endBlock = number
	}

	err = blockchain.GetWithdrawalRewardRecordEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	if endBlock >= number {
		log.Logger.Sugar().Error("扫描区块高度>=最新高度", number)
		return
	}
	// 更新区块高度
	blockchain.SetSTartBlock(int64(endBlock), blockchain.WithdrawEvent)
}

func PollBlockGetSuretyRecordEvent() {
	Cli := blockchain.Client(blockchain.CityNodeConfig)
	err, startBlock := blockchain.GetStartBlock(blockchain.SuretyRecordEvent) // 2306974
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	number, err := Cli.BlockNumber(context.Background())
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	endBlock := startBlock + PoolStep
	if endBlock > number {
		endBlock = number
	}

	err = blockchain.GetSuretyRecordEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	if endBlock >= number {
		log.Logger.Sugar().Error("扫描区块高度>=最新高度", number)
		return
	}
	// 更新区块高度
	blockchain.SetSTartBlock(int64(endBlock), blockchain.SuretyRecordEvent)
}
