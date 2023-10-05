package models

type Pioneer struct {
	Id      int32  `gorm:"column:id;primaryKey"`
	CityId  string `json:"city_id" gorm:"column:city_id"`
	Pioneer string `json:"pioneer" gorm:"column:pioneer"`
}
