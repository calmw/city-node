package models

import (
	"github.com/shopspring/decimal"
)

type CityPioneerToxTx struct {
	Id          int             `gorm:"column:id;primaryKey"`
	From        string          `json:"from_address" gorm:"column:from_address"`
	To          string          `json:"to_address" gorm:"column:to_address"`
	TxHash      string          `json:"tx_hash" gorm:"column:tx_hash"`
	Value       decimal.Decimal `json:"value" gorm:"column:value"`
	BlockHeight int64           `json:"block_height" gorm:"column:block_height"`
	LogIndex    int64           `json:"log_index" gorm:"column:log_index"`
	Ctime       string          `json:"ctime" gorm:"column:ctime"`
}
