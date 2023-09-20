package blockchain

import (
	intoCityNode "city-node-server/binding"
	"city-node-server/blockchain/event"
	"city-node-server/db"
	"city-node-server/log"
	"city-node-server/models"
	"city-node-server/utils"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
	"math/big"
	"strconv"
	"strings"
	"time"
)

// SetNoRepeatCityIds   城市ID数组重构
func SetNoRepeatCityIds(start, end int64) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
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
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
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

// AdminSetPledgeStakeAddressUserLocation 管理员设置获取用户质押量合约地址
func AdminSetPledgeStakeAddressUserLocation() {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
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
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
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
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
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
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
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
func SetCityMapping(countyId, chengShiId, location string) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
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
	_, err = userLocation.SetCityMapping(auth, cityIdBytes32, chengShiIdBytes32, location)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(err)
}

// GetCityId 获取cityId
func GetCityId(cityIdIndex int64) {
	Cli := Client(CityNodeConfig)
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

// GetCountyIdsByChengShiId 获取城市ID对应的所有区县ID
func GetCountyIdsByChengShiId(cityId string) (error, [][32]byte) {
	Cli := Client(CityNodeConfig)
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, [][32]byte{}
	}
	cityIdBytes32 := BytesToByte32(common.Hex2Bytes(cityId))
	res, err := userLocation.GetCountyIdsByChengShiId(nil, cityIdBytes32)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, [][32]byte{}
	}
	return nil, res
}

// GetCountyIdsByChengShiIdBytes32 获取城市ID对应的所有区县ID
func GetCountyIdsByChengShiIdBytes32(cityId [32]byte) (error, [][32]byte) {
	Cli := Client(CityNodeConfig)
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
	Cli := Client(CityNodeConfig)
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

// GetChengShiIdByCityId 获取cityId
func GetChengShiIdByCityId(countyId string) (error, string) {
	Cli := Client(CityNodeConfig)
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ""
	}
	common.Hex2Bytes(countyId)
	fmt.Println(countyId, 8978)
	fmt.Println(common.Hex2Bytes(countyId))
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
	Cli := Client(CityNodeConfig)
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
	Cli := Client(CityNodeConfig)
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
	var areaCode models.AreaCode
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

func GetUserLocationRecord() error {
	Cli := Client(CityNodeConfig)
	startBlock := GetStartBlock()
	number, err := Cli.BlockNumber(context.Background())
	endBlock := startBlock + 999
	if endBlock > number {
		return nil
	}
	err = GetUserLocationRecordEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	return nil
}

func GetUserLocationRecordEvent(Cli *ethclient.Client, startBlock, endBlock int64) error {
	query := event.BuildQuery(
		common.HexToAddress(CityNodeConfig.UserLocationAddress),
		event.UserLocationRecord,
		big.NewInt(startBlock),
		big.NewInt(endBlock),
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
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
		locationCode := utils.ThreeDesDecrypt(locationEncrypt)
		code := strings.Split(locationCode, ",")
		err = InsertUserLocation(userAddress.String(), countyId, code, locationEncrypt, locationCode, countyId2)
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertUserLocation(userAddress, countyId string, code []string, locationEncrypt, locationCode string, countyIdBytes32 [32]byte) error {
	err, cityId := GetChengShiIdByCityIdByte32(countyIdBytes32)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	// 查询明文地址
	var areaCode models.AreaCode
	var whereCondition, locationInfo string
	var locationType int64
	if len(code) == 2 {
		code0, _ := strconv.ParseInt(code[0], 10, 64)
		code1, _ := strconv.ParseInt(code[1], 10, 64)
		locationType = 1
		whereCondition = fmt.Sprintf("country_code=%d and city_code=%d", code0, code1)
	} else if len(code) == 3 {
		code0, _ := strconv.ParseInt(code[0], 10, 64)
		code1, _ := strconv.ParseInt(code[1], 10, 64)
		code2, _ := strconv.ParseInt(code[2], 10, 64)
		locationType = 2
		if code2 == 0 {
			whereCondition = fmt.Sprintf("country_code=%d and city_code=%d", code0, code1)
		} else {
			whereCondition = fmt.Sprintf("country_code=%d and city_code=%d and ad_code=%d", code0, code1, code2)
		}
	}
	err = db.Mysql.Table("area_code").Where(whereCondition).First(&areaCode).Error
	if err == gorm.ErrRecordNotFound {
		locationInfo = ""
	} else {
		locationInfo = areaCode.CountryName + " " + areaCode.CityName + " " + areaCode.AreaName
	}

	var userLocation models.UserLocation
	whereCondition = fmt.Sprintf("user='%s' and city_id='%s' and location_encrypt='%s' and area_code='%s' and county_id='%s'",
		strings.ToLower(userAddress), strings.ToLower(cityId), locationEncrypt, locationCode, strings.ToLower(countyId))
	err = db.Mysql.Table("user_location").Where(whereCondition).First(&userLocation).Error
	if err == gorm.ErrRecordNotFound {
		db.Mysql.Model(&models.UserLocation{}).Create(&models.UserLocation{
			User:            strings.ToLower(userAddress),
			CountyId:        strings.ToLower(countyId),
			CityId:          strings.ToLower(cityId),
			LocationEncrypt: locationEncrypt,
			Location:        locationInfo,
			AreaCode:        locationCode,
			LocationType:    locationType,
		})
	}
	return nil
}
