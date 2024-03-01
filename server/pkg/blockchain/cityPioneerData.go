package blockchain

import (
	"city-node-server/pkg/binding/intoCityNode"
	"city-node-server/pkg/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type CityPioneerData struct {
	Cli      *ethclient.Client
	Contract *intoCityNode.CityPioneerData
}

func NewCityPioneerData(cli *ethclient.Client) *CityPioneerData {
	cityPioneerData, _ := intoCityNode.NewCityPioneerData(common.HexToAddress(CityNodeConfig.CityPioneerDataAddress), cli)

	return &CityPioneerData{
		Cli:      cli,
		Contract: cityPioneerData,
	}
}

// SuretyUSDT 添加先锋，设置先锋所需要缴纳的USDT数量
func (cpd *CityPioneerData) SuretyUSDT(areaId string, usdtNo int64) error {
	cityPioneerDataContract, err := intoCityNode.NewCityPioneerData(common.HexToAddress(CityNodeConfig.CityPioneerDataAddress), cpd.Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}

	err, auth := GetAuth(cpd.Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}

	E18 := big.NewInt(1e18)
	earnestMoneyBigInt := E18.Mul(E18, big.NewInt(usdtNo))
	_, err = cityPioneerDataContract.AdminSetChengShiSuretyUSDT(auth, ConvertAreaIdAtB(areaId), earnestMoneyBigInt)
	if err != nil {
		return err
	}
	return nil
}

func (cpd *CityPioneerData) AdminSetTOX() error {
	err, auth := GetAuth(cpd.Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	_, err = cpd.Contract.AdminSetTOX(auth, common.HexToAddress(CityNodeConfig.ToxAddress))
	if err != nil {
		return err
	}
	return nil
}

func (cpd *CityPioneerData) AdminSetUSDT() error {

	err, auth := GetAuth(cpd.Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	_, err = cpd.Contract.AdminSetUSDT(auth, common.HexToAddress(CityNodeConfig.USDTAddress))
	if err != nil {
		return err
	}
	return nil
}

func (cpd *CityPioneerData) AdminSetCityAddress() error {

	err, auth := GetAuth(cpd.Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	_, err = cpd.Contract.AdminSetCityAddress(auth, common.HexToAddress(CityNodeConfig.CityAddress))
	if err != nil {
		return err
	}
	return nil
}

func (cpd *CityPioneerData) AdminSetCityPioneerAddress() error {

	err, auth := GetAuth(cpd.Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	_, err = cpd.Contract.AdminSetCityPioneerAddress(auth, common.HexToAddress(CityNodeConfig.CityPioneerAddress))
	if err != nil {
		return err
	}
	return nil
}

func (cpd *CityPioneerData) AdminSetIsGlobalNode(globalNodeAddress string) error {

	err, auth := GetAuth(cpd.Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	_, err = cpd.Contract.AdminSetIsGlobalNode(auth, common.HexToAddress(globalNodeAddress), true)
	if err != nil {
		return err
	}
	return nil
}
