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

	for {
		err = cityPioneerData.AdminSetUSDT()
		fmt.Println(err)
		if err == nil {
			break
		}
	}
	for {
		err = cityPioneerData.AdminSetTOX()
		fmt.Println(err)
		if err == nil {
			break
		}
	}
	for {
		err = cityPioneerData.AdminSetCityAddress()
		fmt.Println(err)
		if err == nil {
			break
		}
	}
	for {
		err = cityPioneerData.AdminSetCityPioneerAddress()
		fmt.Println(err)
		if err == nil {
			break
		}
	}
	for {
		err = cityPioneerData.AddAdmin()
		fmt.Println(err)
		if err == nil {
			break
		}
	}

	var globalNode = []string{
		"",
	}

	for _, g := range globalNode {
		for {
			err = cityPioneerData.AdminSetIsGlobalNode(g)
			fmt.Println(err)
			if err == nil {
				break
			}
		}
	}
}
