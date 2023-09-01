package blockchain

import (
	"github.com/ethereum/go-ethereum/ethclient"
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

func CityPioneerDailyTask() error {

	return nil
}
