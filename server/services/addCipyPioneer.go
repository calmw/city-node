package services

import (
	"city-node-server/blockchain"
	"city-node-server/db"
	"city-node-server/models"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/status-im/keycard-go/hexutils"
	"strconv"
	"strings"
)

type PioneerInfo struct {
	Address   string
	CityId    string
	CityLevel int64
	Money     int64
}

func ReadExcel(excelFile string) {
	var pioneers []PioneerInfo
	f, err := excelize.OpenFile(excelFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 取得 Sheet1 表格中所有的行
	rows := f.GetRows("Sheet1")
	for i, row := range rows {
		if i == 1 || i == 0 {
			continue
		}
		pioneer := PioneerInfo{}
		for j, colCell := range row {
			if j == 1 {
				l := strings.Split(colCell, " ")
				var location models.UserLocation
				where := "location like '%" + l[1] + "%'"
				err = db.Mysql.Model(models.UserLocation{}).Where(where).First(&location).Error
				if err != nil {
					return
				}
				pioneer.CityId = location.CityId
			}
			if j == 3 {
				pioneer.Address = strings.ToLower(colCell)
			}
			if j == 8 {
				level, _ := strconv.ParseInt(colCell, 10, 64)
				pioneer.CityLevel = level
				if level == 1 {
					pioneer.Money = 100000
				} else if level == 2 {
					pioneer.Money = 60000
				} else if level == 3 {
					pioneer.Money = 40000
				}
			}
		}
		pioneers = append(pioneers, pioneer)
	}
	for _, p := range pioneers {
		//fmt.Println(p.Address, p.CityId, p.CityLevel, p.Money)
		//blockchain.AdminSetChengShiLevelAndSurety(p.CityId, p.CityLevel, p.Money)
		//time.Sleep(time.Second * 3)
		//blockchain.AdminSetPioneer(p.CityId, p.Address)
		//
		err, city := blockchain.PioneerChengShi(p.Address)
		fmt.Println(p.Address, err, hexutils.BytesToHex(blockchain.Bytes32ToBytes(city)))

	}
}
