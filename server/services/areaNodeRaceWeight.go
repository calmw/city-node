package services

import (
	"city-node-server/api/common/statecode"
	"city-node-server/api/services"
	"city-node-server/pkg/db"
	"city-node-server/pkg/log"
	models2 "city-node-server/pkg/models"
	internal2 "city-node-server/services/internal"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/calmw/fdb"
	json2 "github.com/goccy/go-json"
	"github.com/shopspring/decimal"
	"github.com/xjieinfo/xjgo/xjcore/xjexcel"
	"gorm.io/gorm"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

type RechargeRecord struct {
	Data struct {
		RechargeRecords []struct {
			ID       string `json:"id"`
			User     string `json:"user"`
			CountyID string `json:"countyId"`
			Amount   string `json:"amount"`
			Ctime    string `json:"ctime"`
			TxHash   string `json:"txHash"`
		} `json:"rechargeRecords"`
	} `json:"data"`
}

func SyncStatus(excelFile string) {
	f, err := excelize.OpenFile(excelFile)
	if err != nil {
		fmt.Println(err, 1111)
		return
	}
	// 取得 Sheet1 表格中所有的行
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err, 2222)
		return
	}
	for i, row := range rows {
		if i == 1 || i == 0 {
			continue
		}
		for j, colCell := range row {
			if j == 3 {
				fmt.Println(colCell, 123)
				var location models2.UserLocation
				where := "location like '%" + colCell + "%'"
				err = db.Mysql.Model(models2.UserLocation{}).Where(where).First(&location).Error
				fmt.Println(err)
				if err != nil {
					return
				}
			}

			if j == 4 {
				// 查询下级用户缓存
				userAddress := strings.Trim(strings.TrimSpace(strings.ToLower(colCell)), "\t")
				cacheKey := "team-data" + userAddress
				fmt.Println(cacheKey, 7777)
				_, err = db.FDB.Get([]byte(cacheKey))
				if err != nil {
					if errors.Is(err, fdb.ErrKeyNotFound) {
						services.GetUserSons(userAddress)
					} else {
						fmt.Println(err)
						return
					}
				}
			}
		}
	}
	close(SyncChain)
}

// AreaNodeRace 竞争节点数据
type AreaNodeRace struct {
	AreaName string `json:"area_name" excel:"column:B;desc:地区;width:30"`
	Wallet   string `json:"wallet" excel:"column:C;desc:钱包;width:30"`
	AreaType string `json:"area_type" excel:"column:D;desc:先锋类型;width:30"`
	Weight   string `json:"weight" excel:"column:E;desc:业绩;width:30"`
}

func GetRaceNodeWeight(excelFile string) int {
	var areaNodeRaces []AreaNodeRace
	f, err := excelize.OpenFile(excelFile)
	if err != nil {
		fmt.Println(err)
		return statecode.CommonErrServerErr
	}
	// 取得 Sheet1 表格中所有的行
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return statecode.CommonErrServerErr
	}
	for i, row := range rows {
		if i == 1 || i == 0 {
			continue
		}
		areaNodeRace := AreaNodeRace{}
		for j, colCell := range row {
			if j == 3 {
				areaNodeRace.AreaName = colCell
				fmt.Println(colCell)
			} else if j == 4 {
				areaNodeRace.Wallet = strings.ToLower(colCell)
				fmt.Println(colCell)
			} else if j == 5 {
				areaNodeRace.AreaType = colCell
				fmt.Println(colCell)
			}
		}
		areaNodeRaces = append(areaNodeRaces, areaNodeRace)
	}

	var result []AreaNodeRace
	for _, node := range areaNodeRaces {
		var location models2.UserLocation
		where := "location like '%" + node.AreaName + "%'"
		err = db.Mysql.Model(models2.UserLocation{}).Where(where).First(&location).Error
		if err != nil {
			log.Logger.Error(err.Error())
			return statecode.CommonErrServerErr
		}

		// 查询下级用户缓存
		userAddress := node.Wallet
		cacheKey := "team-data" + userAddress
		data, err := db.FDB.Get([]byte(cacheKey))
		if errors.Is(err, fdb.ErrKeyNotFound) {
			log.Logger.Sugar().Error("异步拉取数据", cacheKey)
			return statecode.SyncingData
		}
		var sons Sons
		err = json.Unmarshal(data, &sons)
		if err != nil {
			log.Logger.Sugar().Error(err)
			return statecode.CommonErrServerErr
		}
		userMap := make(map[string]bool)
		userSet := internal2.NewStringSet()
		for _, user := range sons.Data.Data {
			user = strings.ToLower(user)
			userMap[user] = true
			userSet.Add(user)
		}

		d18 := decimal.NewFromInt(1e18)
		weight := decimal.NewFromInt(0)

		//var lock sync.Mutex
		wg := sync.WaitGroup{}
		var limitMaxNum = 100
		addLock := sync.Mutex{}
		var chData = make(chan int, limitMaxNum)

		areaNodeRaces = []AreaNodeRace{}
		for n, user := range userSet.Data {
			chData <- n
			wg.Add(1)
			go func(u string, index int, n AreaNodeRace, lock *sync.Mutex) {
				_, d, _ := GetWeight(u, location.CityId, []int{2, 8})
				lock.Lock()
				defer lock.Unlock()
				weight = weight.Add(d)
				<-chData
				wg.Done()

			}(user, n, node, &addLock)
		}
		wg.Wait()
		node.Weight = weight.Div(d18).String()
		result = append(result, node)
	}

	f = xjexcel.ListToExcel(result, "团队充值", "详情")
	fileName := fmt.Sprintf("./业绩.xlsx")
	f.SaveAs(fileName)

	return statecode.CommonSuccess
}

