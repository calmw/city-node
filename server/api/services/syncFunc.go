package services

import (
	"city-node-server/pkg/db"
	"city-node-server/pkg/log"
	"city-node-server/pkg/models"
	"city-node-server/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

var SyncChain = make(chan string, 1000)

// GetUserSons 获取并缓存用户下级
func GetUserSons(user string) error {
	// 获取下级数据
	url := fmt.Sprintf("https://subg-api.intoverse.co/api/v1/wallet/relations?wallet_address=%s", user)
	fmt.Println(url)
	err, data := utils.GetWithHeader(url, map[string]string{})
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	var sons Sons
	err = json.Unmarshal(data, &sons)
	if err != nil {
		log.Logger.Sugar().Infof(user, string(data))
		log.Logger.Sugar().Error(err)
		return err
	}
	if sons.Code != 0 {
		log.Logger.Sugar().Error(err)
		return errors.New("获取数据失败")
	}
	// 存入fdb
	cacheKey := fmt.Sprintf("team-data-%s", strings.ToLower(user))
	err = db.FDB.Set([]byte(cacheKey), time.Second*86400, data)
	if err != nil {
		log.Logger.Sugar().Error("获取team数据error：", err)
		return err
	}
	log.Logger.Sugar().Debug("获取team数成功")
	return nil
}

type ToxTxBridge struct {
	TotalOut float64 `json:"totalOut"`
	TotalIn  float64 `json:"totalIn"`
	Data     []struct {
		From      string `json:"from"`
		To        string `json:"to"`
		Value     string `json:"value"`
		Hash      string `json:"hash"`
		TimeStamp string `json:"timeStamp"`
		Type      string `json:"type"`
	} `json:"data"`
	UpdateTime int64 `json:"updateTime"`
}

func InitSyncTask() {
	for {
		select {
		case user, ok := <-SyncChain:
			if ok {
				// 获取下级,并缓存一天
				if GetUserSons(user) != nil {
					SyncChain <- user
				}
			} else {
				return
			}
		}
	}
}

// SyncUserLocationToDb 位置数据同步到本地
func SyncUserLocationToDb() {
	i := 0
	for {
		var userLocations []models.UserLocation
		offset := i * 1000
		db.Mysql.Model(&models.UserLocation{}).Where("id>?", offset).Order("id asc").Limit(1000).Find(&userLocations)
		if len(userLocations) == 0 {
			break
		}

		for _, location := range userLocations {
			locationBytes, err := json.Marshal(location)
			if err != nil {
				panic(err)
			}

			key := fmt.Sprintf("location-%s", strings.ToLower(location.User))
			err = db.FDB.Set([]byte(key), 0, locationBytes)
			if err != nil {
				panic(err)
			}
		}
		i++
	}
}
