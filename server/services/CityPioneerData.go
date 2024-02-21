package services

import (
	"city-node-server/pkg/blockchain"
	"city-node-server/pkg/log"
	"fmt"
)

func InitCityPioneerData() {
	err, cli := blockchain.Client(blockchain.CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	cityPioneerData := blockchain.NewCityPioneerData(cli)

	err = cityPioneerData.AdminSetUSDT()
	fmt.Println(err)
	err = cityPioneerData.AdminSetTOX()
	fmt.Println(err)
	err = cityPioneerData.AdminSetCityAddress()
	fmt.Println(err)
	err = cityPioneerData.AdminSetCityPioneerAddress()
	fmt.Println(err)

}
