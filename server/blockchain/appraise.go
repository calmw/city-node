package blockchain

import (
	intoCityNode "city-node-server/binding"
	"city-node-server/log"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

type Appraise struct{}

func NewAppraise() *Appraise {
	return &Appraise{}
}

// AdminSetStar 管理员设置Star合约地址
func (a Appraise) AdminSetStar() {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	_, auth := GetAuth(Cli)
	fmt.Println(CityNodeConfig.AppraiseAddress, 999)
	fmt.Println(CityNodeConfig.RPC, 999)
	fmt.Println(CityNodeConfig.ChainId, 999)
	appraiseContract, err := intoCityNode.NewAppraise(common.HexToAddress(CityNodeConfig.AppraiseAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	admins, err := appraiseContract.AllAdmins(nil)
	fmt.Println(admins, err)
	if err != nil {
		return
	}

	fmt.Println(CityNodeConfig.StarAddress)
	res, err := appraiseContract.AdminSetStar(auth, common.HexToAddress("0xe8739b502df3A3dC5C0f14c0F27288c06A5ad887"))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res, err)
}

// AdminSetCity 管理员设置城市合约
func (a Appraise) AdminSetCity() {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	_, auth := GetAuth(Cli)
	appraiseContract, err := intoCityNode.NewAppraise(common.HexToAddress(CityNodeConfig.AppraiseAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	res, err := appraiseContract.AdminSetCity(auth, common.HexToAddress(CityNodeConfig.CityAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res, err)
}

// AdminSetPioneer 管理员设置先锋合约
func (a Appraise) AdminSetPioneer() {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	_, auth := GetAuth(Cli)
	appraiseContract, err := intoCityNode.NewAppraise(common.HexToAddress(CityNodeConfig.AppraiseAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	res, err := appraiseContract.AdminSetPioneer(auth, common.HexToAddress(CityNodeConfig.CityPioneerAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res, err)
}
func (a Appraise) AddAdmin() {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	_, auth := GetAuth(Cli)
	appraiseContract, err := intoCityNode.NewAppraise(common.HexToAddress(CityNodeConfig.AppraiseAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	admins, err := appraiseContract.AllAdmins(nil)
	fmt.Println(admins, err, 123)
	if err != nil {
		return
	}
	res, err := appraiseContract.AddAdmin(auth, common.HexToAddress(CityNodeConfig.CityPioneerAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res, err)
}
