package services

import (
	"city-node-server/pkg/blockchain"
	"city-node-server/pkg/log"
)

func InitAppraise() {
	err, cli := blockchain.Client(blockchain.CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	appraise := blockchain.NewAppraise(cli)
	// AdminSetStar 管理员设置Star合约地址
	//appraise.AdminSetStar()
	//// AdminSetCity 管理员设置城市合约
	//appraise.AdminSetCity()
	//// AdminSetPioneer 管理员设置先锋合约
	//appraise.AdminSetPioneer()
	//// AddAdmin 添加管理权限
	//appraise.AddAdmin()
	//// AdminSetPioneerBatch 管理员设置先锋批次
	////appraise.AdminSetPioneerBatch("", 3)
	//// AdminSetWeightByCityLevel 管理员设置第三批考核标准
	appraise.AdminSetWeightByCityLevel()
	// 设置四期区域和城市节点的考核标准，三期的在weightByCityLevel中存储
	{
		// 添加城市考核标准1，2，3线城市
		for j := 1; j < 7; j++ {
			cityWeight := 5000
			for i := 1; i < 4; i++ {
				for {
					// pioneerBatch_, isArea_, areaLevel_, month_, weight_
					err = appraise.AdminSetWeightByAreaLevel(4, 0, int64(i), int64(j), int64(cityWeight))
					if err == nil {
						break
					} else {
						log.Logger.Sugar().Error(err)
					}
				}
				cityWeight = cityWeight / 2
			}
		}

		// 添加城市考核标准1，2线区域节点
		for j := 1; j < 7; j++ {
			areaWeight := 1000
			for i := 1; i < 3; i++ {
				for {
					// pioneerBatch_, isArea_, areaLevel_, month_, weight_
					err = appraise.AdminSetWeightByAreaLevel(4, 1, int64(i), int64(j), int64(areaWeight))
					if err == nil {
						break
					} else {
						log.Logger.Sugar().Error(err)
					}
				}
				areaWeight -= 400
			}
		}

	}

}
