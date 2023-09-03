package models

import (
	"city-node-server/db"
	"errors"
	"gorm.io/gorm"
)

type AdminSetDelegate struct {
	Id      int32  `gorm:"column:id;primaryKey"`
	CityId  string `json:"city_id" gorm:"column:city_id"`
	Amount  string `json:"amount" gorm:"column:amount"`
	SetType int64  `json:"set_type" gorm:"column:set_type"`
}

func NewAdminSetDelegate() *AdminSetDelegate {
	return &AdminSetDelegate{}
}

// Set Multi-Sign
func (m *AdminSetDelegate) Set() error {

	return nil
}

// Get Multi-Sign
func (m *AdminSetDelegate) Get(chainId int) error {
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
