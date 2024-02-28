package blockchain

import (
	"city-node-server/pkg/binding/intoCityNode"
	"city-node-server/pkg/db"
	"city-node-server/pkg/log"
	models2 "city-node-server/pkg/models"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
	"math/big"
	"strings"
	"time"
)

type CityPioneer struct {
	Cli      *ethclient.Client
	Contract *intoCityNode.CityPioneer
}

func NewCityPioneer(cli *ethclient.Client) *CityPioneer {
	cityPioneer, _ := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), cli)

	return &CityPioneer{
		Cli:      cli,
		Contract: cityPioneer,
	}
}

//	func AdminSetAssessmentCriteria(cityLevel, month, point int64) {
//		Cli := Client(CityNodeConfig)
//		_, auth := GetAuth(Cli)
//		cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
//		if err != nil {
//			log.Logger.Sugar().Error(err)
//			return
//		}
//		criteria, err := cityPioneer.AdminSetAssessmentCriteria(auth, big.NewInt(cityLevel), big.NewInt(month), big.NewInt(point))
//		if err != nil {
//			log.Logger.Sugar().Error(err)
//			return
//		}
//		fmt.Println(criteria, err)
//	}
//

// AdminSetAssessmentReturnRate 管理员设置保证金退还比例
//func AdminSetAssessmentReturnRate(cityLevel, month, point int64) {
//	err, Cli := Client(CityNodeConfig)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	err, auth := GetAuth(Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	criteria, err := cityPioneer.AdminSetAssessmentReturnRate(auth, big.NewInt(cityLevel), big.NewInt(month), big.NewInt(point))
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	fmt.Println(criteria, err)
//}

// AdminSetAssessmentReturnRate 管理员设置保证金退还比例
//func (c CityPioneer) AdminSetAppraise() {
//	err, auth := GetAuth(c.Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	criteria, err := c.Contract.AdminSetAppraise(auth, common.HexToAddress(CityNodeConfig.AppraiseAddress))
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	fmt.Println(criteria, err)
//}

//func (c CityPioneer) AdminSetCityPioneerDataAddress() {
//	err, auth := GetAuth(c.Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	criteria, err := c.Contract.AdminSetCityPioneerDataAddress(auth, common.HexToAddress(CityNodeConfig.CityPioneerDataAddress))
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	fmt.Println(criteria.Hash(), err)
//}

//func (c CityPioneer) AdminSetSecondsPerDay(secondsPerDay int64) {
//	err, Cli := Client(CityNodeConfig)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	err, auth := GetAuth(Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	criteria, err := cityPioneer.AdminSetSecondsPerDay(auth, big.NewInt(secondsPerDay))
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	fmt.Println(criteria, err)
//}

//func (c CityPioneer) AdminChangePioneerAddress(newPioneerAddress, oldPioneerAddress string) {
//	err, Cli := Client(CityNodeConfig)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	err, auth := GetAuth(Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//
//	criteria, err := c.Contract.AdminChangePioneerAddress(auth, common.HexToAddress(newPioneerAddress), common.HexToAddress(oldPioneerAddress))
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	fmt.Println(criteria, err)
//}

type Pioneer struct {
	PioneerAddress        common.Address
	Ctime                 *big.Int
	CityLevel             *big.Int
	AssessmentMonthStatus bool
	AssessmentStatus      bool
	ReturnSuretyStatus    bool
	ReturnSuretyRate      *big.Int
	ReturnSuretyTime      *big.Int
	SuretyTime            *big.Int
}

