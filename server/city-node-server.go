package main

import (
	"city-node-server/db"
	"city-node-server/tasks"
)

func main() {
	db.InitMysql()
	tasks.Task()
}
