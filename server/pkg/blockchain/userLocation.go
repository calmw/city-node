package blockchain

import (
	"city-node-server/pkg/binding/intoCityNode"
	"city-node-server/pkg/blockchain/event"
	"city-node-server/pkg/db"
	"city-node-server/pkg/log"
	models2 "city-node-server/pkg/models"
	utils2 "city-node-server/pkg/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/status-im/keycard-go/hexutils"
	"gorm.io/gorm"
	"math/big"
	"strconv"
	"strings"
	"time"
)

type LocationInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		CountryName string `json:"country_name"`
		CityName    string `json:"city_name"`
		CityCode    int    `json:"city_code"`
		AdCode      int    `json:"ad_code"`
		AreaName    string `json:"area_name"`
	} `json:"data"`
}

// SetNoRepeatCityIds   城市ID数组重构
func SetNoRepeatCityIds(start, end int64) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	err, auth := GetAuth(Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	tx, err := userLocation.NoRepeatCityIds(auth, big.NewInt(start), big.NewInt(end))

	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(start, end, tx)
}

// AdminSetCityAddressUserLocation 管理员设置城市合约地址
func AdminSetCityAddressUserLocation() {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	err, auth := GetAuth(Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	res, err := userLocation.AdminSetCityAddress(auth, common.HexToAddress(CityNodeConfig.CityAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res, err)
}

// AddAdmin 管理员设置城市合约地址
func UserLocationAddAdmin() {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	err, auth := GetAuth(Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	res, err := userLocation.AddAdmin(auth, common.HexToAddress("0x94b627F4F829Ac5E97fDc556B5BEeeFf9beF417e"))
	fmt.Println(res, err, 123)
}

// AdminSetPledgeStakeAddressUserLocation 管理员设置获取用户质押量合约地址
func AdminSetPledgeStakeAddressUserLocation() {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	err, auth := GetAuth(Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	res, err := userLocation.AdminSetPledgeStakeAddress(auth, common.HexToAddress(CityNodeConfig.PledgeStakeAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res, err)
}

// SetUserLocation 设置获取用户位置
func SetUserLocation(cityId, location string) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	err, auth := GetAuth(Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	if strings.Contains(cityId, "0x") {
		cityId = strings.ReplaceAll(cityId, "0x", "")
	}
	common.Hex2Bytes(cityId)
	cityIdBytes32 := BytesToByte32(common.Hex2Bytes(cityId))
	res, err := userLocation.SetUserLocation(auth, cityIdBytes32, location)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res, err)
}

// AdminSetSecondsPerDayCityUserLocation 管理员设置每天秒数，用于测试
func AdminSetSecondsPerDayCityUserLocation(seconds int64) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	err, auth := GetAuth(Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	res, err := userLocation.AdminSetSecondsPerDay(auth, big.NewInt(seconds))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res.Hash(), err)
}

// SetUserLocationTest 设置获取用户位置
func SetUserLocationTest(cityId, location, userAddress string) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	err, auth := GetAuth(Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	if strings.Contains(cityId, "0x") {
		cityId = strings.ReplaceAll(cityId, "0x", "")
	}
	common.Hex2Bytes(cityId)
	cityIdBytes32 := BytesToByte32(common.Hex2Bytes(cityId))
	res, err := userLocation.SetUserLocationTest(auth, cityIdBytes32, common.HexToAddress(userAddress), location)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res, err)
}

// SetCityMapping 设置获取用户位置ID映射
func SetCityMapping(countyId, chengShiId, locationChengShi string) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	if strings.Contains(countyId, "0x") {
		countyId = strings.ReplaceAll(countyId, "0x", "")
	}
	common.Hex2Bytes(countyId)
	cityIdBytes32 := BytesToByte32(common.Hex2Bytes(countyId))
	//
	if strings.Contains(chengShiId, "0x") {
		chengShiId = strings.ReplaceAll(chengShiId, "0x", "")
	}
	common.Hex2Bytes(chengShiId)
	chengShiIdBytes32 := BytesToByte32(common.Hex2Bytes(chengShiId))
	err, auth := GetAuth(Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	_, err = userLocation.SetCityMapping(auth, cityIdBytes32, chengShiIdBytes32, locationChengShi)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(err)
}

func GetChengShiIdByCountyId(countyId string) (error, string) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ""
	}
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ""
	}
	if strings.Contains(countyId, "0x") {
		countyId = strings.ReplaceAll(countyId, "0x", "")
	}
	common.Hex2Bytes(countyId)
	countyIdBytes32 := BytesToByte32(common.Hex2Bytes(countyId))
	chengShiIdBytes32, err := userLocation.GetChengShiIdByCountyId(nil, countyIdBytes32)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ""
	}
	return nil, strings.ToLower("0x" + hexutils.BytesToHex(Bytes32ToBytes(chengShiIdBytes32)))
}

