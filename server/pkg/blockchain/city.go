package blockchain

import (
	"city-node-server/pkg/binding/intoCityNode"
	"city-node-server/pkg/db"
	"city-node-server/pkg/log"
	"city-node-server/pkg/models"
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/status-im/keycard-go/hexutils"
	"github.com/xjieinfo/xjgo/xjcore/xjexcel"
	"gorm.io/gorm"
	"math/big"
	"strings"
	"time"
)

type City struct {
	Cli      *ethclient.Client
	Contract *intoCityNode.City
}

func NewCity(cli *ethclient.Client) *City {
	var city *intoCityNode.City
	var err error
	for i := 0; i < 10; i++ {
		city, err = intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), cli)
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}
	if city == nil {
		panic(err)
	}

	return &City{
		Cli:      cli,
		Contract: city,
	}
}

func (c *City) AdminSetChengShiLevelAndSurety(chengShiId string, level int64, surety int64) error {

	err, auth := GetAuth(c.Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}

	E18 := big.NewInt(1e18)
	earnestMoneyBigInt := E18.Mul(E18, big.NewInt(surety))

	_, err = c.Contract.AdminSetChengShiLevelAndSurety(auth, ConvertAreaIdAtB(chengShiId), big.NewInt(level), earnestMoneyBigInt)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	return nil
}

func (c *City) AdminSetPioneer(areaId, pioneer string, pioneerBatch, pioneerType int64) error {
	_, auth := GetAuth(c.Cli)
	_, err := c.Contract.AdminSetPioneer(auth, ConvertAreaIdAtB(areaId), common.HexToAddress(pioneer), big.NewInt(pioneerBatch), big.NewInt(pioneerType))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	fmt.Println("set areaId: ", areaId, "pioneer: ", pioneer)
	return nil
}

func (c *City) AdminRemovePioneer(chengShiId, pioneer string) error {
	err, auth := GetAuth(c.Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}

	if strings.Contains(chengShiId, "0x") {
		chengShiId = strings.ReplaceAll(chengShiId, "0x", "")
	}
	_, err = c.Contract.AdminRemovePioneer(auth, ConvertAreaIdAtB(chengShiId), common.HexToAddress(pioneer))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	return nil
}

func (c *City) PioneerCity(pioneer string) (string, error) {

	AreaIdBytes32, err := c.Contract.PioneerCity(nil, common.HexToAddress(pioneer))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return "", err
	}

	return strings.ToLower("0x" + hexutils.BytesToHex(Bytes32ToBytes(AreaIdBytes32))), nil
}

func (c *City) ChengShiPioneer(areaId string) (string, error) {
	if strings.Contains(areaId, "0x") {
		areaId = strings.ReplaceAll(areaId, "0x", "")
	}
	address, err := c.Contract.ChengShiPioneer(nil, ConvertAreaIdAtB(areaId))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return "", err
	}

	return address.Hex(), nil
}

// AdminSetCityPioneerAddress 管理员设置先锋计划合约地址
func AdminSetCityPioneerAddress() {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	_, auth := GetAuth(Cli)
	city, err := intoCityNode.NewCity(common.Address(common.Address(common.HexToAddress(CityNodeConfig.CityAddress))), Cli)
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
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
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

// AdminSetFoundsAddress 管理员设置获取过去15天社交基金平均值的合约地址
func AdminSetFoundsAddress() {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	_, auth := GetAuth(Cli)
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	res, err := city.AdminSetFoundsAddress(auth, common.HexToAddress(CityNodeConfig.GetFoundsAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res, err)
}

// AdminSetSecondsPerDay 管理员设置获取过去15天社交基金平均值的合约地址
func AdminSetSecondsPerDay(secondsPerDay int64) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	_, auth := GetAuth(Cli)
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	res, err := city.AdminSetSecondsPerDay(auth, big.NewInt(secondsPerDay))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res, err)
}

