package blockchain

import (
	intoCityNode "city-node-server/binding"
	"city-node-server/log"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func ApproveToxCityPioneer() {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	tox, err := intoCityNode.NewBgtToken(common.HexToAddress(CityNodeConfig.ToxAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	E18 := big.NewInt(1e18)
	amount := E18.Mul(E18, big.NewInt(100000))
	tx, err := tox.Approve(auth, common.HexToAddress(CityNodeConfig.CityPioneerAddress), amount)

	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(tx, err)
}
func ApproveTox(contractAddress string) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	tox, err := intoCityNode.NewBgtToken(common.HexToAddress(CityNodeConfig.ToxAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	E18 := big.NewInt(1e18)
	amount := E18.Mul(E18, big.NewInt(100000))
	tx, err := tox.Approve(auth, common.HexToAddress(contractAddress), amount)

	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(tx, err)
}

func CityPioneerBalance() {
	Cli := Client(CityNodeConfig)
	tox, err := intoCityNode.NewBgtToken(common.HexToAddress(CityNodeConfig.ToxAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	balance, err := tox.BalanceOf(nil, common.HexToAddress(CityNodeConfig.CityPioneerAddress))

	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(balance.String(), err)
}