func (c CityPioneer) PioneerInfo(pioneerAddress string) (error, Pioneer) {
	criteria, err := c.Contract.PioneerInfo(nil, common.HexToAddress(pioneerAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, Pioneer{}
	}
	return nil, criteria
}

// AdminSetCheckPioneerDailyStatus 管理员设置保证金退还比例
//func AdminSetCheckPioneerDailyStatus(pioneer string, day int64, status bool) {
//	err, Cli := Client(CityNodeConfig)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	err, auth := GetAuth(Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	criteria, err := cityPioneer.AdminSetCheckPioneerDailyStatus(auth, common.HexToAddress(pioneer), big.NewInt(day), status)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	fmt.Println(criteria, err)
//}

// AdminSetAssessmentReturnCriteria 管理员设置城市先锋保证金退还标准；点数
//func AdminSetAssessmentReturnCriteria(cityLevel, month, point int64) {
//	err, Cli := Client(CityNodeConfig)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	err, auth := GetAuth(Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	criteria, err := cityPioneer.AdminSetAssessmentReturnCriteria(auth, big.NewInt(cityLevel), big.NewInt(month), big.NewInt(point))
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	fmt.Println(criteria, err)
//}

//// AdminSetStartTime 管理员设置开始考核时间
//func AdminSetStartTime(startTime int64) {
//	Cli := Client(CityNodeConfig)
//	_, auth := GetAuth(Cli)
//	cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	criteria, err := cityPioneer.AdminSetStartTime(auth, big.NewInt(startTime))
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	fmt.Println(criteria, err)
//}
//
//// AdminSetTOXAddress  管理员设置TOX代币地址
//func AdminSetTOXAddress() {
//	Cli := Client(CityNodeConfig)
//	_, auth := GetAuth(Cli)
//	cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	criteria, err := cityPioneer.AdminSetTOXAddress(auth, common.HexToAddress(CityNodeConfig.ToxAddress))
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	fmt.Println(criteria, err)
//}
//
//// AdminSetCityAddress 管理员设置城市合约地址
//func AdminSetCityAddress() {
//	Cli := Client(CityNodeConfig)
//	_, auth := GetAuth(Cli)
//	cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	criteria, err := cityPioneer.AdminSetCityAddress(auth, common.HexToAddress(CityNodeConfig.CityAddress))
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	fmt.Println(criteria, err)
//}
//
//// AdminSetMiningAddress 管理员设置增加用户合约余额合约地址
//func AdminSetMiningAddress() {
//	Cli := Client(CityNodeConfig)
//	_, auth := GetAuth(Cli)
//	cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	criteria, err := cityPioneer.AdminSetMiningAddress(auth, common.HexToAddress(CityNodeConfig.MiningAddress))
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	fmt.Println(criteria, err)
//}
//
//// AddCityPioneerAdmin 给城市合约、IntoMining设置用户（增加用户合约余额）添加管理员权限
//func AddCityPioneerAdmin() {
//	Cli := Client(CityNodeConfig)
//	_, auth := GetAuth(Cli)
//	cityPioneer, err := intoCityNode2.NewCity(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//
//	res, err := cityPioneer.AddAdmin(auth, common.HexToAddress("0x94b627F4F829Ac5E97fDc556B5BEeeFf9beF417e"))
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	fmt.Println(err, 222)
//
//	//res, err := cityPioneer.AddAdmin(auth, common.HexToAddress(CityNodeConfig.CityAddress))
//	//if err != nil {
//	//	log.Logger.Sugar().Error(err)
//	//	return
//	//}
//	//
//	//res, err = cityPioneer.AddAdmin(auth, common.HexToAddress(CityNodeConfig.MiningAddress))
//	//if err != nil {
//	//	log.Logger.Sugar().Error(err)
//	//	return
//	//}
//	fmt.Println(res, err)
//}
//
//// AdminSetSecondsPerDayCityPioneer 管理员设置每天秒数，用于测试
//func AdminSetSecondsPerDayCityPioneer(seconds int64) {
//	Cli := Client(CityNodeConfig)
//	_, auth := GetAuth(Cli)
//	cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	res, err := cityPioneer.AdminSetSecondsPerDay(auth, big.NewInt(seconds))
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	fmt.Println(res.Hash(), err)
//}
//
//// AdminSetPresidencyTime 管理员设置任期
//func AdminSetPresidencyTime(seconds int64) {
//	Cli := Client(CityNodeConfig)
//	_, auth := GetAuth(Cli)
//	cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	res, err := cityPioneer.AdminSetPresidencyTime(auth, big.NewInt(seconds))
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	fmt.Println(res.Hash(), err)
//}
//
//// DepositSuretyTest 交保证金成为先锋
//func DepositSuretyTest(pioneerAddress string) {
//	Cli := Client(CityNodeConfig)
//	_, auth := GetAuth(Cli)
//	cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	criteria, err := cityPioneer.DepositSuretyTest(auth, common.HexToAddress(pioneerAddress))
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	fmt.Println(criteria, err)
//}

// DailyTask  每天执行一次，计算奖励和考核
//func DailyTask() {
//	Cli := Client(CityNodeConfig)
//	_, auth := GetAuth(Cli)
//	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	criteria, err := cityPioneer.DailyTask(auth)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	fmt.Println(criteria, err)
//}

// DepositSurety  交保证金
//func DepositSurety() {
//	err, Cli := Client(CityNodeConfig)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	err, auth := GetAuth(Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	criteria, err := c.DepositSurety(auth)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return
//	}
//	fmt.Println(criteria, err)
//}

func GetPioneer(index int64) (error, string) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ""
	}
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ""
	}
	pioneer, err := cityPioneer.Pioneers(nil, big.NewInt(index))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ""
	}
	return nil, pioneer.String()
}

func GetPioneerNumber() (error, int64) {
	err, Cli := Client(CityNodeConfig)
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, 0
	}
	number, err := cityPioneer.GetPioneerNumber(nil)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, 0
	}
	return nil, number.Int64()
}

func GetPioneerTaskStatus(pioneerAddress string) bool {
	var pioneerTask models2.PioneerTask
	t := time.Now().UTC()
	whereCondition := fmt.Sprintf("pioneer='%s' and date='%s'", strings.ToLower(pioneerAddress), t.Format("2006-01-02"))
	err := db.Mysql.Model(&models2.PioneerTask{}).Where(whereCondition).First(&pioneerTask).Error
	if err == gorm.ErrRecordNotFound {
		db.Mysql.Model(&models2.PioneerTask{}).Create(&models2.PioneerTask{
			Pioneer: strings.ToLower(pioneerAddress),
			Date:    t,
		})
		return false
	}
	return pioneerTask.Status > 0
}

func SetPioneerTaskStatus(pioneerAddress string) error {
	whereCondition := fmt.Sprintf("pioneer='%s'", strings.ToLower(pioneerAddress))
	log.Logger.Sugar().Info("pioneer task success ", pioneerAddress)
	return db.Mysql.Model(&models2.PioneerTask{}).Where(whereCondition).Update("status", 1).Error
}