// AdminSetPioneer 管理员设置城市先锋
func AdminSetPioneer(chengShiId, pioneer string, pioneerBatch, pioneerType int64) error {
	//func AdminSetPioneer(chengShiId, pioneer string) error {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	_, auth := GetAuth(Cli)
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}

	res, err := city.AdminSetPioneer(auth, ConvertAreaIdAtB(chengShiId), common.HexToAddress(pioneer), big.NewInt(pioneerBatch), big.NewInt(pioneerType))
	//res, err := city.AdminSetPioneer(auth, ConvertAreaIdAtB(chengShiId), common.HexToAddress(pioneer))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	fmt.Println(res.Hash().String(), err, 567)
	return nil
}

func AdminSetAuthAddress() {
	err, Cli := Client(CityNodeConfig)
	number, err := Cli.BlockNumber(context.Background())

	fmt.Println(number, err, 12233445)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	_, auth := GetAuth(Cli)
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	_, err = city.AdminSetAuthAddress(auth, common.HexToAddress(CityNodeConfig.AuthAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(err)
}

func AdminSetWithdrawLimitAddress() {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	_, auth := GetAuth(Cli)
	fmt.Println(CityNodeConfig.CityAddress, CityNodeConfig.WithdrawLimitAddress)
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	_, err = city.AdminSetWithdrawLimitAddress(auth, common.HexToAddress(CityNodeConfig.WithdrawLimitAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(err)
}

// GetFifteenDayAverageFounds 获取前15天社交基金平均值
func GetFifteenDayAverageFounds() {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	res, err := city.GetFifteenDayAverageFounds(nil)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res.String())
}

// AdminRemovePioneer 管理员删除城市先锋
func AdminRemovePioneer(chengShiId, pioneer string) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
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
	fmt.Println("remove chengShiId: ", common.Bytes2Hex(Bytes32ToBytes(cityIdBytes32)), "pioneer: ", pioneer, city)
	res, err := city.AdminRemovePioneer(auth, cityIdBytes32, common.HexToAddress(pioneer))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res, err, 123)
}

// PioneerChengShi 查看先锋城市
func PioneerChengShi(pioneer string) (error, [32]byte) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, [32]byte{}
	}
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, [32]byte{}
	}
	res, err := city.PioneerChengShi(nil, common.HexToAddress(pioneer))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, [32]byte{}
	}
	fmt.Println(hexutils.BytesToHex(Bytes32ToBytes(res)), err)
	return nil, res
}

// ChengShiLevel 查看先锋城市等级
func ChengShiLevel(cityId [32]byte) (error, int64) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, 0
	}
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, 0
	}
	res, err := city.ChengShiLevel(nil, cityId)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, 0
	}
	//fmt.Println(hexutils.BytesToHex(Bytes32ToBytes(res)), err)
	return nil, res.Int64()
}

// CountyNewlyPioneerDelegateRecord 根据区县ID和天查询新增质押量
func CountyNewlyPioneerDelegateRecord(countyId [32]byte, day int64) (error, *big.Int) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}
	//_, auth := GetAuth(Cli)
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, big.NewInt(0)
	}
	amount, err := city.CountyNewlyPioneerDelegateRecord(nil, countyId, big.NewInt(day))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, big.NewInt(0)
	}
	return nil, amount
}

// RechargeDailyWeightRecord 根据区县ID和天查询新增充值权重
func RechargeDailyWeightRecord(countyId [32]byte, day int64) (error, *big.Int) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}
	//_, auth := GetAuth(Cli)
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, big.NewInt(0)
	}
	amount, err := city.RechargeDailyWeightRecord(nil, countyId, big.NewInt(day))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, big.NewInt(0)
	}
	return nil, amount
}

// GetPioneerCityNumber 查看先锋城市数量
func GetPioneerCityNumber() {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
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
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
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
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	E18 := big.NewInt(1e18)
	amountRes := E18.Mul(E18, big.NewInt(amount))
	_, err = city.AdminSetDelegate(auth, common.HexToAddress(userAddress), amountRes, big.NewInt(setType))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	if setType == 1 {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), userAddress, "增加质押量:", amount, err)
	} else {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), userAddress, "减少质押量:", amount, err)
	}
}

