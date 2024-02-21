package services

import (
	"city-node-server/pkg/blockchain"
	"city-node-server/pkg/log"
)

func AddPioneer(areaShiId, pioneerAddress string, level int64, suretyTox int64, suretyUsdt int64, pioneerBatch int64, isChengShi bool) error {
	if !isChengShi {
		level += 20
	}
	err, cli := blockchain.Client(blockchain.CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	// 设置所需缴纳的TOX，城市等级
	city := blockchain.NewCity(cli)
	err = city.AdminSetChengShiLevelAndSurety(areaShiId, level, suretyTox)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}

	// 设置所需缴纳的USDT
	cityPioneerData := blockchain.NewCityPioneerData(cli)
	err = cityPioneerData.SuretyUSDT(areaShiId, suretyUsdt)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}

	err = city.AdminSetPioneer(areaShiId, pioneerAddress, pioneerBatch)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	return nil
}

func RemovePioneer(areaShiId, pioneerAddress string) {
	err, cli := blockchain.Client(blockchain.CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	city := blockchain.NewCity(cli)
	err = city.AdminRemovePioneer(areaShiId, pioneerAddress)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

}
