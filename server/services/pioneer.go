package services

import (
	"city-node-server/pkg/blockchain"
	"city-node-server/pkg/log"
	"strings"
	"time"
)

func AddPioneerBatch4(areaShiId, pioneerAddress string, level int64, suretyTox int64, suretyUsdt int64, pioneerBatch int64, pioneerType int64) error {
	err, cli := blockchain.Client(blockchain.CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	// 设置所需缴纳的TOX，城市等级
	city := blockchain.NewCity(cli)
	err = city.AdminSetChengShiLevelAndSurety(areaShiId, level, suretyTox)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}

	// 设置所需缴纳的USDT
	cityPioneerData := blockchain.NewCityPioneerData(cli)
	err = cityPioneerData.SuretyUSDT(areaShiId, suretyUsdt)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}

	err = city.AdminSetPioneer(areaShiId, pioneerAddress, pioneerBatch, pioneerType)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	return nil
}

func RemovePioneer(areaShiId, pioneerAddress string) error {
	err, cli := blockchain.Client(blockchain.CityNodeConfig)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}
	city := blockchain.NewCity(cli)
	err = city.AdminRemovePioneer(areaShiId, pioneerAddress)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err
	}

	return nil
}

func AddPioneerTest() {
	// 我的
	// "0x3edf72f7ab938a14e6aae3701a1e5acbb2512c840a019c6ead9d01415dbac864",
	// "0x365Cae736D93Ad3e388919d0E4d3EE6Ed364b060",

	// 玉斌
	// 0xe48baef0767f2198d4a783075148f0d7650294f840652960e9ea74f56c9171a6
	// 0xFf2B12085c7F7B8133eEf5006703A6c055a3ed7d

	//services.AddPioneerBeth4()
	//err := services.AddPioneerBatch4(
	//	"0xe48baef0767f2198d4a783075148f0d7650294f840652960e9ea74f56c9171a6",
	//	"0xFf2B12085c7F7B8133eEf5006703A6c055a3ed7d",
	//	1,
	//	10, //5000 6
	//	5, // 1000 3
	//	4,
	//	1)
	//fmt.Println(err, 123456)

	//err := services.AddPioneerBatch4(
	//	"0xe48baef0767f2198d4a783075148f0d7650294f840652960e9ea74f56c9171a6",
	//	"0xFf2B12085c7F7B8133eEf5006703A6c055a3ed7d",
	//	1,
	//	6, //5000 6
	//	3, // 1000 3
	//	4,
	//	0)
	//fmt.Println(err, 123456)

	//0x0a23fb256dc6340f9287a2edf424eb5cb92f37f39a8f58cc6c1a9fa349d51b34

	//services.RemovePioneer(
	//	"0xe48baef0767f2198d4a783075148f0d7650294f840652960e9ea74f56c9171a6",
	//	"0xFf2B12085c7F7B8133eEf5006703A6c055a3ed7d",
	//)
}

func ResetYuBin() {
	// 玉斌
	// 0x3ebfe141b23e28220361b813ee0fd5b18eb15db214c58d20fa66c0875dbbb281  // 湛江 徐闻县
	// 0xFf2B12085c7F7B8133eEf5006703A6c055a3ed7d

	for {
		err := RemovePioneer(
			"0x3ebfe141b23e28220361b813ee0fd5b18eb15db214c58d20fa66c0875dbbb281",
			"0xFf2B12085c7F7B8133eEf5006703A6c055a3ed7d",
		)
		if err == nil {
			break
		}
	}
	time.Sleep(time.Second * 2)
	for {
		err := AddPioneerBatch4(
			"0x3ebfe141b23e28220361b813ee0fd5b18eb15db214c58d20fa66c0875dbbb281",
			"0xFf2B12085c7F7B8133eEf5006703A6c055a3ed7d",
			1,
			10, //5000 6
			5,  // 1000 3
			4,
			1)
		if err == nil {
			break
		}
	}
}

func ResetYeHui() {
	// 业辉
	// 0x4c26fda7d3080732ee56f0d2bfd089ae1ff8040a568cbddfabc7df759a320ca9
	// 0x287229fD161feaf69a14b6848Da2690C29d4a0Ea

	for {
		err := RemovePioneer(
			"0x4c26fda7d3080732ee56f0d2bfd089ae1ff8040a568cbddfabc7df759a320ca9",
			"0x287229fD161feaf69a14b6848Da2690C29d4a0Ea",
		)
		if err == nil {
			break
		}
	}
	time.Sleep(time.Second * 2)
	for {
		err := AddPioneerBatch4(
			"0x4c26fda7d3080732ee56f0d2bfd089ae1ff8040a568cbddfabc7df759a320ca9",
			"0x287229fD161feaf69a14b6848Da2690C29d4a0Ea",
			1,
			10, //5000 6
			5,  // 1000 3
			4,
			1)
		if err == nil {
			break
		}
	}
}

