package models

import (
	"time"
)

type PioneerTask struct {
	Id      int       `gorm:"column:id;primaryKey"`
	Pioneer string    `json:"pioneer" gorm:"column:pioneer"`
	Status  int64     `json:"status" gorm:"column:status"`
	Date    time.Time `json:"date" gorm:"column:date"`
}
