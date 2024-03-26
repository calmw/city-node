package services

import (
	"bufio"
	"city-node-server/api/utils"
	blockchain2 "city-node-server/pkg/blockchain"
	"city-node-server/pkg/db"
	"city-node-server/pkg/log"
	"city-node-server/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/status-im/keycard-go/hexutils"
	"github.com/xjieinfo/xjgo/xjcore/xjexcel"
	"os"
	"strconv"
	"strings"
	"time"
)

type PioneerInfo struct {
	Address   string
	AreaId    string
	CityLevel int64
	Money     int64
}
type PioneerSurety struct {
	Address string
	Surety  bool
}

type Pioneer struct {
	AreaName       string `json:"area_name" excel:"column:B;desc:先锋地区;width:30"`
	AreaLevel      int64  `json:"area_level" excel:"column:C;desc:先锋等级;width:30"`
	PioneerBatch   int64  `json:"pioneer_batch" excel:"column:D;desc:批次;width:30"`
	AreaId         string `json:"area_id" excel:"column:E;desc:地区ID;width:30"`
	IsAreaNode     int64  `json:"is_area_node" excel:"column:F;desc:是否是区域先锋;width:30"`
	PioneerAddress string `json:"pioneer_address" excel:"column:G;desc:先锋钱包地址;width:30"`
}

// AddPioneerBeth3 添加三期先锋,四期上线前海需要用
func AddPioneerBeth3() {
	pioneerBytes, err := os.ReadFile("./assets/新节点添加.json")
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	var pioneer Pioneer
	err = json.Unmarshal(pioneerBytes, &pioneer)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	var location models.UserLocation
	where := "location like '%" + pioneer.AreaName + "%'"
	err = db.Mysql.Model(models.UserLocation{}).Where(where).First(&location).Error
	if err != nil {
		return
	}
	cid := location.CityId

	pioneer.PioneerAddress = strings.ToLower(pioneer.PioneerAddress)
	pioneer.AreaId = strings.ToLower(pioneer.AreaId)

	fmt.Println("表格中CityID和数据库中是否相等", cid == strings.ToLower(pioneer.AreaId))
	if pioneer.IsAreaNode != 0 {
		panic("IsAreaNode error")
	}
	var suretyTOX int64
	level := pioneer.AreaLevel
	if level == 1 {
		suretyTOX = 100000
	} else if level == 2 {
		suretyTOX = 60000
	} else if level == 3 {
		suretyTOX = 40000
	} else {
		panic("level error")
	}

	// 根据city_id判断该城市是否已经添加过先锋
	var pioneerModel models.Pioneer
	err = db.Mysql.Model(models.Pioneer{}).Where("city_id=?", pioneer.AreaId).First(&pioneerModel).Debug().Error
	if err == nil {
		panic("该城市先锋已经存在")
	} else {
		fmt.Println("该城市先锋不存在")
	}
	fmt.Println(fmt.Sprintf("address:%s,cityId:%s,cityLevel:%s,suretyTOX:%d", pioneer.PioneerAddress, pioneer.AreaId, pioneer.AreaLevel, suretyTOX))

	//for true {
	//	err = blockchain2.AdminSetChengShiLevelAndSurety(pioneer.AreaId, level, suretyTOX)
	//	if err == nil {
	//		break
	//	}
	//}
	//for true {
	//	err = blockchain2.AdminSetPioneer(pioneer.AreaId, pioneer.PioneerAddress)
	//	if err == nil {
	//		break
	//	}
	//}

	err, cityIdBytes32 := blockchain2.PioneerChengShi(pioneer.PioneerAddress)
	err, levelChain := blockchain2.ChengShiLevel(cityIdBytes32)
	fmt.Println(err, level, strings.ToLower("0x"+hexutils.BytesToHex(blockchain2.Bytes32ToBytes(cityIdBytes32))), 555)
	var ok bool
	if pioneer.AreaId == strings.ToLower("0x"+hexutils.BytesToHex(blockchain2.Bytes32ToBytes(cityIdBytes32))) {
		ok = true
	}
	fmt.Println(pioneer.PioneerAddress, err, strings.ToLower("0x"+hexutils.BytesToHex(blockchain2.Bytes32ToBytes(cityIdBytes32))), ok, levelChain)

}