func ResetJingJing(pioneer string) {
	/// 景景
	//	0xd35a0ba3b4048cea0b00f3659353645b31affe150409591b8292de2933df2f1f, // 新乡红旗区，区域节点，一级县区
	//	0x360C815e8C5F130913113801D0c57611Ee95723A,
	//  0x9882a6038f5037928f8ceabccbbbe73652f94cda789484427957542073dbb2f5  // 新乡，城市节点，一级城市
	//  0xb3b640E65eA83BC683115348bb63b78470097A30 定位区县 0x64a5b80bf352b0ede5c59ffe909636b4b9dfeb3314717d3edcf260be634d25c5 郑州管城区

	// 焦作市
	//blockchain.AdminRemovePioneer("0x85202e610a59bdbc7559cabd41c28f4ecfe18675575572ac5822ca697df5103d", "0x81F61e0b02d899956d2d96ACc2c6F9Cb43D6b99d")
	//blockchain.AdminSetPioneer("0x85202e610a59bdbc7559cabd41c28f4ecfe18675575572ac5822ca697df5103d", "0x81F61e0b02d899956d2d96ACc2c6F9Cb43D6b99d")
	// 杭州
	//blockchain.AdminRemovePioneer("0xe48baef0767f2198d4a783075148f0d7650294f840652960e9ea74f56c9171a6", "0x28ddE14d31Bd7025b3Db1FA8eC7C5707E4FFE1e8")
	//blockchain.AdminSetPioneer("0xe48baef0767f2198d4a783075148f0d7650294f840652960e9ea74f56c9171a6", "0x28ddE14d31Bd7025b3Db1FA8eC7C5707E4FFE1e8")

	// 新乡
	//blockchain.AdminRemovePioneer("0x9882a6038f5037928f8ceabccbbbe73652f94cda789484427957542073dbb2f5", "0xc1CfC68d1CD8CE6571ee8ee167b2c80E249bCe32")

	// 新乡原阳，二级先锋
	// 0xcD1f731A1529d5F8e8f8cA94dF6092B680C88e2E 0xe5abc510cd002e604dff20e595f121e073dad93b136574a0cd571f1f407404a8

	if strings.ToLower("0x360C815e8C5F130913113801D0c57611Ee95723A") == strings.ToLower(pioneer) {
		for {
			err := RemovePioneer(
				"0xd35a0ba3b4048cea0b00f3659353645b31affe150409591b8292de2933df2f1f",
				"0x360C815e8C5F130913113801D0c57611Ee95723A",
			)
			if err == nil {
				break
			}
		}
		time.Sleep(time.Second * 2)
		for {
			err := AddPioneerBatch4(
				"0xd35a0ba3b4048cea0b00f3659353645b31affe150409591b8292de2933df2f1f",
				"0x360C815e8C5F130913113801D0c57611Ee95723A",
				1,
				10, //5000 6
				5,  // 1000 3
				4,
				1)
			if err == nil {
				break
			}
		}
	} else if strings.ToLower("0xb3b640E65eA83BC683115348bb63b78470097A30") == strings.ToLower(pioneer) {
		for {
			err := RemovePioneer(
				"0x9882a6038f5037928f8ceabccbbbe73652f94cda789484427957542073dbb2f5",
				"0xb3b640E65eA83BC683115348bb63b78470097A30",
			)
			if err == nil {
				break
			}
		}
		time.Sleep(time.Second * 2)
		for {
			err := AddPioneerBatch4(
				"0x9882a6038f5037928f8ceabccbbbe73652f94cda789484427957542073dbb2f5",
				"0xb3b640E65eA83BC683115348bb63b78470097A30",
				1,
				100, //5000 6
				100, // 1000 3
				4,
				0)
			if err == nil {
				break
			}
		}
	} else if strings.ToLower("0x81F61e0b02d899956d2d96ACc2c6F9Cb43D6b99d") == strings.ToLower(pioneer) {
		for {
			err := RemovePioneer(
				"0x85202e610a59bdbc7559cabd41c28f4ecfe18675575572ac5822ca697df5103d",
				"0x81F61e0b02d899956d2d96ACc2c6F9Cb43D6b99d",
			)
			if err == nil {
				break
			}
		}
		time.Sleep(time.Second * 2)
		for {
			err := AddPioneerBatch4(
				"0x85202e610a59bdbc7559cabd41c28f4ecfe18675575572ac5822ca697df5103d",
				"0x81F61e0b02d899956d2d96ACc2c6F9Cb43D6b99d",
				3,
				60, //5000 6
				40, // 1000 3
				4,
				0)
			if err == nil {
				break
			}
		}
	} else if strings.ToLower("0x28ddE14d31Bd7025b3Db1FA8eC7C5707E4FFE1e8") == strings.ToLower(pioneer) {
		for {
			err := RemovePioneer(
				"0xe48baef0767f2198d4a783075148f0d7650294f840652960e9ea74f56c9171a6",
				"0x28ddE14d31Bd7025b3Db1FA8eC7C5707E4FFE1e8",
			)
			if err == nil {
				break
			}
		}
		time.Sleep(time.Second * 2)
		for {
			err := AddPioneerBatch4(
				"0xe48baef0767f2198d4a783075148f0d7650294f840652960e9ea74f56c9171a6",
				"0x28ddE14d31Bd7025b3Db1FA8eC7C5707E4FFE1e8",
				2,
				80, //5000 6
				60, // 1000 3
				4,
				0)
			if err == nil {
				break
			}
		}
	} else if strings.ToLower("0xc1CfC68d1CD8CE6571ee8ee167b2c80E249bCe32") == strings.ToLower(pioneer) {
		for {
			err := RemovePioneer(
				"0x9882a6038f5037928f8ceabccbbbe73652f94cda789484427957542073dbb2f5",
				"0xc1CfC68d1CD8CE6571ee8ee167b2c80E249bCe32",
			)
			if err == nil {
				break
			}
		}
		time.Sleep(time.Second * 2)
		for {
			err := AddPioneerBatch4(
				"0x9882a6038f5037928f8ceabccbbbe73652f94cda789484427957542073dbb2f5",
				"0xc1CfC68d1CD8CE6571ee8ee167b2c80E249bCe32",
				1,
				100, //5000 6
				100, // 1000 3
				4,
				0)
			if err == nil {
				break
			}
		}
	} else if strings.ToLower("0xcD1f731A1529d5F8e8f8cA94dF6092B680C88e2E") == strings.ToLower(pioneer) {
		for {
			err := RemovePioneer(
				"0xe5abc510cd002e604dff20e595f121e073dad93b136574a0cd571f1f407404a8",
				"0xcD1f731A1529d5F8e8f8cA94dF6092B680C88e2E",
			)
			if err == nil {
				break
			}
		}
		time.Sleep(time.Second * 2)
		for {
			err := AddPioneerBatch4(
				"0xe5abc510cd002e604dff20e595f121e073dad93b136574a0cd571f1f407404a8",
				"0xcD1f731A1529d5F8e8f8cA94dF6092B680C88e2E",
				2,
				6, //5000 6
				3, // 1000 3
				4,
				1)
			if err == nil {
				break
			}
		}
	}
}

