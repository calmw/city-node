package services

import (
	"city-node-server/api/common/statecode"
	"city-node-server/api/models/request"
	"city-node-server/api/services/internal"
	"city-node-server/api/utils"
	"city-node-server/pkg/db"
	"city-node-server/pkg/log"
	models2 "city-node-server/pkg/models"
	utils2 "city-node-server/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xjieinfo/xjgo/xjcore/xjexcel"
	"gorm.io/gorm"
	"io"
	"net/http"
	"sort"
	"strings"
	"sync"
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
	User      string          `json:"user" excel:"column:B;desc:钱包地址;width:30"`
	ChainName string          `json:"chain_name" excel:"column:C;desc:链名称;width:30"`
	Direction string          `json:"direction" excel:"column:D;desc:方向;width:30"`
	Amount    decimal.Decimal `json:"amount" excel:"column:E;desc:数额;width:30"`
	Ctime     string          `json:"ctime" excel:"column:F;desc:时间;width:30"`
	TimeStamp int64           `json:"time_stamp" excel:"column:G;desc:时间戳;width:30"`
}

type LedgerSum struct {
	User     string          `json:"user" excel:"column:B;desc:钱包地址;width:30"`
	In       decimal.Decimal `json:"in" excel:"column:C;desc:总充值;width:30"`
	Out      decimal.Decimal `json:"out" excel:"column:D;desc:总提现;width:30"`
	MatchIn  decimal.Decimal `json:"match_in" excel:"column:E;desc:Match充值;width:30"`
	MatchOut decimal.Decimal `json:"match_out" excel:"column:F;desc:Match提现;width:30"`
	BscIn    decimal.Decimal `json:"bsc_in" excel:"column:G;desc:Bsc充值;width:30"`
	BscOut   decimal.Decimal `json:"bsc_out" excel:"column:H;desc:Bsc提现;width:30"`
}

type Recharge struct {
	User   string          `json:"user" excel:"column:B;desc:钱包地址;width:30"`
	Amount decimal.Decimal `json:"in" excel:"column:C;desc:总充值;width:30"`
}

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

type StudioInfo struct {
	UserName  string `json:"user_name" excel:"column:B;desc:报名人;width:30"`
	Level     string `json:"level" excel:"column:C;desc:级别;width:30"`
	User      string `json:"user" excel:"column:D;desc:钱包地址;width:30"`
	Super     string `json:"super" excel:"column:E;desc:上级领导人;width:30"`
	Community string `json:"community" excel:"column:F;desc:隶属社区;width:30"`
	City      string `json:"city" excel:"column:G;desc:城市;width:30"`
	Address   string `json:"address" excel:"column:H;desc:工作室地址;width:30"`
	//ApplyTime      string `json:"apply_time" excel:"column:I;desc:申请日期;width:30"`
	//AuditeTime     string `json:"audite_time" excel:"column:J;desc:审核日期;width:30"`
	//AuditeContent  string `json:"audite_content" excel:"column:K;desc:审核内容;width:30"`
	//ApprovedTime   string `json:"approved_time" excel:"column:L;desc:审核通过日期;width:30"`
	CountTime      string `json:"count_time" excel:"column:I;desc:业绩计算起止时间;width:30"`
	RechargeWeight string `json:"recharge_weight" excel:"column:J;desc:充值权重;width:30"`
}

