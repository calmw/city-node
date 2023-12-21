package main

import (
	"city-node-server/pkg/db"
	"city-node-server/services"
)

func main() {

	db.InitMysql()
	services.InitMainNet()
	//services.InitUserLocation()
	services.InitCityPioneer(86400)
	//services.InitCity(86400)
	services.InitAppraise()

}
