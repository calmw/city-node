package models

import (
	"github.com/shopspring/decimal"
)

type RechargeWeight struct {
	Id             int             `gorm:"column:id;primaryKey"`
	Pioneer        string          `json:"pioneer" gorm:"column:pioneer"`
	CountyId       string          `json:"county_id" gorm:"column:county_id"`
	CityId         string          `json:"city_id" gorm:"column:city_id"`
	CityLocation   string          `json:"city_location" gorm:"column:city_location"`
	CountyLocation string          `json:"county_location" gorm:"column:county_location"`
	Weight         decimal.Decimal `json:"weight" gorm:"column:weight"`
	Day            string          `json:"day" gorm:"column:day"`
}
