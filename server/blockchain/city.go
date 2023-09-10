package blockchain

import (
	intoCityNode "city-node-server/binding"
	"city-node-server/blockchain/event"
	"city-node-server/db"
	"city-node-server/log"
	"city-node-server/models"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/status-im/keycard-go/hexutils"
	"math/big"
	"strings"
	"time"
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
func AdminSetPioneer(chengShiId, pioneer string) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	if strings.Contains(chengShiId, "0x") {
		chengShiId = strings.ReplaceAll(chengShiId, "0x", "")
	}
	common.Hex2Bytes(chengShiId)
	cityIdBytes32 := BytesToByte32(common.Hex2Bytes(chengShiId))
	fmt.Println("set chengShiId: ", common.Bytes2Hex(Bytes32ToBytes(cityIdBytes32)), "pioneer: ", pioneer)
	res, err := city.AdminSetPioneer(auth, cityIdBytes32, common.HexToAddress(pioneer))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res.Hash(), err)
}

// AdminRemovePioneer 管理员删除城市先锋
func AdminRemovePioneer(chengShiId, pioneer string) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	if strings.Contains(chengShiId, "0x") {
		chengShiId = strings.ReplaceAll(chengShiId, "0x", "")
	}
	common.Hex2Bytes(chengShiId)
	cityIdBytes32 := BytesToByte32(common.Hex2Bytes(chengShiId))
	fmt.Println("remove chengShiId: ", common.Bytes2Hex(Bytes32ToBytes(cityIdBytes32)), "pioneer: ", pioneer)
	res, err := city.AdminRemovePioneer(auth, cityIdBytes32, common.HexToAddress(pioneer))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res.Hash(), err)
}

// PioneerChengShi 查看先锋城市
func PioneerChengShi(pioneer string) {
	Cli := Client(CityNodeConfig)
	//_, auth := GetAuth(Cli)
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	res, err := city.PioneerChengShi(nil, common.HexToAddress(pioneer))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(hexutils.BytesToHex(Bytes32ToBytes(res)), err)
}

// GetPioneerCityNumber 查看先锋城市数量
func GetPioneerCityNumber() {
	Cli := Client(CityNodeConfig)
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	res, err := city.GetPioneerCityNumber(nil)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res.String(), err)
}

// PioneerChengShiId 查看先锋城市数量
func PioneerChengShiId(chengShiIndex int64) {
	Cli := Client(CityNodeConfig)
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	res, err := city.PioneerChengShiIds(nil, big.NewInt(chengShiIndex))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(hexutils.BytesToHex(Bytes32ToBytes(res)), err)
}

// AdminPopPioneerChengShiId 删除城市ID
func AdminPopPioneerChengShiId() {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	_, err = city.AdminPopPioneerChengShiId(auth)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(err)
}

// AdminSetDelegate 管理员设置用户增加或减少质押量
func AdminSetDelegate(userAddress string, amount, setType int64) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	E18 := big.NewInt(1e18)
	amountRes := E18.Mul(E18, big.NewInt(amount))
	res, err := city.AdminSetDelegate(auth, common.HexToAddress(userAddress), amountRes, big.NewInt(setType))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res.Hash(), setType, err)
}

// AddCityAdmin 给城市先锋合约、用户定位合约、设置质押量合约添加管理员权限
func AddCityAdmin() {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	_, err = city.AddAdmin(auth, common.HexToAddress(CityNodeConfig.CityPioneerAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	_, err = city.AddAdmin(auth, common.HexToAddress(CityNodeConfig.UserLocationAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	_, err = city.AddAdmin(auth, common.HexToAddress(CityNodeConfig.SetDelegateAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	_, err = city.AddAdmin(auth, common.HexToAddress(CityNodeConfig.RechargeWeightAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
}

// AdminSetChengShiLevelAndSurety 管理员设置先锋计划，城市等级以及该等级城市所需缴纳的保证金数额
func AdminSetChengShiLevelAndSurety(cityId string, level, earnestMoney int64) {
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
	fmt.Println("chengShiId: ", common.Bytes2Hex(Bytes32ToBytes(cityIdBytes32)), "Surety: ", earnestMoneyBigInt.String())
	res, err := city.AdminSetChengShiLevelAndSurety(auth, cityIdBytes32, big.NewInt(level), earnestMoneyBigInt)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res.Hash(), err)
}

// AdminEditSurety 管理员修改先锋计划，城市等级以及该等级城市所需缴纳的保证金数额
func AdminEditSurety(cityId string, level, earnestMoney int64) {
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
	res, err := city.AdminEditSurety(auth, cityIdBytes32, big.NewInt(level), earnestMoneyBigInt)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res, err)
}

func GetIncreaseCityDelegateEvent(Cli *ethclient.Client, startBlock, endBlock int64) error {
	query := event.BuildQuery(
		common.HexToAddress("0xebD06631510A66968f0379A4deB896d3eE7DD6ED"),
		event.IncreaseCityDelegate,
		big.NewInt(startBlock),
		big.NewInt(endBlock),
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	logs, err := Cli.FilterLogs(ctx, query)
	fmt.Println(len(logs))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	cancel()
	abi, _ := intoCityNode.CityMetaData.GetAbi()
	for _, logE := range logs {
		logData, err := abi.Unpack(event.IncreaseCityDelegateEvent.EventName, logE.Data)
		if err != nil {
			log.Logger.Sugar().Error(err)
		}
		cityId := "0x" + common.Bytes2Hex(Bytes32ToBytes(logData[0].([32]uint8)))
		amount := logData[1].(*big.Int).String()
		CreateAdminSetDelegate(models.AdminSetDelegate{
			CityId:  cityId,
			Amount:  amount,
			SetType: 1,
		})
	}
	return nil
}

func GetDecreaseCityDelegateEvent(Cli *ethclient.Client, startBlock, endBlock int64) error {
	query := event.BuildQuery(common.HexToAddress(
		"0xebD06631510A66968f0379A4deB896d3eE7DD6ED"),
		event.DecreaseCityDelegate,
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
	abi, _ := intoCityNode.CityMetaData.GetAbi()
	for _, logE := range logs {
		logData, err := abi.Unpack(event.DecreaseCityDelegateEvent.EventName, logE.Data)
		if err != nil {
			log.Logger.Sugar().Error(err)
		}
		cityId := "0x" + common.Bytes2Hex(Bytes32ToBytes(logData[0].([32]uint8)))
		amount := logData[1].(*big.Int).String()
		CreateAdminSetDelegate(models.AdminSetDelegate{
			CityId:  cityId,
			Amount:  amount,
			SetType: 2,
		})
	}
	return nil
}

func CreateAdminSetDelegate(adminSetDelegate models.AdminSetDelegate) {
	db.Mysql.Table("admin_set_delegate").Create(&adminSetDelegate)
}
