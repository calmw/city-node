package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type RewardWithdraw struct {
	Id          int             `gorm:"column:id;primaryKey"`
	Pioneer     string          `json:"pioneer" gorm:"column:pioneer"`
	TxHash      string          `json:"tx_hash" gorm:"column:tx_hash"`
	Amount      decimal.Decimal `json:"amount" gorm:"column:amount"`
	RewardType  decimal.Decimal `json:"reward_type" gorm:"column:reward_type"`
	BlockHeight int64           `json:"block_height" gorm:"column:block_height"`
	Ctime       time.Time       `json:"ctime" gorm:"column:ctime"`
}
