package models

import "github.com/shopspring/decimal"

type Pioneer struct {
	Id        int32           `gorm:"column:id;primaryKey"`
	CityId    string          `json:"city_id" gorm:"column:city_id"`
	Location  string          `json:"location" gorm:"column:location"`
	Pioneer   string          `json:"pioneer" gorm:"column:pioneer"`
	IsDeposit int             `json:"is_deposit" gorm:"column:is_deposit"`
	CityLevel int64           `json:"city_level" gorm:"column:city_level"`
	Surety    decimal.Decimal `json:"surety" gorm:"column:surety"`
}