// AdminSetRechargeAmount 管理员设置用户充值量
func AdminSetRechargeAmount(userAddress string, amount int64) {
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
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	E18 := big.NewInt(1e18)
	amountRes := E18.Mul(E18, big.NewInt(amount))
	_, err = city.AdminSetRechargeAmount(auth, common.HexToAddress(userAddress), amountRes)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), userAddress, "增加充值量:", amount, err)
}

// AdminSetSecondsPerDayCity 管理员设置每天秒数，用于测试
func AdminSetSecondsPerDayCity(seconds int64) {
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
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	res, err := city.AdminSetSecondsPerDay(auth, big.NewInt(seconds))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res, err)
}

// AddCityAdmin 给城市先锋合约、用户定位合约、设置质押量合约添加管理员权限
func AddCityAdmin() {
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
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	_, err = city.AddAdmin(auth, common.HexToAddress("0x94b627F4F829Ac5E97fDc556B5BEeeFf9beF417e"))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(err, 33)

	//_, err = city.AddAdmin(auth, common.HexToAddress(CityNodeConfig.CityPioneerAddress))
	//if err != nil {
	//	log.Logger.Sugar().Error(err)
	//	return
	//}
	//
	//_, err = city.AddAdmin(auth, common.HexToAddress(CityNodeConfig.UserLocationAddress))
	//if err != nil {
	//	log.Logger.Sugar().Error(err)
	//	return
	//}
	//
	//_, err = city.AddAdmin(auth, common.HexToAddress(CityNodeConfig.SetDelegateAddress))
	//if err != nil {
	//	log.Logger.Sugar().Error(err)
	//	return
	//}
	//
	//_, err = city.AddAdmin(auth, common.HexToAddress(CityNodeConfig.RechargeWeightAddress))
	//if err != nil {
	//	log.Logger.Sugar().Error(err)
	//	return
	//}
}

// AdminSetChengShiLevelAndSurety 管理员设置先锋计划，城市等级以及该等级城市所需缴纳的保证金数额
func AdminSetChengShiLevelAndSurety(cityId string, level, earnestMoney int64) error {
	err, Cli := Client(CityNodeConfig)
	fmt.Println(CityNodeConfig.RPC, CityNodeConfig.CityAddress)

	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	err, auth := GetAuth(Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
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
		return err
	}
	fmt.Println(res.Hash().String(), err, 123)
	return nil
}

// AdminEditSurety 管理员修改先锋计划，城市等级以及该等级城市所需缴纳的保证金数额
func AdminEditSurety(cityId string, level, earnestMoney int64) {
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

// PioneerCity 获取先锋对应的城市ID
func PioneerCity(pioneerAddress string) (error, string) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ""
	}
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ""
	}
	res, err := city.PioneerChengShi(nil, common.HexToAddress(pioneerAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ""
	}

	return nil, string(Bytes32ToBytes(res))
}

func GetAuth2(nonce_ int64) (error, *bind.TransactOpts) {
	//privateKeyEcdsa, err := crypto.HexToECDSA(CityNodeConfig.PrivateKey)
	privateKeyEcdsa, err := crypto.HexToECDSA("73b283875c5c44836e6e377b9285fd00778cdea8da3a1a102d5c0f56bbc40974")
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, big.NewInt(CityNodeConfig.ChainId))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}
	//nonce, _ := cli.NonceAt(context.Background(), common.HexToAddress("0x89876D12A4cB4d19957cEBE3663EA485E05fD3f2"), nil)
	//gasLimit := uint64(21000)
	return nil, &bind.TransactOpts{
		From: auth.From,
		//Nonce:     big.NewInt(int64(nonce) + 1),
		Nonce: big.NewInt(nonce_),
		//Nonce:     nil,
		Signer:    auth.Signer, // Method to use for signing the transaction (mandatory)
		Value:     big.NewInt(0),
		GasPrice:  nil,
		GasFeeCap: nil,
		GasTipCap: nil,
		GasLimit:  0,
		Context:   context.Background(),
		NoSend:    false, // Do all transact steps but do not send the transaction
	}
}

