package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type Delegate struct {
	Id      int32           `gorm:"column:id;primaryKey"`
	CityId  string          `json:"city_id" gorm:"column:city_id"`
	TxHash  string          `json:"tx_hash" gorm:"column:tx_hash"`
	Amount  decimal.Decimal `json:"amount" gorm:"column:amount"`
	SetType int64           `json:"set_type" gorm:"column:set_type"`
	Ctime   time.Time       `json:"ctime" gorm:"column:ctime"`
}
