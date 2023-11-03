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
	fmt.Println(err, 5)
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
		var cid string
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
				cid = location.CityId
			}
			if j == 3 {
				pioneer.Address = strings.ToLower(colCell)
			}
			if j == 5 {
				fmt.Println(i, "表格中CityID和数据库中是否相等", cid == strings.ToLower(colCell))
			}
			if j == 4 {
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
	for i, p := range pioneers {
		//fmt.Println("")
		//fmt.Println(p.Address, p.CityId, p.CityLevel, p.Money, i)
		//if i == 7 {
		//	blockchain.AdminSetChengShiLevelAndSurety(p.CityId, p.CityLevel, p.Money)
		//	time.Sleep(time.Second * 5)
		//}
		//if i == 11 || i == 14 {
		//	blockchain.AdminSetPioneer(p.CityId, p.Address)
		//	time.Sleep(time.Second * 8)
		//}
		//blockchain.AdminSetPioneer(p.CityId, p.Address)

		//time.Sleep(time.Second * 8)
		//
		err, cityIdBytes32 := blockchain.PioneerChengShi(p.Address)
		err, level := blockchain.ChengShiLevel(cityIdBytes32)
		var ok bool
		if p.CityId == strings.ToLower("0x"+hexutils.BytesToHex(blockchain.Bytes32ToBytes(cityIdBytes32))) {
			ok = true
		}
		fmt.Println(i, p.Address, err, strings.ToLower("0x"+hexutils.BytesToHex(blockchain.Bytes32ToBytes(cityIdBytes32))), ok, level)

	}
}
