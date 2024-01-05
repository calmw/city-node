package db

import (
	"github.com/syndtr/goleveldb/leveldb"
	"gorm.io/gorm"
)

var Mysql *gorm.DB
var LevelDb *leveldb.DB