type StudioUserInfo struct {
	UserName       string `json:"user_name" excel:"column:B;desc:报名人;width:30"`
	User           string `json:"user" excel:"column:C;desc:钱包地址;width:30"`
	InCity         bool   `json:"in_city" excel:"column:D;desc:是否在该城市;width:30"`
	BindCity       string `json:"bind_city" excel:"column:E;desc:查询城市;width:30"`
	UserCity       string `json:"user_city" excel:"column:F;desc:用户定位城市;width:30"`
	RechargeWeight string `json:"recharge_weight" excel:"column:G;desc:充值权重;width:30"`
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
					User:      bsc.To,
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
		var toxDayData []models2.ToxDayData
		tx := db.Mysql.Model(&models2.ToxDayData{}).Where("status=1")
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
		var pledgeBalanceTransferable []models2.PledgeBalanceTransferable
		tx = db.Mysql.Model(&models2.PledgeBalanceTransferable{}).Where("types=3")
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
		toxDayData = make([]models2.ToxDayData, 0)
		tx = db.Mysql.Model(&models2.ToxDayData{}).Where("status=2")
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
		matchOut = matchOut.Sub(bscOut)
		matchIn = matchIn.Sub(bscIn)
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
	var toxDayData []models2.ToxDayData
	tx := db.Mysql.Model(&models2.ToxDayData{}).Where("status=1")
	if req.Start > 0 {
		tx = tx.Where("date>=?", req.Start)
	}
	if req.End > 0 {
		tx = tx.Where("date<=?", req.End)
	}
	tx.Order("date desc")
	tx.Find(&toxDayData)
	for _, mT := range toxDayData {
		user := strings.ToLower(mT.Addr)
		if userMap[user] {
			result = append(result, Ledger{
				User:      user,
				ChainName: "Match",
				Amount:    mT.Amount,
				Direction: "In",
				TimeStamp: time.Unix(mT.Date, 0).Unix(),
				Ctime:     time.Unix(mT.Date, 0).Format("2006-01-02 15:04:05"),
			})
			matchIn = matchIn.Add(mT.Amount)
		}
	}
	var pledgeBalanceTransferable []models2.PledgeBalanceTransferable
	tx = db.Mysql.Model(&models2.PledgeBalanceTransferable{}).Where("types=3")
	if req.Start > 0 {
		tx = tx.Where("timestamp>=?", req.Start)
	}
	if req.End > 0 {
		tx = tx.Where("timestamp<=?", req.End)
	}
	//tx.Where("sender in ?", sons.Data)
	tx.Order("timestamp desc")
	tx.Find(&pledgeBalanceTransferable)
	for _, pT := range pledgeBalanceTransferable {
		user := strings.ToLower(pT.Sender)
		if userMap[user] {
			result = append(result, Ledger{
				User:      user,
				ChainName: "Match",
				Amount:    pT.Amount,
				Direction: "In",
				TimeStamp: pT.Timestamp,
				Ctime:     time.Unix(pT.Timestamp, 0).Format("2006-01-02 15:04:05"),
			})
			matchIn = matchIn.Add(pT.Amount)
		}
	}
	// 获取Match提现
	toxDayData = []models2.ToxDayData{}
	tx = db.Mysql.Model(&models2.ToxDayData{}).Where("status=2")
	if req.Start > 0 {
		tx = tx.Where("date>=?", req.Start)
	}
	if req.End > 0 {
		tx = tx.Where("date<=?", req.End)
	}
	tx.Order("date desc")
	tx.Find(&toxDayData)
	for _, mT := range toxDayData {
		user := strings.ToLower(mT.Addr)
		if userMap[user] {
			result = append(result, Ledger{
				User:      user,
				ChainName: "Match",
				Amount:    mT.Amount,
				Direction: "Out",
				TimeStamp: time.Unix(mT.Date, 0).Unix(),
				Ctime:     time.Unix(mT.Date, 0).Format("2006-01-02 15:04:05"),
			})
			matchOut = matchOut.Add(mT.Amount)
		}
	}
	sort.Slice(result, func(i, j int) bool { return result[i].TimeStamp > result[j].TimeStamp })
	matchOut = matchOut.Sub(bscOut)
	matchIn = matchIn.Sub(bscIn)
	// 保存excel
	// "github.com/xjieinfo/xjgo/xjcore/xjexcel"
	//f := xjexcel.ListToExcel(result, "团队充值提现", "下级成员详情")
	//fileName := fmt.Sprintf("./团队充值提现详情-%s.xls", req.User)
	//f.SaveAs(fileName)
	return statecode.CommonSuccess, result, bscIn, bscOut, matchIn, matchOut
}

