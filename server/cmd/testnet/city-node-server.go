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

	//taskTest()
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
	//services.ResetJingJing("0x360C815e8C5F130913113801D0c57611Ee95723A")
	//services.ResetJingJing("0xb3b640E65eA83BC683115348bb63b78470097A30")
	//services.ResetWangPing("0x74bBa41049B40803cBcF97ea6F60d0e064d24B72")
	//services.ResetWangPing("0x010E293425F6Ad6D498893267C096853603D0d42")

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
	_ = s.Every(25).Seconds().From(gocron.NextTick()).Do(blockchain.TriggerAllPioneerTaskTestNet)
	//_ = s.Every(10).Seconds().From(gocron.NextTick()).Do(services.AdminSetRechargeAmountTask)
	<-s.Start() // Start all the pending jobs
}