// AddPioneerBeth4 添加四期先锋，保证金USDT+TOX
func AddPioneerBeth4() {
	pioneerBytes, err := os.ReadFile("./assets/新节点添加.json")
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	var pioneer Pioneer
	err = json.Unmarshal(pioneerBytes, &pioneer)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	if pioneer.PioneerBatch != 4 {
		panic("pioneer batch error")
	}

	var location models.UserLocation
	where := "location like '%" + pioneer.AreaName + "%'"
	if pioneer.IsAreaNode == 1 {
		where = "county like '%" + pioneer.AreaName + "%'"
	}
	err = db.Mysql.Model(models.UserLocation{}).Where(where).First(&location).Error
	if err != nil {
		return
	}
	aid := location.CityId
	if pioneer.IsAreaNode == 1 {
		aid = location.CountyId
	}
	pioneer.PioneerAddress = strings.ToLower(pioneer.PioneerAddress)
	pioneer.AreaId = strings.ToLower(pioneer.AreaId)

	if pioneer.IsAreaNode == 1 {
		fmt.Println("表格中CountyID和数据库中是否相等", aid == pioneer.AreaId)
	} else {

		fmt.Println("表格中CityID和数据库中是否相等", aid == pioneer.AreaId)
	}
	if aid != pioneer.AreaId {
		panic("表格中CountyID和数据库中不相等")
	}

	var suretyTOX int64
	var suretyUSDT int64
	level := pioneer.AreaLevel
	if pioneer.IsAreaNode == 1 {
		if level == 1 {
			suretyTOX = 5000 // 主网
			suretyUSDT = 1000
			//suretyTOX = 10 // 测试网
			//suretyUSDT = 5
		} else if level == 2 {
			suretyTOX = 3000
			suretyUSDT = 600
			//suretyTOX = 6
			//suretyUSDT = 3
		} else {
			panic("level error")
		}
	} else {
		if level == 1 {
			suretyTOX = 100000
			suretyUSDT = 10000
			//suretyTOX = 100
			//suretyUSDT = 100
		} else if level == 2 {
			suretyTOX = 60000
			suretyUSDT = 8000
			//suretyTOX = 80
			//suretyUSDT = 60
		} else if level == 3 {
			suretyTOX = 40000
			suretyUSDT = 6000
			//suretyTOX = 60
			//suretyUSDT = 40
		} else {
			panic("level error")
		}
	}

	// 根据city_id判断该城市是否已经添加过先锋
	var pioneerModel models.Pioneer
	err = db.Mysql.Model(models.Pioneer{}).Where("city_id=?", pioneer.AreaId).First(&pioneerModel).Debug().Error
	if err == nil {
		panic("该城市先锋已经存在")
	} else {
		fmt.Println("该城市先锋不存在")
	}
	fmt.Println(fmt.Sprintf("address:%s,areaId:%s,cityLevel:%s,suretyTOX:%d", pioneer.PioneerAddress, pioneer.AreaId, pioneer.AreaLevel, suretyTOX))

	for {
		err = AddPioneerBatch4(
			pioneer.AreaId,
			pioneer.PioneerAddress,
			level,
			suretyTOX,
			suretyUSDT,
			4,
			pioneer.IsAreaNode)
		if err == nil {
			break
		}
	}

	err, cityIdBytes32 := blockchain2.PioneerChengShi(pioneer.PioneerAddress)
	err, levelChain := blockchain2.ChengShiLevel(cityIdBytes32)
	var ok bool
	if pioneer.AreaId == strings.ToLower(blockchain2.ConvertAreaIdBtA(cityIdBytes32)) {
		ok = true
	}
	fmt.Println(pioneer.PioneerAddress, err, strings.ToLower("0x"+hexutils.BytesToHex(blockchain2.Bytes32ToBytes(cityIdBytes32))), ok, levelChain)
}

