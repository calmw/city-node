package models

type Children struct {
	Id    int    `gorm:"column:id;primaryKey"`
	User  string `json:"user" gorm:"column:user"`
	Child string `json:"child" gorm:"column:child"`
}
