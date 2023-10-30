package models

type UserLocation struct {
	Id              int    `gorm:"column:id;primaryKey"`
	User            string `json:"user" gorm:"column:user"`
	CountyId        string `json:"county_id" gorm:"column:county_id"`
	CityId          string `json:"city_id" gorm:"column:city_id"`
	Location        string `json:"location" gorm:"column:location"`
	AreaCode        string `json:"area_code" gorm:"column:area_code"`
	LocationEncrypt string `json:"location_encrypt" gorm:"column:location_encrypt"`
	BlockHeight     int64  `json:"block_height" gorm:"column:block_height"`
	LogIndex        int64  `json:"log_index" gorm:"column:log_index"`
	TxHash          string `json:"tx_hash" gorm:"column:tx_hash"`
	Ctime           string `json:"ctime" gorm:"column:ctime"`
}
