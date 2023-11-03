package blockchain

import (
	intoCityNode2 "city-node-server/binding/intoCityNode"
	"city-node-server/blockchain/event"
	"city-node-server/db"
	"city-node-server/log"
	"city-node-server/models"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"github.com/status-im/keycard-go/hexutils"
	"gorm.io/gorm"
	"math/big"
	"strings"
	"time"
)

// AdminSetCityPioneerAddress 管理员设置先锋计划合约地址
func AdminSetCityPioneerAddress() {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	_, auth := GetAuth(Cli)
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
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
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
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
	_, auth := GetAuth2(Cli)
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
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

// AdminSetPioneer 管理员设置城市先锋
func AdminSetPioneer(chengShiId, pioneer string) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	_, auth := GetAuth(Cli)
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
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

// GetFifteenDayAverageFounds 获取前15天社交基金平均值
func GetFifteenDayAverageFounds() {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
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
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
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
func PioneerChengShi(pioneer string) (error, [32]byte) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, [32]byte{}
	}
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, [32]byte{}
	}
	res, err := city.PioneerChengShi(nil, common.HexToAddress(pioneer))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, [32]byte{}
	}
	//fmt.Println(hexutils.BytesToHex(Bytes32ToBytes(res)), err)
	return nil, res
}

// ChengShiLevel 查看先锋城市等级
func ChengShiLevel(cityId [32]byte) (error, int64) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, 0
	}
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
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
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
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
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
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
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
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
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
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
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
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
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
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
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
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
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	res, err := city.AdminSetSecondsPerDay(auth, big.NewInt(seconds))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res.Hash(), err)
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
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	_, err = city.AddAdmin(auth, common.HexToAddress("0x73A8f49C231ffBF9D190C623361c332bEc59F95A"))
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
func AdminSetChengShiLevelAndSurety(cityId string, level, earnestMoney int64) {
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
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
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
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
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
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
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

func GetAuth2(cli *ethclient.Client) (error, *bind.TransactOpts) {
	//privateKeyEcdsa, err := crypto.HexToECDSA(CityNodeConfig.PrivateKey)
	privateKeyEcdsa, err := crypto.HexToECDSA("5c21103ec19752593b49662f539f410fa1a2efccbef21f304ea73b94a84be045")
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, big.NewInt(CityNodeConfig.ChainId))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}
	nonce, _ := cli.NonceAt(context.Background(), common.HexToAddress("0x94b627F4F829Ac5E97fDc556B5BEeeFf9beF417e"), nil)
	//gasLimit := uint64(21000)
	return nil, &bind.TransactOpts{
		From:      auth.From,
		Nonce:     big.NewInt(int64(nonce) + 1),
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
	Cli, err := ethclient.Dial("https://rpc-sen.matchscan.io")
	if err != nil {
		log.Logger.Sugar().Error("dail failed")
	}

	err, auth := GetAuth2(Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)

	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	pioneerNumber, err := cityPioneer.GetPioneerNumber(nil)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	for i := 0; i < int(pioneerNumber.Int64()); i++ {
		time.Sleep(time.Second * 5)
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
		fmt.Println(pioneer.String(), err, day, i, status)
		if err != nil {
			log.Logger.Sugar().Error(err)
			continue
		}
		if status {
			SetPioneerTaskStatus(pioneer.String())
			log.Logger.Sugar().Error("已经执行成功，跳过")
			continue
		} else {
			_, err = city.PioneerDailyTask(auth, pioneer)
			if err != nil {
				log.Logger.Sugar().Error("定时任务执行失败：", err)
				continue
			} else {
				log.Logger.Sugar().Info("定时任务执行成功")
				SetPioneerTaskStatus(pioneer.String())
			}
		}
	}
}

