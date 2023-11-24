package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type Reward struct {
	Id             int             `gorm:"column:id;primaryKey"`
	Pioneer        string          `json:"pioneer" gorm:"column:pioneer"`
	TxHash         string          `json:"tx_hash" gorm:"column:tx_hash"`
	City           string          `json:"city" gorm:"column:city"`
	FoundsReward   decimal.Decimal `json:"founds_reward" gorm:"column:founds_reward"`
	DelegateReward decimal.Decimal `json:"delegate_reward" gorm:"column:delegate_reward"`
	NodeReward     decimal.Decimal `json:"node_reward" gorm:"column:node_reward"`
	BlockHeight    int64           `json:"block_height" gorm:"column:block_height"`
	Ctime          time.Time       `json:"ctime" gorm:"column:ctime"`
}
