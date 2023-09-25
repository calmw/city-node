package models

import (
	"city-node-server/db"
	"errors"
	"gorm.io/gorm"
)

type BlockStore struct {
	Id         int32  `gorm:"column:id;primaryKey"`
	StartBlock uint64 `json:"start_block" gorm:"column:start_block"`
}

func NewBlockStore() *BlockStore {
	return &BlockStore{}
}

// Set Multi-Sign
func (m *BlockStore) Set() error {

	return nil
}

// Get Multi-Sign
func (m *BlockStore) Get(chainId int) error {
	err := db.Mysql.Table("multi_sign").Where("chain_id", chainId).First(&m).Debug().Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			return errors.New("record select err " + err.Error())
		}
	}
	return nil
}