func GetAllPioneer() {
	Cli, err := ethclient.Dial("https://rpc-sen.matchscan.io")
	if err != nil {
		log.Logger.Sugar().Error("dail failed")
	}
	cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)

	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	pioneerNumber, err := cityPioneer.GetPioneerNumber(nil)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
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
		if isPioneer {
			pioneerInfo := models.Pioneer{}
			err = db.Mysql.Model(models.Pioneer{}).Where("pioneer=?", strings.ToLower(pioneer.String())).First(&pioneerInfo).Error
			if err == gorm.ErrRecordNotFound {
				city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
				if err != nil {
					log.Logger.Sugar().Error(err)
					continue
				}
				res, err := city.PioneerChengShi(nil, common.HexToAddress(pioneer.String()))
				if err != nil {
					log.Logger.Sugar().Error(err)
					continue
				}
				cityId := strings.ToLower("0x" + hexutils.BytesToHex(Bytes32ToBytes(res)))
				cityLevel, err := city.ChengShiLevel(nil, res)
				if err != nil {
					log.Logger.Sugar().Error(err)
					continue
				}
				// 根据城市ID查询location
				local := models.UserLocation{}

				db.Mysql.Model(models.UserLocation{}).Where("city_id=?", cityId).First(&local)

				db.Mysql.Model(models.Pioneer{}).Create(&models.Pioneer{
					CityId:    strings.ToLower("0x" + hexutils.BytesToHex(Bytes32ToBytes(res))),
					Location:  local.Location,
					CityLevel: cityLevel.Int64(),
					Pioneer:   strings.ToLower(pioneer.String()),
				})
			}
		}
	}
}

// TriggerAllPioneerTaskTestNet 触发所有先锋分红和考核
func TriggerAllPioneerTaskTestNet() {
	defer func() {
		if r := recover(); r != nil {
			log.Logger.Sugar().Error("recovery")
		}
	}()
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
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	pioneerNumber, err := cityPioneer.GetPioneerNumber(nil)
	fmt.Println(pioneerNumber, "-----")
	for i := 0; i < int(pioneerNumber.Int64()); i++ {
		pioneer, err := cityPioneer.Pioneers(nil, big.NewInt(int64(i)))
		fmt.Println(pioneer, "-----==")
		_, err = city.PioneerDailyTask(auth, pioneer)
		fmt.Println(pioneer.String(), i, err)
		time.Sleep(time.Second)
	}
}

// CityRechargeTotal 获取区县对应的累计充值权重
func CityRechargeTotal(countyId [32]byte) (error, *big.Int) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}
	city, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
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

//// RechargeDailyWeightRecord 获取区县某一天对应的充值权重
//func RechargeDailyWeightRecord(countyId [32]byte) (error, *big.Int) {
//	Cli := Client(CityNodeConfig)
//	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return err, nil
//	}
//	res, err := city.RechargeDailyWeightRecord(nil, countyId)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return err, nil
//	}
//
//	return nil, res
//}

