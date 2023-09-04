package main

import (
	"city-node-server/blockchain"
	"log"
)

func main() {
	//db.InitMysql()
	err := blockchain.CityDailyTask()
	log.Println(err)
	//blockchain.UserLocationNoRepeatCityIds()
	//err, cityNum := blockchain.()
	//for i := 0; i < int(cityNum); i++ {
	//	err, s := blockchain.UserLocationGetCity(big.NewInt(int64(i)))
	//	log.Println(err, s)
	//}
	//tasks.Task()
}
