package main

import (
	"city-node-server/blockchain"
	"log"
)

func main() {
	//db.InitMysql()
	err, cityNum := blockchain.UserLocationGetCityNum()
	//blockchain.UserLocationNoRepeatCityIds()
	//err, cityNum := blockchain.()
	log.Println(err, cityNum)
	//for i := 0; i < int(cityNum); i++ {
	//	err, s := blockchain.UserLocationGetCity(big.NewInt(int64(i)))
	//	log.Println(err, s)
	//}
	//tasks.Task()
}
