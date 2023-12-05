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
