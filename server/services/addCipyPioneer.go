package services

import (
	"bufio"
	blockchain2 "city-node-server/pkg/blockchain"
	"city-node-server/pkg/db"
	models2 "city-node-server/pkg/models"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/status-im/keycard-go/hexutils"
	"github.com/xjieinfo/xjgo/xjcore/xjexcel"
	"os"
	"strconv"
	"strings"
	"time"
)

type PioneerInfo struct {
	Address   string
	CityId    string
	CityLevel int64
	Money     int64
}
type PioneerSurety struct {
	Address string
	Surety  bool
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
				var location models2.UserLocation
				where := "location like '%" + l[1] + "%'"
				err = db.Mysql.Model(models2.UserLocation{}).Where(where).First(&location).Error
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
		// 根据city_id判断该城市是否已经添加过先锋
		//var pioneer models2.Pioneer
		//err := db.Mysql.Model(models2.Pioneer{}).Where("city_id=?", p.CityId).First(&pioneer).Debug().Error
		//if err == nil {
		//	fmt.Println("该城市先锋已经存在", i)
		//	continue
		//} else {
		//	fmt.Println("该城市先锋不存在", i)
		//}
		//fmt.Println("")
		//fmt.Println(p.Address, p.CityId, p.CityLevel, p.Money, i)
		//if i == 11 {
		//blockchain2.AdminSetChengShiLevelAndSurety(p.CityId, p.CityLevel, p.Money)
		//time.Sleep(time.Second * 5)
		//}
		//if i == 15 {
		//blockchain2.AdminSetPioneer(p.CityId, p.Address)
		//time.Sleep(time.Second * 10)
		//}
		//blockchain.AdminSetPioneer(p.CityId, p.Address)

		//time.Sleep(time.Second * 8)
		//
		err, cityIdBytes32 := blockchain2.PioneerChengShi(p.Address)
		err, level := blockchain2.ChengShiLevel(cityIdBytes32)
		var ok bool
		if p.CityId == strings.ToLower("0x"+hexutils.BytesToHex(blockchain2.Bytes32ToBytes(cityIdBytes32))) {
			ok = true
		}
		fmt.Println(i, p.Address, err, strings.ToLower("0x"+hexutils.BytesToHex(blockchain2.Bytes32ToBytes(cityIdBytes32))), ok, level)

	}
}

