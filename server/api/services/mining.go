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
	"sort"
	"strings"
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
	Ctime     string          `json:"ctime"`
	TimeStamp int64           `json:"time_stamp"`
}

func (c *Mining) LedgerDetails(req *request.LedgerDetails) (int, []Ledger, decimal.Decimal, decimal.Decimal, decimal.Decimal, decimal.Decimal) {
	var result []Ledger
	var bscIn, bscOut, matchIn, matchOut decimal.Decimal
	decimalZero, _ := decimal.NewFromString("0")
	/// 查全网数据
	if req.User == "" {
		// 获取bsc bridge数据
		err, txBridgeBsc := GetToxTxBridgeBsc()
		if err != nil {
			log.Logger.Sugar().Error(err)
			return statecode.CommonErrServerErr, result, decimalZero, decimalZero, decimalZero, decimalZero
		}
		for _, bsc := range txBridgeBsc.Data {
			amount, _ := decimal.NewFromString(bsc.Value)
			timeStamp := utils.StringToInt64(bsc.TimeStamp)
			if req.Start > 0 {
				if timeStamp < req.Start {
					continue
				}
			}
			if req.End > 0 {
				if timeStamp > req.End {
					continue
				}
			}
			if bsc.Type == "in" {
				result = append(result, Ledger{
					User:      bsc.From,
					ChainName: "BSC",
					Amount:    amount,
					Direction: "In",
					TimeStamp: timeStamp,
					Ctime:     time.Unix(timeStamp, 0).Format("2006-01-02 15:04:05"),
				})
				bscIn = bscIn.Add(amount)
			} else if bsc.Type == "out" {
				result = append(result, Ledger{
					User:      bsc.From,
					ChainName: "BSC",
					Amount:    amount,
					Direction: "Out",
					TimeStamp: timeStamp,
					Ctime:     time.Unix(timeStamp, 0).Format("2006-01-02 15:04:05"),
				})
				bscOut = bscOut.Add(amount)
			}
		}
		// 获取Match充值【合约余额（get_tox_types 1）+ 可质押（get_pledge_balance_transferable 3）】
		var toxDayData []models.ToxDayData
		tx := db.Mysql.Model(&models.ToxDayData{}).Where("status=1")
		if req.Start > 0 {
			tx = tx.Where("date>=?", req.Start)
		}
		if req.End > 0 {
			tx = tx.Where("date<=?", req.End)
		}
		tx.Order("date desc")
		tx.Find(&toxDayData)
		for _, mT := range toxDayData {
			result = append(result, Ledger{
				User:      mT.Addr,
				ChainName: "Match",
				Amount:    mT.Amount,
				Direction: "In",
				TimeStamp: time.Unix(mT.Date, 0).Unix(),
				Ctime:     time.Unix(mT.Date, 0).Format("2006-01-02 15:04:05"),
			})
			matchIn = matchIn.Add(mT.Amount)
		}
		var pledgeBalanceTransferable []models.PledgeBalanceTransferable
		tx = db.Mysql.Model(&models.PledgeBalanceTransferable{}).Where("types=3")
		if req.Start > 0 {
			tx = tx.Where("timestamp>=?", req.Start)
		}
		if req.End > 0 {
			tx = tx.Where("timestamp<=?", req.End)
		}
		tx.Order("timestamp desc")
		tx.Find(&pledgeBalanceTransferable)
		for _, pT := range pledgeBalanceTransferable {
			result = append(result, Ledger{
				User:      pT.Sender,
				ChainName: "Match",
				Amount:    pT.Amount,
				Direction: "In",
				TimeStamp: pT.Timestamp,
				Ctime:     time.Unix(pT.Timestamp, 0).Format("2006-01-02 15:04:05"),
			})
			matchIn = matchIn.Add(pT.Amount)
		}
		// 获取Match提现
		toxDayData = make([]models.ToxDayData, 0)
		tx = db.Mysql.Model(&models.ToxDayData{}).Where("status=2")
		if req.Start > 0 {
			tx = tx.Where("date>=?", req.Start)
		}
		if req.End > 0 {
			tx = tx.Where("date<=?", req.End)
		}
		tx.Order("date desc")
		tx.Find(&toxDayData)
		for _, mT := range toxDayData {
			result = append(result, Ledger{
				User:      mT.Addr,
				ChainName: "Match",
				Amount:    mT.Amount,
				Direction: "Out",
				TimeStamp: time.Unix(mT.Date, 0).Unix(),
				Ctime:     time.Unix(mT.Date, 0).Format("2006-01-02 15:04:05"),
			})
			matchOut = matchOut.Add(mT.Amount)
		}
		return statecode.CommonSuccess, result, bscIn, bscOut, matchIn, matchOut
	}
	/// 查指定网体数据
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
		user = strings.ToLower(user)
		userMap[user] = true
	}
	// 获取bsc bridge数据
	err, txBridgeBsc := GetToxTxBridgeBsc()
	for _, bsc := range txBridgeBsc.Data {
		timeStamp := utils.StringToInt64(bsc.TimeStamp)
		if req.Start > 0 {
			if timeStamp < req.Start {
				continue
			}
		}
		if req.End > 0 {
			if timeStamp > req.End {
				continue
			}
		}
		amount, _ := decimal.NewFromString(bsc.Value)
		from := strings.ToLower(bsc.From)
		to := strings.ToLower(bsc.To)
		if userMap[from] {
			result = append(result, Ledger{
				User:      from,
				ChainName: "BSC",
				Amount:    amount,
				Direction: "In",
				TimeStamp: utils.StringToInt64(bsc.TimeStamp),
				Ctime:     time.Unix(utils.StringToInt64(bsc.TimeStamp), 0).Format("2006-01-02 15:04:05"),
			})
			bscIn = bscIn.Add(amount)
		} else if userMap[to] {
			result = append(result, Ledger{
				User:      to,
				ChainName: "BSC",
				Amount:    amount,
				Direction: "Out",
				TimeStamp: utils.StringToInt64(bsc.TimeStamp),
				Ctime:     time.Unix(utils.StringToInt64(bsc.TimeStamp), 0).Format("2006-01-02 15:04:05"),
			})
			bscOut = bscOut.Add(amount)
		}
	}
	// 获取Match充值【合约余额（get_tox_types 1）+ 可质押（get_pledge_balance_transferable 3）】
	var toxDayData []models.ToxDayData
	tx := db.Mysql.Model(&models.ToxDayData{}).Where("status=1")
	if req.Start > 0 {
		tx = tx.Where("date>=?", req.Start)
	}
	if req.End > 0 {
		tx = tx.Where("date<=?", req.End)
	}
	tx.Order("date desc")
	tx.Find(&toxDayData)
	for _, mT := range toxDayData {
		result = append(result, Ledger{
			User:      mT.Addr,
			ChainName: "Match",
			Amount:    mT.Amount,
			Direction: "In",
			TimeStamp: time.Unix(mT.Date, 0).Unix(),
			Ctime:     time.Unix(mT.Date, 0).Format("2006-01-02 15:04:05"),
		})
		matchIn = matchIn.Add(mT.Amount)
	}
	var pledgeBalanceTransferable []models.PledgeBalanceTransferable
	tx = db.Mysql.Model(&models.PledgeBalanceTransferable{}).Where("types=3")
	if req.Start > 0 {
		tx = tx.Where("timestamp>=?", req.Start)
	}
	if req.End > 0 {
		tx = tx.Where("timestamp<=?", req.End)
	}
	tx.Order("timestamp desc")
	tx.Find(&pledgeBalanceTransferable)
	for _, pT := range pledgeBalanceTransferable {
		result = append(result, Ledger{
			User:      pT.Sender,
			ChainName: "Match",
			Amount:    pT.Amount,
			Direction: "In",
			TimeStamp: pT.Timestamp,
			Ctime:     time.Unix(pT.Timestamp, 0).Format("2006-01-02 15:04:05"),
		})
		matchIn = matchIn.Add(pT.Amount)
	}
	// 获取Match提现
	toxDayData = make([]models.ToxDayData, 0)
	tx = db.Mysql.Model(&models.ToxDayData{}).Where("status=2")
	if req.Start > 0 {
		tx = tx.Where("date>=?", req.Start)
	}
	if req.End > 0 {
		tx = tx.Where("date<=?", req.End)
	}
	tx.Order("date desc")
	tx.Find(&toxDayData)
	for _, mT := range toxDayData {
		result = append(result, Ledger{
			User:      mT.Addr,
			ChainName: "Match",
			Amount:    mT.Amount,
			Direction: "Out",
			TimeStamp: time.Unix(mT.Date, 0).Unix(),
			Ctime:     time.Unix(mT.Date, 0).Format("2006-01-02 15:04:05"),
		})
		matchOut = matchOut.Add(mT.Amount)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].TimeStamp > result[j].TimeStamp })

	return statecode.CommonSuccess, result, bscIn, bscOut, matchIn, matchOut
}

// GetToxTxBridgeBsc 获取bsc bridge数据
func GetToxTxBridgeBsc() (error, ToxTxBridge) {
	url := "https://intocache.intoverse.co/tox_tx_bridge_bsc.json"
	fmt.Println(url)
	err, data := utils2.GetWithHeader(url, map[string]string{})
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ToxTxBridge{}
	}
	var toxTxBridge ToxTxBridge
	err = json.Unmarshal(data, &toxTxBridge)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, ToxTxBridge{}
	}
	if toxTxBridge.TotalIn == 0 {
		log.Logger.Sugar().Error(err)
		return errors.New("获取数据失败"), ToxTxBridge{}
	}
	return nil, toxTxBridge
}