func GetCountyIdsByChengShiId(chengShiId string) (error, []string) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, []string{}
	}
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, []string{}
	}
	if strings.Contains(chengShiId, "0x") {
		chengShiId = strings.ReplaceAll(chengShiId, "0x", "")
	}
	chengShiIdBytes32 := BytesToByte32(common.Hex2Bytes(chengShiId))
	countyIds, err := userLocation.GetCountyIdsByChengShiId(nil, chengShiIdBytes32)
	if err != nil {
		return err, []string{}
	}
	var cIds []string
	for _, id := range countyIds {
		cIds = append(cIds, strings.ToLower("0x"+hexutils.BytesToHex(Bytes32ToBytes(id))))
	}

	return nil, cIds
}

// GetCityId 获取cityId
func GetCityId(cityIdIndex int64) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	res, err := userLocation.CityIdsNoRepeat(nil, big.NewInt(cityIdIndex))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	info, err := userLocation.CityInfo(nil, res)
	if err != nil {
		return
	}
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	fmt.Println("0x"+common.Bytes2Hex(Bytes32ToBytes(res)), info, err)
}

//// GetCountyIdsByChengShiId 获取城市ID对应的所有区县ID
//func GetCountyIdsByChengShiId(cityId string) (error, [][32]byte) {
//	err, Cli := Client(CityNodeConfig)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return err, nil
//	}
//	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return err, [][32]byte{}
//	}
//	cityIdBytes32 := BytesToByte32(common.Hex2Bytes(cityId))
//	res, err := userLocation.GetCountyIdsByChengShiId(nil, cityIdBytes32)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return err, [][32]byte{}
//	}
//	return nil, res
//}

// GetCountyIdsByChengShiIdBytes32 获取城市ID对应的所有区县ID
func GetCountyIdsByChengShiIdBytes32(cityId [32]byte) (error, [][32]byte) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, [][32]byte{}
	}
	res, err := userLocation.GetCountyIdsByChengShiId(nil, cityId)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, [][32]byte{}
	}
	return nil, res
}

// GetDay 获取天
func GetDay() (error, int64) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, 0
	}
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, 0
	}
	res, err := userLocation.GetDay(nil)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, 0
	}
	return nil, res.Int64()
}

