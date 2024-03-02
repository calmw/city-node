package db

import (
	"city-node-server/pkg/log"
	"fmt"
	"github.com/calmw/fdb"
	"github.com/calmw/fdb/redis"
)

func InitFdb() {
	log.Logger.Info("Init fdb db")
	// 打开fdb redis
	var err error
	FDB, err = redis.NewRedisDataStructure(fdb.DefaultOption)
	if err != nil {
		panic(fmt.Sprintf("Init fdb db failed:%v", err))
	}
	log.Logger.Info("Init fdb success")
}
