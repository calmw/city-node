package blockchain

import (
	intoCityNode "city-node-server/binding"
	"city-node-server/blockchain/event"
	"city-node-server/db"
	"city-node-server/log"
	"city-node-server/models"
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

// AdminSetAssessmentReturnRate 管理员设置保证金退还比例
func AdminSetAssessmentReturnRate(cityLevel, month, point int64) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	criteria, err := cityPioneer.AdminSetAssessmentReturnRate(auth, big.NewInt(cityLevel), big.NewInt(month), big.NewInt(point))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(criteria, err)
}

// AdminSetAssessmentReturnCriteria 管理员设置城市先锋保证金退还标准；点数
func AdminSetAssessmentReturnCriteria(cityLevel, month, point int64) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	criteria, err := cityPioneer.AdminSetAssessmentReturnCriteria(auth, big.NewInt(cityLevel), big.NewInt(month), big.NewInt(point))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(criteria, err)
}

// AdminSetStartTime 管理员设置开始考核时间
func AdminSetStartTime(startTime int64) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	criteria, err := cityPioneer.AdminSetStartTime(auth, big.NewInt(startTime))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(criteria, err)
}

// AdminSetTOXAddress  管理员设置TOX代币地址
func AdminSetTOXAddress() {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	criteria, err := cityPioneer.AdminSetTOXAddress(auth, common.HexToAddress(CityNodeConfig.ToxAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(criteria, err)
}

// AdminSetCityAddress 管理员设置城市合约地址
func AdminSetCityAddress() {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	criteria, err := cityPioneer.AdminSetCityAddress(auth, common.HexToAddress(CityNodeConfig.CityAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(criteria, err)
}

// AdminSetMiningAddress 管理员设置增加用户合约余额合约地址
func AdminSetMiningAddress() {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	criteria, err := cityPioneer.AdminSetMiningAddress(auth, common.HexToAddress(CityNodeConfig.MiningAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(criteria, err)
}

// AddCityPioneerAdmin 给城市合约、IntoMining设置用户（增加用户合约余额）添加管理员权限
func AddCityPioneerAdmin() {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	cityPioneer, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	res, err := cityPioneer.AddAdmin(auth, common.HexToAddress(CityNodeConfig.CityAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	res, err = cityPioneer.AddAdmin(auth, common.HexToAddress(CityNodeConfig.MiningAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res, err)
}

// AdminSetSecondsPerDayCityPioneer 管理员设置每天秒数，用于测试
func AdminSetSecondsPerDayCityPioneer(seconds int64) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	res, err := cityPioneer.AdminSetSecondsPerDay(auth, big.NewInt(seconds))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res.Hash(), err)
}

// AdminSetPresidencyTime 管理员设置任期
func AdminSetPresidencyTime(seconds int64) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	res, err := cityPioneer.AdminSetPresidencyTime(auth, big.NewInt(seconds))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(res.Hash(), err)
}

// DepositSuretyTest 交保证金成为先锋
func DepositSuretyTest(pioneerAddress string) {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	criteria, err := cityPioneer.DepositSuretyTest(auth, common.HexToAddress(pioneerAddress))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(criteria, err)
}

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
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(CityNodeConfig.CityPioneerAddress), Cli)
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
	Cli := Client(CityNodeConfig)
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
	Cli := Client(CityNodeConfig)
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

func GetDailyRewardRecordEvent(Cli *ethclient.Client, startBlock, endBlock int64) error {
	query := event.BuildQuery(
		common.HexToAddress(CityNodeConfig.CityPioneerAddress),
		event.DailyRewardRecord,
		big.NewInt(startBlock),
		big.NewInt(endBlock),
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	logs, err := Cli.FilterLogs(ctx, query)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	cancel()
	abi, _ := intoCityNode.CityPioneerMetaData.GetAbi()

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
		err = InsertDailyReward(pioneerAddress, header.TxHash.String(), foundsReward, delegateReward, nodeReward, int64(logE.BlockNumber), timestamp)
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertDailyReward(pioneerAddress, tx_hash string, foundsReward, delegateReward, nodeReward decimal.Decimal, blockHeight, timestamp int64) error {
	InsertDailyRewardLock.Lock()
	defer InsertDailyRewardLock.Unlock()
	var reward models.Reward
	whereCondition := fmt.Sprintf("pioneer='%s' and block_height='%d'", strings.ToLower(pioneerAddress), blockHeight)
	err := db.Mysql.Table("reward").Where(whereCondition).First(&reward).Error
	if err == gorm.ErrRecordNotFound {
		db.Mysql.Table("reward").Create(&models.Reward{
			Pioneer:        pioneerAddress,
			TxHash:         tx_hash,
			FoundsReward:   foundsReward,
			DelegateReward: delegateReward,
			NodeReward:     nodeReward,
			BlockHeight:    blockHeight,
			Ctime:          time.Unix(timestamp, 0),
		})
	}
	return nil
}