func (c *Mining) Ledger(req *request.LedgerDetails) (int, []LedgerSum, decimal.Decimal, decimal.Decimal, decimal.Decimal, decimal.Decimal) {
	start := time.Now()
	var ledger []Ledger
	var result []LedgerSum
	var bscIn, bscOut, matchIn, matchOut decimal.Decimal
	decimalZero, _ := decimal.NewFromString("0")
	//decimalE18, _ := decimal.NewFromString("1000000000000000000")
	userSet := internal.NewStringSet()
	userDetails := map[string]map[string]decimal.Decimal{}
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
				u := strings.ToLower(bsc.From)
				ledger = append(ledger, Ledger{
					User:      u,
					ChainName: "BSC",
					Amount:    amount,
					Direction: "In",
				})
				userSet.Add(u)
				bscIn = bscIn.Add(amount)
				if userDetails[u] == nil {
					userDetails[u] = make(map[string]decimal.Decimal)
				}
				userDetails[u]["BscIn"] = userDetails[u]["BscIn"].Add(amount)
			} else if bsc.Type == "out" {
				u := strings.ToLower(bsc.To)
				ledger = append(ledger, Ledger{
					User:      u,
					ChainName: "BSC",
					Amount:    amount,
					Direction: "Out",
				})
				userSet.Add(u)
				bscOut = bscOut.Add(amount)
				if userDetails[u] == nil {
					userDetails[u] = make(map[string]decimal.Decimal)
				}
				userDetails[u]["BscOut"] = userDetails[u]["BscOut"].Add(amount)
			}
		}
		// 获取Match充值【合约余额（get_tox_types 1）+ 可质押（get_pledge_balance_transferable 3）】
		var toxDayData []models2.ToxDayData
		tx := db.Mysql.Model(&models2.ToxDayData{}).Where("status=1")
		if req.Start > 0 {
			tx = tx.Where("date>=?", req.Start)
		}
		if req.End > 0 {
			tx = tx.Where("date<=?", req.End)
		}
		tx.Order("date desc")
		tx.Find(&toxDayData)
		for _, mT := range toxDayData {
			u := strings.ToLower(mT.Addr)
			ledger = append(ledger, Ledger{
				User:      u,
				ChainName: "Match",
				Amount:    mT.Amount,
				Direction: "In",
			})
			userSet.Add(u)
			matchIn = matchIn.Add(mT.Amount)
			if userDetails[u] == nil {
				userDetails[u] = make(map[string]decimal.Decimal)
			}
			userDetails[u]["In"] = userDetails[u]["In"].Add(mT.Amount)
		}
		var pledgeBalanceTransferable []models2.PledgeBalanceTransferable
		tx = db.Mysql.Model(&models2.PledgeBalanceTransferable{}).Where("types=3")
		if req.Start > 0 {
			tx = tx.Where("timestamp>=?", req.Start)
		}
		if req.End > 0 {
			tx = tx.Where("timestamp<=?", req.End)
		}
		tx.Order("timestamp desc")
		tx.Find(&pledgeBalanceTransferable)
		for _, pT := range pledgeBalanceTransferable {
			u := strings.ToLower(pT.Sender)
			ledger = append(ledger, Ledger{
				User:      u,
				ChainName: "Match",
				Amount:    pT.Amount,
				Direction: "In",
			})
			userSet.Add(u)
			matchIn = matchIn.Add(pT.Amount)
			if userDetails[u] == nil {
				userDetails[u] = make(map[string]decimal.Decimal)
			}
			userDetails[u]["In"] = userDetails[u]["In"].Add(pT.Amount)
		}
		// 获取Match提现
		toxDayData = make([]models2.ToxDayData, 0)
		tx = db.Mysql.Model(&models2.ToxDayData{}).Where("status=2")
		if req.Start > 0 {
			tx = tx.Where("date>=?", req.Start)
		}
		if req.End > 0 {
			tx = tx.Where("date<=?", req.End)
		}
		tx.Order("date desc")
		tx.Find(&toxDayData)
		for _, mT := range toxDayData {
			u := strings.ToLower(mT.Addr)
			ledger = append(ledger, Ledger{
				User:      u,
				ChainName: "Match",
				Amount:    mT.Amount,
				Direction: "Out",
			})
			userSet.Add(u)
			matchOut = matchOut.Add(mT.Amount)
			if userDetails[u] == nil {
				userDetails[u] = make(map[string]decimal.Decimal)
			}
			userDetails[u]["Out"] = userDetails[u]["Out"].Add(mT.Amount)
		}
		matchOut = matchOut.Sub(bscOut)
		matchIn = matchIn.Sub(bscIn)
		fmt.Println(time.Since(start))
		for _, userAddress := range userSet.Data {
			result = append(result, LedgerSum{
				User:     userAddress,
				In:       userDetails[userAddress]["In"],
				Out:      userDetails[userAddress]["Out"],
				BscIn:    userDetails[userAddress]["BscIn"],
				BscOut:   userDetails[userAddress]["BscOut"],
				MatchIn:  userDetails[userAddress]["In"].Sub(userDetails[userAddress]["BscIn"]),
				MatchOut: userDetails[userAddress]["Out"].Sub(userDetails[userAddress]["BscOut"]),
			})
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
		userSet.Add(user)
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
			ledger = append(ledger, Ledger{
				User:      from,
				ChainName: "BSC",
				Amount:    amount,
				Direction: "In",
			})
			bscIn = bscIn.Add(amount)
			if userDetails[from] == nil {
				userDetails[from] = make(map[string]decimal.Decimal)
			}
			userDetails[from]["BscIn"] = userDetails[from]["BscIn"].Add(amount)
		} else if userMap[to] {
			ledger = append(ledger, Ledger{
				User:      to,
				ChainName: "BSC",
				Amount:    amount,
				Direction: "Out",
			})
			bscOut = bscOut.Add(amount)
			if userDetails[to] == nil {
				userDetails[to] = make(map[string]decimal.Decimal)
			}
			userDetails[to]["BscOut"] = userDetails[to]["BscOut"].Add(amount)
		}
	}
	// 获取Match充值【合约余额（get_tox_types 1）+ 可质押（get_pledge_balance_transferable 3）】
	var toxDayData []models2.ToxDayData
	tx := db.Mysql.Model(&models2.ToxDayData{}).Where("status=1")
	if req.Start > 0 {
		tx = tx.Where("date>=?", req.Start)
	}
	if req.End > 0 {
		tx = tx.Where("date<=?", req.End)
	}
	tx.Order("date desc")
	tx.Find(&toxDayData)
	for _, mT := range toxDayData {
		u := strings.ToLower(mT.Addr)
		if userMap[u] {
			ledger = append(ledger, Ledger{
				User:      u,
				ChainName: "Match",
				Amount:    mT.Amount,
				Direction: "In",
				TimeStamp: time.Unix(mT.Date, 0).Unix(),
				Ctime:     time.Unix(mT.Date, 0).Format("2006-01-02 15:04:05"),
			})
			matchIn = matchIn.Add(mT.Amount)
		}
		if userDetails[u] == nil {
			userDetails[u] = make(map[string]decimal.Decimal)
		}
		userDetails[u]["In"] = userDetails[u]["In"].Add(mT.Amount)
	}
	var pledgeBalanceTransferable []models2.PledgeBalanceTransferable
	tx = db.Mysql.Model(&models2.PledgeBalanceTransferable{}).Where("types=3")
	if req.Start > 0 {
		tx = tx.Where("timestamp>=?", req.Start)
	}
	if req.End > 0 {
		tx = tx.Where("timestamp<=?", req.End)
	}
	//tx.Where("sender in ?", sons.Data)
	tx.Order("timestamp desc")
	tx.Find(&pledgeBalanceTransferable)
	for _, pT := range pledgeBalanceTransferable {
		u := strings.ToLower(pT.Sender)
		if userMap[u] {
			ledger = append(ledger, Ledger{
				User:      u,
				ChainName: "Match",
				Amount:    pT.Amount,
				Direction: "In",
			})
			matchIn = matchIn.Add(pT.Amount)
		}
		if userDetails[u] == nil {
			userDetails[u] = make(map[string]decimal.Decimal)
		}
		userDetails[u]["In"] = userDetails[u]["In"].Add(pT.Amount)
	}
	// 获取Match提现
	toxDayData = []models2.ToxDayData{}
	tx = db.Mysql.Model(&models2.ToxDayData{}).Where("status=2")
	if req.Start > 0 {
		tx = tx.Where("date>=?", req.Start)
	}
	if req.End > 0 {
		tx = tx.Where("date<=?", req.End)
	}
	tx.Order("date desc")
	tx.Find(&toxDayData)
	for _, mT := range toxDayData {
		u := strings.ToLower(mT.Addr)
		if userMap[u] {
			ledger = append(ledger, Ledger{
				User:      u,
				ChainName: "Match",
				Amount:    mT.Amount,
				Direction: "Out",
			})
			matchOut = matchOut.Add(mT.Amount)
		}
		if userDetails[u] == nil {
			userDetails[u] = make(map[string]decimal.Decimal)
		}
		userDetails[u]["Out"] = userDetails[u]["Out"].Add(mT.Amount)
	}
	matchOut = matchOut.Sub(bscOut)
	matchIn = matchIn.Sub(bscIn)

	for _, userAddress := range userSet.Data {
		result = append(result, LedgerSum{
			User: userAddress,
			//In:       userDetails[userAddress]["In"].Div(decimalE18),
			In:       userDetails[userAddress]["In"],
			Out:      userDetails[userAddress]["Out"],
			BscIn:    userDetails[userAddress]["BscIn"],
			BscOut:   userDetails[userAddress]["BscOut"],
			MatchIn:  userDetails[userAddress]["In"].Sub(userDetails[userAddress]["BscIn"]),
			MatchOut: userDetails[userAddress]["Out"].Sub(userDetails[userAddress]["BscOut"]),
		})
	}
	// 保存excel
	// "github.com/xjieinfo/xjgo/xjcore/xjexcel"
	//f := xjexcel.ListToExcel(result, "团队充值提现", "下级成员详情")
	//fileName := fmt.Sprintf("./团队充值提现详情-%s.xls", req.User)
	//f.SaveAs(fileName)
	return statecode.CommonSuccess, result, bscIn, bscOut, matchIn, matchOut
}

