package blockchain

import (
	intoCityNode "city-node-server/binding"
	"city-node-server/log"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

// AdminSetCityPioneerAddress 管理员设置先锋计划合约地址
func AdminSetCityPioneerAddress() {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	res, err := city.AdminSetCityPioneerAddress(auth, common.HexToAddress(CityNodeConfig.CityPioneerAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res, err)
}

// AdminSetUserLocationAddress 管理员设置用户定位合约地址
func AdminSetUserLocationAddress() {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	res, err := city.AdminSetUserLocationAddress(auth, common.HexToAddress(CityNodeConfig.UserLocationAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res, err)
}

// AdminSetPioneer 管理员设置城市先锋
func AdminSetPioneer(cityId, pioneer string) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	if strings.Contains(cityId, "0x") {
		cityId = strings.ReplaceAll(cityId, "0x", "")
	}
	common.Hex2Bytes(cityId)
	cityIdBytes32 := BytesToByte32(common.Hex2Bytes(cityId))
	fmt.Println("cityId: ", common.Bytes2Hex(Bytes32ToBytes(cityIdBytes32)))
	res, err := city.AdminSetPioneer(auth, cityIdBytes32, common.HexToAddress(pioneer))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res, err)
}

// AdminSetCityLevel 管理员设置城市等级和保证金
func AdminSetCityLevel(cityId string, level, earnestMoney int64) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	if strings.Contains(cityId, "0x") {
		cityId = strings.ReplaceAll(cityId, "0x", "")
	}
	common.Hex2Bytes(cityId)
	cityIdBytes32 := BytesToByte32(common.Hex2Bytes(cityId))
	E18 := big.NewInt(1e18)
	earnestMoneyBigInt := E18.Mul(E18, big.NewInt(earnestMoney))
	fmt.Println("cityId: ", common.Bytes2Hex(Bytes32ToBytes(cityIdBytes32)), "earnestMoney: ", earnestMoneyBigInt.String())
	res, err := city.AdminSetCityLevel(auth, cityIdBytes32, big.NewInt(level), earnestMoneyBigInt)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res, err)
}
