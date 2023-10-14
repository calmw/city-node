package services

import (
	"city-node-server/blockchain"
	"city-node-server/db"
	"city-node-server/models"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/status-im/keycard-go/hexutils"
	"time"
)

type CountData struct {
	Pioneer string `json:"pioneer" gorm:"column:pioneer"`
	Count   int64  `json:"count" gorm:"column:count"`
}
type CtimeData struct {
	Ctime time.Time `json:"ctime" gorm:"column:ctime"`
}
type YingFaData struct {
	Pioneer string          `json:"pioneer" gorm:"column:pioneer"`
	Yingfa  decimal.Decimal `json:"yingfa" gorm:"column:yingfa"`
}
type ShijiFafangData struct {
	CountyId string `json:"county_id" gorm:"column:county_id"`
	Count    int64  `json:"count" gorm:"column:count"`
}

func CheckData() {
	db.InitMysql()
	var pioneers []models.Pioneer
	db.Mysql.Model(&models.Pioneer{}).Find(&pioneers)

	for i, d := range pioneers {
		err, cityIdBytes32Arr := blockchain.GetCountyIdsByPioneer(d.Pioneer)
		if err != nil {
			fmt.Println(err, i)
			continue
		}

		for ia, bytesJ := range cityIdBytes32Arr {
			for ib, bytesL := range cityIdBytes32Arr {
				if hexutils.BytesToHex(blockchain.Bytes32ToBytes(bytesL)) == hexutils.BytesToHex(blockchain.Bytes32ToBytes(bytesJ)) && ia != ib {
					fmt.Println(hexutils.BytesToHex(blockchain.Bytes32ToBytes(bytesL)), hexutils.BytesToHex(blockchain.Bytes32ToBytes(bytesJ)), "----------")
				}
			}
		}
		fmt.Println(d.Pioneer, i, "没有重复")
		time.Sleep(time.Second)
	}

}
