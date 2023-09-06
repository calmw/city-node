package blockchain

import (
	intoCityNode "city-node-server/binding"
	"city-node-server/log"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
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