// TriggerAllPioneerTask 触发所有先锋分红和考核
func TriggerAllPioneerTask() {
	Cli, err := ethclient.Dial(CityNodeConfig.RPC)
	if err != nil {
		log.Logger.Sugar().Error("dail failed")
		return
	}

	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)

	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	var pioneerNumber *big.Int
	for {
		pioneerNumber, err = cityPioneer.GetPioneerNumber(nil)
		if err == nil || pioneerNumber != nil {
			log.Logger.Sugar().Error(err)
			break
		}
		time.Sleep(time.Second)
	}

	for i := 0; i < int(pioneerNumber.Int64()); i++ {
		pioneer, err := cityPioneer.Pioneers(nil, big.NewInt(int64(i)))
		if err != nil {
			log.Logger.Sugar().Error(err)
			continue
		}
		isPioneer, err := cityPioneer.IsPioneer(nil, pioneer)
		if err != nil {
			log.Logger.Sugar().Error(err)
			continue
		}
		if !isPioneer {
			log.Logger.Sugar().Error(i, pioneer, ", 未交保证金")
			continue
		}
		failed, err := cityPioneer.FailedAt(nil, pioneer)
		if err != nil {
			log.Logger.Sugar().Error(err)
			continue
		} else {
			fmt.Println(i, pioneer, ", 考核：", failed.Int64() == 0)
		}
		//AdminSetCheckPioneerDailyStatus(pioneer.String(), int64(19643), false) // 重置定时任务的执行状态
		GetPioneerTaskStatus(pioneer.String())

		day, err := cityPioneer.GetDay(nil)
		if err != nil {
			log.Logger.Sugar().Error(err)
			continue
		}
		status, err := cityPioneer.CheckPioneerDailyStatus(nil, day, pioneer)
		if err != nil {
			log.Logger.Sugar().Error(err)
			continue
		}
		if status {
			SetPioneerTaskStatus(pioneer.String())
			log.Logger.Sugar().Error("已经执行成功，跳过")
			continue
		} else {
			for k := 0; k < 10; k++ {
				err, auth := GetPioneerDailyTaskAuth(Cli)
				if err != nil {
					log.Logger.Sugar().Error(err)
					continue
				}
				_, err = city.PioneerDailyTask(auth, pioneer)
				if err != nil {
					log.Logger.Sugar().Error("定时任务执行失败：", err)
					time.Sleep(time.Second * 2)
					continue
				} else {
					log.Logger.Sugar().Info("定时任务执行成功")
					SetPioneerTaskStatus(pioneer.String())
					break
				}
			}

		}
	}
}

