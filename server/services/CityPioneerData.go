package services

import (
	"city-node-server/pkg/blockchain"
	"city-node-server/pkg/log"
	"fmt"
	"time"
)

func InitCityPioneerData() {
	err, cli := blockchain.Client(blockchain.CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	cityPioneerData := blockchain.NewCityPioneerData(cli)

	//for {
	//	err = cityPioneerData.AdminSetUSDT()
	//	fmt.Println(err)
	//	if err == nil {
	//		break
	//	}
	//}
	//for {
	//	err = cityPioneerData.AdminSetTOX()
	//	fmt.Println(err)
	//	if err == nil {
	//		break
	//	}
	//}
	//for {
	//	err = cityPioneerData.AdminSetCityAddress()
	//	fmt.Println(err)
	//	if err == nil {
	//		break
	//	}
	//}
	//for {
	//	err = cityPioneerData.AdminSetCityPioneerAddress()
	//	fmt.Println(err)
	//	if err == nil {
	//		break
	//	}
	//}
	//for {
	//	err = cityPioneerData.AddAdmin()
	//	fmt.Println(err)
	//	if err == nil {
	//		break
	//	}
	//}

	var globalNode = []string{
		"0x2907d73b5B693330d41F0bff8D7e827F8F94C70c",
		"0x26AAB384B0e46F53B3f51DBeaF552cB9185cA854",
		"0xb7C6EE5a0Bd1133Ed6f2eb6434cbAB798057d62c",
		"0x56B873652250AeaBb298dbC463FDc8C614F6d6f6",
		"0x04d0f21724b1Da9565c1FB27aeeC9c1BC6FFC0F3",
		"0x74E552c1591A9Bc767057B270090Fd88b82Ea0a1",
		"0x53616e17A4A1817429Cc26B00A420cb3eb270cd0",
		"0x3B317214994bAcA6195A27262eCa2B47b6cD6e4f",
		"0x155B0aa72F4aE1c3f466F303049B069E31a12542",
		"0xb9DE4f123A330bd4B178a93AfDDa34A66843B69D",
		"0x11FA12f176568541f67Caff5dEffD606F8786c71",
		"0x774b8D0Db5c42BC8153e3884016b7B36855fab88",
		"0xE17a7c7898Ae1B13eE82F8476c28Bc1C0309c4F6",
	}

	for _, g := range globalNode {
		for {
			err = cityPioneerData.AdminSetIsGlobalNode(g)
			if err == nil {
				break
			}
			fmt.Println(err)
			time.Sleep(time.Second)
		}
	}
}
