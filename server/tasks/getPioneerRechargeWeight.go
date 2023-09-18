package tasks

import (
	"city-node-server/blockchain"
	"city-node-server/db"
	"city-node-server/models"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/status-im/keycard-go/hexutils"
	"gorm.io/gorm"
	"strings"
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
	fmt.Println(today, number, 898978)
	for i := 0; i < int(number); i++ {
		err, pioneerAddress := blockchain.GetPioneer(int64(i))
		if err != nil {
			continue
		}
		fmt.Println("pioneer:", pioneerAddress)
		// 获取先锋对应的城市ID
		err, cityId := blockchain.PioneerChengShi(pioneerAddress)
		if err != nil {
			continue
		}
		// 获取城市位置信息
		err, cityLocation := GetCityInfo("0x" + hexutils.BytesToHex(blockchain.Bytes32ToBytes(cityId)))
		if err != nil {
			continue
		}
		fmt.Println(1212)
		// 获取先锋城市所有的区县ID
		err, countyIds := blockchain.GetCountyIdsByChengShiIdBytes32(cityId)
		if err != nil {
			continue
		}
		for _, countyId := range countyIds {
			// 获取区县位置信息
			err, countyLocation := GetCountyInfo("0x" + hexutils.BytesToHex(blockchain.Bytes32ToBytes(countyId)))
			if err != nil {
				continue
			}
			// 获取区县ID某天的充值权重
			for day := int(today); day > int(today)-20; day-- {
				err, weight := blockchain.RechargeDailyWeightRecord(countyId, int64(day))
				fmt.Println(weight.String(), "---------")
				if err != nil {
					continue
				}
				//E18 := big.NewInt(1e18)
				//newWeight := E18.Div(weight, E18)
				newWeight := decimal.NewFromBigInt(weight, 0)
				err = InsertRechargeWeight(pioneerAddress, "0x"+hexutils.BytesToHex(blockchain.Bytes32ToBytes(cityId)), "0x"+hexutils.BytesToHex(blockchain.Bytes32ToBytes(countyId)), cityLocation, countyLocation,
					int64(day),
					newWeight)
				if err != nil {
					continue
				}
			}

		}
	}
}

func GetCountyInfo(countyId string) (error, string) {
	var userLocation models.UserLocation
	countyId = strings.ToLower(countyId)
	whereCondition := fmt.Sprintf("county_id='%s'", countyId)
	err := db.Mysql.Table("user_location").Where(whereCondition).First(&userLocation).Error
	return err, userLocation.Location
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

func InsertRechargeWeight(pioneer, cityId, countyId, cityLocation, countyLocation string, day int64, weight decimal.Decimal) error {
	var rechargeWeight models.RechargeWeight
	t := time.Unix(day*86400+1, 0)
	whereCondition := fmt.Sprintf("city_id='%s' and day='%s'", cityId, t.Format("2006-01-02 15:04:05"))
	err := db.Mysql.Table("recharge_weight").Where(whereCondition).First(&rechargeWeight).Error

	t.Format("2006-01-02 15:04:05")

	if err == gorm.ErrRecordNotFound {
		db.Mysql.Table("recharge_weight").Create(&models.RechargeWeight{
			Pioneer:        pioneer,
			CountyId:       countyId,
			CityId:         cityId,
			CityLocation:   cityLocation,
			CountyLocation: countyLocation,
			Weight:         weight,
			//Date:     t.Format("2006-01-02 15:04:05"),
			Day: t.Format("2006-01-02 15:04:05"),
		})
	}
	return err
}
