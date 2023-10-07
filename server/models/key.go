package models

type PrivateKey struct {
	Id  int32  `gorm:"column:id;primaryKey"`
	Key string `json:"key" gorm:"column:key"`
}