func GetIncreaseCityDelegateEvent(Cli *ethclient.Client, startBlock, endBlock int64) error {
	defer func() {
		recover()
	}()
	query := event.BuildQuery(
		common.HexToAddress(CityNodeConfig.CityAddress),
		event.IncreaseCityDelegate,
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
	abi, _ := intoCityNode2.CityMetaData.GetAbi()
	for _, logE := range logs {
		logData, err := abi.Unpack(event.IncreaseCityDelegateEvent.EventName, logE.Data)
		if err != nil {
			log.Logger.Sugar().Error(err)
		}
		if len(logData) <= 0 {
			log.Logger.Sugar().Error("logData len 0")
			continue
		}
		var timestamp int64
		header, err := Cli.HeaderByNumber(context.Background(), big.NewInt(int64(logE.BlockNumber)))
		if err == nil {
			timestamp = int64(header.Time)
		}
		cityId := "0x" + common.Bytes2Hex(Bytes32ToBytes(logData[0].([32]uint8)))
		amount := decimal.NewFromBigInt(logData[1].(*big.Int), 0)
		if amount.IsZero() {
			continue
		}
		InsertAdminSetDelegateRecord(cityId, logE.TxHash.String(), amount, timestamp)
		time.Sleep(time.Second * 3)
	}
	return nil
}

func InsertAdminSetDelegateRecord(cityId, txHash string, amount decimal.Decimal, ctime int64) error {
	InsertDelegateLock.Lock()
	defer InsertDelegateLock.Unlock()
	var delegate models.Delegate
	whereCondition := fmt.Sprintf("city_id='%s' and tx_hash=%d", strings.ToLower(cityId), txHash)
	err := db.Mysql.Model(&models.Delegate{}).Where(whereCondition).First(&delegate).Error
	if err == gorm.ErrRecordNotFound {
		db.Mysql.Model(&models.Delegate{}).Create(models.Delegate{
			CityId:  cityId,
			TxHash:  txHash,
			Ctime:   time.Unix(ctime, 0),
			Amount:  amount,
			SetType: 1,
		})
	}
	return nil
}

//func GetDecreaseCityDelegateEvent(Cli *ethclient.Client, startBlock, endBlock int64) error {
//	query := event.BuildQuery(common.HexToAddress(
//		"0xebD06631510A66968f0379A4deB896d3eE7DD6ED"),
//		event.DecreaseCityDelegate,
//		big.NewInt(startBlock),
//		big.NewInt(endBlock),
//	)
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(20))
//	logs, err := Cli.FilterLogs(ctx, query)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return err
//	}
//	cancel()
//	abi, _ := intoCityNode.CityMetaData.GetAbi()
//	for _, logE := range logs {
//		logData, err := abi.Unpack(event.DecreaseCityDelegateEvent.EventName, logE.Data)
//		if err != nil {
//			log.Logger.Sugar().Error(err)
//		}
//		cityId := "0x" + common.Bytes2Hex(Bytes32ToBytes(logData[0].([32]uint8)))
//		amount := logData[1].(*big.Int).String()
//		CreateAdminSetDelegate(models.AdminSetDelegate{
//			CityId:  cityId,
//			Amount:  amount,
//			SetType: 2,
//		})
//	}
//	return nil
//}

func GetRechargeRecordEvent(Cli *ethclient.Client, startBlock, endBlock int64) error {
	query := event.BuildQuery(
		common.HexToAddress(CityNodeConfig.CityAddress),
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
	abi, _ := intoCityNode2.CityMetaData.GetAbi()
	for _, logE := range logs {
		logData, err := abi.Unpack(event.RechargeRecordEvent.EventName, logE.Data)
		if err != nil {
			log.Logger.Sugar().Error(err)
			return err
		}
		userAddress := logData[0].(common.Address)
		countyId := "0x" + common.Bytes2Hex(Bytes32ToBytes(logData[1].([32]uint8)))
		amount := logData[2].(big.Int)
		ctime := logData[3].(int64)
		err, cityId := GetCityIdBytes32ByCountyId(countyId)
		InsertRechargeRecordEvent(
			userAddress.String(),
			countyId,
			cityId,
			decimal.NewFromBigInt(&amount, 0),
			ctime,
		)
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 3)
	}
	return nil
}

func InsertRechargeRecordEvent(userAddress, countyId string, countyIdBytes32 [32]byte, amount decimal.Decimal, ctime int64) error {
	InsertRechargeRecordEventLock.Lock()
	defer InsertRechargeRecordEventLock.Unlock()
	err, cityId := GetChengShiIdByCityIdByte32(countyIdBytes32)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}

	var rechargeRecord models.RechargeRecord
	whereCondition := fmt.Sprintf("user='%s' and county_id='%s' and city_id='%s' and amount='%d' and ctime='%d'",
		strings.ToLower(userAddress), strings.ToLower(countyId), strings.ToLower(cityId), amount, ctime)
	err = db.Mysql.Model(&models.RechargeRecord{}).Where(whereCondition).First(&rechargeRecord).Error
	if err == gorm.ErrRecordNotFound {
		db.Mysql.Model(&models.RechargeRecord{}).Create(&models.RechargeRecord{
			User:     strings.ToLower(userAddress),
			CountyId: strings.ToLower(countyId),
			CityId:   strings.ToLower(cityId),
			Amount:   amount,
			Ctime:    time.Unix(ctime, 0),
		})
	}
	return nil
}
