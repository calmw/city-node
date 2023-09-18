package services

import (
	"city-node-server/api/common/statecode"
	"city-node-server/api/models/request"
	"city-node-server/db"
	models2 "city-node-server/models"
	"fmt"
)

type UserLocationService struct{}

func NewUserLocation() *UserLocationService {
	return &UserLocationService{}
}

func (c *UserLocationService) GetLocationByUserAddress(userReq *request.GetLocationByUserAddress) (int, []models2.UserLocation) {
	var userLocation []models2.UserLocation
	whereCondition := fmt.Sprintf("user='%s'", userReq.User)
	db.Mysql.Table("user_location").Where(whereCondition).Find(&userLocation)

	return statecode.CommonSuccess, userLocation
}
