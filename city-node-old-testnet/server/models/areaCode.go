package models

import "time"

type AreaCode struct {
	Id          int       `gorm:"column:id;primaryKey"`
	CountryName string    `json:"country_name" gorm:"column:country_name"`
	CountryCode uint64    `json:"country_code" gorm:"column:country_code"`
	CityName    string    `json:"city_name" gorm:"column:city_name"`
	CityCode    uint      `json:"city_code" gorm:"column:city_code"`
	AreaName    string    `json:"area_name" gorm:"column:area_name"`
	AdCode      uint      `json:"ad_code" gorm:"column:ad_code"`
	Deleted     string    `json:"deleted" gorm:"column:deleted"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
}
