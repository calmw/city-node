package services

import (
	utils2 "city-node-server/api/utils"
	"city-node-server/pkg/log"
	"city-node-server/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var SyncChain = make(chan string, 1000)

// GetUserSons 获取并缓存用户下级
func GetUserSons(user string) error {
	// 获取下级数据
	url := fmt.Sprintf("https://baodao-api.baodao.app/api/v1/pledge/relation/child?address=%s", user)
	err, data := utils.GetWithHeader(url, map[string]string{})
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	var sons Sons
	err = json.Unmarshal(data, &sons)
	if err != nil {
		log.Logger.Sugar().Infof(string(data))
		log.Logger.Sugar().Error(err)
		return err
	}
	if sons.Code != 200 {
		log.Logger.Sugar().Error(err)
		return errors.New("获取数据失败")
	}
	// 存入缓存
	yesterday := time.Now().Add(-time.Hour * 24).Format("2006-01-02")
	cacheKey := "LedgerDetails-" + yesterday + user
	utils2.EventCache.Set(cacheKey, data, 86405)
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
	for true {
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
