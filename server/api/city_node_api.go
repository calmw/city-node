package main

import (
	"city-node-server/api/middlewares"
	"city-node-server/api/routes"
	"city-node-server/api/services"
	"city-node-server/api/validate"
	"city-node-server/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {

	//init mysql
	db.InitMysql()

	// init level db
	db.InitLevelDb()

	// init fdb
	db.InitFdb()

	// init cache
	//go utils.InitCache()

	go services.InitSyncTask()

	//gin bind go-playground-validator
	validate.BindingValidator()

	// gin start
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	app.Use(middlewares.Cors()) // 「 Cross domain Middleware 」
	routes.InitRoute(app)
	_ = app.Run(":8000")

}

/*
 If you change the version, you need to modify the following files'
 config/init.go
*/