func AddPioneerBeth4FromExcel(pioneer Pioneer) {
	if pioneer.PioneerBatch != 4 {
		panic("pioneer batch error")
	}

	var suretyTOX int64
	var suretyUSDT int64
	level := pioneer.AreaLevel
	if pioneer.IsAreaNode == 1 {
		if level == 1 {
			suretyTOX = 5000 // 主网
			suretyUSDT = 1000
		} else if level == 2 {
			suretyTOX = 3000
			suretyUSDT = 600
		} else {
			panic("level error")
		}
	} else {
		if level == 1 {
			suretyTOX = 100000
			suretyUSDT = 10000
		} else if level == 2 {
			suretyTOX = 60000
			suretyUSDT = 8000
		} else if level == 3 {
			suretyTOX = 40000
			suretyUSDT = 6000
		} else {
			panic("level error")
		}
	}

	// 根据city_id判断该城市是否已经添加过先锋
	var pioneerModel models.Pioneer
	err := db.Mysql.Model(models.Pioneer{}).Where("city_id=?", pioneer.AreaId).First(&pioneerModel).Debug().Error
	if err == nil {
		panic("该城市先锋已经存在")
	} else {
		fmt.Println("该城市先锋不存在")
	}

	fmt.Println(fmt.Sprintf("地区：%s,address:%s,areaId:%s,cityLevel:%d,suretyTOX:%d,suretyUSDT:%d,IsAreaNode:%d", pioneer.AreaName, pioneer.PioneerAddress, pioneer.AreaId, pioneer.AreaLevel, suretyTOX, suretyUSDT, pioneer.IsAreaNode))

	for {
		err = AddPioneerBatch4(
			pioneer.AreaId,
			pioneer.PioneerAddress,
			level,
			suretyTOX,
			suretyUSDT,
			4,
			pioneer.IsAreaNode,
		)

		//fmt.Println(err, pioneer.AreaName, strings.Contains(err.Error(), "can not set any more"), 6666666)

		if err == nil || strings.Contains(err.Error(), "can not set any more") {
			break
		}
		if err == nil || strings.Contains(err.Error(), "you are global node") {
			log.Logger.Sugar().Info("全球节点", pioneer)
			break
		}
	}

	//err, cityIdBytes32 := blockchain2.PioneerChengShi(pioneer.PioneerAddress)
	//err, levelChain := blockchain2.ChengShiLevel(cityIdBytes32)
	//var ok bool
	//if pioneer.AreaId == strings.ToLower(blockchain2.ConvertAreaIdBtA(cityIdBytes32)) {
	//	ok = true
	//}
	//fmt.Println(pioneer.PioneerAddress, err, strings.ToLower("0x"+hexutils.BytesToHex(blockchain2.Bytes32ToBytes(cityIdBytes32))), ok, levelChain)
}

