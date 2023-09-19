package models

type User struct {
	UserAddress string `json:"user_address" gorm:"column:user_address"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
}
