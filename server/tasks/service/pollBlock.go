package service

import (
	"city-node-server/blockchain"
	"city-node-server/log"
	"context"
	"fmt"
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
		return
	}
	fmt.Println(startBlock, endBlock)
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
		return
	}
	fmt.Println(startBlock, endBlock)
	// 用户定位事件处理
	err = blockchain.GetDailyRewardRecordEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
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
		return
	}
	fmt.Println(startBlock, endBlock)
	// 用户定位事件处理
	err = blockchain.GetRechargeRecordEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
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
		return
	}
	fmt.Println(startBlock, endBlock)
	// 用户定位事件处理
	err = blockchain.GetIncreaseCityDelegateEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
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
		return
	}
	fmt.Println(startBlock, endBlock)
	// 用户定位事件处理
	err = blockchain.GetWithdrawalRewardRecordEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		return
	}
	// 更新区块高度
	blockchain.SetSTartBlock(int64(startBlock+1000), blockchain.WithdrawEvent)
}

//func PollBlockTask() {
//	Cli := blockchain.Client(blockchain.CityNodeConfig)
//	err, startBlock := blockchain.GetStartBlock() // 2306974
//	if err != nil {
//		return
//	}
//	number, err := Cli.BlockNumber(context.Background())
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	endBlock := startBlock + PoolStep
//	if endBlock > number {
//		return
//	}
//
//	fmt.Println(startBlock, endBlock)
//	wg := sync.WaitGroup{}
//	wg.Add(5)
//
//	go func() {
//		defer wg.Done()
//		// 用户定位事件处理
//		_ = blockchain.GetUserLocationRecordEvent(Cli, int64(startBlock), int64(endBlock))
//	}()
//	go func() {
//		defer wg.Done()
//		// 城市先锋奖励事件
//		_ = blockchain.GetDailyRewardRecordEvent(Cli, int64(startBlock), int64(endBlock))
//	}()
//	go func() {
//		defer wg.Done()
//		// 增加充值事件
//		_ = blockchain.GetRechargeRecordEvent(Cli, int64(startBlock), int64(endBlock))
//	}()
//	go func() {
//		defer wg.Done()
//		//获取新增质押事件
//		_ = blockchain.GetIncreaseCityDelegateEvent(Cli, int64(startBlock), int64(endBlock))
//	}()
//	go func() {
//		defer wg.Done()
//		//获取奖励领取事件
//		_ = blockchain.GetWithdrawalRewardRecordEvent(Cli, int64(startBlock), int64(endBlock))
//	}()
//	wg.Wait()
//
//	// 更新区块高度
//	blockchain.SetSTartBlock(int64(startBlock + 1000))
//}
