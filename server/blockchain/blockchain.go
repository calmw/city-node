package blockchain

import (
	intoCityNode "city-node-server/binding"
	"city-node-server/log"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type CityNodeConfig struct {
	ChainId            int64
	RPC                string
	CityPioneerAddress string
	PrivateKey         string
}

var cityNodeConfig = CityNodeConfig{
	ChainId:            9001,
	RPC:                "",
	CityPioneerAddress: "",
	PrivateKey:         "",
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
	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, big.NewInt(int64(cityNodeConfig.ChainId)))
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
