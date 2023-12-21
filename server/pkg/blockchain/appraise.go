package blockchain

import (
	"city-node-server/pkg/binding/intoCityNode"
	"city-node-server/pkg/log"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
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
	appraiseContract, err := intoCityNode.NewAppraise(common.HexToAddress(CityNodeConfig.AppraiseAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	res, err := appraiseContract.AdminSetStar(auth, common.HexToAddress(CityNodeConfig.StarAddress))
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

func (a Appraise) AdminSetPioneerBatch(pioneerAddress string, batch int64) error {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	_, auth := GetAuth(Cli)
	appraiseContract, err := intoCityNode.NewAppraise(common.HexToAddress(CityNodeConfig.AppraiseAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	_, err = appraiseContract.AdminSetPioneerBatch(auth, common.HexToAddress(pioneerAddress), big.NewInt(batch))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	return nil
}

func (a Appraise) PioneerBatch(pioneerAddress string) (error, int64) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, 0
	}
	appraiseContract, err := intoCityNode.NewAppraise(common.HexToAddress(CityNodeConfig.AppraiseAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, 0
	}
	beath, err := appraiseContract.PioneerBatch(nil, common.HexToAddress(pioneerAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, 0
	}
	return nil, beath.Int64()
}

func (a Appraise) AdminSetWeightByCityLevel() {
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
	bigE18 := big.NewInt(1e18)
	res, err := appraiseContract.AdminSetWeightByCityLevel(auth, big.NewInt(1), bigE18.Mul(bigE18, big.NewInt(10000)))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	bigE18 = big.NewInt(1e18)
	res, err = appraiseContract.AdminSetWeightByCityLevel(auth, big.NewInt(2), bigE18.Mul(bigE18, big.NewInt(5000)))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	bigE18 = big.NewInt(1e18)
	res, err = appraiseContract.AdminSetWeightByCityLevel(auth, big.NewInt(3), bigE18.Mul(bigE18, big.NewInt(2500)))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res, err)
}