func ResetWangPing(pioneer string) {
	/// 王平
	// 0xfcbd8b9c3139222ae5d427bdf278de3d9f96ffa721e92fbfb125ecb6707998cf, // 渭南市 （三线城市）
	// 0x74bBa41049B40803cBcF97ea6F60d0e064d24B72,
	// 0x4de9096348869d087c953a1b6df5267276b15929bf8c3eedd52332bb946dadda  // 渭南市韩城市（二级区域）
	// 0x010E293425F6Ad6D498893267C096853603D0d42

	// 0xa4394cA3985598Fff63E9551e81DD4D8602eA68B
	//0x9181a959c9c4a16284eadf62405b56b5a4cf39ff14f8b9efbe09feeb37d99e25
	//0xe48baef0767f2198d4a783075148f0d7650294f840652960e9ea74f56c9171a6
	//
	//0xA2C100a4C098D3D05Ff8c0AE2CE4Ea93eCe550A8
	//0xed725125d1de69c4aad9d3efd40a85622632d196a0dbf11935f0fb278aed2057 合肥 蜀山区
	//0x730ad915d608f40a7cea880f5abc7fe6bb253f903667be02ec00dd3028f7d24a
	//
	//0xbF0D360879Cc0AED8479058fbf8f86F510512188
	//0xdb60dad7d29002a0b4379895f5be941e8b8b9c073ff54d746b28d23b60291cb6 福州 鼓楼区
	//0x491f0fa71e9db46360a7ac880539f923cda458800f22fb344e10e98491f1f39a
	//
	//0xC250bFC206Ba2d9b761Eb51537ef8D665e0a54d7
	//0xdb60dad7d29002a0b4379895f5be941e8b8b9c073ff54d746b28d23b60291cb6 福州 鼓楼区
	//0xc50db5c60dc66068637ffc825954152fbce8b73a290285f252697c31b005a382  佛山

	if strings.ToLower("0x010E293425F6Ad6D498893267C096853603D0d42") == strings.ToLower(pioneer) {
		for {
			err := RemovePioneer(
				"0x4de9096348869d087c953a1b6df5267276b15929bf8c3eedd52332bb946dadda",
				"0x010E293425F6Ad6D498893267C096853603D0d42",
			)
			if err == nil {
				break
			}
		}
		time.Sleep(time.Second * 2)
		for {
			err := AddPioneerBatch4(
				"0x4de9096348869d087c953a1b6df5267276b15929bf8c3eedd52332bb946dadda",
				"0x010E293425F6Ad6D498893267C096853603D0d42",
				2,
				10, //5000 6
				5,  // 1000 3
				4,
				1)
			if err == nil {
				break
			}
		}
	} else if strings.ToLower("0x74bBa41049B40803cBcF97ea6F60d0e064d24B72") == strings.ToLower(pioneer) {
		for {
			err := RemovePioneer(
				"0xfcbd8b9c3139222ae5d427bdf278de3d9f96ffa721e92fbfb125ecb6707998cf",
				"0x74bBa41049B40803cBcF97ea6F60d0e064d24B72",
			)
			if err == nil {
				break
			}
		}
		time.Sleep(time.Second * 2)
		for {
			err := AddPioneerBatch4(
				"0xfcbd8b9c3139222ae5d427bdf278de3d9f96ffa721e92fbfb125ecb6707998cf",
				"0x74bBa41049B40803cBcF97ea6F60d0e064d24B72",
				3,
				40, //5000 6
				60, // 1000 3
				4,
				0)
			if err == nil {
				break
			}
		}
	} else if strings.ToLower("0xa4394cA3985598Fff63E9551e81DD4D8602eA68B") == strings.ToLower(pioneer) {
		for {
			err := RemovePioneer(
				"0xee72b2ea43e56148d63a0adf23653790ab0e0ab31bc0f8a4326ab9f377ee931a",
				"0xa4394cA3985598Fff63E9551e81DD4D8602eA68B",
			)
			if err == nil {
				break
			}
		}
		time.Sleep(time.Second * 2)
		for {
			err := AddPioneerBatch4(
				"0xee72b2ea43e56148d63a0adf23653790ab0e0ab31bc0f8a4326ab9f377ee931a",
				"0xa4394cA3985598Fff63E9551e81DD4D8602eA68B",
				3,
				40, //5000 6
				60, // 1000 3
				1,
				0)
			if err == nil {
				break
			}
		}
	} else if strings.ToLower("0xA2C100a4C098D3D05Ff8c0AE2CE4Ea93eCe550A8") == strings.ToLower(pioneer) {
		for {
			err := RemovePioneer(
				"0x730ad915d608f40a7cea880f5abc7fe6bb253f903667be02ec00dd3028f7d24a",
				"0xA2C100a4C098D3D05Ff8c0AE2CE4Ea93eCe550A8",
			)
			if err == nil {
				break
			}
		}
		time.Sleep(time.Second * 2)
		for {
			err := AddPioneerBatch4(
				"0x730ad915d608f40a7cea880f5abc7fe6bb253f903667be02ec00dd3028f7d24a",
				"0xA2C100a4C098D3D05Ff8c0AE2CE4Ea93eCe550A8",
				3,
				40, //5000 6
				60, // 1000 3
				2,
				0)
			if err == nil {
				break
			}
		}
	} else if strings.ToLower("0xbF0D360879Cc0AED8479058fbf8f86F510512188") == strings.ToLower(pioneer) {
		for {
			err := RemovePioneer(
				"0x491f0fa71e9db46360a7ac880539f923cda458800f22fb344e10e98491f1f39a",
				"0xbF0D360879Cc0AED8479058fbf8f86F510512188",
			)
			if err == nil {
				break
			}
		}
		time.Sleep(time.Second * 2)
		for {
			err := AddPioneerBatch4(
				"0x491f0fa71e9db46360a7ac880539f923cda458800f22fb344e10e98491f1f39a",
				"0xbF0D360879Cc0AED8479058fbf8f86F510512188",
				3,
				40, //5000 6
				60, // 1000 3
				3,
				0)
			if err == nil {
				break
			}
		}
	} else if strings.ToLower("0xC250bFC206Ba2d9b761Eb51537ef8D665e0a54d7") == strings.ToLower(pioneer) {
		for {
			err := RemovePioneer(
				"0xc50db5c60dc66068637ffc825954152fbce8b73a290285f252697c31b005a382",
				"0xC250bFC206Ba2d9b761Eb51537ef8D665e0a54d7",
			)
			if err == nil {
				break
			}
		}
		time.Sleep(time.Second * 2)
		for {
			err := AddPioneerBatch4(
				"0xc50db5c60dc66068637ffc825954152fbce8b73a290285f252697c31b005a382",
				"0xC250bFC206Ba2d9b761Eb51537ef8D665e0a54d7",
				1,
				40, //5000 6
				60, // 1000 3
				3,
				0)
			if err == nil {
				break
			}
		}
	}

}
