package service

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
	endBlock := startBlock + PoolStep
	if endBlock > number {
		return
	}

	fmt.Println(startBlock, endBlock)
	//wg := sync.WaitGroup{}
	//wg.Add(5)

	//go func() {
	//	// 用户定位事件处理
	//	_ = blockchain.GetUserLocationRecordEvent(Cli, int64(startBlock), int64(endBlock))
	//	wg.Done()
	//}()
	//go func() {
	//	// 城市先锋奖励事件
	//	_ = blockchain.GetDailyRewardRecordEvent(Cli, int64(startBlock), int64(endBlock))
	//	wg.Done()
	//}()
	//go func() {
	//	// 增加充值事件
	//	_ = blockchain.GetRechargeRecordEvent(Cli, int64(startBlock), int64(endBlock))
	//	wg.Done()
	//}()
	//go func() {
	// 获取新增质押事件
	//_ = blockchain.GetIncreaseCityDelegateEvent(Cli, int64(startBlock), int64(endBlock))
	//wg.Done()
	//}()
	//go func() {
	// 获取奖励领取事件
	_ = blockchain.GetWithdrawalRewardRecordEvent(Cli, int64(startBlock), int64(endBlock))
	//wg.Done()
	//}()
	//wg.Wait()
	// 用户定位事件处理
	//_ = blockchain.GetUserLocationRecordEvent(Cli, int64(startBlock), int64(endBlock))

	// 城市先锋奖励事件
	//_ = blockchain.GetDailyRewardRecordEvent(Cli, int64(startBlock), int64(endBlock))

	// 增加充值事件
	//_ = blockchain.GetRechargeRecordEvent(Cli, int64(startBlock), int64(endBlock))

	// 获取新增质押事件
	//_ = blockchain.GetIncreaseCityDelegateEvent(Cli, int64(startBlock), int64(endBlock))

	// 更新区块高度
	blockchain.SetSTartBlock(int64(startBlock + 1000))
}