// RestoreUserLocation 更新单个用户地址
func RestoreUserLocation(user string) error {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	user = strings.ToLower(user)
	userLocationContract, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	locationEncrypt, err := userLocationContract.UserLocationInfo(nil, common.HexToAddress(user))
	if err != nil || locationEncrypt == "" {
		log.Logger.Sugar().Error(err)
		return err
	}

	countyIdBytes32, err := userLocationContract.UserCityId(nil, common.HexToAddress(user))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}

	cityIdBytes32, err := userLocationContract.CityIdChengShiID(nil, countyIdBytes32)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	cityId := "0x" + common.Bytes2Hex(Bytes32ToBytes(cityIdBytes32))
	countyId := "0x" + common.Bytes2Hex(Bytes32ToBytes(countyIdBytes32))
	locationCode := utils2.ThreeDesDecrypt(locationEncrypt)
	code := strings.Split(locationCode, ",")
	if code[0] == "0" && code[2] == "" { // 兼容map增量更新的位置信息
		locationEncrypt, err = userLocationContract.CityInfo(nil, cityIdBytes32)
		if err != nil || locationEncrypt == "" {
			log.Logger.Sugar().Error(err)
			return err
		}
		locationCode = utils2.ThreeDesDecrypt(locationEncrypt)
		code = strings.Split(locationCode, ",")
	}
	if len(code) == 2 {
		code = append(code, "0")
	}
	if len(code) < 3 {
		log.Logger.Sugar().Warn("用户位置加密信息解析错误", user, locationCode, locationEncrypt, code)
		return errors.New("用户位置加密信息解析错误")
	}
	// 国外用户
	codeSecond := code[1]
	codeSecondSlice := strings.Split(codeSecond, "")
	if code[0] != "0" && codeSecondSlice[0] != "0" {
		log.Logger.Sugar().Warn("国外用户位置信息", user, locationCode, locationEncrypt, code)
		code[2] = "0"
	}
	// 容错，国内城市code首位是0的情况
	if codeSecondSlice[0] == "0" {
		// 获取城市code,根据区县code
		var areaCode models2.AreaCode
		whereCondition := fmt.Sprintf("ad_code=%s", code[2])
		err = db.Mysql.Table("area_code").Where(whereCondition).First(&areaCode).Error
		code[1] = fmt.Sprintf("%d", areaCode.CityCode)
	}

	err = InsertUserLocation(user, countyId, code, locationEncrypt, locationCode, countyIdBytes32, 0)
	if err != nil {
		return err
	}
	var userLocation models2.UserLocation
	whereCondition := fmt.Sprintf("user='%s' and city_id='%s' and area_code='%s' and county_id='%s'",
		user, strings.ToLower(cityId), locationCode, strings.ToLower(countyId))
	err = db.Mysql.Table("user_location").Where(whereCondition).First(&userLocation).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println(code, "========++++")
		// 查询明文地址
		uri := fmt.Sprintf("https://wallet-api-v2.intowallet.io/api/v1/city_node/geographic_info?city_code=%s&ad_code=%s", code[1], code[2])
		err, location := utils2.HttpGet(uri)
		fmt.Println(err, location, "========")
		if err != nil {
			return err
		}
		var locationInfo LocationInfo
		err = json.Unmarshal(location, &locationInfo)
		if err != nil {
			return err
		}
		db.Mysql.Model(&models2.UserLocation{}).Create(&models2.UserLocation{
			// 	Id              int    `gorm:"column:id;primaryKey"`
			//	User            string `json:"user" gorm:"column:user"`
			//	CountyId        string `json:"county_id" gorm:"column:county_id"`
			//	CityId          string `json:"city_id" gorm:"column:city_id"`
			//	Location        string `json:"location" gorm:"column:location"`
			//	AreaCode        string `json:"area_code" gorm:"column:area_code"`
			//	LocationEncrypt string `json:"location_encrypt" gorm:"column:location_encrypt"`
			//	BlockHeight     int64  `json:"block_height" gorm:"column:block_height"`
			//	LogIndex        int64  `json:"log_index" gorm:"column:log_index"`
			//	TxHash          string `json:"tx_hash" gorm:"column:tx_hash"`
			//	Ctime           string `json:"ctime" gorm:"column:ctime"`
			User:            user,
			CountyId:        strings.ToLower(countyId),
			CityId:          strings.ToLower(cityId),
			LocationEncrypt: locationEncrypt,
			//BlockHeight:
			Location: locationInfo.Data.CountryName + " " + locationInfo.Data.CityName + " " + locationInfo.Data.AreaName,
			AreaCode: locationCode,
		})
	}
	return nil
}

// ReSaveUserLocation 更新CityID和location为空的所有地址
func ReSaveUserLocation() {
	var locations []models2.UserLocation
	db.Mysql.Model(models2.UserLocation{}).Where("location=''").Find(&locations)
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	userLocationContract, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	for _, l := range locations {
		fmt.Println(l.User, l.Id)
		countyIdBytes32, err := userLocationContract.UserCityId(nil, common.HexToAddress(l.User))
		if err != nil {
			log.Logger.Sugar().Error(err)
			return
		}

		cityIdBytes32, err := userLocationContract.CityIdChengShiID(nil, countyIdBytes32)
		if err != nil {
			log.Logger.Sugar().Error(err)
			return
		}

		code := strings.Split(l.AreaCode, ",")
		if code[0] == "0" && code[2] == "" { // 兼容map增量更新的位置信息
			locationEncrypt, err := userLocationContract.CityInfo(nil, cityIdBytes32)
			if err != nil || locationEncrypt == "" {
				log.Logger.Sugar().Error(err)
				return
			}
			locationCode := utils2.ThreeDesDecrypt(locationEncrypt)
			code = strings.Split(locationCode, ",")
		}
		if len(code) == 2 {
			code = append(code, "0")
		}
		if len(code) < 3 {
			log.Logger.Sugar().Warn("用户位置加密信息解析错误", code)
			return
		}
		// 国外用户
		codeSecond := code[1]
		codeSecondSlice := strings.Split(codeSecond, "")
		if code[0] != "0" && codeSecondSlice[0] != "0" {
			log.Logger.Sugar().Warn("国外用户位置信息", code)
			code[2] = "0"
		}
		// 容错，国内城市code首位是0的情况
		if codeSecondSlice[0] == "0" {
			// 获取城市code,根据区县code
			var areaCode models2.AreaCode
			whereCondition := fmt.Sprintf("ad_code=%s", code[2])
			err = db.Mysql.Table("area_code").Where(whereCondition).First(&areaCode).Error
			code[1] = fmt.Sprintf("%d", areaCode.CityCode)
		}

		// 查询明文地址
		uri := fmt.Sprintf("https://wallet-api-v2.intowallet.io/api/v1/city_node/geographic_info?city_code=%s&ad_code=%s", code[1], code[2])
		err, location := utils2.HttpGet(uri)
		if err != nil {
			return
		}
		var locationInfo LocationInfo
		err = json.Unmarshal(location, &locationInfo)
		if err != nil {
			return
		}
		db.Mysql.Model(&models2.UserLocation{}).Where("user=?", strings.ToLower(l.User)).Updates(map[string]interface{}{
			"city_id":  strings.ToLower("0x" + hexutils.BytesToHex(Bytes32ToBytes(cityIdBytes32))),
			"location": locationInfo.Data.CountryName + " " + locationInfo.Data.CityName + " " + locationInfo.Data.AreaName,
		})
	}

}

