package blockchain

import (
	intoCityNode2 "city-node-server/pkg/binding/intoCityNode"
	"city-node-server/pkg/blockchain/event"
	"city-node-server/pkg/db"
	"city-node-server/pkg/log"
	models2 "city-node-server/pkg/models"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"math/big"
	"strings"
	"time"
)

type CityPioneer struct{}

func NewCityPioneer() *CityPioneer {
	return &CityPioneer{}
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
func (c CityPioneer) AdminSetAppraise() {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	err, auth := GetAuth(Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	criteria, err := cityPioneer.AdminSetAppraise(auth, common.HexToAddress(CityNodeConfig.AppraiseAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(criteria, err)
}

func (c CityPioneer) AdminSetSecondsPerDay(secondsPerDay int64) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	err, auth := GetAuth(Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	criteria, err := cityPioneer.AdminSetSecondsPerDay(auth, big.NewInt(secondsPerDay))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(criteria, err)
}

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
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, Pioneer{}
	}
	cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, Pioneer{}
	}
	criteria, err := cityPioneer.PioneerInfo(nil, common.HexToAddress(pioneerAddress))
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
func DepositSurety() {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	err, auth := GetAuth(Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	criteria, err := cityPioneer.DepositSurety(auth)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(criteria, err)
}
func GetPioneer(index int64) (error, string) {
	err, Cli := Client(CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ""
	}
	cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
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
	cityPioneer, err := intoCityNode2.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
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

func GetDailyRewardRecordEvent(Cli *ethclient.Client, startBlock, endBlock int64) error {
	time.Sleep(time.Second * 2)
	InsertDailyRewardLock.Lock()
	defer InsertDailyRewardLock.Unlock()
	query := event.BuildQuery(
		common.HexToAddress(CityNodeConfig.CityPioneerAddress),
		event.DailyRewardRecord,
		big.NewInt(startBlock),
		big.NewInt(endBlock),
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(20))
	logs, err := Cli.FilterLogs(ctx, query)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	cancel()
	abi, _ := intoCityNode2.CityPioneerMetaData.GetAbi()

	for _, logE := range logs {
		logData, err := abi.Unpack(event.DailyRewardRecordEvent.EventName, logE.Data)
		if err != nil {
			log.Logger.Sugar().Error(err)
			return err
		}
		var timestamp int64
		pioneerAddress := strings.ToLower(logData[0].(common.Address).String())
		nodeReward := decimal.NewFromBigInt(logData[1].(*big.Int), 0)
		foundsReward := decimal.NewFromBigInt(logData[2].(*big.Int), 0)
		delegateReward := decimal.NewFromBigInt(logData[3].(*big.Int), 0)
		header, err := Cli.HeaderByNumber(context.Background(), big.NewInt(int64(logE.BlockNumber)))
		if err == nil {
			timestamp = int64(header.Time)
		}

		err = InsertDailyReward(pioneerAddress, logE.TxHash.String(), foundsReward, delegateReward, nodeReward, int64(logE.BlockNumber), timestamp)
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 3)
	}
	return nil
}

func InsertDailyReward(pioneerAddress, txHash string, foundsReward, delegateReward, nodeReward decimal.Decimal, blockHeight, timestamp int64) error {

	var reward models2.Reward
	whereCondition := fmt.Sprintf("pioneer='%s' and block_height=%d", strings.ToLower(pioneerAddress), blockHeight)
	err := db.Mysql.Table("reward").Where(whereCondition).First(&reward).Error
	if err == gorm.ErrRecordNotFound {
		// 获取城市信息
		var userLocation models2.UserLocation
		db.Mysql.Model(&models2.UserLocation{}).Where("user=?", strings.ToLower(pioneerAddress)).First(&userLocation)
		// 插入数据
		db.Mysql.Table("reward").Create(&models2.Reward{
			Pioneer:        pioneerAddress,
			TxHash:         txHash,
			City:           userLocation.Location,
			FoundsReward:   foundsReward,
			DelegateReward: delegateReward,
			NodeReward:     nodeReward,
			BlockHeight:    blockHeight,
			Ctime:          time.Unix(timestamp, 0),
		})
	}
	return nil
}

//func GetWithdrawalRewardRecordEvent(Cli *ethclient.Client, startBlock, endBlock int64) error {
//	query := event.BuildQuery(
//		common.HexToAddress(CityNodeConfig.CityPioneerAddress),
//		event.WithdrawalRewardRecord,
//		big.NewInt(startBlock),
//		big.NewInt(endBlock),
//	)
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(20))
//	logs, err := Cli.FilterLogs(ctx, query)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return err
//	}
//	cancel()
//	abi, _ := intoCityNode2.CityPioneerMetaData.GetAbi()
//
//	for _, logE := range logs {
//		logData, err := abi.Unpack(event.WithdrawalRewardRecordEvent.EventName, logE.Data)
//		if err != nil {
//			log.Logger.Sugar().Error(err)
//			return err
//		}
//		var timestamp int64
//		pioneerAddress := strings.ToLower(logData[0].(common.Address).String())
//		amount := decimal.NewFromBigInt(logData[1].(*big.Int), 0)
//		wType := decimal.NewFromBigInt(logData[2].(*big.Int), 0)
//		header, err := Cli.HeaderByNumber(context.Background(), big.NewInt(int64(logE.BlockNumber)))
//		if err == nil {
//			timestamp = int64(header.Time)
//		}
//		err = InsertWithdrawalRewardRecord(pioneerAddress, logE.TxHash.String(), amount, wType, int64(logE.BlockNumber), int64(logE.Index), timestamp)
//		if err != nil {
//			return err
//		}
//		time.Sleep(time.Second * 3)
//	}
//	return nil
//}

//func InsertWithdrawalRewardRecord(pioneerAddress, txHash string, amount, rewardType decimal.Decimal, blockHeight, logIndex, timestamp int64) error {
//	InsertWithdrawalRewardRecordLock.Lock()
//	defer InsertWithdrawalRewardRecordLock.Unlock()
//	var rewardWithdraw models.RewardWithdraw
//	whereCondition := fmt.Sprintf("pioneer='%s' and block_height=%d and reward_type=%s", strings.ToLower(pioneerAddress), blockHeight, rewardType.String())
//	//whereCondition := fmt.Sprintf("pioneer='%s' and block_height=%d and log_index=%d", strings.ToLower(pioneerAddress), blockHeight, logIndex)
//	err := db.Mysql.Model(&models.RewardWithdraw{}).Where(whereCondition).First(&rewardWithdraw).Error
//	if err == gorm.ErrRecordNotFound {
//		db.Mysql.Model(&models.RewardWithdraw{}).Create(&models.RewardWithdraw{
//			Pioneer:     pioneerAddress,
//			TxHash:      txHash,
//			Amount:      amount,
//			RewardType:  rewardType,
//			BlockHeight: blockHeight,
//			LogIndex:    logIndex,
//			Ctime:       time.Unix(timestamp, 0),
//		})
//	}
//	return nil
//}

func GetSuretyRecordEvent(Cli *ethclient.Client, startBlock, endBlock int64) error {
	fmt.Println(startBlock, endBlock)
	query := event.BuildQuery(
		common.HexToAddress(CityNodeConfig.CityPioneerAddress),
		event.SuretyRecord,
		big.NewInt(startBlock),
		big.NewInt(endBlock),
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(20))
	logs, err := Cli.FilterLogs(ctx, query)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	cancel()
	abi, _ := intoCityNode2.CityPioneerMetaData.GetAbi()

	for _, logE := range logs {
		logData, err := abi.Unpack(event.SuretyRecordEvent.EventName, logE.Data)
		if err != nil {
			log.Logger.Sugar().Error(err)
			return err
		}
		var timestamp int64
		pioneerAddress := strings.ToLower(logData[0].(common.Address).String())
		amount := decimal.NewFromBigInt(logData[1].(*big.Int), 0)
		month := logData[2].(*big.Int).Int64()
		header, err := Cli.HeaderByNumber(context.Background(), big.NewInt(int64(logE.BlockNumber)))
		if err == nil {
			timestamp = int64(header.Time)
		}
		err = InsertSuretyRecordRecord(pioneerAddress, logE.TxHash.String(), amount, int64(logE.BlockNumber), month, timestamp)
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 3)
	}
	return nil
}

func InsertSuretyRecordRecord(pioneerAddress, txHash string, amount decimal.Decimal, blockHeight, month, timestamp int64) error {
	InsertSuretyRecordRecordLock.Lock()
	defer InsertSuretyRecordRecordLock.Unlock()
	var suretyRecord models2.SuretyRecord
	whereCondition := fmt.Sprintf("pioneer='%s' and block_height=%d", strings.ToLower(pioneerAddress), blockHeight)
	err := db.Mysql.Model(&models2.SuretyRecord{}).Where(whereCondition).First(&suretyRecord).Error
	if err == gorm.ErrRecordNotFound {
		db.Mysql.Model(&models2.SuretyRecord{}).Create(&models2.SuretyRecord{
			Pioneer:     pioneerAddress,
			TxHash:      txHash,
			Amount:      amount,
			Month:       month,
			BlockHeight: blockHeight,
			Ctime:       time.Unix(timestamp, 0),
		})
	}
	return nil
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
