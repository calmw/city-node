package blockchain

import (
	"city-node-server/log"
	"city-node-server/pkg/binding/intoCityNode"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
	"time"
)

// GetNewlyDelegateRecordByChengId   根据城市ID和天查询新增质押量
func GetNewlyDelegateRecordByChengId(chengShiId string, day int64) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	if strings.Contains(chengShiId, "0x") {
		chengShiId = strings.ReplaceAll(chengShiId, "0x", "")
	}
	common.Hex2Bytes(chengShiId)
	chengShiIdBytes32 := BytesToByte32(common.Hex2Bytes(chengShiId))
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	countyIds, err := userLocation.GetCountyIdsByChengShiId(nil, chengShiIdBytes32)

	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	amountTotal := big.NewInt(0)
	for _, countyId := range countyIds {
		err, amount := CountyNewlyPioneerDelegateRecord(countyId, day)
		if err == nil {
			amountTotal.Add(amountTotal, amount)
		}
	}
	fmt.Println(amountTotal.String())
}

// GetRechargeDailyWeightRecordByChengId   根据区县ID和天查询新增充值权重
func GetRechargeDailyWeightRecordByChengId(chengShiId string, day int64) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	if strings.Contains(chengShiId, "0x") {
		chengShiId = strings.ReplaceAll(chengShiId, "0x", "")
	}
	common.Hex2Bytes(chengShiId)
	chengShiIdBytes32 := BytesToByte32(common.Hex2Bytes(chengShiId))
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	countyIds, err := userLocation.GetCountyIdsByChengShiId(nil, chengShiIdBytes32)

	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	amountTotal := big.NewInt(0)
	for _, countyId := range countyIds {
		err, amount := RechargeDailyWeightRecord(countyId, day)
		if err == nil {
			amountTotal.Add(amountTotal, amount)
		}
		time.Sleep(time.Second * 3)
	}
	fmt.Println(amountTotal.String())
}
