package main

import (
	"city-node-server/blockchain"
)

func main() {
	//db.InitMysql()
	//err := blockchain.CityDailyTask()
	//log.Println(err)
	//blockchain.UserLocationNoRepeatCityIds()
	//err, cityNum := blockchain.()
	//for i := 0; i < int(cityNum); i++ {
	//	err, s := blockchain.UserLocationGetCity(big.NewInt(int64(i)))
	//	log.Println(err, s)
	//}
	//tasks.Task()
	// 设置城市先锋考核地址
	blockchain.AdminSetAssessmentCriteria(1, 1, 2500)
	blockchain.AdminSetAssessmentCriteria(1, 2, 5000)
	blockchain.AdminSetAssessmentCriteria(1, 3, 10000)
	blockchain.AdminSetAssessmentCriteria(2, 1, 1250)
	blockchain.AdminSetAssessmentCriteria(2, 2, 2500)
	blockchain.AdminSetAssessmentCriteria(2, 3, 5000)
	blockchain.AdminSetAssessmentCriteria(3, 1, 625)
	blockchain.AdminSetAssessmentCriteria(3, 2, 125)
	blockchain.AdminSetAssessmentCriteria(3, 3, 2500)
}
