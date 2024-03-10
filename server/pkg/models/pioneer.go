package models

type Pioneer struct {
	Id             int32  `gorm:"column:id;primaryKey"`
	AreaId         string `json:"area_id" gorm:"column:area_id"`
	Location       string `json:"location" gorm:"column:location"`
	Pioneer        string `json:"pioneer" gorm:"column:pioneer"`
	FailedDelegate string `json:"failed_delegate" gorm:"column:failed_delegate"`
	FailedAt       string `json:"failed_at" gorm:"column:failed_at"`
	AreaLevel      int64  `json:"area_level" gorm:"column:area_level"`
	IsAreaNode     int64  `json:"is_area_node" gorm:"column:is_area_node"`
	PioneerBatch   int64  `json:"pioneer_batch" gorm:"column:pioneer_batch"`
	EndTime        string `json:"end_time" gorm:"column:end_time"`
	IsOverTime     int64  `json:"is_over_time" gorm:"column:is_over_time"`
	SuretyUsdt     int64  `json:"surety_usdt" gorm:"column:surety_usdt"`
	SuretyTox      int64  `json:"surety_tox" gorm:"column:surety_tox"`
}
