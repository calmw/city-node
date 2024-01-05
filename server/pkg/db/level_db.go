package db

import (
	"city-node-server/pkg/log"
	"github.com/syndtr/goleveldb/leveldb"
)

func InitLevelDb() {
	log.Logger.Info("Init level db")
	db, err := leveldb.OpenFile("./leveldb/db", nil)
	if err != nil {
		panic("Init level db failed")
	}
	LevelDb = db
}