func GetWeight(user, cityId string, countTime []int) (error, decimal.Decimal, bool) {
	inCity := false
	key := fmt.Sprintf("%s-%s", user, cityId)
	var val string
	valBytes, err := db.LevelDb.Get([]byte(key), nil)
	if err != nil {
		locationBytes, err := db.FDB.Get([]byte(fmt.Sprintf("user_location_%s", user)))
		if err != nil {
			return err, decimal.Zero, inCity
		}
		var userLocation models2.UserLocation
		err = json2.Unmarshal(locationBytes, &userLocation)
		if err != nil {
			return err, decimal.Zero, inCity
		}

		//userLocation := models2.UserLocation{}
		//err = db.Mysql.Model(models2.UserLocation{}).Where("user=?", user).First(&userLocation).Error
		if err == nil {
			key = fmt.Sprintf("%s-%s", user, cityId)
			if userLocation.CityId == cityId {
				inCity = true
				val = "1"
				err = db.LevelDb.Put([]byte(key), []byte(val), nil)
				if err != nil {
					log.Logger.Sugar().Error("LevelDb.Put error:", err)
				}
				db.LevelDb.Put([]byte(fmt.Sprintf("%s_city", user)), []byte(userLocation.Location), nil)
			} else {
				val = "2"
				err = db.LevelDb.Put([]byte(key), []byte(val), nil)
				if err != nil {
					log.Logger.Sugar().Error("LevelDb.Put error:", err)
				}
				log.Logger.Sugar().Error(err)
				return err, decimal.Zero, inCity
			}
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			val = "3" // 用户未定位，下次统计重新查询
			err = db.LevelDb.Put([]byte(key), []byte(val), nil)
			if err != nil {
				log.Logger.Sugar().Error("LevelDb.Put error:", err)
			}
			return err, decimal.Zero, inCity
		} else {
			log.Logger.Sugar().Error(err)
			return err, decimal.Zero, inCity
		}
	}
	if string(valBytes) == "1" {
		inCity = true
	}
	// 每次统计打开一次，用户定位数据会更新
	if string(valBytes) == "3" {
		InCity(user, cityId)
	}
	if !inCity {
		return err, decimal.Zero, inCity
	}

	// 获取该用户网体充值权重，审核通过后到2月29日的充值权重
	err, d := GetRechargeWeight(user, fmt.Sprintf("%d-%02d-%02d 00:00:00", 2024, countTime[0], countTime[1]), "2024-03-01 00:00:00")
	if err != nil {
		log.Logger.Sugar().Debugf("是否在绑定城市：%v,下级：%s,获取权重错误：%v", inCity, user, err)
		return err, decimal.Zero, inCity
	}
	log.Logger.Sugar().Debugf("是否在绑定城市：%v,下级：%s,权重：%s", inCity, user, d.String())
	return nil, d, inCity
}

