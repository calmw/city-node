package blockchain

import (
	"city-node-server/pkg/binding/intoCityNode"
	"city-node-server/pkg/db"
	"city-node-server/pkg/log"
	"city-node-server/pkg/models"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xjieinfo/xjgo/xjcore/xjexcel"
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

// GetRechargeDailyWeightRecordByChengId   根据城市ID和天查询新增充值权重
func GetRechargeDailyWeightRecordByChengId(chengShiId string, day int64) error {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	if strings.Contains(chengShiId, "0x") {
		chengShiId = strings.ReplaceAll(chengShiId, "0x", "")
	}
	//common.Hex2Bytes(chengShiId)
	//chengShiIdBytes32 := BytesToByte32(common.Hex2Bytes(chengShiId))
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	countyIds, err := userLocation.GetCountyIdsByChengShiId(nil, ConvertAreaIdAtB(chengShiId))

	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
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
	return nil
}

type PioneerRecharge struct {
	Pioneer   string `json:"pioneer" excel:"column:B;desc:先锋地址;width:30"`
	City      string `json:"city" excel:"column:C;desc:城市;width:30"`
	Weight    string `json:"weight" excel:"column:D;desc:总充值权重;width:30"`
	StartDate string `json:"start_date" excel:"column:E;desc:开始日期;width:30"`
}

type CountyRecharge struct {
	Pioneer   string `json:"pioneer" excel:"column:B;desc:所属先锋;width:30"`
	City      string `json:"city" excel:"column:C;desc:城市;width:30"`
	Weight    string `json:"weight" excel:"column:D;desc:总充值权重;width:30"`
	StartDate string `json:"start_date" excel:"column:E;desc:日期;width:30"`
}

// RechargeWeightRecordByChengId 生成最近三个月的充值权重,并生成Excel数据，针对一二期用户
func RechargeWeightRecordByChengId(pioneerAddressArr []string) error {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		panic(err)
	}

	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}

	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}

	var pioneerRecharges []PioneerRecharge

	for j, pioneerAddress := range pioneerAddressArr {
		fmt.Println(j)
		var chengShi [32]byte
		for true {
			chengShi, err = city.PioneerChengShi(nil, common.HexToAddress(pioneerAddress))
			if err == nil {
				break
			} else {
				log.Logger.Sugar().Error(err)
			}
		}
		// 获取城市名称
		var location models.UserLocation
		where := fmt.Sprintf("city_id='%s'", ConvertAreaIdBtA(chengShi))
		err = db.Mysql.Model(models.UserLocation{}).Where(where).First(&location).Error
		if err != nil {
			panic(err)
		}

		// 开始时间
		startData := time.Time{}
		date := time.Time{}

		weightPioneer := big.NewInt(0)
		for i := 90; i > 0; i-- {
			date = time.Now().Add(-time.Hour * 24 * time.Duration(i))
			fmt.Println(date.Format("2006-01-02 15:04:05"))
			if i == 90 {
				startData = time.Now().Add(-time.Hour * 24 * time.Duration(i))
				weightPioneer = big.NewInt(0)

			} else if i == 60 {
				pioneerRecharges = append(pioneerRecharges, PioneerRecharge{
					Pioneer:   pioneerAddress,
					City:      location.City,
					Weight:    weightPioneer.String(),
					StartDate: startData.Format("2006-01-02") + "-" + startData.Add(time.Hour*24*30).Format("2006-01-02"),
				})
				startData = time.Now().Add(-time.Hour * 24 * time.Duration(i))
				weightPioneer = big.NewInt(0)

			} else if i == 30 {
				pioneerRecharges = append(pioneerRecharges, PioneerRecharge{
					Pioneer:   pioneerAddress,
					City:      location.City,
					Weight:    weightPioneer.String(),
					StartDate: startData.Format("2006-01-02") + "-" + startData.Add(time.Hour*24*30).Format("2006-01-02"),
				})
				startData = time.Now().Add(-time.Hour * 24 * time.Duration(i))
				weightPioneer = big.NewInt(0)

			} else if i == 1 {
				pioneerRecharges = append(pioneerRecharges, PioneerRecharge{
					Pioneer:   pioneerAddress,
					City:      location.City,
					Weight:    weightPioneer.String(),
					StartDate: startData.Format("2006-01-02") + "-" + startData.Add(time.Hour*24*30).Format("2006-01-02"),
				})
			}

			var countyIds [][32]byte
			for {
				countyIds, err = userLocation.GetCountyIdsByChengShiId(nil, chengShi)
				if err == nil {
					break
				} else {
					log.Logger.Sugar().Error(err)
				}
			}

			countyWeight := big.NewInt(0)
			day := date.Unix() / 86400
			//wg := &sync.WaitGroup{}
			//countyLock := &sync.Mutex{}
			for _, countyId := range countyIds {
				//	go func(c *intoCityNode.City, w *sync.WaitGroup, l *sync.Mutex) {
				//		w.Add(1)
				//		for {
				//			amount, err := c.RechargeDailyWeightRecord(nil, countyId, big.NewInt(day))
				//			if err == nil {
				//				l.Lock()
				//				countyWeight = countyWeight.Add(countyWeight, amount.Div(amount, big.NewInt(1e18)))
				//				//fmt.Println(amount.String(), countyWeight.String())
				//				l.Unlock()
				//				break
				//			} else {
				//				log.Logger.Sugar().Error(err)
				//			}
				//		}
				//		w.Done()
				//	}(city, wg, countyLock)
				for {
					amount, err := city.RechargeDailyWeightRecord(nil, countyId, big.NewInt(day))
					if err == nil {
						countyWeight = countyWeight.Add(countyWeight, amount.Div(amount, big.NewInt(1e18)))
						break
					} else {
						log.Logger.Sugar().Error(err)
					}
				}
			}
			//wg.Wait()
			fmt.Println(weightPioneer.String(), countyWeight.String(), 67878908)
			weightPioneer = weightPioneer.Add(weightPioneer, countyWeight)
			fmt.Println(weightPioneer.String(), 678789080)
		}
	}

	f := xjexcel.ListToExcel(pioneerRecharges, "最近三个月充值权重", "详情")
	fileName := fmt.Sprintf("充值权重.xlsx")
	f.SaveAs(fileName)

	return nil
}