func ReadCityFIle(cityFile string) {
	file, err := os.Open(cityFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i, line := range lines {
		fmt.Println(i, line)
		var userLocation models2.UserLocation
		//whererCondition := "location = 'Viet Nam " + line + "'"
		whererCondition := "location = 'Thailand " + line + "'"
		err = db.Mysql.Model(models2.UserLocation{}).Where(whererCondition).First(&userLocation).Debug().Error
		// 0x7b43556372050ba40a01e40a2bf9b20514e71977d04f780e38c628c7f14ac40d
		if err == nil {
			fmt.Println("该区县已经存在", i, line, userLocation.CountyId, userLocation.CityId)
			//if i == 12 || i == 17 || i == 23 || i == 25 {
			//
			//	fmt.Println("该区县已经存在", i, line, userLocation.CountyId, userLocation.CityId)
			//	blockchain2.SetCityMapping(userLocation.CountyId, "0x7b43556372050ba40a01e40a2bf9b20514e71977d04f780e38c628c7f14ac40d", "pYnBembRMi+rRswBxe+wsg==")
			//	time.Sleep(time.Second * 15)
			//}
		} else {
			fmt.Println("该城市不存在", i, line)
		}

	}
}

func CheckPioneer(excelFile string) {

	type Export struct {
		City      string `json:"city" excel:"column:B;desc:城市;width:30"`
		CityLevel string `json:"city_level" excel:"column:C;desc:城市等级;width:30"`
		Pioneer   string `json:"pioneer" excel:"column:D;desc:先锋地址;width:30"`
		Deopsit   string `json:"deopsit" excel:"column:E;desc:是否交保证金;width:30"`
	}
	f, err := excelize.OpenFile(excelFile)
	fmt.Println(err, 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 取得 Sheet1 表格中所有的行
	rows := f.GetRows("Sheet1")
	var exports []Export
	for i, row := range rows {
		if i == 0 {
			continue
		}
		export := Export{}
		for j, colCell := range row {
			if j == 0 {
				export.City = colCell
			} else if j == 1 {
				export.CityLevel = colCell
			} else if j == 2 {
				export.Pioneer = strings.ToLower(colCell)
				var pioneer models2.Pioneer
				err := db.Mysql.Model(models2.Pioneer{}).Where("pioneer=?", export.Pioneer).First(&pioneer).Debug().Error
				if err == nil {
					fmt.Println("该城市先锋已经存在", i)
					export.Deopsit = "是"
				} else {
					fmt.Println("该城市先锋不存在", i)
					export.Deopsit = "未交保证金"
				}
			}
		}
		exports = append(exports, export)
	}
	// 检查数据库中交过保证金的，是否在表格中
	var pioneers []models2.Pioneer
	err = db.Mysql.Model(models2.Pioneer{}).Find(&pioneers).Debug().Error
	if err == nil {
		for _, p := range pioneers {
			exits := false
			for _, e := range exports {
				if p.Pioneer == e.Pioneer {
					exits = true
				}
			}
			fmt.Println(p.Pioneer, exits)
		}
	}
	// 保存excel
	//"github.com/xjieinfo/xjgo/xjcore/xjexcel"
	excel := xjexcel.ListToExcel(exports, "城市先锋-用户信息", "先锋详情")
	fileName := fmt.Sprintf("./城市先锋-%s.xls", time.Now().Format("2006-01-02"))
	excel.SaveAs(fileName)
}

func CheckPioneer2(excelFile, excelFile2 string) {

	f, err := excelize.OpenFile(excelFile)
	fmt.Println(err)
	if err != nil {
		return
	}
	f2, err := excelize.OpenFile(excelFile2)
	fmt.Println(err, 2)
	if err != nil {
		return
	}
	// 取得 Sheet1 表格中所有的行
	rows := f.GetRows("Sheet1")
	var pioneers []string
	for i, row := range rows {
		if i == 0 {
			continue
		}
		for j, colCell := range row {
			if j == 2 {
				pioneers = append(pioneers, strings.ToLower(colCell))
			}
		}
	}
	// 取得 Sheet1 表格中所有的行
	rows = f2.GetRows("Sheet1")
	var pioneers2 []string
	for i, row := range rows {
		if i == 0 || i == 1 {
			continue
		}
		for j, colCell := range row {
			if j == 3 {
				pioneers2 = append(pioneers2, strings.ToLower(colCell))
			}
		}
	}
	exitMap := map[string]bool{}
	var p3 []string
	//for _, p2 := range pioneers2 {
	//	exitMap[p2] = true
	//}
	// 检查2中先锋是否在1中存在
	for i, p2 := range pioneers2 {
		_, ok := exitMap[p2]
		if !ok {
			p3 = append(p3, p2)
			exitMap[p2] = true
		} else {
			fmt.Println(p2)

		}
		exits := false
		for _, p1 := range pioneers {
			if p2 == p1 {
				exits = true
			}
		}
		if !exits {
			fmt.Println(i, p2, exits)
		}
		fmt.Println(i, p2, exits)
	}
	fmt.Println("---------", len(pioneers2), len(pioneers), len(p3))
	// 检查1中先锋是否在2中存在
	for i, p1 := range pioneers {
		exits := false
		for _, p2 := range pioneers2 {
			if p2 == p1 {
				exits = true
			}
		}
		if !exits {
			fmt.Println(i, p1, exits)
		}
	}

}

func CheckLocation(excelFile string) {
	f, err := excelize.OpenFile(excelFile)
	fmt.Println(err, 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 取得 Sheet1 表格中所有的行
	rows := f.GetRows("整理后")
	fmt.Println(len(rows), 6)
loop:
	for i, row := range rows {
		if i == 0 {
			continue
		}
		for j, colCell := range row {
			if j == 3 {
				if colCell != "Bangkok" {
					continue loop
				}
			}
			if j == 4 {
				//fmt.Println(colCell)
				var userLocation models2.UserLocation
				whererCondition := "location = 'Thailand " + colCell + "'"
				err := db.Mysql.Model(models2.UserLocation{}).Where(whererCondition).First(&userLocation).Debug().Error
				if err == nil {
					fmt.Println("该区县已经存在", i, colCell)
				} else {
					//fmt.Println("该城市不存在", i, colCell)
				}
			}
		}
	}
}