func (c *Mining) RechargeSum(req *request.RechargeSum) (int, []Recharge) {
	var result []Recharge
	decimalE16, _ := decimal.NewFromString("10000000000000000")
	userSet := internal.NewStringSet()
	// 查全网数据
	if req.User == "" {
		return statecode.CommonErrServerErr, result
	}
	/// 查指定网体数据
	// 查询下级用户缓存
	yesterday := time.Now().Add(-time.Hour * 24).Format("2006-01-02")
	cacheKey := "LedgerDetails-" + yesterday + req.User
	ok, data := utils.EventCache.Get(cacheKey)
	if !ok {
		// 异步拉取数据
		SyncChain <- req.User
		return statecode.SyncingData, result
	}
	var sons Sons
	err := json.Unmarshal(data, &sons)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return statecode.CommonErrServerErr, result
	}
	userMap := make(map[string]bool)
	for _, user := range sons.Data {
		user = strings.ToLower(user)
		userMap[user] = true
		userSet.Add(user)
	}

	// 获取Match充值【合约余额（get_tox_types 1）+ 可质押（get_pledge_balance_transferable 3）】
	type WeightSum struct {
		Total decimal.Decimal
	}
	rechargeWeights := []models2.RechargeWeight{}
	db.Mysql.Model(&models2.RechargeWeight{}).Where("pioneer=?", strings.ToLower("0x77bd41fdE654FE0054b771Ec6985dC9d5247BAfe")).Find(&rechargeWeights)
	fmt.Println(len(rechargeWeights))
	for _, userAddress := range userSet.Data {
		// 获取该用户数据
		//rechargeWeight := models2.RechargeWeight{}

		decimalAmount, _ := decimal.NewFromString("0")
		for _, r := range rechargeWeights {
			if strings.ToLower(userAddress) == r.User {
				decimalAmount = decimalAmount.Add(r.Weight)
			}
		}
		//total := WeightSum{}
		//whereCondation := fmt.Sprintf("pioneer='%s' and user='%s' and day>='2023-11-01 00:00:00' and day<'2023-12-01 00:00:00'", strings.ToLower("0x77bd41fdE654FE0054b771Ec6985dC9d5247BAfe"), strings.ToLower(userAddress))
		//db.Mysql.Model(&models2.RechargeWeight{}).Where(whereCondation).Select("sum(weight) as total").Scan(&total)
		//fmt.Println(total.Total)
		result = append(result, Recharge{
			User:   userAddress,
			Amount: decimalAmount.Div(decimalE16),
		})
		//fmt.Println(decimalAmount.Mul(decimalE10))
	}

	// 保存excel
	//"github.com/xjieinfo/xjgo/xjcore/xjexcel"
	f := xjexcel.ListToExcel(result, "团队充值", "成员详情")
	fileName := fmt.Sprintf("./成都团队充值详情-团队长-%s.xls", req.User)
	f.SaveAs(fileName)
	return statecode.CommonSuccess, result
}

