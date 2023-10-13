package services

import (
	"city-node-server/db"
	"fmt"
	"github.com/shopspring/decimal"
	"time"
)

type CountData struct {
	Pioneer string `json:"pioneer" gorm:"column:pioneer"`
	Count   int64  `json:"count" gorm:"column:count"`
}
type CtimeData struct {
	Ctime time.Time `json:"ctime" gorm:"column:ctime"`
}
type YingFaData struct {
	Pioneer string          `json:"pioneer" gorm:"column:pioneer"`
	Yingfa  decimal.Decimal `json:"yingfa" gorm:"column:yingfa"`
}
type ShijiFafangData struct {
	Pioneer     string          `json:"pioneer" gorm:"column:pioneer"`
	ShijiFafang decimal.Decimal `json:"shiji_fafang" gorm:"column:shiji_fafang"`
}

func CheckData() {
	db.InitMysql()
	var countData []CountData
	db.Mysql.Raw("SELECT pioneer, COUNT(*) as count FROM reward where ctime>='2023-10-13 00:00:00' and  ctime<='2023-10-13 23:59:59' GROUP BY pioneer HAVING COUNT(*) >1").Scan(&countData)

	for _, d := range countData {
		ctimeData := CtimeData{}
		sql := fmt.Sprintf("SELECT ctime FROM reward where pioneer='%s' and ctime>='2023-10-13 00:00:00' and  ctime<='2023-10-13 23:59:59' ORDER BY ctime asc limit 1", d.Pioneer)
		db.Mysql.Raw(sql).Scan(&ctimeData)

		yingFaData := YingFaData{}
		sql = fmt.Sprintf("SELECT pioneer,SUM(founds_reward+delegate_reward+node_reward) as yingfa FROM reward where pioneer='%s' and ctime='%s' ", d.Pioneer, ctimeData.Ctime)
		db.Mysql.Raw(sql).Scan(&yingFaData)

		shijiFafangData := ShijiFafangData{}
		sql = fmt.Sprintf("SELECT pioneer,SUM(founds_reward+delegate_reward+node_reward) as shiji_fafang  FROM reward where pioneer='%s' and ctime>='2023-10-13 00:00:00' and  ctime<='2023-10-13 23:59:59'", d.Pioneer)
		db.Mysql.Raw(sql).Scan(&shijiFafangData)

		fmt.Println(d.Pioneer, shijiFafangData.ShijiFafang, yingFaData.Yingfa, shijiFafangData.ShijiFafang.Sub(yingFaData.Yingfa))

	}

}
