package models

type PioneerAddInfo struct {
	Id            int32  `gorm:"column:id;primaryKey"`
	AreaId        string `json:"area_id" gorm:"column:area_id"`
	OldAreaId     string `json:"old_area_id" gorm:"column:old_area_id"`
	Location      string `json:"location" gorm:"column:location"`
	Pioneer       string `json:"pioneer" gorm:"column:pioneer"`
	AreaLevel     int64  `json:"area_level" gorm:"column:area_level"`
	IsAreaNode    int64  `json:"is_area_node" gorm:"column:is_area_node"`
	PioneerBatch  int64  `json:"pioneer_batch" gorm:"column:pioneer_batch"`
	SuretyUsdt    int64  `json:"surety_usdt" gorm:"column:surety_usdt"`
	SuretyTox     int64  `json:"surety_tox" gorm:"column:surety_tox"`
	IsSet         int64  `json:"is_set" gorm:"column:is_set"`
	IsReset       int64  `json:"is_reset" gorm:"column:is_reset"`
	PioneerStatus int64  `json:"pioneer_status" gorm:"column:pioneer_status"`
	Remark        string `json:"remark" gorm:"column:remark"`
	Ctime         string `json:"ctime" gorm:"column:ctime"`
}
