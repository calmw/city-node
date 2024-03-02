package db

import (
	"github.com/calmw/fdb/redis"
	"github.com/syndtr/goleveldb/leveldb"
	"gorm.io/gorm"
)

var Mysql *gorm.DB
var LevelDb *leveldb.DB
var FDB *redis.RedisDataStructure
