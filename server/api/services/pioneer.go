package services

import (
	"city-node-server/api/common/statecode"
	"city-node-server/api/models/request"
	"city-node-server/db"
	models2 "city-node-server/models"
	"fmt"
	"github.com/shopspring/decimal"
)

type Pioneer struct{}

func NewPioneer() *Pioneer {
	return &Pioneer{}
}

type PioneerWeight struct {
	CityId        string              `json:"city_id"`
	Location      string              `json:"location"`
	TotalWeight   decimal.Decimal     `json:"total_weight"`
	PioneerCounty []PioneerCountyData `json:"pioneer_county_data"`
}

type PioneerCountyData struct {
	CountyId           string               `json:"county_id"`
	Location           string               `json:"location"`
	TotalWeight        decimal.Decimal      `json:"total_weight"`
	PioneerWeightDaily []PioneerWeightDaily `json:"county_weight_daily"`
}

type PioneerWeightDaily struct {
	Day    string          `json:"day"`
	Weight decimal.Decimal `json:"weight"`
}

type FHSum struct {
	Total float64
}

func (c *Pioneer) GetRechargeWeightByPioneerAddress(userReq *request.GetRechargeWeightByPioneerAddress) (int, PioneerWeight) {
	var pioneer models2.RechargeWeight
	whereCondition := fmt.Sprintf("pioneer='%s'", userReq.Pioneer)
	db.Mysql.Table("recharge_weight").Where(whereCondition).First(&pioneer)
	var rechargeWeight []models2.RechargeWeight
	whereCondition = fmt.Sprintf("pioneer='%s'", userReq.Pioneer)
	db.Mysql.Table("recharge_weight").Where(whereCondition).Group("county_id").Find(&rechargeWeight)
	pioneerWeight := PioneerWeight{
		CityId:   pioneer.CityId,
		Location: pioneer.CityLocation,
	}
	for _, rc := range rechargeWeight {
		var rechargeWeightDaily []models2.RechargeWeight
		var pioneerWeightDailys []PioneerWeightDaily
		whereCondition = fmt.Sprintf("county_id='%s'", rc.CountyId)
		db.Mysql.Table("recharge_weight").Where(whereCondition).Order("day desc").Find(&rechargeWeightDaily)
		for _, dw := range rechargeWeightDaily {
			pioneerWeightDailys = append(pioneerWeightDailys, PioneerWeightDaily{
				Day:    dw.Day,
				Weight: dw.Weight,
			})
		}
		var total FHSum
		db.Mysql.Table("recharge_weight").Where(whereCondition).Select("sum(weight) as total").Scan(&total)
		whereCondition = fmt.Sprintf("county_id='%s'", rc.CountyId)
		pioneerWeight.PioneerCounty = append(pioneerWeight.PioneerCounty, PioneerCountyData{
			CountyId:           rc.CountyId,
			Location:           rc.CountyLocation,
			TotalWeight:        decimal.NewFromFloat(total.Total),
			PioneerWeightDaily: pioneerWeightDailys,
		})
	}

	return statecode.CommonSuccess, pioneerWeight
}
