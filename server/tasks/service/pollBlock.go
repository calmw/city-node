package service

import (
	"city-node-server/log"
	blockchain2 "city-node-server/pkg/blockchain"
	"context"
	"golang.org/x/exp/rand"
)

// const PoolStep = 999
const PoolStep = 500

func PollBlockTaskGetUserLocationRecordEvent() {
	err, Cli := blockchain2.Client(blockchain2.CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	err, startBlock := blockchain2.GetStartBlock(blockchain2.LocationEvent) // 2306974
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
	err = blockchain2.GetUserLocationRecordEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	if endBlock >= number {
		log.Logger.Sugar().Error("扫描区块高度>=最新高度", number)
		return
	}
	// 更新区块高度
	blockchain2.SetSTartBlock(int64(endBlock), blockchain2.LocationEvent)
}

func PollBlockTaskGetUserLocationRecordEvent2() {
	err, Cli := blockchain2.Client(blockchain2.CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	err, startBlock := blockchain2.GetStartBlock(blockchain2.LocationEvent2) // 2306974
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
	err = blockchain2.GetUserLocationRecordEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	if endBlock >= number {
		log.Logger.Sugar().Error("扫描区块高度>=最新高度", number)
		return
	}
	// 更新区块高度
	blockchain2.SetSTartBlock(int64(startBlock+1000), blockchain2.LocationEvent2)
}

func PollBlockTaskGetDailyRewardRecordEvent() {
	err, Cli := blockchain2.Client(blockchain2.CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	err, startBlock := blockchain2.GetStartBlock(blockchain2.RewardEvent) // 2430927
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
	err = blockchain2.GetDailyRewardRecordEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	if endBlock >= number {
		log.Logger.Sugar().Error("扫描区块高度>=最新高度", number)
		return
	}
	// 更新区块高度
	blockchain2.SetSTartBlock(int64(endBlock), blockchain2.RewardEvent)
}

func PollBlockTaskGetRechargeRecordEvent() {
	err, Cli := blockchain2.Client(blockchain2.CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	err, startBlock := blockchain2.GetStartBlock(blockchain2.RechargeEvent) // 2306974
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

	err = blockchain2.GetRechargeRecordEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	if endBlock >= number {
		log.Logger.Sugar().Error("扫描区块高度>=最新高度", number)
		return
	}
	// 更新区块高度
	blockchain2.SetSTartBlock(int64(endBlock), blockchain2.RechargeEvent)
}

func PollBlockTaskGetIncreaseCityDelegateEvent() {
	err, Cli := blockchain2.Client(blockchain2.CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	err, startBlock := blockchain2.GetStartBlock(blockchain2.DelegateEvent) // 2306974
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

	err = blockchain2.GetIncreaseCityDelegateEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	if endBlock >= number {
		log.Logger.Sugar().Error("扫描区块高度>=最新高度", number)
		return
	}
	// 更新区块高度
	blockchain2.SetSTartBlock(int64(endBlock), blockchain2.DelegateEvent)
}

//func PollBlockTaskGetWithdrawalRewardRecordEvent() {
//	err, Cli := blockchain.Client(blockchain.CityNodeConfig)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	err, startBlock := blockchain.GetStartBlock(blockchain.WithdrawEvent) // 2306974
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	number, err := Cli.BlockNumber(context.Background())
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	endBlock := startBlock + PoolStep
//	if endBlock > number {
//		endBlock = number
//	}
//
//	err = blockchain.GetWithdrawalRewardRecordEvent(Cli, int64(startBlock), int64(endBlock))
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//
//	if endBlock >= number {
//		log.Logger.Sugar().Error("扫描区块高度>=最新高度", number)
//		return
//	}
//	// 更新区块高度
//	blockchain.SetSTartBlock(int64(endBlock), blockchain.WithdrawEvent)
//}

func PollBlockGetSuretyRecordEvent() {
	err, Cli := blockchain2.Client(blockchain2.CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	err, startBlock := blockchain2.GetStartBlock(blockchain2.SuretyRecordEvent) // 2306974
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

	err = blockchain2.GetSuretyRecordEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	if endBlock >= number {
		log.Logger.Sugar().Error("扫描区块高度>=最新高度", number)
		return
	}
	// 更新区块高度
	blockchain2.SetSTartBlock(int64(endBlock), blockchain2.SuretyRecordEvent)
}

func ChangeRpc() {
	index := rand.Int() % 4
	blockchain2.CityNodeConfig.RPC = blockchain2.RpcArr[index]
}
