package services

import (
	"city-node-server/api/common/statecode"
	"city-node-server/api/models/request"
	"city-node-server/db"
	models2 "city-node-server/models"
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"strings"
	"time"
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
	Day    time.Time       `json:"day"`
	Weight decimal.Decimal `json:"weight"`
}

type FHSum struct {
	Total float64
}

func (c *Pioneer) Reward(req *request.Reward) (int, int64, []models2.Reward) {
	var rewards []models2.Reward

	tx := db.Mysql.Model(&models2.Reward{}).Order("ctime desc")
	if req.Pioneer != "" {
		tx = tx.Where("pioneer=?", strings.ToLower(req.Pioneer))
	}
	if req.Start != "" {
		tx = tx.Where("ctime>=? and ctime<=?", req.Start, req.End)
	}
	var total int64
	tx.Count(&total)
	page := 1 // 第5页
	if req.Page > 0 {
		page = req.Page
	}
	pageSize := 10
	if req.PageSize > 0 {
		pageSize = req.PageSize
	}
	offset := (page - 1) * pageSize // 计算偏移量
	err := tx.Debug().Limit(pageSize).Offset(offset).Find(&rewards).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return statecode.CommonErrServerErr, 0, rewards
	}
	return statecode.CommonSuccess, total, rewards
}

//func (c *Pioneer) GetRechargeWeightByPioneerAddress(userReq *request.GetRechargeWeightByPioneerAddress) (int, PioneerWeight) {
//	var pioneer models2.RechargeWeight
//	var total FHSum
//	whereCondition := fmt.Sprintf("pioneer='%s'", strings.ToLower(userReq.Pioneer))
//	db.Mysql.Table("recharge_weight").Where(whereCondition).First(&pioneer)
//	var rechargeWeight []models2.RechargeWeight
//	whereCondition = fmt.Sprintf("pioneer='%s'", strings.ToLower(userReq.Pioneer))
//
//	db.Mysql.Table("recharge_weight").Select("`county_id`,`pioneer`").Where(whereCondition).Group("county_id").Find(&rechargeWeight)
//	db.Mysql.Table("recharge_weight").Where(whereCondition).Select("sum(weight) as total").Scan(&total)
//	pioneerWeight := PioneerWeight{
//		CityId:      pioneer.CityId,
//		Location:    pioneer.CityLocation,
//		TotalWeight: decimal.NewFromFloat(total.Total),
//	}
//	for _, rc := range rechargeWeight {
//		var rechargeWeightDaily []models2.RechargeWeight
//		var pioneerWeightDailys []PioneerWeightDaily
//		whereCondition = fmt.Sprintf("county_id='%s'", rc.CountyId)
//		db.Mysql.Table("recharge_weight").Where(whereCondition).Order("day desc").Find(&rechargeWeightDaily)
//		for _, dw := range rechargeWeightDaily {
//			pioneerWeightDailys = append(pioneerWeightDailys, PioneerWeightDaily{
//				Day:    dw.Day,
//				Weight: dw.Weight,
//			})
//		}
//
//		// 获取区县位置信息
//		_, countyLocation := GetCountyInfo(rc.CountyId)
//		total.Total = 0
//		db.Mysql.Table("recharge_weight").Where(whereCondition).Select("sum(weight) as total").Scan(&total)
//		whereCondition = fmt.Sprintf("county_id='%s'", rc.CountyId)
//		pioneerWeight.PioneerCounty = append(pioneerWeight.PioneerCounty, PioneerCountyData{
//			CountyId:           rc.CountyId,
//			Location:           countyLocation,
//			TotalWeight:        decimal.NewFromFloat(total.Total),
//			PioneerWeightDaily: pioneerWeightDailys,
//		})
//	}
//
//	return statecode.CommonSuccess, pioneerWeight
//}

func (c *Pioneer) RechargeWeight(req *request.RechargeWeight) (int, int64, []models2.RechargeWeight) {
	var rechargeWeights []models2.RechargeWeight

	tx := db.Mysql.Model(&models2.RechargeWeight{}).Order("day desc")
	if req.Pioneer != "" {
		tx = tx.Where("pioneer=?", strings.ToLower(req.Pioneer))
	}
	if req.User != "" {
		tx = tx.Where("user=?", strings.ToLower(req.User))
	}
	if req.Start != "" {
		tx = tx.Where("day>=? and day<=?", req.Start, req.End)
	}
	var total int64
	tx.Count(&total)
	page := 1 // 第5页
	if req.Page > 0 {
		page = req.Page
	}
	pageSize := 10
	if req.PageSize > 0 {
		pageSize = req.PageSize
	}
	offset := (page - 1) * pageSize // 计算偏移量
	tx.Limit(pageSize).Offset(offset).Find(&rechargeWeights)

	return statecode.CommonSuccess, total, rechargeWeights
}

func (c *Pioneer) RewardWithdraw(req *request.RewardWithdraw) (int, int64, []models2.RewardWithdraw) {
	var rewardWithdraw []models2.RewardWithdraw

	tx := db.Mysql.Model(&models2.RewardWithdraw{}).Order("ctime desc")
	if req.Pioneer != "" {
		tx = tx.Where("pioneer=?", strings.ToLower(req.Pioneer))
	}

	if req.Start != "" {
		tx = tx.Where("ctime>=? and ctime<=?", req.Start, req.End)
	}
	var total int64
	tx.Count(&total)
	page := 1 // 第5页
	if req.Page > 0 {
		page = req.Page
	}
	pageSize := 10
	if req.PageSize > 0 {
		pageSize = req.PageSize
	}
	offset := (page - 1) * pageSize // 计算偏移量
	tx.Limit(pageSize).Offset(offset).Find(&rewardWithdraw)

	return statecode.CommonSuccess, total, rewardWithdraw
}

func (c *Pioneer) Pioneer(req *request.Pioneer) (int, int64, []models2.Pioneer) {
	var rewardWithdraw []models2.Pioneer

	tx := db.Mysql.Model(&models2.Pioneer{}).Order("id asc")
	if req.Pioneer != "" {
		tx = tx.Where("pioneer=?", strings.ToLower(req.Pioneer))
	}

	var total int64
	tx.Count(&total)
	page := 1 // 第5页
	if req.Page > 0 {
		page = req.Page
	}
	pageSize := 10
	if req.PageSize > 0 {
		pageSize = req.PageSize
	}
	offset := (page - 1) * pageSize // 计算偏移量
	tx.Limit(pageSize).Offset(offset).Find(&rewardWithdraw)

	return statecode.CommonSuccess, total, rewardWithdraw
}

func GetCountyInfo(countyId string) (error, string) {
	var userLocation models2.UserLocation
	countyId = strings.ToLower(countyId)
	whereCondition := fmt.Sprintf("county_id='%s'", countyId)
	err := db.Mysql.Table("user_location").Where(whereCondition).First(&userLocation).Error
	return err, userLocation.Location
}
