package services

import (
	"city-node-server/api/common/statecode"
	"city-node-server/api/models/request"
	"city-node-server/db"
	models2 "city-node-server/models"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type UserLocationService struct{}

func NewUserLocation() *UserLocationService {
	return &UserLocationService{}
}

func (c *UserLocationService) GetLocationByUserAddress(userReq *request.GetLocationByUserAddress) (int, []models2.UserLocation) {
	var userLocation []models2.UserLocation
	whereCondition := fmt.Sprintf("user='%s'", strings.ToLower(userReq.User))
	db.Mysql.Table("user_location").Where(whereCondition).Find(&userLocation)

	return statecode.CommonSuccess, userLocation
}

func (c *UserLocationService) UserLocation(req *request.Location) (int, []models2.UserLocation) {
	var userLocations []models2.UserLocation

	tx := db.Mysql.Model(&models2.UserLocation{})
	if req.User != "" {
		tx = tx.Where("user=?", req.User)
	}

	page := 1 // 第5页
	if req.Page > 0 {
		page = req.Page
	}
	pageSize := 10
	if req.PageSize > 0 {
		pageSize = req.PageSize
	}
	offset := (page - 1) * pageSize // 计算偏移量
	err := tx.Debug().Limit(pageSize).Offset(offset).Find(&userLocations).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return statecode.CommonErrServerErr, userLocations
	}
	return statecode.CommonSuccess, userLocations
}
