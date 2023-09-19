package blockchain

import (
	intoCityNode "city-node-server/binding"
	"city-node-server/log"
	"github.com/ethereum/go-ethereum/common"
)

func GetUidsWithAddress(addresses []common.Address) (error, []string) {
	Cli := Client(CityNodeConfig)
	intoBind, err := intoCityNode.NewIntoBind(common.HexToAddress("0x15f414e10101afDeDA12Ba5608795eb4e14f2D3d"), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}
	tx, err := intoBind.GetUidsWithAddress(nil, addresses)

	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}
	return nil, tx
}
