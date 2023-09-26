package service

import (
	"city-node-server/blockchain"
	"city-node-server/log"
	"context"
	"fmt"
	"sync"
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

	fmt.Println(startBlock, endBlock)
	wg := sync.WaitGroup{}
	wg.Add(5)

	go func() {
		defer wg.Done()
		// 用户定位事件处理
		_ = blockchain.GetUserLocationRecordEvent(Cli, int64(startBlock), int64(endBlock))
	}()
	go func() {
		defer wg.Done()
		// 城市先锋奖励事件
		_ = blockchain.GetDailyRewardRecordEvent(Cli, int64(startBlock), int64(endBlock))
	}()
	go func() {
		defer wg.Done()
		// 增加充值事件
		_ = blockchain.GetRechargeRecordEvent(Cli, int64(startBlock), int64(endBlock))
	}()
	go func() {
		defer wg.Done()
		//获取新增质押事件
		_ = blockchain.GetIncreaseCityDelegateEvent(Cli, int64(startBlock), int64(endBlock))
	}()
	go func() {
		defer wg.Done()
		//获取奖励领取事件
		_ = blockchain.GetWithdrawalRewardRecordEvent(Cli, int64(startBlock), int64(endBlock))
	}()
	wg.Wait()

	// 更新区块高度
	blockchain.SetSTartBlock(int64(startBlock + 1000))
}

func PollBlockTaskTest() {
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

	fmt.Println(startBlock, endBlock)
	// 用户定位事件处理
	_ = blockchain.GetUserLocationRecordEvent(Cli, int64(startBlock), int64(endBlock))
	//wg := sync.WaitGroup{}
	//wg.Add(2)
	//
	//go func() {
	//	defer wg.Done()
	//	// 用户定位事件处理
	//	_ = blockchain.GetUserLocationRecordEvent(Cli, int64(startBlock), int64(endBlock))
	//}()
	//go func() {
	//	defer wg.Done()
	//	// 城市先锋奖励事件
	//	_ = blockchain.GetDailyRewardRecordEvent(Cli, int64(startBlock), int64(endBlock))
	//}()
	////go func() {
	////	defer wg.Done()
	////	// 增加充值事件
	////	_ = blockchain.GetRechargeRecordEvent(Cli, int64(startBlock), int64(endBlock))
	////}()
	////go func() {
	////	defer wg.Done()
	////	//获取新增质押事件
	////	_ = blockchain.GetIncreaseCityDelegateEvent(Cli, int64(startBlock), int64(endBlock))
	////}()
	////go func() {
	////	defer wg.Done()
	////	//获取奖励领取事件
	////	_ = blockchain.GetWithdrawalRewardRecordEvent(Cli, int64(startBlock), int64(endBlock))
	////}()
	//wg.Wait()

	// 更新区块高度
	blockchain.SetSTartBlock(int64(startBlock + 1000))
}
