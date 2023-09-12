package blockchain

import (
	intoCityNode "city-node-server/binding"
	"city-node-server/log"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
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

//
//// GetCityId 获取cityId
//func GetCityId(cityIdIndex int64) {
//	Cli := Client(CityNodeConfig)
//	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	res, err := userLocation.CityInfo(nil, big.NewInt(cityIdIndex))
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//
//	fmt.Println("0x"+common.Bytes2Hex(Bytes32ToBytes(res)), err)
//}
