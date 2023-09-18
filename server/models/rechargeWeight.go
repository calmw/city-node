package models

import (
	"github.com/shopspring/decimal"
)

//CREATE TABLE `recharge_weight` (
//  `id` int unsigned NOT NULL AUTO_INCREMENT,
//  `county_id` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
//  `city_id` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
//  `weight` decimal(40,0) DEFAULT NULL,
//  `location` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
//  `date` date DEFAULT NULL,
//  PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

type RechargeWeight struct {
	Id             int             `gorm:"column:id;primaryKey"`
	Pioneer        string          `json:"pioneer" gorm:"column:pioneer"`
	CountyId       string          `json:"county_id" gorm:"column:county_id"`
	CityId         string          `json:"city_id" gorm:"column:city_id"`
	CityLocation   string          `json:"city_location" gorm:"column:city_location"`
	CountyLocation string          `json:"county_location" gorm:"column:county_location"`
	Weight         decimal.Decimal `json:"weight" gorm:"column:weight"`
	Day            string          `json:"day" gorm:"column:day"`
}