func GetAllPioneer() {
	Cli, err := ethclient.Dial(CityNodeConfig.RPC)
	if err != nil {
		log.Logger.Sugar().Error("dail failed")
		return
	}
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	cityPioneerData, err := intoCityNode.NewCityPioneerData(common.HexToAddress(CityNodeConfig.CityPioneerDataAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	appraise, err := intoCityNode.NewAppraise(common.HexToAddress(CityNodeConfig.AppraiseAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	var pioneerNumber *big.Int
	for {
		pioneerNumber, err = cityPioneer.GetPioneerNumber(nil)
		if err == nil || pioneerNumber == nil {
			break
		}
		log.Logger.Sugar().Error(err)
		time.Sleep(time.Second)
	}

	for i := 0; i < int(pioneerNumber.Int64()); i++ {
		pioneer, err := cityPioneer.Pioneers(nil, big.NewInt(int64(i)))
		if err != nil {
			log.Logger.Sugar().Error(err)
			continue
		}
		isPioneer, err := cityPioneer.IsPioneer(nil, pioneer)
		if err != nil {
			log.Logger.Sugar().Error(err)
			continue
		}
		if !isPioneer {
			continue
		}
		pioneerInfo := models.Pioneer{}
		create := false
		err = db.Mysql.Model(models.Pioneer{}).Where("pioneer=?", strings.ToLower(pioneer.String())).Order("id desc").First(&pioneerInfo).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Logger.Sugar().Error(err)
			continue
		}
		/// 查询状态
		// 查询批次
		batch, err := appraise.PioneerBatch(nil, common.HexToAddress(pioneer.String()))
		if err != nil {
			log.Logger.Sugar().Error(err)
			continue
		}
		// 获取先锋类型
		pioneerType, err := appraise.PioneerType(nil, common.HexToAddress(pioneer.String()))
		if err != nil {
			log.Logger.Sugar().Error(err)
			continue
		}
		// 查询地区ID
		areaID, err := city.PioneerChengShi(nil, common.HexToAddress(pioneer.String()))
		if err != nil {
			log.Logger.Sugar().Error(err)
			continue
		}

		// 查询考核状态
		info, err := cityPioneer.PioneerInfo(nil, common.HexToAddress(pioneer.String()))
		if err != nil {
			log.Logger.Sugar().Error(err)
			return
		}
		failedDelegate := ""
		failedAt := ""
		if !info.AssessmentMonthStatus {
			failedDelegateRes, _ := cityPioneer.FailedDelegate(nil, areaID)
			failedDelegate = failedDelegateRes.String()
			failedAtRes, _ := cityPioneer.FailedAt(nil, common.HexToAddress(pioneer.String()))
			failedAt = failedAtRes.String()
		}
		// 查询到期时间
		endTimestamp := info.Ctime.Int64() + 86400*(180-1)
		endTime := time.Unix(endTimestamp, 0).Format("2006-01-02")
		isOverTime := int64(0)
		if time.Now().Unix()/86400 > endTimestamp/86400 {
			isOverTime = 1
		}

		//fmt.Println(errors.Is(err, gorm.ErrRecordNotFound), endTime, pioneerInfo.EndTime, pioneerInfo.Pioneer)
		fmt.Println(errors.Is(err, gorm.ErrRecordNotFound), endTime, pioneerInfo.EndTime, pioneerInfo.Pioneer)
		if errors.Is(err, gorm.ErrRecordNotFound) || pioneerInfo.EndTime != endTime {
			create = true
		}
		var upgrade bool
		if (create == false && (isOverTime == 1 && pioneerInfo.IsOverTime == 0)) || (failedAt != "" && create == false) {
			upgrade = true
		}
		// 获取先锋绑定城市信息
		areaIdStr := strings.ToLower("0x" + hexutils.BytesToHex(Bytes32ToBytes(areaID)))
		areaLevel, err := city.ChengShiLevel(nil, areaID)
		if err != nil {
			log.Logger.Sugar().Error(err)
			continue
		}
		suretyUsdt := int64(0)
		suretyTox := int64(0)
		if batch.Int64() > 3 {
			tox, err := cityPioneerData.SuretyDepositTOX(nil, common.HexToAddress(pioneer.String()))
			if err != nil {
				log.Logger.Sugar().Error(err)
				return
			}
			suretyTox = tox.Div(tox, big.NewInt(1e18)).Int64()
			usdt, err := cityPioneerData.SuretyDepositUSDT(nil, common.HexToAddress(pioneer.String()))
			if err != nil {
				log.Logger.Sugar().Error(err)
				return
			}
			suretyUsdt = usdt.Div(usdt, big.NewInt(1e18)).Int64()
		} else {
			if areaLevel.Int64() == 1 {
				suretyTox = 100000
			} else if areaLevel.Int64() == 2 {
				suretyTox = 60000
			} else if areaLevel.Int64() == 3 {
				suretyTox = 40000
			}
		}

		// 根据城市ID查询location
		local := models.UserLocation{}
		var whereCondition string
		if pioneerType.Int64() == 1 {
			whereCondition = fmt.Sprintf("county_id='%s'", areaIdStr)
		} else {
			whereCondition = fmt.Sprintf("city_id='%s'", areaIdStr)
		}
		db.Mysql.Model(models.UserLocation{}).Where(whereCondition).First(&local)
		if create {
			// 不存在创建
			db.Mysql.Model(models.Pioneer{}).Create(&models.Pioneer{
				AreaId:         areaIdStr,
				Location:       local.Location,
				Pioneer:        strings.ToLower(pioneer.String()),
				FailedDelegate: failedDelegate,
				FailedAt:       failedAt,
				AreaLevel:      areaLevel.Int64(),
				IsAreaNode:     pioneerType.Int64(),
				PioneerBatch:   batch.Int64(),
				EndTime:        endTime,
				IsOverTime:     isOverTime,
				SuretyUsdt:     suretyUsdt,
				SuretyTox:      suretyTox,
			})
		}
		if upgrade {
			// 未到期，未考核失败的更新
			//if failedAt != "" || isOverTime > 0 {
			//	continue
			//}
			db.Mysql.Model(models.Pioneer{}).Where("pioneer=? and is_over_time=0", strings.ToLower(pioneer.String())).Updates(&models.Pioneer{
				AreaId:         areaIdStr,
				Location:       local.Location,
				Pioneer:        strings.ToLower(pioneer.String()),
				FailedDelegate: failedDelegate,
				FailedAt:       failedAt,
				AreaLevel:      areaLevel.Int64(),
				IsAreaNode:     pioneerType.Int64(),
				PioneerBatch:   batch.Int64(),
				EndTime:        endTime,
				IsOverTime:     isOverTime,
				SuretyUsdt:     suretyUsdt,
				SuretyTox:      suretyTox,
			})
		}
	}
}

// TriggerAllPioneerTaskTestNet 触发所有先锋分红和考核
func TriggerAllPioneerTaskTestNet() {

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
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	var pioneerNumber *big.Int
	for {
		pioneerNumber, err = cityPioneer.GetPioneerNumber(nil)
		if err == nil && pioneerNumber != nil {
			break
		}
	}
	fmt.Println(pioneerNumber, "-----", err, pioneerNumber)
	for i := 0; i < int(pioneerNumber.Int64()); i++ {
		pioneer, err := cityPioneer.Pioneers(nil, big.NewInt(int64(i)))
		res, err := city.PioneerDailyTask(auth, pioneer)
		if err == nil {
			if res != nil {
				fmt.Println(pioneer.String(), i, err, res.Hash())
			} else {
				fmt.Println(pioneer.String(), i, "res nil:", res)
			}
		} else {
			fmt.Println(pioneer.String(), i, "error:", err)
		}
	}
}

// 0xcD1f731A1529d5F8e8f8cA94dF6092B680C88e2E
// CityRechargeTotal 获取区县对应的累计充值权重
func CityRechargeTotal(countyId [32]byte) (error, *big.Int) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}
	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}
	res, err := city.CityRechargeTotal(nil, countyId)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}

	return nil, res
}

