package api

import (
	"city-node-server/api/middlewares"
	"city-node-server/api/routes"
	"city-node-server/api/validate"
	"city-node-server/db"
	"github.com/gin-gonic/gin"
)

func Start() {

	//init mysql
	db.InitMysql()

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
