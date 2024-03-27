package main

import (
	"city-node-server/pkg/blockchain"
	"city-node-server/pkg/db"
	"city-node-server/services"
	"github.com/jasonlvhit/gocron"
	"time"
)

func main() {
	db.InitMysql()
	services.InitTestNet()
	//services.InitUserLocation()
	//services.InitAppraise()
	//services.InitCityPioneer(20)
	//services.InitCity(20)
	//services.InitCityPioneerData()

	//db.InitMysql()

	//services.AddPioneerBeth3()

	taskTest()
	// 四期设置
	// 我的
	// "0x3edf72f7ab938a14e6aae3701a1e5acbb2512c840a019c6ead9d01415dbac864",
	// "0x365Cae736D93Ad3e388919d0E4d3EE6Ed364b060",

	//services.AddPioneerBeth4()

	//err := services.AddPioneerBatch4(
	//	"0xe48baef0767f2198d4a783075148f0d7650294f840652960e9ea74f56c9171a6",
	//	"0xb125f8E9888CBf3598A24E5FF96C71e002cfCD81",
	//	1,
	//	10, //5000 6
	//	5,  // 1000 3
	//	4,
	//	1)
	//fmt.Println(err)

	//services.ResetYeHui()
	//services.ResetYuBin()
	//services.ResetJingJing("0x81F61e0b02d899956d2d96ACc2c6F9Cb43D6b99d") // 焦作市先锋 三线
	//services.ResetJingJing("0x28ddE14d31Bd7025b3Db1FA8eC7C5707E4FFE1e8") // 杭州先锋 二线
	//services.ResetJingJing("0xc1CfC68d1CD8CE6571ee8ee167b2c80E249bCe32") // 新乡先锋 一线
	//services.ResetJingJing("0x360C815e8C5F130913113801D0c57611Ee95723A") // 新乡一级县区
	//services.ResetJingJing("0xcD1f731A1529d5F8e8f8cA94dF6092B680C88e2E") // 新乡原阳，二级区域
	//services.AdminSetRechargeAmountTask()
	//services.AdminSetDelegateTask()
	//services.ResetWangPing("0x74bBa41049B40803cBcF97ea6F60d0e064d24B72")
	//services.ResetWangPing("0x010E293425F6Ad6D498893267C096853603D0d42")
	//services.ResetWangPing("0x010E293425F6Ad6D498893267C096853603D0d42")
	//services.ResetWangPing("0xa4394cA3985598Fff63E9551e81DD4D8602eA68B")
	//services.ResetWangPing("0xA2C100a4C098D3D05Ff8c0AE2CE4Ea93eCe550A8")
	//services.ResetWangPing("0xbF0D360879Cc0AED8479058fbf8f86F510512188")
	//services.ResetWangPing("0xC250bFC206Ba2d9b761Eb51537ef8D665e0a54d7")

	//taskTest()

	//err := services.RemovePioneer(
	//	"0xd35a0ba3b4048cea0b00f3659353645b31affe150409591b8292de2933df2f1f",
	//	"0x010E293425F6Ad6D498893267C096853603D0d42",
	//)
	//fmt.Println(err)

}

func taskTest() {
	services.InitTestNet()

	s := gocron.NewScheduler()
	s.ChangeLoc(time.UTC)
	//_ = s.Every(3).Seconds().From(gocron.NextTick()).Do(services.AdminSetDelegateTask)
	_ = s.Every(9).Seconds().From(gocron.NextTick()).Do(blockchain.TriggerAllPioneerTaskTestNet)
	//_ = s.Every(10).Seconds().From(gocron.NextTick()).Do(services.AdminSetRechargeAmountTask)
	<-s.Start() // Start all the pending jobs
}