// GetChengShiIdByCityId 获取cityId
func GetChengShiIdByCityId(countyId string) (error, string) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ""
	}
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ""
	}
	common.Hex2Bytes(countyId)
	countyIdBytes32 := BytesToByte32(common.Hex2Bytes(countyId))
	res, err := userLocation.CityIdChengShiID(nil, countyIdBytes32)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ""
	}
	return nil, "0x" + common.Bytes2Hex(Bytes32ToBytes(res))
}

// GetChengShiIdByCityIdByte32 获取cityId
func GetChengShiIdByCityIdByte32(countyId [32]byte) (error, string) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ""
	}
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ""
	}
	res, err := userLocation.CityIdChengShiID(nil, countyId)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ""
	}
	return nil, "0x" + common.Bytes2Hex(Bytes32ToBytes(res))
}

// CityInfo 获取区县对应的加密信息
func CityInfo(countyId [32]byte) (error, string) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ""
	}
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ""
	}
	location, err := userLocation.CityInfo(nil, countyId)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ""
	}
	code := strings.Split(location, ",")
	// 查询明文地址
	var areaCode models2.AreaCode
	var whereCondition string
	if len(code) == 2 {
		code0, _ := strconv.ParseInt(code[0], 10, 64)
		code1, _ := strconv.ParseInt(code[1], 10, 64)
		whereCondition = fmt.Sprintf("country_code=%d and city_code=%d", code0, code1)
	} else if len(code) == 3 {
		code0, _ := strconv.ParseInt(code[0], 10, 64)
		code1, _ := strconv.ParseInt(code[1], 10, 64)
		code2, _ := strconv.ParseInt(code[2], 10, 64)
		if code2 == 0 {
			whereCondition = fmt.Sprintf("country_code=%d and city_code=%d", code0, code1)
		} else {
			whereCondition = fmt.Sprintf("country_code=%d and city_code=%d and ad_code=%d", code0, code1, code2)
		}
	}
	err = db.Mysql.Table("area_code").Where(whereCondition).First(&areaCode).Error

	return nil, areaCode.CountryName + " " + areaCode.CityName + " " + areaCode.AreaName
}

// GetCityIdBytes32ByCountyId 获取区县对应的加密信息
func GetCityIdBytes32ByCountyId(countyId string) (error, [32]byte) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, [32]byte{}
	}
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, [32]byte{}
	}
	cityId, err := userLocation.CityIdChengShiID(nil, BytesToByte32(hexutils.HexToBytes(countyId)))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, [32]byte{}
	}

	return nil, cityId
}

// GetCountyIdsByPioneer 获取区县对应的加密信息
func GetCountyIdsByPioneer(pioneer_ string) (error, [][32]byte) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, [][32]byte{}
	}
	cityId, err := userLocation.GetChengShiIdByAddress(nil, common.HexToAddress(pioneer_))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, [][32]byte{}
	}
	countyIds, err := userLocation.GetCountyIdsByChengShiId(nil, cityId)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, [][32]byte{}
	}

	return nil, countyIds
}

// GetChengShiIdByAddress 获取区县对应的加密信息
func GetChengShiIdByAddress(pioneer_ string) (error, [32]byte) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, [32]byte{}
	}
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, [32]byte{}
	}
	cityId, err := userLocation.GetChengShiIdByAddress(nil, common.HexToAddress(pioneer_))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, [32]byte{}
	}

	return nil, cityId
}

//func GetUserLocationRecord() error {
//	Cli := Client(CityNodeConfig)
//	startBlock := GetStartBlock()
//	number, err := Cli.BlockNumber(context.Background())
//	endBlock := startBlock + 999
//	if endBlock > number {
//		return nil
//	}
//	err = GetUserLocationRecordEvent(Cli, int64(startBlock), int64(endBlock))
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return err
//	}
//	return nil
//}

