package models

import (
	"github.com/shopspring/decimal"
)

type PledgeBalanceTransferable struct {
	Id          int             `gorm:"column:id;primaryKey"`
	Sender      string          `json:"sender" gorm:"column:sender"`
	Amount      decimal.Decimal `json:"amount" gorm:"column:amount"`
	Types       int64           `json:"types" gorm:"column:types"`
	Timestamp   int64           `json:"timestamp" gorm:"column:timestamp"`
	BlockHeight int64           `json:"block_height" gorm:"column:block_height"`
	LogIndex    int64           `json:"log_index" gorm:"column:log_index"`
	Ctime       string          `json:"ctime" gorm:"column:ctime"`
}
