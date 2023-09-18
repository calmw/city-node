package models

type UserLocation struct {
	Id              int    `gorm:"column:id;primaryKey"`
	User            string `json:"user" gorm:"column:user"`
	CountyId        string `json:"county_id" gorm:"column:county_id"`
	CityId          string `json:"city_id" gorm:"column:city_id"`
	Location        string `json:"location" gorm:"column:location"`
	LocationEncrypt string `json:"location_encrypt" gorm:"column:location_encrypt"`
	AreaCode        string `json:"area_code" gorm:"column:area_code"`
}