// RechargeSumInCity 查询此用户所在城市的网体业绩，网体不在这个城市的不做計算
func (c *Mining) RechargeSumInCity() int {

	ReadExcel("./assets/副本INTO工作室申请统计表(审核12月31日)发给技术.xlsx")
	//ReadExcel("./assets/单个INTO工作室申请统计表.xlsx")

	return statecode.CommonSuccess
}

// SyncUserData 查询此用户所在城市的网体业绩，网体不在这个城市的不做計算
func (c *Mining) SyncUserData() int {

	//ReadExcel("./assets/单个INTO工作室申请统计表.xlsx")
	SyncStatus("./assets/副本INTO工作室申请统计表(审核12月31日)发给技术.xlsx")

	return statecode.CommonSuccess
}

func ReadExcel(excelFile string) int {
	var studioInfos []StudioInfo
	f, err := excelize.OpenFile(excelFile)
	if err != nil {
		return statecode.CommonErrServerErr
	}
	// 取得 Sheet1 表格中所有的行
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return statecode.CommonErrServerErr
	}
	for i, row := range rows {
		if i == 1 || i == 0 {
			continue
		}
		studioInfo := StudioInfo{}
		for j, colCell := range row {
			if j == 1 {
				studioInfo.UserName = colCell
			} else if j == 2 {
				studioInfo.Level = colCell
			} else if j == 3 {
				studioInfo.User = strings.ToLower(colCell)
			} else if j == 4 {
				studioInfo.Super = colCell
			} else if j == 5 {
				studioInfo.Community = colCell
			} else if j == 6 {
				studioInfo.City = colCell
			} else if j == 7 {
				studioInfo.Address = colCell
			} else if j == 12 {
				studioInfo.CountTime = colCell
			}
		}
		studioInfos = append(studioInfos, studioInfo)
	}

	var result []StudioInfo
	for _, studioInfo := range studioInfos {
		userSet := internal.NewStringSet()
		var location models2.UserLocation
		where := "location like '%" + studioInfo.City + "%'"
		err = db.Mysql.Model(models2.UserLocation{}).Where(where).First(&location).Error
		if err != nil {
			log.Logger.Error(err.Error())
			return statecode.CommonErrServerErr
		}

		// 查询下级用户缓存
		userAddress := studioInfo.User
		yesterday := time.Now().Add(-time.Hour * 24).Format("2006-01-02")
		cacheKey := "LedgerDetails-" + yesterday + userAddress
		ok, data := utils.EventCache.Get(cacheKey)
		if !ok {
			log.Logger.Error("异步拉取数据")
			return statecode.SyncingData
		}
		var sons Sons
		err = json.Unmarshal(data, &sons)
		if err != nil {
			log.Logger.Sugar().Error(err)
			return statecode.CommonErrServerErr
		}
		userMap := make(map[string]bool)
		for _, user := range sons.Data {
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
		studioUserInfos := []StudioUserInfo{}
		for n, user := range userSet.Data {
			chData <- n
			wg.Add(1)
			go func(u string, index int, studio StudioInfo, studioUserInfos *[]StudioUserInfo, lock *sync.Mutex) {
				_, d, inCity := GetWeight(u, location.CityId, studio.CountTime, studio.UserName)
				lock.Lock()
				defer lock.Unlock()
				weight = weight.Add(d)
				<-chData
				wg.Done()
				userCityBytes, _ := db.LevelDb.Get([]byte(fmt.Sprintf("%s_city", u)), nil)

				*studioUserInfos = append(*studioUserInfos, StudioUserInfo{
					UserName:       studio.UserName,
					User:           u,
					InCity:         inCity,
					BindCity:       studio.City,
					UserCity:       string(userCityBytes),
					RechargeWeight: d.String(),
				})
			}(user, n, studioInfo, &studioUserInfos, &addLock)
		}
		wg.Wait()
		log.Logger.Sugar().Infof("工作室:%s,错误:%v,网体用户:%s,用户数量:%d,weight:%s,%s", studioInfo.UserName, err, userAddress, len(userSet.Data), weight.String(), "------------------------------------------------------------")
		studioInfo.RechargeWeight = weight.Div(d18).String()
		result = append(result, studioInfo)

		countTimeSlice := strings.Split(studioInfo.CountTime, "--")
		countTimeSlice = strings.Split(countTimeSlice[0], "/")
		f = xjexcel.ListToExcel(studioUserInfos, "团队充值", "详情")
		fileName := fmt.Sprintf("./工作室(%s)-%s[%s_%s]%s[%v].xlsx",
			studioInfo.UserName,
			studioInfo.City,
			fmt.Sprintf("%s-%s-%s", countTimeSlice[0], countTimeSlice[1], countTimeSlice[2]),
			"2024-01-01",
			studioInfo.User,
			studioInfo.RechargeWeight,
		)
		f.SaveAs(fileName)
	}

	f = xjexcel.ListToExcel(result, "团队充值", "详情")
	fileName := fmt.Sprintf("./工作室充值权重.xlsx")
	f.SaveAs(fileName)

	return statecode.CommonSuccess
}

