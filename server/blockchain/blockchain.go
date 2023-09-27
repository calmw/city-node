package blockchain

import (
	"city-node-server/db"
	"city-node-server/log"
	"city-node-server/models"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

// 1 用户定位事件处理，2 城市先锋奖励事件，3增加充值事件，4获取新增质押事件，5获取奖励领取事件
const (
	LocationEvent = iota + 1
	RewardEvent
	RechargeEvent
	DelegateEvent
	WithdrawEvent
)

type CityNodeConfigs struct {
	ChainId               int64
	RPC                   string
	CityAddress           string
	CityPioneerAddress    string
	UserLocationAddress   string
	MiningAddress         string
	SetDelegateAddress    string
	PledgeStakeAddress    string // 获取用户质押量的合约地址
	RechargeWeightAddress string // 获取用户质押量的合约地址
	GetFoundsAddress      string // 获取社交基金值的合约
	ToxAddress            string
	PrivateKey            string
}

var CityNodeConfig CityNodeConfigs

func Client(c CityNodeConfigs) *ethclient.Client {
	client, err := ethclient.Dial(c.RPC)
	if err != nil {
		panic("dail failed")
	}
	return client
}

func GetAuth(cli *ethclient.Client) (error, *bind.TransactOpts) {
	privateKeyEcdsa, err := crypto.HexToECDSA(CityNodeConfig.PrivateKey)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, big.NewInt(CityNodeConfig.ChainId))
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

//func GetCityDelegateEvent() error {
//	Cli := Client(CityNodeConfig)
//	err, startBlock := GetStartBlock()
//	if err != nil {
//		return err
//	}
//	number, err := Cli.BlockNumber(context.Background())
//	endBlock := startBlock + 999
//	if endBlock > number {
//		return nil
//	}
//	err = GetIncreaseCityDelegateEvent(Cli, int64(startBlock), int64(endBlock))
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//	}
//	//err = GetDecreaseCityDelegateEvent(Cli, int64(startBlock), int64(endBlock))
//	//if err != nil {
//	//	log.Logger.Sugar().Error(err)
//	//}
//	//SetSTartBlock(int64(startBlock + 1000))
//	return nil
//}

func Bytes32ToBytes(bytes32 [32]byte) []byte {
	var by []byte
	for i := 0; i < 32; i++ {
		by = append(by, bytes32[i])
	}
	return by
}

func BytesToByte32(bytes []byte) [32]byte {
	var by [32]byte
	for i := 0; i < 32; i++ {
		by[i] = bytes[i]
	}
	return by
}

func GetStartBlock(id int64) (error, uint64) {
	var blockStore = models.BlockStore{}
	err := db.Mysql.Table("block_store").Where("id=?", id).First(&blockStore).Error
	return err, uint64(blockStore.StartBlock)
}

func SetSTartBlock(startBlock, id int64) {
	fmt.Println("更新区块高度", startBlock)
	db.Mysql.Model(&models.BlockStore{}).Where("id=?", id).Update("start_block", startBlock)
}
