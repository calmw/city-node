package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type SuretyRecord struct {
	Id          int             `gorm:"column:id;primaryKey"`
	Pioneer     string          `json:"pioneer" gorm:"column:pioneer"`
	Amount      decimal.Decimal `json:"amount" gorm:"column:amount"`
	Month       int64           `json:"month" gorm:"column:month"`
	TxHash      string          `json:"tx_hash" gorm:"column:tx_hash"`
	BlockHeight int64           `json:"block_height" gorm:"column:block_height"`
	Ctime       time.Time       `json:"ctime" gorm:"column:ctime"`
}
