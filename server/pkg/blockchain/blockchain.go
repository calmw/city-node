package blockchain

import (
	"city-node-server/pkg/db"
	"city-node-server/pkg/log"
	"city-node-server/pkg/models"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/status-im/keycard-go/hexutils"
	"math/big"
	"strings"
)

// 1 用户定位事件处理，2 城市先锋奖励事件，3增加充值事件，4获取新增质押事件，5获取奖励领取事件
const (
	LocationEvent = iota + 1
	RewardEvent
	RechargeEvent
	DelegateEvent
	WithdrawEvent
	LocationEvent2
	SuretyRecordEvent
)
const LayoutTime = "2006-01-02 15:04:05"

type CityNodeConfigs struct {
	ChainId                int64
	RPC                    string
	AuthAddress            string
	WithdrawLimitAddress   string
	AppraiseAddress        string
	StarAddress            string
	CityAddress            string
	CityPioneerAddress     string
	CityPioneerDataAddress string
	UserLocationAddress    string
	MiningAddress          string
	SetDelegateAddress     string
	PledgeStakeAddress     string // 获取用户质押量的合约地址
	RechargeWeightAddress  string // 获取用户质押量的合约地址
	GetFoundsAddress       string // 获取社交基金值的合约
	ToxAddress             string
	USDTAddress            string
	PrivateKey             string
}

var CityNodeConfig CityNodeConfigs

func Client(c CityNodeConfigs) (error, *ethclient.Client) {
	client, err := ethclient.Dial(c.RPC)
	if err != nil {
		log.Logger.Sugar().Error("dail failed")
		return err, nil
	}
	return nil, client
}

func GetAuth(cli *ethclient.Client) (error, *bind.TransactOpts) {
	privateKeyEcdsa, err := crypto.HexToECDSA(CityNodeConfig.PrivateKey)
	//fmt.Println(CityNodeConfig.PrivateKey)
	//fmt.Println(CityNodeConfig.RPC)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, big.NewInt(CityNodeConfig.ChainId))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}

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

func ConvertAreaIdAtB(cityId string) [32]byte {
	if strings.Contains(cityId, "0x") {
		cityId = strings.ReplaceAll(cityId, "0x", "")
	}
	common.Hex2Bytes(cityId)
	return BytesToByte32(common.Hex2Bytes(cityId))
}

func ConvertAreaIdBtA(cityId [32]byte) string {
	cityIdBytes := Bytes32ToBytes(cityId)
	hexutils.BytesToHex(cityIdBytes)
	return strings.ToLower("0x" + hexutils.BytesToHex(cityIdBytes))
}
