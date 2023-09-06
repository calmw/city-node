package blockchain

import (
	intoCityNode "city-node-server/binding"
	"city-node-server/log"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func AdminSetAssessmentCriteria(cityLevel, month, point int64) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	criteria, err := cityPioneer.AdminSetAssessmentCriteria(auth, big.NewInt(cityLevel), big.NewInt(month), big.NewInt(point))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(criteria, err)
}

// AdminSetTOXAddress  管理员设置TOX代币地址
func AdminSetTOXAddress() {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	criteria, err := cityPioneer.AdminSetTOXAddress(auth, common.HexToAddress(CityNodeConfig.ToxAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(criteria, err)
}

// AdminSetCityAddress 管理员设置城市合约地址
func AdminSetCityAddress() {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	criteria, err := cityPioneer.AdminSetCityAddress(auth, common.HexToAddress(CityNodeConfig.CityAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(criteria, err)
}

// AdminSetMiningAddress 管理员设置增加用户合约余额合约地址
func AdminSetMiningAddress() {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	criteria, err := cityPioneer.AdminSetMiningAddress(auth, common.HexToAddress(CityNodeConfig.MiningAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(criteria, err)
}

// AddCityPioneerAdmin 给城市合约、IntoMining设置用户（增加用户合约余额）添加管理员权限
func AddCityPioneerAdmin() {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	cityPioneer, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	res, err := cityPioneer.AddAdmin(auth, common.HexToAddress(CityNodeConfig.CityAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	res, err = cityPioneer.AddAdmin(auth, common.HexToAddress(CityNodeConfig.MiningAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res, err)
}

// DepositSuretyTest 交保证金成为先锋
func DepositSuretyTest(pioneerAddress string) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	criteria, err := cityPioneer.DepositSuretyTest(auth, common.HexToAddress(pioneerAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(criteria, err)
}
