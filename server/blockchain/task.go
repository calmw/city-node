package blockchain

import (
	intoCityNode "city-node-server/binding"
	"city-node-server/log"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

// CityDailyTask 城市合约定时任务
func CityDailyTask() error {
	Cli := Client(CityNodeConfig)
	_, auth := GetAuth(Cli)
	userLocation, err := intoCityNode.NewUserLocation(common.HexToAddress(CityNodeConfig.UserLocationAddress), Cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	cityNumber, err := userLocation.GetCityNumber(nil)
	fmt.Println(cityNumber.Int64(), err)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	numberOfTimes := 1000
	forTimes := int(cityNumber.Int64() / int64(numberOfTimes))
	remainder := int(cityNumber.Int64() % int64(numberOfTimes))

	city, err := intoCityNode.NewCity(common.HexToAddress(CityNodeConfig.CityAddress), Cli)
	for i := 0; i < forTimes; i++ {
		//fmt.Println(int64(forTimes*numberOfTimes), int64((forTimes+1)*numberOfTimes))
		task, err := city.DailyTask(auth, big.NewInt(int64(i*numberOfTimes)), big.NewInt(int64((i+1)*numberOfTimes)))
		fmt.Println(int64(i*numberOfTimes), int64((i+1)*numberOfTimes), err)
		if err != nil {
			log.Logger.Sugar().Error(err)
			return err
		}
		log.Logger.Sugar().Info("hash: ", task.Hash())
		time.Sleep(time.Second * 10)
	}
	if remainder > 0 {
		task, err := city.DailyTask(auth, big.NewInt(cityNumber.Int64()-int64(remainder)), cityNumber)
		if err != nil {
			log.Logger.Sugar().Error(err)
			return err
		}
		fmt.Println(cityNumber.Int64()-int64(remainder), cityNumber.Int64())
		log.Logger.Sugar().Info("hash: ", task.Hash())
	}
	log.Logger.Sugar().Info("success")
	return nil
}

//func CityPioneerDailyTask() error {
//	Cli := Client(CityNodeConfig)
//	_, auth := GetAuth(Cli)
//	cityPioneer, err := intoCityNode.NewCityPioneer(common.HexToAddress(""), Cli)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return err
//	}
//	task, err := cityPioneer.DailyTask(auth)
//	if err != nil {
//		log.Logger.Sugar().Error(err)
//		return err
//	}
//	log.Logger.Sugar().Info("hash: ", task.Hash())
//	return nil
//}
