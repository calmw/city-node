package services

import (
	"city-node-server/pkg/blockchain"
	"city-node-server/pkg/db"
	"city-node-server/pkg/log"
	models2 "city-node-server/pkg/models"
	"city-node-server/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

func InitCityPioneer(secondsPerDay int64) {

	cityPioneer := blockchain.NewCityPioneer()
	//cityPioneer.AdminSetAppraise()
	cityPioneer.AdminChangePioneerAddress("0x5db860601869Dad7Eb2961341056b389C3149e5f", "0x12Cc2278b37c2751D11B5A64712Ff439b57F6E6a")
	//cityPioneer.AdminSetSecondsPerDay(secondsPerDay)
	// 管理员设置TOX代币地址
	//AdminSetTOXAddress
	// 管理员设置城市合约地址
	//blockchain.AdminSetCityAddress()
	// 管理员设置增加用户合约余额合约地址
	//blockchain.AdminSetMiningAddress()
	//
	////// 设置城市先锋考核标准
	//blockchain.AdminSetAssessmentCriteria(1, 1, 2500)
	//blockchain.AdminSetAssessmentCriteria(1, 2, 5000)
	//blockchain.AdminSetAssessmentCriteria(1, 3, 10000)
	//blockchain.AdminSetAssessmentCriteria(2, 1, 1250)
	//blockchain.AdminSetAssessmentCriteria(2, 2, 2500)
	//blockchain.AdminSetAssessmentCriteria(2, 3, 5000)
	//blockchain.AdminSetAssessmentCriteria(3, 1, 625)
	//blockchain.AdminSetAssessmentCriteria(3, 2, 1250)
	//blockchain.AdminSetAssessmentCriteria(3, 3, 2500)
	//// 管理员设置城市先锋保证金退还标准,点数
	//blockchain.AdminSetAssessmentReturnCriteria(1, 1, 10000)
	//blockchain.AdminSetAssessmentReturnCriteria(1, 2, 20000)
	//blockchain.AdminSetAssessmentReturnCriteria(1, 3, 30000)
	//blockchain.AdminSetAssessmentReturnCriteria(1, 4, 10000)
	//blockchain.AdminSetAssessmentReturnCriteria(1, 5, 20000)
	//blockchain.AdminSetAssessmentReturnCriteria(1, 6, 30000)
	//blockchain.AdminSetAssessmentReturnCriteria(2, 1, 5000)
	//blockchain.AdminSetAssessmentReturnCriteria(2, 2, 10000)
	//blockchain.AdminSetAssessmentReturnCriteria(2, 3, 15000)
	//blockchain.AdminSetAssessmentReturnCriteria(2, 4, 5000)
	//blockchain.AdminSetAssessmentReturnCriteria(2, 5, 10000)
	//blockchain.AdminSetAssessmentReturnCriteria(2, 6, 15000)
	//blockchain.AdminSetAssessmentReturnCriteria(3, 1, 2500)
	//blockchain.AdminSetAssessmentReturnCriteria(3, 2, 5000)
	//blockchain.AdminSetAssessmentReturnCriteria(3, 3, 7500)
	//blockchain.AdminSetAssessmentReturnCriteria(3, 4, 2500)
	//blockchain.AdminSetAssessmentReturnCriteria(3, 5, 5000)
	//blockchain.AdminSetAssessmentReturnCriteria(3, 6, 7500)
	//// 管理员设置保证金退还比例
	//blockchain.AdminSetAssessmentReturnRate(1, 1, 50)
	//blockchain.AdminSetAssessmentReturnRate(1, 2, 75)
	//blockchain.AdminSetAssessmentReturnRate(1, 3, 100)
	//blockchain.AdminSetAssessmentReturnRate(1, 4, 25)
	//blockchain.AdminSetAssessmentReturnRate(1, 5, 50)
	//blockchain.AdminSetAssessmentReturnRate(1, 6, 75)
	//blockchain.AdminSetAssessmentReturnRate(2, 1, 50)
	//blockchain.AdminSetAssessmentReturnRate(2, 2, 75)
	//blockchain.AdminSetAssessmentReturnRate(2, 3, 100)
	//blockchain.AdminSetAssessmentReturnRate(2, 4, 25)
	//blockchain.AdminSetAssessmentReturnRate(2, 5, 50)
	//blockchain.AdminSetAssessmentReturnRate(2, 6, 75)
	//blockchain.AdminSetAssessmentReturnRate(3, 1, 50)
	//blockchain.AdminSetAssessmentReturnRate(3, 2, 75)
	//blockchain.AdminSetAssessmentReturnRate(3, 3, 100)
	//blockchain.AdminSetAssessmentReturnRate(3, 4, 25)
	//blockchain.AdminSetAssessmentReturnRate(3, 5, 50)
	//blockchain.AdminSetAssessmentReturnRate(3, 6, 75)
	//
	//// AddCityPioneerAdmin 给城市合约、IntoMining设置用户（增加用户合约余额）添加管理员权限
	//blockchain.AddCityPioneerAdmin()

	// 管理员设置开始考核时间,先交保证金，后考核
	//blockchain.AdminSetStartTime(time.Now().Unix())
}

type ToxTx struct {
	Message string `json:"message"`
	Result  []struct {
		Value             string `json:"value"`
		BlockHash         string `json:"blockHash"`
		BlockNumber       string `json:"blockNumber"`
		Confirmations     string `json:"confirmations"`
		ContractAddress   string `json:"contractAddress"`
		CumulativeGasUsed string `json:"cumulativeGasUsed"`
		From              string `json:"from"`
		Gas               string `json:"gas"`
		GasPrice          string `json:"gasPrice"`
		GasUsed           string `json:"gasUsed"`
		Hash              string `json:"hash"`
		Input             string `json:"input"`
		LogIndex          string `json:"logIndex"`
		Nonce             string `json:"nonce"`
		TimeStamp         string `json:"timeStamp"`
		To                string `json:"to"`
		TokenDecimal      string `json:"tokenDecimal"`
		TokenName         string `json:"tokenName"`
		TokenSymbol       string `json:"tokenSymbol"`
		TransactionIndex  string `json:"transactionIndex"`
	} `json:"result"`
	Status string `json:"status"`
}

func GetToxTx() {
	url := "https://lisbon.matchscan.io/api?module=account&action=tokentx&address=0x60C541388077d524178521A7ceD95d0c7a016B72"
	err, res := utils.HttpGet(url)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	var toxTx ToxTx
	json.Unmarshal(res, &toxTx)
	fmt.Println(len(toxTx.Result))
	for _, t := range toxTx.Result {
		if t.ContractAddress != "0x96397347ea2bee29713bc48749eb277d6a36a407" {
			fmt.Println(`ERC20地址不是：0x96397347ea2bee29713bc48749eb277d6a36a407`)
			continue
		}
		if t.From != "0x60c541388077d524178521a7ced95d0c7a016b72" && t.To != "0x60c541388077d524178521a7ced95d0c7a016b72" {
			fmt.Println(`from和to都不为0x60c541388077d524178521a7ced95d0c7a016b72`)
			continue
		}
		if t.To == "0x60c541388077d524178521a7ced95d0c7a016b72" {
			var poneer models2.Pioneer
			err = db.Mysql.Model(models2.Pioneer{}).Where("pioneer=?", strings.ToLower(t.From)).First(&poneer).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				fmt.Println(t.From, `非先锋交保证金`)
				continue
			}
		}

		var cityPioneerToxTx models2.CityPioneerToxTx
		err = db.Mysql.Model(models2.CityPioneerToxTx{}).Where("block_height=? and log_index=?", t.BlockNumber, t.LogIndex).First(&cityPioneerToxTx).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			value, _ := decimal.NewFromString(t.Value)
			value = value.Div(decimal.NewFromInt(1e18))
			blockNumber, _ := strconv.ParseInt(t.BlockNumber, 10, 64)
			logIndex, _ := strconv.ParseInt(t.LogIndex, 10, 64)
			timeStamp, _ := strconv.ParseInt(t.TimeStamp, 10, 64)
			err = db.Mysql.Model(&models2.CityPioneerToxTx{}).Create(&models2.CityPioneerToxTx{
				From:        t.From,
				To:          t.To,
				TxHash:      t.Hash,
				Value:       value,
				BlockHeight: blockNumber,
				LogIndex:    logIndex,
				Ctime:       time.Unix(timeStamp, 0).Format(blockchain.LayoutTime),
			}).Error
		}
		if t.To == "0x60c541388077d524178521a7ced95d0c7a016b72" {
			var pioneer models2.Pioneer
			//err = db.Mysql.Model(models.Pioneer{}).Where("pioneer=? and surety=0", t.From).First(&pioneer).Error
			err = db.Mysql.Model(models2.Pioneer{}).Where("pioneer=?", t.From).First(&pioneer).Error
			if err == gorm.ErrRecordNotFound {
				fmt.Println("非先锋转入", t.From)
			}
			value, _ := decimal.NewFromString(t.Value)
			db.Mysql.Model(models2.Pioneer{}).Where("pioneer=?", t.From).UpdateColumns(map[string]interface{}{
				"surety": pioneer.Surety.Add(value),
			})
			fmt.Println(t.From, pioneer.CityLevel, t.Value)
		} else {
			fmt.Println(t.From, t.To)
		}
	}

}