func AddPioneerBeth4FromDb() {

	var pioneerAddInfos []models.PioneerAddInfo

	err := db.Mysql.Model(&models.PioneerAddInfo{}).Where("is_set=? and pioneer_status=?", 0, 0).Find(&pioneerAddInfos).Error

	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	err, cli := blockchain2.Client(blockchain2.CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	city := blockchain2.NewCity(cli)

	for _, info := range pioneerAddInfos {
		if info.PioneerBatch != 4 {
			panic("pioneer batch error")
		}

		var suretyTOX int64
		var suretyUSDT int64
		level := info.AreaLevel
		if info.IsAreaNode == 1 {
			if level == 1 {
				suretyTOX = 5000 // 主网
				suretyUSDT = 1000
			} else if level == 2 {
				suretyTOX = 3000
				suretyUSDT = 600
			} else {
				panic("level error")
			}
		} else {
			if level == 1 {
				suretyTOX = 100000
				suretyUSDT = 10000
			} else if level == 2 {
				suretyTOX = 60000
				suretyUSDT = 8000
			} else if level == 3 {
				suretyTOX = 40000
				suretyUSDT = 6000
			} else {
				panic("level error")
			}
		}

		fmt.Println(fmt.Sprintf("地区：%s,address:%s,areaId:%s,areaLevel:%d,suretyTOX:%d,suretyUSDT:%d,IsAreaNode:%d", info.Location, info.Pioneer, info.AreaId, info.AreaLevel, suretyTOX, suretyUSDT, info.IsAreaNode))

		// 重置,需要重置，并且改先锋地址之前是已经失效的先锋
		if info.IsReset == 1 && info.OldAreaId != "" {
			for {
				err = RemovePioneer(
					info.OldAreaId,
					info.Pioneer,
				)
				//err = RemovePioneer(
				//	"0x8791cf7226e6b3c30f0557f8f0aaa201d606061d431d70458876f22dca6bb210",
				//	"0xeaddeef3df4687d26bf6e160515abc424c55d4f4",
				//)
				if err == nil {
					fmt.Println("重置先锋成功", info.OldAreaId, info.Pioneer)
					break
				}
			}
			//return
		}

		// 添加
		for {
			err = AddPioneerBatch4(
				info.AreaId,
				info.Pioneer,
				level,
				suretyTOX,
				suretyUSDT,
				4,
				info.IsAreaNode,
			)

			if err == nil || strings.Contains(err.Error(), "can not set any more") {
				fmt.Println("设置先锋成功", info.OldAreaId, info.Pioneer, err)

				for i := 0; i < 10; i++ {
					pioneerCity, err := city.PioneerCity(info.Pioneer)
					if err == nil && pioneerCity != "0x0000000000000000000000000000000000000000000000000000000000000000" {
						fmt.Println("检测先锋成功", info.Location, info.Pioneer)
						break
					}
					if i == 9 {
						fmt.Println("检测先锋失败", info.Location, info.Pioneer)
					}
				}

				break
			}
			if err == nil || strings.Contains(err.Error(), "you are global node") {
				log.Logger.Sugar().Info("全球节点")
				fmt.Println("设置先锋成功", info.OldAreaId, info.Pioneer, err)
				for i := 0; i < 10; i++ {
					pioneerCity, err := city.PioneerCity(info.Pioneer)
					if err == nil && pioneerCity != "0x0000000000000000000000000000000000000000000000000000000000000000" {
						fmt.Println("检测先锋成功", info.Location, info.Pioneer)
						break
					}
					if i == 9 {
						fmt.Println("检测先锋失败", info.Location, info.Pioneer)
					}
				}

				break
			}

			fmt.Println("设置先锋失败", info.OldAreaId, info.Pioneer, err)
		}

		db.Mysql.Model(&models.PioneerAddInfo{}).Where("pioneer=?", info.Pioneer).Update("is_set", 1)

	}
}

type PioneerNotPay struct {
	User      string `json:"user" excel:"column:B;desc:用户地址;width:30"`
	Area      string `json:"area" excel:"column:C;desc:用户地区;width:30"`
	IsArea    string `json:"is_area" excel:"column:D;desc:先锋类型;width:30"`
	AreaLevel int64  `json:"area_level" excel:"column:E;desc:地区等级;width:30"`
}

func AddPioneerBeth4FromDb2() {

	var pioneers []models.Pioneer

	err := db.Mysql.Model(&models.Pioneer{}).Where("surety_usdt>0").Find(&pioneers).Error

	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	//err, cli := blockchain2.Client(blockchain2.CityNodeConfig)
	//if err != nil {
	//	log.Logger.Sugar().Error(err)
	//	return
	//}
	// 设置所需缴纳的TOX，城市等级
	//cityPioneer := blockchain2.NewCityPioneer(cli)
	var pioneerNotPays []PioneerNotPay
	for _, info := range pioneers {
		pt := "城市先锋"
		if info.IsAreaNode > 0 {
			pt = "区域先锋"
		}
		pioneerNotPays = append(pioneerNotPays, PioneerNotPay{
			User:      info.Pioneer,
			Area:      info.Location,
			IsArea:    pt,
			AreaLevel: info.AreaLevel,
		})
	}

	// 保存excel
	f := xjexcel.ListToExcel(pioneerNotPays, "四期交保证金用户", "数据详情")
	f.SaveAs(fmt.Sprintf("四期交保证金用户.xlsx"))
}

func ReadExcel(excelFile string) {
	var pioneers []Pioneer
	var noLocationPioneers []Pioneer
	f, err := excelize.OpenFile(excelFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 取得 Sheet1 表格中所有的行
	rows, err := f.GetRows("总表")

forOut:
	for i, row := range rows {
		if i == 1 || i == 0 {
			continue
		}
		hasLocation := false
		pioneer := Pioneer{PioneerBatch: 4}
		for j, colCell := range row {
			var location models.UserLocation
			if j == 3 {
				colCell = strings.TrimSpace(colCell)
				area := strings.Split(colCell, " ")
				var where string
				if len(area) == 1 {
					pioneer.AreaName = area[0]
					where = "city like '%" + area[0] + "%'"
				} else if len(area) == 2 {
					pioneer.AreaName = area[0] + " " + area[1]
					where = "city like '%" + area[0] + "%' and county like '%" + area[1] + "%'"
				} else {
					panic("地区错误")
				}

				err = db.Mysql.Model(models.UserLocation{}).Where(where).First(&location).Error
				if err == nil {
					hasLocation = true
					if len(area) == 1 {
						pioneer.AreaId = location.CityId
					} else if len(area) == 2 {
						pioneer.AreaId = location.CountyId
					}
				} else {
					fmt.Println(colCell, "该地区有问题", 123)
				}

				if colCell == "德阳市" {
					fmt.Println(pioneer.AreaName, len(area), pioneer.AreaId, 99990)
					continue forOut
				}

			}
			if j == 4 {
				pioneer.PioneerAddress = strings.ToLower(colCell)
			}
			if j == 5 {
				if colCell == "一线区县节点" {
					pioneer.IsAreaNode = 1
					pioneer.AreaLevel = 1
				} else if colCell == "二线区县节点" {
					pioneer.IsAreaNode = 1
					pioneer.AreaLevel = 2
				} else if colCell == "一线城市节点" {
					pioneer.IsAreaNode = 0
					pioneer.AreaLevel = 1
				} else if colCell == "二线城市节点" {
					pioneer.IsAreaNode = 0
					pioneer.AreaLevel = 2
				} else if colCell == "三线城市节点" {
					pioneer.IsAreaNode = 0
					pioneer.AreaLevel = 3
				}
			}
		}
		if hasLocation {
			pioneers = append(pioneers, pioneer)
		} else {
			noLocationPioneers = append(noLocationPioneers, pioneer)
		}
	}

	f2 := xjexcel.ListToExcel(pioneers, "先锋信息", "先锋详情")
	fileName := fmt.Sprintf("./先锋信息.xlsx")
	f2.SaveAs(fileName)

	f3 := xjexcel.ListToExcel(noLocationPioneers, "先锋信息", "先锋详情")
	fileName = fmt.Sprintf("./更改后依然有问题地区信息.xlsx")
	f3.SaveAs(fileName)

	for i, p := range pioneers {
		if p.PioneerAddress != strings.ToLower("0x17E56C5f4E271a2Ce0920580784C6397e247C9d9") {
			continue
		}
		// 根据city_id判断该城市是否已经添加过先锋
		var pioneer models.Pioneer
		err = db.Mysql.Model(models.Pioneer{}).Where("city_id=?", p.AreaId).First(&pioneer).Debug().Error
		if err == nil {
			fmt.Println("该城市先锋已经存在", i)
			continue
		} else {
			fmt.Println("该城市先锋不存在", i)
		}
		fmt.Println(p.AreaName, p.AreaLevel, p.AreaId, p.PioneerBatch, p.PioneerAddress, p.IsAreaNode, i)
		//if i < 12 {
		//	continue
		//}

		AddPioneerBeth4FromExcel(p)
	}
}

func ReadExcel5(excelFile string) {
	f, err := excelize.OpenFile(excelFile)
	fmt.Println(err, 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 取得 Sheet1 表格中所有的行
	rows, err := f.GetRows("Sheet1")
	for i, row := range rows {
		if i == 1 || i == 0 {
			continue
		}
		for j, colCell := range row {
			if j == 5 {
				fmt.Println(colCell)
				var location models.UserLocation
				where := "location like '%" + colCell + "%'"
				err = db.Mysql.Model(models.UserLocation{}).Where(where).First(&location).Error
				fmt.Println(err)
				if err != nil {
					return
				}
			}

			//if j == 4 {
			//	fmt.Println(colCell)
			//}
		}
		//pioneers = append(pioneers, pioneer)
	}
	//for i, p := range pioneers {
	//	// 根据city_id判断该城市是否已经添加过先锋
	//	var pioneer models.Pioneer
	//	err = db.Mysql.Model(models.Pioneer{}).Where("city_id=?", p.AreaId).First(&pioneer).Debug().Error
	//	if err == nil {
	//		fmt.Println("该城市先锋已经存在", i)
	//		continue
	//	} else {
	//		fmt.Println("该城市先锋不存在", i)
	//	}
	//	//fmt.Println("")
	//	fmt.Println(p.Address, p.AreaId, p.CityLevel, p.Money, i)
	//	//if i == 11 {
	//	//blockchain2.AdminSetChengShiLevelAndSurety(p.AreaId, p.CityLevel, p.Money)
	//	//time.Sleep(time.Second * 5)
	//	//}
	//	//if i == 15 {
	//	//blockchain2.AdminSetPioneer(p.AreaId, p.Address)
	//	//time.Sleep(time.Second * 10)
	//	//}
	//	//blockchain.AdminSetPioneer(p.AreaId, p.Address)
	//
	//	//time.Sleep(time.Second * 8)
	//	//
	//	//err, cityIdBytes32 := blockchain2.PioneerChengShi(p.Address)
	//	//err, level := blockchain2.ChengShiLevel(cityIdBytes32)
	//	//var ok bool
	//	//if p.AreaId == strings.ToLower("0x"+hexutils.BytesToHex(blockchain2.Bytes32ToBytes(cityIdBytes32))) {
	//	//	ok = true
	//	//}
	//	//fmt.Println(i, p.Address, err, strings.ToLower("0x"+hexutils.BytesToHex(blockchain2.Bytes32ToBytes(cityIdBytes32))), ok, level)
	//
	//}
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
		var userLocation models.UserLocation
		//whererCondition := "location = 'Viet Nam " + line + "'"
		whererCondition := "location = 'Thailand " + line + "'"
		err = db.Mysql.Model(models.UserLocation{}).Where(whererCondition).First(&userLocation).Debug().Error
		// 0x7b43556372050ba40a01e40a2bf9b20514e71977d04f780e38c628c7f14ac40d
		if err == nil {
			fmt.Println("该区县已经存在", i, line, userLocation.CountyId, userLocation.CityId)
			//if i == 12 || i == 17 || i == 23 || i == 25 {
			//
			//	fmt.Println("该区县已经存在", i, line, userLocation.CountyId, userLocation.AreaId)
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
		City           string `json:"city" excel:"column:B;desc:城市;width:30"`
		CityLevel      string `json:"city_level" excel:"column:C;desc:城市等级;width:30"`
		Pioneer        string `json:"pioneer" excel:"column:D;desc:先锋地址;width:60"`
		Deposit        string `json:"deposit" excel:"column:E;desc:是否交保证金;width:30"`
		Batch          string `json:"batch" excel:"column:F;desc:批次;width:30"`
		IsReturnMargin string `json:"is_return_margin" excel:"column:G;desc:是否退保证金;width:30"`
	}
	f, err := excelize.OpenFile(excelFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 取得 Sheet1 表格中所有的行
	rows, err := f.GetRows("Sheet1")
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
				var pioneer models.Pioneer
				err = db.Mysql.Model(models.Pioneer{}).Where("pioneer=?", export.Pioneer).First(&pioneer).Debug().Error
				if err == nil {
					fmt.Println("该城市先锋已经存在", i)
					export.Deposit = "是"
				} else {
					fmt.Println("该城市先锋不存在", i)
					export.Deposit = "未交"
				}
			} else if j == 6 {
				export.Batch = colCell
				if colCell == "1" {
					export.IsReturnMargin = "是"
				} else {
					export.IsReturnMargin = "否"
				}
			}
		}
		exports = append(exports, export)
	}
	// 检查数据库中交过保证金的，是否在表格中
	var pioneers []models.Pioneer
	err = db.Mysql.Model(models.Pioneer{}).Find(&pioneers).Debug().Error
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

func CheckPioneer4(excelFile string) {

	type Export struct {
		City           string `json:"city" excel:"column:B;desc:城市;width:30"`
		CityLevel      string `json:"city_level" excel:"column:C;desc:城市等级;width:30"`
		Pioneer        string `json:"pioneer" excel:"column:D;desc:先锋地址;width:60"`
		Deposit        string `json:"deposit" excel:"column:E;desc:是否交保证金;width:30"`
		Batch          string `json:"batch" excel:"column:F;desc:批次;width:30"`
		IsReturnMargin string `json:"is_return_margin" excel:"column:G;desc:是否退保证金;width:30"`
		UserCount      string `json:"user_count" excel:"column:H;desc:绑定城市的用户数量;width:30"`
	}
	f, err := excelize.OpenFile(excelFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 取得 Sheet1 表格中所有的行
	rows, err := f.GetRows("Sheet1")
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
				var pioneer models.Pioneer
				err = db.Mysql.Model(models.Pioneer{}).Where("pioneer=?", export.Pioneer).First(&pioneer).Debug().Error
				if err == nil {
					fmt.Println("该城市先锋已经存在", i)
					export.Deposit = "是"
				} else {
					fmt.Println("该城市先锋不存在", i)
					export.Deposit = "未交"
				}
				//fmt.Println(pioneer.CityId, "----")
				err, chengShiIdBytes32 := blockchain2.PioneerChengShi(export.Pioneer)
				if err != nil {
					fmt.Println("获取先锋城市错误", i)
					return
				}
				err, cIds := blockchain2.GetCountyIdsByChengShiId(strings.ToLower("0x" + hexutils.BytesToHex(utils.Bytes32ToBytes(chengShiIdBytes32))))
				if err != nil {
					fmt.Println("获取先锋城市所对应区县错误", i, err)
					return
				}
				var userTotal int64
				for _, cId := range cIds {
					var userCount int64
					err = db.Mysql.Model(models.UserLocation{}).Where("county_id=?", cId).Count(&userCount).Error
					fmt.Printf("county_id:%s,count:%d /n", cId, userCount)
					userTotal += userCount
				}
				export.UserCount = strconv.FormatInt(userTotal, 10)
			} else if j == 6 {
				export.Batch = colCell
				if colCell == "1" {
					export.IsReturnMargin = "是"
				} else {
					export.IsReturnMargin = "否"
				}
			}
		}
		exports = append(exports, export)
	}
	// 检查数据库中交过保证金的，是否在表格中
	var pioneers []models.Pioneer
	err = db.Mysql.Model(models.Pioneer{}).Find(&pioneers).Debug().Error
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
	fileName := fmt.Sprintf("./城市先锋-%s.xlsx", time.Now().Format("2006-01-02"))
	excel.SaveAs(fileName)
}

