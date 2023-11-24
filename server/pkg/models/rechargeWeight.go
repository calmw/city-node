package models

import (
	"github.com/shopspring/decimal"
)

type RechargeWeight struct {
	Id           int             `gorm:"column:id;primaryKey"`
	Pioneer      string          `json:"pioneer" gorm:"column:pioneer"`
	User         string          `json:"user" gorm:"column:user"`
	CountyId     string          `json:"county_id" gorm:"column:county_id"`
	CityId       string          `json:"city_id" gorm:"column:city_id"`
	UserLocation string          `json:"county_location" gorm:"column:county_location"`
	Weight       decimal.Decimal `json:"weight" gorm:"column:weight"`
	Day          string          `json:"day" gorm:"column:day"`
	TxHash       string          `json:"tx_hash" gorm:"column:tx_hash"`
	BlockHeight  int64           `json:"block_height" gorm:"column:block_height"`
	LogIndex     int64           `json:"log_index" gorm:"column:log_index"`
}
