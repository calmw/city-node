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