func CheckPioneer3(excelFile string) {

	f, err := excelize.OpenFile(excelFile)
	fmt.Println(err, 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 取得 Sheet1 表格中所有的行
	rows, err := f.GetRows("Sheet1")
	err, cli := blockchain2.Client(blockchain2.CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	appraise := blockchain2.NewAppraise(cli)
	for i, row := range rows {
		if i == 0 {
			continue
		}
		var pioneerAddress string
		for j, colCell := range row {
			if j == 2 {
				var pioneer models.Pioneer
				pioneerAddress = strings.ToLower(colCell)
				err := db.Mysql.Model(models.Pioneer{}).Where("pioneer=?", pioneerAddress).First(&pioneer).Debug().Error
				if err != nil {
					fmt.Println("该城市先锋不存在", i)
					continue
				}
			} else if j == 6 {
				//if colCell != "1" {
				//	continue
				//}
				//// AdminSetPioneerBatch 管理员设置先锋批次
				//fmt.Println(pioneerAddress, utils.StringToInt64(colCell))
				//for i := 0; i < 20; i++ {
				//	err = appraise.AdminSetPioneerBatch(pioneerAddress, utils.StringToInt64(colCell))
				//	if err == nil {
				//		fmt.Println("该先锋批次", i, colCell, pioneerAddress, "成功------------------------------")
				//		break
				//	}
				//}

				err, beath := appraise.PioneerBatch(pioneerAddress)
				fmt.Println("该先锋批次", i, err, colCell, beath)

				//if beath != utils.StringToInt64(colCell) {
				//	fmt.Println("++++++++++++++++++++++ 错误", err)
				//} else {
				//	fmt.Println("正确", err, beath, "--", i, colCell)
				//
				//}

				//cityPioneer := blockchain2.NewCityPioneer()
				//err, pioneerInfo := cityPioneer.PioneerInfo(pioneerAddress)
				//if err != nil {
				//	fmt.Println("获取先锋信息错误", i, colCell)
				//	continue
				//} else {
				//	fmt.Println("先锋交保证金时间-----------------------------------", time.Unix(pioneerInfo.Ctime.Int64(), 0).Format(time.Layout))
				//
				//}
			}
		}
	}

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
	rows, err := f.GetRows("Sheet1")
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
	rows, err = f2.GetRows("Sheet1")
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
	rows, err := f.GetRows("Sheet1")
	fmt.Println(len(rows), 6)
loop:
	for i, row := range rows {
		if i == 0 {
			continue
		}
		for j, colCell := range row {
			if j == 2 {
				if colCell != "Da Nang" {
					continue loop
				}
			}
			if j == 4 {
				//fmt.Println(colCell)
				var userLocation models.UserLocation
				whereCondition := "location = 'Viet Nam " + colCell + "'"
				err := db.Mysql.Model(models.UserLocation{}).Where(whereCondition).First(&userLocation).Debug().Error
				if err == nil {
					fmt.Println("该区县已经存在", i, colCell, userLocation.CountyId)
					// 做映射
					blockchain2.SetCityMapping(userLocation.CountyId, "0xe07e8a8a22b4a0f18a804512e305a2de9ec91d71a54cdc8a050ab376641401e8", "pYnBembRMi8D2I+GR3bQkg==")

				} else {
					fmt.Println("该城市不存在", i, colCell)
				}
			}
		}
	}
}

func SyncAddPioneerInfoFromExcel(excelFile string) {
	var pioneers []Pioneer
	f, err := excelize.OpenFile(excelFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	err, cli := blockchain2.Client(blockchain2.CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	// 设置所需缴纳的TOX，城市等级
	city := blockchain2.NewCity(cli)

	// 取得 Sheet1 表格中所有的行
	rows, err := f.GetRows("Sheet1")

forOut:
	for i, row := range rows {
		if i == 1 || i == 0 {
			continue
		}
		pioneer := Pioneer{PioneerBatch: 4}
		for j, colCell := range row {
			var location models.UserLocation
			if j == 0 {
				colCell = strings.TrimSpace(colCell)
				area := strings.Split(colCell, " ")
				var where string
				if len(area) == 1 {
					where = "city like '%" + area[0] + "%'"
				} else {
					where = "city like '%" + area[0] + "%' and county like '%" + area[1] + "%'"
				}
				pioneer.AreaName = colCell
				err = db.Mysql.Model(models.UserLocation{}).Where(where).First(&location).Error
				if err == nil {
					if len(area) == 1 {
						pioneer.AreaId = location.CityId
					} else {
						pioneer.AreaId = location.CountyId
					}
				} else {
					fmt.Println(colCell, "该地区有问题", 123)
					continue forOut
				}

			}
			if j == 1 {
				pioneer.PioneerAddress = strings.ToLower(colCell)
			}
			if j == 2 {
				if colCell == "一线区县节点" {
					pioneer.IsAreaNode = 1
					pioneer.AreaLevel = 1
				} else if colCell == "二线区县节点" {
					pioneer.IsAreaNode = 1
					pioneer.AreaLevel = 2
				} else if colCell == "一线城市节点" {
					pioneer.IsAreaNode = 0
					pioneer.AreaLevel = 1
				} else if colCell == "二线城市节点" {
					pioneer.IsAreaNode = 0
					pioneer.AreaLevel = 2
				} else if colCell == "三线城市节点" {
					pioneer.IsAreaNode = 0
					pioneer.AreaLevel = 3
				} else {
					panic("节点级别错误")
				}
			}
		}
		//fmt.Println(pioneer)
		//if pioneer.PioneerAddress == "0x9d8213d2e8b3277c970e0d4dc99f2202c6e75919" ||
		//	pioneer.PioneerAddress == "0x89aa0acbdc006383fc26cb84ad5488cc6283217e" {
		//	continue
		//}
		pioneers = append(pioneers, pioneer)
	}

	//fmt.Println(pioneers, city)
	SyncAddPioneerInfoToDb(pioneers, city)
}

func SyncAddPioneerInfoToDb(pioneers []Pioneer, city *blockchain2.City) {
	for i, p := range pioneers {
		isSet := int64(0)
		isReset := int64(0)
		var remark, oldAreaId, oldPioneer string
		pioneerStatus := int64(0)
		// 根据city_id判断该城市是否已经添加过先锋
		var pioneer models.Pioneer
		err := db.Mysql.Model(models.Pioneer{}).Where("area_id=?", p.AreaId).Order("id desc").First(&pioneer).Debug().Error
		if err == nil {
			oldPioneer = pioneer.Pioneer
			if pioneer.FailedAt != "" {
				fmt.Println("该地区的先锋已经存在,但考核失败", i)
				remark = "该地区的先锋已经存在,但考核失败"
				isReset = 1
			} else if pioneer.IsOverTime > 0 {
				fmt.Println("该地区的先锋已经存在,但已经到期", i)
				remark = "该地区的先锋已经存在,但已经到期"
				isReset = 1
			} else {
				fmt.Println("该地区的先锋已经存在,且状态正常", i)
				remark = "该地区的先锋已经存在,且状态正常"
				pioneerStatus = 1
			}
			if remark != "" {
				remark += ";"
			}
		}
		pioneer = models.Pioneer{}
		err = db.Mysql.Model(models.Pioneer{}).Where("pioneer=?", p.PioneerAddress).Order("id desc").First(&pioneer).Debug().Error
		if err == nil {
			oldAreaId = pioneer.AreaId
			if pioneer.FailedAt != "" {
				fmt.Println("该先锋已经存在,但考核失败", i)
				remark += "该先锋已经存在,但考核失败"
				isReset = 1
			} else if pioneer.IsOverTime > 0 {
				fmt.Println("该先锋已经存在,但已经到期", i)
				remark += "该先锋已经存在,但已经到期"
				isReset = 1
			} else {
				fmt.Println("该先锋已经存在,且状态正常", i)
				remark += "该先锋已经存在,且状态正常"
				pioneerStatus = 1
			}
		}
		existAreaId, err2 := city.PioneerCity(p.PioneerAddress)
		if err2 != nil {
			log.Logger.Sugar().Error(err)
			return
		}
		if existAreaId != "0x0000000000000000000000000000000000000000000000000000000000000000" {
			isSet = 1
			remark += "该先锋已经录入过"
			fmt.Println("该先锋已经设置过", i)
		}
		existPioneerAddress, err2 := city.ChengShiPioneer(p.AreaId)
		if err2 != nil {
			log.Logger.Sugar().Error(err)
			return
		}
		fmt.Println(existPioneerAddress, 98790)
		if existPioneerAddress != "0x0000000000000000000000000000000000000000" {
			isSet = 1
			remark += "该地区已经录入过"
			fmt.Println("该地区已经录入过", i)
		}
		var pioneerAddInfo models.PioneerAddInfo
		err = db.Mysql.Model(&models.PioneerAddInfo{}).Where("area_id=?", p.AreaId).First(&pioneerAddInfo).Debug().Error
		if err == nil {
			fmt.Println("该地区先锋已设置过", i)
		}

		fmt.Println(p.AreaName, p.AreaLevel, p.AreaId, p.PioneerBatch, p.PioneerAddress, p.IsAreaNode, i)
		pioneerAddInfo = models.PioneerAddInfo{
			AreaId:        p.AreaId,
			OldAreaId:     oldAreaId,
			OldPioneer:    oldPioneer,
			Location:      p.AreaName,
			Pioneer:       p.PioneerAddress,
			AreaLevel:     p.AreaLevel,
			IsAreaNode:    p.IsAreaNode,
			PioneerBatch:  p.PioneerBatch,
			IsSet:         isSet,
			IsReset:       isReset,
			PioneerStatus: pioneerStatus,
			Remark:        remark,
			Ctime:         time.Now().Format("2006-01-02 15:04:05"),
		}
		fmt.Println(pioneerAddInfo)
		err = db.Mysql.Model(&models.PioneerAddInfo{}).Create(&pioneerAddInfo).Error
		if err != nil {
			fmt.Println("插入数据出错", i)
		}
	}
}
