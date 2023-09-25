package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type RechargeRecord struct {
	Id       int             `gorm:"column:id;primaryKey"`
	User     string          `json:"user" gorm:"column:user"`
	CountyId string          `json:"county_id" gorm:"column:county_id"`
	CityId   string          `json:"city_id" gorm:"column:city_id"`
	Amount   decimal.Decimal `json:"amount" gorm:"column:amount"`
	Ctime    time.Time       `json:"ctime" gorm:"column:ctime"`
}
