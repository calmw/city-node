package blockchain

import (
	intoCityNode "city-node-server/binding"
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
	"math/big"
	"time"
)

type CityNodeConfig struct {
	ChainId             int64
	RPC                 string
	CityAddress         string
	CityPioneerAddress  string
	UserLocationAddress string
	PrivateKey          string
}

var cityNodeConfig = CityNodeConfig{
	ChainId:             9001,
	RPC:                 "https://rpc.matchscan.io/",
	CityAddress:         "0xebD06631510A66968f0379A4deB896d3eE7DD6ED",
	CityPioneerAddress:  "",
	UserLocationAddress: "0x1B535f616B0465891Bc0bb71307A8781A8cCB8f2",
	PrivateKey:          "a12dc8efdc993a8a7e67700c471f4ef85ddd7d8dceb781c9104637ec194b7ed2",
}

func Client(c CityNodeConfig) *ethclient.Client {
	client, err := ethclient.Dial(c.RPC)
	if err != nil {
		panic("dail failed")
	}
	return client
}

func GetAuth(cli *ethclient.Client) (error, *bind.TransactOpts) {
	privateKeyEcdsa, err := crypto.HexToECDSA(cityNodeConfig.PrivateKey)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, big.NewInt(cityNodeConfig.ChainId))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}

	//gasLimit := uint64(21000)
	return nil, &bind.TransactOpts{
		From:      auth.From,
		Nonce:     nil,
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

func CityPioneerDailyTask() error {
	Cli := Client(cityNodeConfig)
	_, auth := GetAuth(Cli)
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(""), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	task, err := cityPioneer.DailyTask(auth)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	log.Logger.Sugar().Info("hash: ", task.Hash())
	return nil
}

// CityDailyTask 城市合约定时任务
func CityDailyTask() error {
	Cli := Client(cityNodeConfig)
	_, auth := GetAuth(Cli)
	city, err := intoCityNode.NewCity(common.HexToAddress(cityNodeConfig.CityAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	task, err := city.DailyTask(auth)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	log.Logger.Sugar().Info("hash: ", task.Hash())
	return nil
}

func GetCityDelegateEvent() error {
	Cli := Client(cityNodeConfig)
	startBlock := GetStartBlock()
	number, err := Cli.BlockNumber(context.Background())
	endBlock := startBlock + 999
	if endBlock > number {
		return nil
	}
	err = GetIncreaseCityDelegateEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		log.Logger.Sugar().Error(err)
	}
	err = GetDecreaseCityDelegateEvent(Cli, int64(startBlock), int64(endBlock))
	if err != nil {
		log.Logger.Sugar().Error(err)
	}
	SetSTartBlock(int64(startBlock + 1000))
	return nil
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

func Bytes32ToBytes(bytes32 [32]byte) []byte {
	var by []byte
	for i := 0; i < 32; i++ {
		by = append(by, bytes32[i])
	}
	return by
}

func GetStartBlock() uint64 {
	// 2307995
	var blockStore = models.BlockStore{}
	db.Mysql.Table("block_store").Where("id=1").First(&blockStore)
	return blockStore.StartBlock
}

func SetSTartBlock(startBlock int64) {
	db.Mysql.Table("block_store").Where("id=1").Update("start_block", startBlock)
}

func CreateAdminSetDelegate(adminSetDelegate models.AdminSetDelegate) {
	db.Mysql.Table("admin_set_delegate").Create(&adminSetDelegate)
}

// UserLocationGetCityNum  城市合约定时任务
func UserLocationGetCityNum() (error, int64) {
	Cli := Client(cityNodeConfig)
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(cityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, 0
	}
	cityNum, err := userLocation.CityIdNum(nil)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, 0
	}

	for i := 0; i < int(cityNum.Int64()); i++ {
		err, cityByte32 := UserLocationGetCity(big.NewInt(int64(i)))
		if err == nil {
			UserLocationSetNoRepeatCityIds(cityByte32)
		}
	}
	return nil, cityNum.Int64()
}

// UserLocationGetCity  城市合约定时任务
func UserLocationGetCity(index *big.Int) (error, [32]byte) {
	Cli := Client(cityNodeConfig)
	//_, auth := GetAuth(Cli)
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(cityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, [32]byte{}
	}
	cityId, err := userLocation.CityIds(nil, index)

	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, [32]byte{}
	}
	fmt.Println("0x" + common.Bytes2Hex(Bytes32ToBytes(cityId)))
	return nil, [32]byte{}
}

// UserLocationNoRepeatCityIds  城市合约定时任务
func UserLocationNoRepeatCityIds() {
	Cli := Client(cityNodeConfig)
	_, auth := GetAuth(Cli)
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(cityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	tx, err := userLocation.NoRepeatCityIds(auth)

	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(tx)
}

// UserLocationSetNoRepeatCityIds  城市合约定时任务
func UserLocationSetNoRepeatCityIds(cityId [32]byte) {
	Cli := Client(cityNodeConfig)
	_, auth := GetAuth(Cli)
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(cityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	tx, err := userLocation.SetNoRepeatCityIds(auth, cityId)
	fmt.Println(tx, err, 666)
	time.Sleep(time.Second)
}