func GetUserLocationRecordEvent(Cli *ethclient.Client, startBlock, endBlock int64) error {
	query := event.BuildQuery(
		common.HexToAddress(CityNodeConfig.UserLocationAddress),
		event.UserLocationRecord,
		big.NewInt(startBlock),
		big.NewInt(endBlock),
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(20))
	logs, err := Cli.FilterLogs(ctx, query)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	cancel()
	abi, _ := intoCityNode.UserLocationMetaData.GetAbi()
	for _, logE := range logs {
		logData, err := abi.Unpack(event.UserLocationRecordEvent.EventName, logE.Data)
		if err != nil {
			log.Logger.Sugar().Error(err)
			return err
		}
		userAddress := logData[0].(common.Address)
		countyId := "0x" + common.Bytes2Hex(Bytes32ToBytes(logData[1].([32]uint8)))
		countyId2 := logData[1].([32]uint8)
		locationEncrypt := logData[2].(string)
		locationCode := utils2.ThreeDesDecrypt(locationEncrypt)
		code := strings.Split(locationCode, ",")
		if len(code) == 2 {
			code = append(code, "0")
		}
		if len(code) < 3 {
			log.Logger.Sugar().Warn("用户位置加密信息解析错误", userAddress.String(), locationCode, locationEncrypt, code)
			continue
		}
		//if code[0] != "0" {
		//	log.Logger.Sugar().Warnln("国外用户位置信息", userAddress.String(), locationCode, locationEncrypt, code)
		//	code[2] = "0"
		//}
		// 国外用户
		codeSecond := code[1]
		codeSecondSlice := strings.Split(codeSecond, "")
		if code[0] != "0" && codeSecondSlice[0] != "0" {
			log.Logger.Sugar().Warn("国外用户位置信息", userAddress.String(), locationCode, locationEncrypt, code)
			code[2] = "0"
		}
		// 容错，国内城市code首位是0的情况
		if codeSecondSlice[0] == "0" {
			// 获取城市code,根据区县code
			var areaCode models2.AreaCode
			whereCondition := fmt.Sprintf("ad_code=%s", code[2])
			err = db.Mysql.Table("area_code").Where(whereCondition).First(&areaCode).Error
			code[1] = fmt.Sprintf("%d", areaCode.CityCode)
		}

		var timestamp int64
		header, err := Cli.HeaderByNumber(context.Background(), big.NewInt(int64(logE.BlockNumber)))
		if err == nil {
			timestamp = int64(header.Time)
		}
		err = InsertUserLocation(userAddress.String(), countyId, code, locationEncrypt, locationCode, countyId2, timestamp)
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
	return nil
}

func InsertUserLocation(userAddress, countyId string, code []string, locationEncrypt, locationCode string, countyIdBytes32 [32]byte, timestamp int64) error {
	InsertUserLocationLock.Lock()
	defer InsertUserLocationLock.Unlock()
	err, cityId := GetChengShiIdByCityIdByte32(countyIdBytes32)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	var userLocation models2.UserLocation
	whereCondition := fmt.Sprintf("user='%s' and city_id='%s' and area_code='%s' and county_id='%s'",
		strings.ToLower(userAddress), strings.ToLower(cityId), locationCode, strings.ToLower(countyId))
	err = db.Mysql.Table("user_location").Where(whereCondition).First(&userLocation).Error
	if err == gorm.ErrRecordNotFound {
		// 查询明文地址
		uri := fmt.Sprintf("https://wallet-api-v2.intowallet.io/api/v1/city_node/geographic_info?city_code=%s&ad_code=%s", code[1], code[2])
		log.Logger.Sugar().Info(uri)
		err, location := utils2.HttpGet(uri)
		if err != nil {
			log.Logger.Sugar().Error(err)
			return err
		}
		var locationInfo LocationInfo
		err = json.Unmarshal(location, &locationInfo)
		if err != nil {
			log.Logger.Sugar().Error(err)
			return err
		}
		locationStr := locationInfo.Data.CountryName + " " + locationInfo.Data.CityName + " " + locationInfo.Data.AreaName
		if locationStr == "" {
			RestoreUserLocation(strings.ToLower(userAddress))
		} else {
			db.Mysql.Model(&models2.UserLocation{}).Create(&models2.UserLocation{
				User:            strings.ToLower(userAddress),
				CountyId:        strings.ToLower(countyId),
				CityId:          strings.ToLower(cityId),
				LocationEncrypt: locationEncrypt,
				Location:        locationStr,
				AreaCode:        locationCode,
				Ctime:           time.Unix(timestamp, 0).Format(LayoutTime),
			})
		}

	}
	return nil
}