func GetRechargeWeight(user, start, end string) (error, decimal.Decimal) {
	var result = decimal.Zero
	url := "https://subgraph.intoverse.co/subgraphs/name/city_node"
	method := "POST"
	startTime, _ := time.Parse("2006-01-02 15:04:05", start)
	endTime, _ := time.Parse("2006-01-02 15:04:05", end)
	index := 0
	for {
		index++
		query := fmt.Sprintf(`{"query": "query myQuery {\n rechargeRecords(first: 1000, skip:%d, orderBy: timestamp,  orderDirection: asc,  where: {timestamp_gte: \"%d\" timestamp_lt: \"%d\" user: \"%s\"}) { id\n  user\n    countyId\n    amount\n    ctime\n    timestamp\n    txHash\n }\n}\n    ", "variables": null, "operationName": "MyQuery", "extensions": {"headers": null}}`, (index-1)*1000, startTime.Unix(), endTime.Unix(), user)
		payload := strings.NewReader(query)
		client := &http.Client{Timeout: time.Second * 30}
		req, err := http.NewRequest(method, url, payload)
		if err != nil {
			log.Logger.Error(err.Error())
			return err, result
		}
		req.Header.Add("Content-Type", "application/json")
		var res *http.Response
		for k := 0; k < 30; k++ {
			res, err = client.Do(req)
			if err == nil {
				break
			}
		}
		if err != nil {
			log.Logger.Error(err.Error())
			return err, result
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Logger.Error(err.Error())
			return err, result
		}
		var rechargeRecord RechargeRecord
		err = json.Unmarshal(body, &rechargeRecord)
		if err != nil {
			log.Logger.Error(err.Error())
			return err, result
		}
		if len(rechargeRecord.Data.RechargeRecords) <= 0 {
			break
		}

		for j := 0; j < len(rechargeRecord.Data.RechargeRecords); j++ {
			record := rechargeRecord.Data.RechargeRecords[j]
			result = result.Add(decimal.RequireFromString(record.Amount))
		}
	}
	return nil, result
}

func InCity(user, cityId string) bool {
	var inCity bool
	var key, val string
	locationBytes, err := db.FDB.Get([]byte(fmt.Sprintf("user_location_%s", user)))
	if err != nil {
		return false
	}
	var userLocation models2.UserLocation
	err = json2.Unmarshal(locationBytes, &userLocation)
	if err != nil {
		return false
	}
	//userLocation := models2.UserLocation{}
	//err := db.Mysql.Model(models2.UserLocation{}).Where("user=?", user).First(&userLocation).Error
	if err == nil {
		key = fmt.Sprintf("%s-%s", user, cityId)
		if userLocation.CityId == cityId {
			inCity = true
			val = "1"
			err = db.LevelDb.Put([]byte(key), []byte(val), nil)
			if err != nil {
				log.Logger.Sugar().Error("LevelDb.Put error:", err)
			}
		} else {
			val = "2"
			err = db.LevelDb.Put([]byte(key), []byte(val), nil)
			if err != nil {
				log.Logger.Sugar().Error("LevelDb.Put error:", err)
			}
			log.Logger.Sugar().Error(err)
			return inCity
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		val = "3" // 用户未定位，下次统计重新查询
		err = db.LevelDb.Put([]byte(key), []byte(val), nil)
		if err != nil {
			log.Logger.Sugar().Error("LevelDb.Put error:", err)
		}
		return inCity
	} else {
		log.Logger.Sugar().Error(err)
		return inCity
	}
	return inCity
}

func SyncUserLocation() {
	index := 0
	for {
		var locations []models2.UserLocation
		db.Mysql.Model(&models2.UserLocation{}).Where("id>?", index*1000).Limit(1000).Find(&locations)

		if len(locations) == 0 {
			break
		}

		for _, location := range locations {

			marshal, err := json.Marshal(location)
			if err != nil {
				log.Logger.Sugar().Error(err)
				return
			}
			err = db.FDB.Set([]byte(fmt.Sprintf("user_location_%s", location.User)), time.Hour*24, marshal)
			if err != nil {
				log.Logger.Sugar().Error(err)
				return
			}
		}
		index++
	}
}
