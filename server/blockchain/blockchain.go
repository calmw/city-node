package blockchain

import (
	"city-node-server/db"
	"city-node-server/log"
	"city-node-server/models"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type CityNodeConfigs struct {
	ChainId             int64
	RPC                 string
	CityAddress         string
	CityPioneerAddress  string
	UserLocationAddress string
	MiningAddress       string
	SetDelegateAddress  string
	PledgeStakeAddress  string // 获取用户质押量的合约地址
	ToxAddress          string
	PrivateKey          string
}

var CityNodeConfig = CityNodeConfigs{
	ChainId:             9001,
	RPC:                 "https://testnet-rpc.d2ao.com/",
	CityAddress:         "0x133d15CeBfaf0fb2cA91f8EA6F2c2E8B26D2B310",
	CityPioneerAddress:  "0x425C48382F30D6331E096dd5B720F231C7a9114F",
	UserLocationAddress: "0xcd507929f9f8B79f02192837eaD33B30c89752Ce",
	MiningAddress:       "0xD8c1d40a6FF4E53577389C8008343081949C373D",
	SetDelegateAddress:  "0xD8c1d40a6FF4E53577389C8008343081949C373D",
	PledgeStakeAddress:  "0x575493F35AA4926decF877004056380538C8eB52",
	ToxAddress:          "0x05171e5C88b43ef35D223f64E1304D3D5210701D",
	PrivateKey:          "a12dc8efdc993a8a7e67700c471f4ef85ddd7d8dceb781c9104637ec194b7ed2",
}

//var CityNodeConfig = CityNodeConfigs{
//	ChainId: 9001,
//	RPC:     "https://rpc.matchscan.io/",
//	//RPC:                 "https://testnet-rpc.d2ao.com/",https://testnet.matchscan.io/, 9001 // test net
//	CityAddress:         "0xebD06631510A66968f0379A4deB896d3eE7DD6ED",
//	CityPioneerAddress:  "",
//	UserLocationAddress: "0x1B535f616B0465891Bc0bb71307A8781A8cCB8f2",
//	PrivateKey:          "a12dc8efdc993a8a7e67700c471f4ef85ddd7d8dceb781c9104637ec194b7ed2",
//}

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

func GetCityDelegateEvent() error {
	Cli := Client(CityNodeConfig)
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

func GetStartBlock() uint64 {
	// 2307995
	var blockStore = models.BlockStore{}
	db.Mysql.Table("block_store").Where("id=1").First(&blockStore)
	return blockStore.StartBlock
}

func SetSTartBlock(startBlock int64) {
	db.Mysql.Table("block_store").Where("id=1").Update("start_block", startBlock)
}