type UserLocationExcel struct {
	Id      int    `gorm:"column:id;primaryKey" excel:"column:B;desc:ID;width:30"`
	User    string `json:"user" gorm:"column:user" excel:"column:C;desc:用户地址;width:30"`
	Country string `json:"country" gorm:"column:country" excel:"column:D;desc:国家;width:30"`
	City    string `json:"city" gorm:"column:city" excel:"column:E;desc:城市;width:30"`
	County  string `json:"county" gorm:"column:county" excel:"column:F;desc:区县;width:30"`
	TxHash  string `json:"tx_hash" gorm:"column:tx_hash" excel:"column:G;desc:交易hash;width:30"`
	Ctime   string `json:"ctime" gorm:"column:ctime" excel:"column:H;desc:d定位时间;width:30"`
}

func GetUserLocation() {
	var userLocationExcels []UserLocationExcel

	for i := 0; i < 240; i++ {
		for {
			var locations []models.UserLocation
			err := db.Mysql.Model(models.UserLocation{}).Where("id>? and id<=?", i*1000, (i+1)*1000).Find(&locations).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				fmt.Println(i, "无数据")
				break
			}
			if (err != nil) && (!errors.Is(err, gorm.ErrRecordNotFound)) {
				fmt.Println(i, "重试")
				continue
			}
			for _, l := range locations {
				userLocationExcels = append(userLocationExcels, UserLocationExcel{
					Id:      l.Id,
					User:    l.User,
					Country: l.Country,
					City:    l.City,
					County:  l.County,
					TxHash:  l.TxHash,
					Ctime:   l.Ctime,
				})
			}
			fmt.Println(i)
			break
		}
	}

	f := xjexcel.ListToExcel(userLocationExcels, "用户定位信息", "详情")
	fileName := fmt.Sprintf("./用户定位信息.xlsx")
	f.SaveAs(fileName)
}
