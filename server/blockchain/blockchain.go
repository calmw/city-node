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
	ChainId               int64
	RPC                   string
	CityAddress           string
	CityPioneerAddress    string
	UserLocationAddress   string
	MiningAddress         string
	SetDelegateAddress    string
	PledgeStakeAddress    string // 获取用户质押量的合约地址
	RechargeWeightAddress string // 获取用户质押量的合约地址
	ToxAddress            string
	PrivateKey            string
}

var CityNodeConfig = CityNodeConfigs{
	ChainId:               9001,
	RPC:                   "https://testnet-rpc.d2ao.com/",
	CityAddress:           "0xDfFA9bfB4D6376DB617fD2Fc56682cC7646cCb9C", // 城市合约地址
	CityPioneerAddress:    "0x11Ef2544A6f62a2cF589390325ab0A56df1eeF15", // 成市先锋合约地址
	UserLocationAddress:   "0x12CEC1A760E1e25Ce5b143cDB9D115D83603fAba", // 用户定位合约地址
	MiningAddress:         "0xD8c1d40a6FF4E53577389C8008343081949C373D", // 需要杨森设置IntoMining合约的管理员权限，权限给到cityPioneer
	SetDelegateAddress:    "0xD8c1d40a6FF4E53577389C8008343081949C373D", // 该合约为增加或减少质押量的合约，艳杰，在city给管理员权限
	PledgeStakeAddress:    "0x575493F35AA4926decF877004056380538C8eB52", // 获取质押量合约。不需要权限
	RechargeWeightAddress: "0xD8c1d40a6FF4E53577389C8008343081949C373D", // 需要杨森给合约地址，在cityPioneer给管理员权限
	ToxAddress:            "0x3eE243ff68074502b1D9D65443fa97b99f634570", // TOX代币合约地址
	PrivateKey:            "a12dc8efdc993a8a7e67700c471f4ef85ddd7d8dceb781c9104637ec194b7ed2",
}

//var CityNodeConfig = CityNodeConfigs{
//	ChainId:             9001,
//	RPC:                 "https://rpc.matchscan.io/",
//	CityAddress:         "0xebD06631510A66968f0379A4deB896d3eE7DD6ED",
//	CityPioneerAddress:  "0x60C541388077d524178521A7ceD95d0c7a016B72",
//	UserLocationAddress: "0x1B535f616B0465891Bc0bb71307A8781A8cCB8f2",
//
//	//	MiningAddress:         "0x0007B44b6Ca810EBff3ED4560cD7d997b08BA104",
//	//	SetDelegateAddress:    "0xD8c1d40a6FF4E53577389C8008343081949C373D",
//	PledgeStakeAddress: "0x909448fBb09880AF2812d263f7E5C77dcfC2AB53",
//	//	RechargeWeightAddress: "0x0007B44b6Ca810EBff3ED4560cD7d997b08BA104",
//	//	ToxAddress:            "0x96397347Ea2beE29713Bc48749eB277D6A36A407",
//
//	PrivateKey: "a12dc8efdc993a8a7e67700c471f4ef85ddd7d8dceb781c9104637ec194b7ed2",
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
