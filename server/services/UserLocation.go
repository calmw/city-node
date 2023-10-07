package services

import (
	"city-node-server/blockchain"
	"city-node-server/db"
	"city-node-server/models"
	"city-node-server/utils"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

func InitUserLocation() {

	// 管理员设置城市合约地址
	//blockchain.AdminSetCityAddressUserLocation()
	blockchain.UserLocationAddAdmin()
	// 管理员设置获取用户质押量合约地址
	//blockchain.AdminSetPledgeStakeAddressUserLocation()
	//blockchain.SetUserLocationTest("0xaa3f3f93c743005d497d2a21cf5ac5132960e42496d9b45437722351f8524496", "05C0TS6p86w=", "0x77bd41fdE654FE0054b771Ec6985dC9d5247BAfe")
	//blockchain.SetUserLocationTest("0xb3b348b5d62537a0d87eb58f09045cd2a2d2631fb72a5c58d912bf8f751f1ede", "7gZ4qrL761s=", "0x17dC6411D638672A073f23267C4735ca877AA623")
	//blockchain.SetUserLocationTest("0xa3c9a043a733fc013604651b53591dec8cc6a09c7ebebb413fa826d0a3388cfe", "SLFWsm931/M=", "0xef3E99B0A86284e57c73d096f20c3B983cd4c18e")
	//blockchain.SetUserLocationTest("0x895e139e0a4c629694d8c590f4c0f05ada11904b7997cf30756d69022c64301c", "V+gXgb0En48=", "0xE137fF4FCdDA90C3665562F52491B511155e19FF")
	//blockchain.SetUserLocationTest("0x2cc69c7ca78ccbc134cbd3eef57f3c882263ecadf071a0f7b14fc8cc4cc1dc44", "owonZp4vrhQ=", "0xF22fc6b430C265Fc6eC39d434897cE4E0bC2fD21")
	//blockchain.SetUserLocationTest("0x5d2e8bf23f54504a2f6afbe8abe5eafc331b7da7c166287da97c40935c2a49b2", "HterW8zfFlo=", "0x2E418BF2d36eE6464E7Dcf08fBb994C78971F50A")
	//blockchain.SetUserLocationTest("0x59eaad55ef447a968257797d75e289157b780ef63027267b85c07b1309b6edee", "3sMx/LrnTOE=", "0x7295f5a83337163aA06c70dcdd51903aCc5E27fa")
	//blockchain.SetUserLocationTest("0x80af9895d0010ac6650f3048ff3d21fcbf3e2583faf72a01c57d377981f91181", "Yl8LVnAELT8=", "0xFc542654708477B585a8C219397C34f921Ed0089")
	//blockchain.SetUserLocationTest("0xdf94a1319cc8f5fb67fdfe859ef61d7f883589f463cbd0ef14a95f4bcbf03159", "8EfpiWTnZVY=", "0x2Fde6c8b270dEeDb011F54eC8abDf08CA8cA5dfF")
	//blockchain.SetUserLocationTest("0x890d4253b0eaef52eb154158ccbc49aafe7d53db49eb15215f8874024a9fbd91", "rGlTv2yBVxk=", "0xcE492Ae90D030DB5C833296F1ABdE6c088013E50")
	//blockchain.SetUserLocationTest("0x91fa5e0e901e28b581707c10829194cd85d1d281e4d7dbf0dfd0d78463cd2e87", "pYnBembRMi+XVCx630fpWw==", "0xF032A7e9556f3a9688278e995c1Ac9CC7A676eD5")
	//blockchain.SetUserLocationTest("0x95fecb63031dfc09623ae52f647df2e749533e3acda9a0ea73f34c29cd07d667", "hJfPa76M9sE=", "0x7763aE1D565Dcc422b7Bc07D5567281338a92A25")
	//blockchain.SetUserLocationTest("0x3d77372f2d82ba9f5d0b972086def85527c58999c521f33bb3047ccac794d06a", "x4Kmjs9f7w4=", "0x1763271960D827D1DE4654D9E754021A0f5f94A5")
	//blockchain.SetUserLocationTest("0x4209dce01c4f06b1b93a35e021ec2efea127ef4ae29a82be536a376504bef0cc", "XxOqpfI6Eo4=", "0x06EfCb450059517f5b34a7cC4eFD80298825124E")
	//blockchain.SetUserLocationTest("0xbb2d35b76dc3358825b55db25f3696363681ccd940f5d39cc6b1767b8e069404", "19Oc48iM8WQ=", "0x9a04896807Aa61458cB3f996F811022086C61698")
	//blockchain.SetUserLocationTest("0x0a23fb256dc6340f9287a2edf424eb5cb92f37f39a8f58cc6c1a9fa349d51b34", "Dk4xzONyA9T2EdPpiSjVqw==", "0xC31Bd725D0C15c085a0860Be0c054F60B22c937F")
}

func UserCityChange() {
	var userLocation []models.UserLocation
	db.Mysql.Model(&models.UserLocation{}).Find(&userLocation)
	for _, l := range userLocation {
		codeSlic := strings.Split(l.AreaCode, ",")
		if len(codeSlic[2]) != 6 { // code最后一段是六位数的按国内用户
			//fmt.Println("跳过：", l.User)
			continue
		}
		codeSecondSlic := strings.Split(codeSlic[1], "")
		if codeSecondSlic[0] == "0" { // code第二段开头是0的用户
			err, oldCityId := blockchain.GetChengShiIdByCountyId(l.CountyId)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(l.User, codeSlic[2], "+++++")
			err, cityCode := GetCityCodeByAdCode(codeSlic[2])
			if err != nil {
				fmt.Println(err)
				continue
			}
			encrypt := fmt.Sprintf("%d,%d", 0, cityCode)
			newCityId := strings.ToLower(utils.Keccak256(encrypt))
			//blockchain.SetCityMapping(l.CountyId, newCityId, l.LocationEncrypt)
			InsertUserCItyIdChangeLog(l.User, l.CountyId, oldCityId, newCityId)

		}
	}
}

func GetCityCodeByAdCode(adCode string) (error, uint) {
	var areaCode models.AreaCode
	whereCondition := fmt.Sprintf("ad_code=%s", adCode)
	err := db.Mysql.Model(&models.AreaCode{}).Where(whereCondition).First(&areaCode).Error
	if err != nil {
		// 查询明文地址
		uri := fmt.Sprintf("https://wallet-api-v2.intowallet.io/api/v1/city_node/geographic_info?city_code=%s&ad_code=%s", "0", adCode)
		err, location := utils.HttpGet(uri)
		if err != nil {
			return err, 0
		}
		var locationInfo blockchain.LocationInfo
		err = json.Unmarshal(location, &locationInfo)
		if err != nil {
			return err, 0
		}
		fmt.Println(err, uri, locationInfo)

		return err, 0
	}
	return nil, areaCode.CityCode
}

func InsertUserCItyIdChangeLog(userAddress, countyId, cityId, newCityId string) error {
	var userCityIdChangeLog models.UserCityIdChangeLog
	whereCondition := fmt.Sprintf("user='%s'", strings.ToLower(userAddress))
	err := db.Mysql.Model(&models.UserCityIdChangeLog{}).Where(whereCondition).First(&userCityIdChangeLog).Error
	if err == gorm.ErrRecordNotFound {

		db.Mysql.Model(&models.UserCityIdChangeLog{}).Create(&models.UserCityIdChangeLog{
			User:        strings.ToLower(userAddress),
			CountyId:    strings.ToLower(countyId),
			CityId:      strings.ToLower(cityId),
			PioneerCity: strings.ToLower(newCityId),
		})
	}
	return nil
}