func GetWeight(user, cityId, countTime, userName string) (error, decimal.Decimal, bool) {
	inCity := false
	key := fmt.Sprintf("%s-%s", user, cityId)
	var val string
	valBytes, err := db.LevelDb.Get([]byte(key), nil)
	if err != nil {
		userLocation := models2.UserLocation{}
		err = db.Mysql.Model(models2.UserLocation{}).Where("user=?", user).First(&userLocation).Error
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
	//if string(valBytes) == "3" {
	//	InCity(user, cityId)
	//}
	if !inCity {
		return err, decimal.Zero, inCity
	}

	// 获取该用户网体充值权重，审核通过后到12月31日的充值权重
	countTimeSlice := strings.Split(countTime, "--")
	countTimeSlice = strings.Split(countTimeSlice[0], "/")
	err, d := GetRechargeWeight(user, fmt.Sprintf("%s-%s-%s 00:00:00", countTimeSlice[0], countTimeSlice[1], countTimeSlice[2]), "2024-01-01 00:00:00")
	//err, d := GetRechargeWeightFromDb(user, fmt.Sprintf("%s-%s-%s 00:00:00", countTimeSlice[0], countTimeSlice[1], countTimeSlice[2]), "2024-01-01 00:00:00")
	if err != nil {
		log.Logger.Sugar().Debugf("工作室：%s，是否在绑定城市：%v,下级：%s,获取权重错误：%v", userName, inCity, user, err)
		return err, decimal.Zero, inCity
	}
	log.Logger.Sugar().Debugf("工作室：%s，是否在绑定城市：%v,下级：%s,权重：%s", userName, inCity, user, d.String())
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
		for k := 0; k < 20; k++ {
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

func GetRechargeWeightFromDb(user, start, end string) (error, decimal.Decimal) {
	type FHSum struct {
		Total decimal.Decimal
	}
	var total FHSum

	for {
		err := db.Mysql.Model(&models2.RechargeWeightAll{}).Where("user=? and day>=? and day<?", strings.ToLower(user), start, end).Select("sum(weight) as total").Scan(&total).Error
		if err == nil {
			return nil, total.Total
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, decimal.NewFromInt(0)
		}
	}
}

func InCity(user, cityId string) bool {
	var inCity bool
	var key, val string
	userLocation := models2.UserLocation{}
	err := db.Mysql.Model(models2.UserLocation{}).Where("user=?", user).First(&userLocation).Error
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

func SyncStatus(excelFile string) {
	f, err := excelize.OpenFile(excelFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 取得 Sheet1 表格中所有的行
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, row := range rows {
		if i == 1 || i == 0 {
			continue
		}
		for j, colCell := range row {
			if j == 6 {
				fmt.Println(colCell)
				var location models2.UserLocation
				where := "location like '%" + colCell + "%'"
				err = db.Mysql.Model(models2.UserLocation{}).Where(where).First(&location).Error
				fmt.Println(err)
				if err != nil {
					return
				}
			}

			if j == 3 {
				// 查询下级用户缓存
				userAddress := strings.ToLower(colCell)
				yesterday := time.Now().Add(-time.Hour * 24).Format("2006-01-02")
				cacheKey := "LedgerDetails-" + yesterday + userAddress
				ok, _ := utils.EventCache.Get(cacheKey)
				if !ok {
					// 异步拉取数据
					SyncChain <- userAddress
					fmt.Println("异步拉取数据", i)
					continue
				}
			}
		}
	}
}

// GetToxTxBridgeBsc 获取bsc bridge数据
func GetToxTxBridgeBsc() (error, ToxTxBridge) {
	url := "https://intocache.intoverse.co/tox_tx_bridge_bsc.json"
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
