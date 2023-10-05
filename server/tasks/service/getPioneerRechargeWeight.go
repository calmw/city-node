package service

import (
	"city-node-server/blockchain"
	"city-node-server/db"
	"city-node-server/log"
	"city-node-server/models"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/status-im/keycard-go/hexutils"
	"gorm.io/gorm"
	"strings"
	"sync"
	"time"
)

func GetPioneerRechargeWeight() {
	err, number := blockchain.GetPioneerNumber()
	if err != nil {
		return
	}
	err, today := blockchain.GetDay()
	if err != nil {
		return
	}
	yesterday := today - 1
	for i := 0; i < int(number); i++ {
		err, pioneerAddress := blockchain.GetPioneer(int64(i))
		if err != nil {
			continue
		}
		// 获取先锋对应的城市ID
		err, cityId := blockchain.PioneerChengShi(pioneerAddress)
		if err != nil {
			continue
		}
		// 获取城市位置信息
		_, cityLocation := GetCityInfo("0x" + hexutils.BytesToHex(blockchain.Bytes32ToBytes(cityId)))
		// 获取先锋城市所有的区县ID
		err, countyIds := blockchain.GetCountyIdsByChengShiIdBytes32(cityId)
		if err != nil {
			continue
		}

		for _, countyId := range countyIds {
			// 获取区县位置信息
			//_, countyLocation := blockchain.CityInfo(countyId)
			_, countyLocation := GetCountyInfo("0x" + hexutils.BytesToHex(blockchain.Bytes32ToBytes(countyId)))
			if countyLocation == "" { //
				continue
			}
			// 获取区县ID某天的充值权重
			for day := int(yesterday); day > int(yesterday)-12; day-- {
				err, weight := blockchain.RechargeDailyWeightRecord(countyId, int64(day))
				if err != nil {
					continue
				}
				newWeight := decimal.NewFromBigInt(weight, 0)
				lock := sync.RWMutex{}
				lock.Lock()
				defer lock.Unlock()
				_ = InsertRechargeWeight(pioneerAddress, "0x"+hexutils.BytesToHex(blockchain.Bytes32ToBytes(cityId)), "0x"+hexutils.BytesToHex(blockchain.Bytes32ToBytes(countyId)), cityLocation, countyLocation,
					int64(day),
					newWeight)
			}
		}
	}
}

func GetCityInfo(cityId string) (error, string) {
	var userLocation models.UserLocation
	cityId = strings.ToLower(cityId)
	whereCondition := fmt.Sprintf("city_id='%s'", cityId)
	err := db.Mysql.Table("user_location").Where(whereCondition).First(&userLocation).Error
	var location string
	if err == nil {
		l := strings.Split(userLocation.Location, " ")
		if len(l) >= 2 {
			location = l[0] + "-" + l[1]
		}
	}
	return err, location
}

func GetCountyInfo(countyId string) (error, string) {
	var userLocation models.UserLocation
	countyId = strings.ToLower(countyId)
	whereCondition := fmt.Sprintf("county_id='%s'", countyId)
	err := db.Mysql.Table("user_location").Where(whereCondition).First(&userLocation).Error
	var location string
	if err == nil {
		l := strings.Split(userLocation.Location, " ")
		if len(l) >= 2 {
			location = l[0] + "-" + l[1] + "-" + l[2]
		}
	}
	return err, location
}

func InsertRechargeWeight(pioneer, cityId, countyId, cityLocation, countyLocation string, day int64, weight decimal.Decimal) error {
	blockchain.InsertRechargeWeightLock.Lock()
	defer blockchain.InsertRechargeWeightLock.Unlock()
	var rechargeWeight models.RechargeWeight
	t := time.Unix(day*86400+1, 0)
	whereCondition := fmt.Sprintf("city_id='%s' and county_id='%s' and day='%s' and pioneer='%s'",
		strings.ToLower(cityId), strings.ToLower(countyId), t.Format("2006-01-02 15:04:05"), strings.ToLower(pioneer))
	err := db.Mysql.Table("recharge_weight").Where(whereCondition).First(&rechargeWeight).Error
	if err == gorm.ErrRecordNotFound {
		db.Mysql.Table("recharge_weight").Create(&models.RechargeWeight{
			Pioneer:        strings.ToLower(pioneer),
			CountyId:       strings.ToLower(countyId),
			CityId:         strings.ToLower(cityId),
			CityLocation:   strings.ToLower(cityLocation),
			CountyLocation: strings.ToLower(countyLocation),
			Weight:         weight,
			Day:            t.Format("2006-01-02 15:04:05"),
		})
	}
	return err
}

func SavePioneerInDb(pioneerAddress, cityId string) error {
	var pioneer models.Pioneer
	whereCondition := fmt.Sprintf("city_id='%s' and pioneer='%s'",
		strings.ToLower(cityId), strings.ToLower(pioneerAddress))
	err := db.Mysql.Model(&models.Pioneer{}).Where(whereCondition).First(&pioneer).Error
	if err == gorm.ErrRecordNotFound {
		db.Mysql.Model(&models.Pioneer{}).Create(&models.Pioneer{
			Pioneer: strings.ToLower(pioneerAddress),
			CityId:  strings.ToLower(cityId),
		})
	}
	return err
}

func UpdatePioneer() {
	err, number := blockchain.GetPioneerNumber()
	if err != nil {
		return
	}
	for i := 0; i < int(number); i++ {
		err, pioneerAddress := blockchain.GetPioneer(int64(i))
		if err != nil {
			log.Logger.Sugar().Error(err)
			continue
		}
		// 获取先锋对应的城市ID
		err, cityId := blockchain.PioneerChengShi(pioneerAddress)
		if err != nil {
			log.Logger.Sugar().Error(err)
			continue
		}
		_ = SavePioneerInDb(pioneerAddress, "0x"+hexutils.BytesToHex(blockchain.Bytes32ToBytes(cityId)))
	}
}
