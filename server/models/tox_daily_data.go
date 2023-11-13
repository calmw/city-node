package models

import (
	"github.com/shopspring/decimal"
)

type ToxDayData struct {
	Id     int             `gorm:"column:id;primaryKey"`
	LogId  string          `json:"log_id" gorm:"column:log_id"`
	From   string          `json:"from" gorm:"column:from"`
	Addr   string          `json:"addr" gorm:"column:addr"`
	Amount decimal.Decimal `json:"amount" gorm:"column:amount"`
	Status int64           `json:"status" gorm:"column:status"`
	Date   int64           `json:"date" gorm:"column:date"`
}
