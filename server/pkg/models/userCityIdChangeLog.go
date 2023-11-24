package models

type UserCityIdChangeLog struct {
	Id          int32  `gorm:"column:id;primaryKey"`
	User        string `json:"user" gorm:"column:user"`
	CityId      string `json:"city_id" gorm:"column:city_id"`
	CountyId    string `json:"county_id" gorm:"column:county_id"`
	PioneerCity string `json:"pioneer_city" gorm:"column:pioneer_city"`
}
