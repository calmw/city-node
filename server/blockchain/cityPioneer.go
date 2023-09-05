package blockchain

import (
	intoCityNode "city-node-server/binding"
	"city-node-server/log"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func AdminSetAssessmentCriteria(cityLevel, month, point int64) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	criteria, err := cityPioneer.AdminSetAssessmentCriteria(auth, big.NewInt(cityLevel), big.NewInt(month), big.NewInt(point))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(criteria, err)
}
