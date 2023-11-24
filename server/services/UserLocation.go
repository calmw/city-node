package services

import (
	"city-node-server/db"
	"city-node-server/pkg/blockchain"
	models2 "city-node-server/pkg/models"
	utils2 "city-node-server/pkg/utils"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

func InitUserLocation() {

	// 管理员设置城市合约地址
	//blockchain.AdminSetCityAddressUserLocation()
	//blockchain.UserLocationAddAdmin()
	blockchain.AdminSetPledgeStakeAddressUserLocation()
	// 管理员设置获取用户质押量合约地址
	//blockchain.AdminSetPledgeStakeAddressUserLocation()
	//blockchain.SetUserLocationTest("0xaa3f3f93c743005d497d2a21cf5ac5132960e42496d9b45437722351f8524496", "05C0TS6p86w=", "0x77bd41fdE654FE0054b771Ec6985dC9d5247BAfe")

}

func UserCityChange() {
	var userLocation []models2.UserLocation
	db.Mysql.Model(&models2.UserLocation{}).Find(&userLocation)
	for _, l := range userLocation {
		codeSlic := strings.Split(l.AreaCode, ",")
		if len(codeSlic[2]) != 6 { // code最后一段是六位数的按国内用户
			//fmt.Println("跳过：", l.User)
			continue
		}
		codeSecondSlic := strings.Split(codeSlic[1], "")
		if codeSecondSlic[0] == "0" { // code第二段开头是0的用户
			err, oldCityId := blockchain.GetChengShiIdByCountyId(l.CountyId)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(l.User, codeSlic[2], "+++++")
			err, cityCode := GetCityCodeByAdCode(codeSlic[2])
			if err != nil {
				fmt.Println(err)
				continue
			}
			encrypt := fmt.Sprintf("%d,%d", 0, cityCode)
			newCityId := strings.ToLower(utils2.Keccak256(encrypt))
			//blockchain.SetCityMapping(l.CountyId, newCityId, l.LocationEncrypt)
			InsertUserCItyIdChangeLog(l.User, l.CountyId, oldCityId, newCityId)

		}
	}
}

func GetCityCodeByAdCode(adCode string) (error, uint) {
	var areaCode models2.AreaCode
	whereCondition := fmt.Sprintf("ad_code=%s", adCode)
	err := db.Mysql.Model(&models2.AreaCode{}).Where(whereCondition).First(&areaCode).Error
	if err != nil {
		// 查询明文地址
		uri := fmt.Sprintf("https://wallet-api-v2.intowallet.io/api/v1/city_node/geographic_info?city_code=%s&ad_code=%s", "0", adCode)
		err, location := utils2.HttpGet(uri)
		if err != nil {
			return err, 0
		}
		var locationInfo blockchain.LocationInfo
		err = json.Unmarshal(location, &locationInfo)
		if err != nil {
			return err, 0
		}
		fmt.Println(err, uri, locationInfo)

		return err, 0
	}
	return nil, areaCode.CityCode
}

func InsertUserCItyIdChangeLog(userAddress, countyId, cityId, newCityId string) error {
	var userCityIdChangeLog models2.UserCityIdChangeLog
	whereCondition := fmt.Sprintf("user='%s'", strings.ToLower(userAddress))
	err := db.Mysql.Model(&models2.UserCityIdChangeLog{}).Where(whereCondition).First(&userCityIdChangeLog).Error
	if err == gorm.ErrRecordNotFound {

		db.Mysql.Model(&models2.UserCityIdChangeLog{}).Create(&models2.UserCityIdChangeLog{
			User:        strings.ToLower(userAddress),
			CountyId:    strings.ToLower(countyId),
			CityId:      strings.ToLower(cityId),
			PioneerCity: strings.ToLower(newCityId),
		})
	}
	return nil
}
