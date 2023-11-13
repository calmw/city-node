package services

import (
	"city-node-server/api/common/statecode"
	"city-node-server/api/models/request"
	"city-node-server/api/utils"
	"city-node-server/db"
	"city-node-server/log"
	"city-node-server/models"
	utils2 "city-node-server/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"time"
)

type Mining struct{}

func NewMining() *Mining {
	return &Mining{}
}

type Sons struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data []string `json:"data"`
}

type Res struct {
	Date     int             `json:"date"`
	Recharge decimal.Decimal `json:"recharge"`
	Withdraw decimal.Decimal `json:"withdraw"`
}
type Ledger struct {
	User      string          `json:"user"`
	ChainName string          `json:"chain_name"`
	Direction string          `json:"direction"`
	Amount    decimal.Decimal `json:"amount"`
	//Withdraw  decimal.Decimal `json:"withdraw"`
	Ctime string `json:"ctime"`
}

func (c *Mining) LedgerDetails(req *request.LedgerDetails) (int, []Ledger, decimal.Decimal, decimal.Decimal, decimal.Decimal, decimal.Decimal) {
	var result []Ledger
	decimalZero, _ := decimal.NewFromString("0")
	// 查询下级用户缓存
	yesterday := time.Now().Add(-time.Hour * 24).Format("2006-01-02")
	cacheKey := "LedgerDetails-" + yesterday + req.User
	ok, data := utils.EventCache.Get(cacheKey)
	if !ok {
		// 异步拉取数据
		SyncChain <- req.User
		return statecode.SyncingData, result, decimalZero, decimalZero, decimalZero, decimalZero
	}
	var sons Sons
	err := json.Unmarshal(data, &sons)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return statecode.CommonErrServerErr, result, decimalZero, decimalZero, decimalZero, decimalZero
	}
	userMap := make(map[string]bool)
	for _, user := range sons.Data {
		userMap[user] = true
	}
	var bscIn, bscOut, matchIn, matchOut decimal.Decimal
	// 获取bsc bridge数据
	err, txBridgeBsc := GetToxTxBridgeBsc()
	for _, bsc := range txBridgeBsc.Data {
		amount, _ := decimal.NewFromString(bsc.Value)
		if userMap[bsc.From] {
			result = append(result, Ledger{
				User:      bsc.From,
				ChainName: "BSC",
				Amount:    amount,
				Direction: "In",
				Ctime:     time.Unix(utils.StringToInt64(bsc.TimeStamp), 0).Format("2006-01-02 15:04:05"),
			})
			bscIn = bscIn.Add(amount)
		} else if userMap[bsc.To] {
			result = append(result, Ledger{
				User:      bsc.From,
				ChainName: "BSC",
				Amount:    amount,
				Direction: "Out",
				Ctime:     time.Unix(utils.StringToInt64(bsc.TimeStamp), 0).Format("2006-01-02 15:04:05"),
			})
			bscOut = bscOut.Add(amount)
		}
	}
	// 获取Match充值【合约余额（get_tox_types 1）+ 可质押（get_pledge_balance_transferable 3）】
	var toxDayData []models.ToxDayData
	db.Mysql.Model(&models.ToxDayData{}).Where("status=1").Find(&toxDayData)
	for _, mT := range toxDayData {
		result = append(result, Ledger{
			User:      mT.Addr,
			ChainName: "Match",
			Amount:    mT.Amount,
			Direction: "In",
			Ctime:     time.Unix(mT.Date, 0).Format("2006-01-02 15:04:05"),
		})
		matchIn = matchIn.Add(mT.Amount)
	}
	var pledgeBalanceTransferable []models.PledgeBalanceTransferable
	db.Mysql.Model(&models.PledgeBalanceTransferable{}).Where("types=3").Find(&pledgeBalanceTransferable)
	for _, pT := range pledgeBalanceTransferable {
		result = append(result, Ledger{
			User:      pT.Sender,
			ChainName: "Match",
			Amount:    pT.Amount,
			Direction: "In",
			Ctime:     time.Unix(pT.Timestamp, 0).Format("2006-01-02 15:04:05"),
		})
		matchIn = matchIn.Add(pT.Amount)
	}
	// 获取Match提现
	toxDayData = make([]models.ToxDayData, 0)
	db.Mysql.Model(&models.ToxDayData{}).Where("status=2").Find(&toxDayData)
	for _, mT := range toxDayData {
		result = append(result, Ledger{
			User:      mT.Addr,
			ChainName: "Match",
			Amount:    mT.Amount,
			Direction: "Out",
			Ctime:     time.Unix(mT.Date, 0).Format("2006-01-02 15:04:05"),
		})
		matchOut = matchOut.Add(mT.Amount)
	}
	fmt.Println(data)
	return statecode.CommonSuccess, result, bscIn, bscOut, matchIn, matchOut
}

// GetToxTxBridgeBsc 获取bsc bridge数据
func GetToxTxBridgeBsc() (error, ToxTxBridge) {
	url := "https://intocache.intoverse.co/tox_tx_bridge_bsc.json"
	fmt.Println(url)
	err, data := utils2.GetWithHeader(url, map[string]string{})
	fmt.Println(err, 789)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ToxTxBridge{}
	}
	fmt.Println(123, string(data))
	var toxTxBridge ToxTxBridge
	err = json.Unmarshal(data, &toxTxBridge)
	if err != nil {
		log.Logger.Sugar().Error(err)
		fmt.Println(66, err)
		return err, ToxTxBridge{}
	}
	fmt.Println(77, err)
	if toxTxBridge.TotalIn == 0 {
		log.Logger.Sugar().Error(err)
		return errors.New("获取数据失败"), ToxTxBridge{}
	}
	// 存入缓存
	//yesterday := time.Now().Add(-time.Hour * 24).Format("2006-01-02")
	//cacheKey := "LedgerDetails-" + yesterday + user
	//utils2.EventCache.Set(cacheKey, data, 86405)
	return nil, toxTxBridge
}
