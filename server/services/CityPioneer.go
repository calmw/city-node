package services

import "city-node-server/blockchain"

func InitCityPioneer() {

	// 管理员设置TOX代币地址
	blockchain.AdminSetTOXAddress()
	// 管理员设置城市合约地址
	blockchain.AdminSetCityAddress()
	// 管理员设置增加用户合约余额合约地址
	blockchain.AdminSetMiningAddress()
	//
	//// 设置城市先锋考核标准
	//blockchain.AdminSetAssessmentCriteria(1, 1, 2500)
	//blockchain.AdminSetAssessmentCriteria(1, 2, 5000)
	//blockchain.AdminSetAssessmentCriteria(1, 3, 10000)
	//blockchain.AdminSetAssessmentCriteria(2, 1, 1250)
	//blockchain.AdminSetAssessmentCriteria(2, 2, 2500)
	//blockchain.AdminSetAssessmentCriteria(2, 3, 5000)
	//blockchain.AdminSetAssessmentCriteria(3, 1, 625)
	//blockchain.AdminSetAssessmentCriteria(3, 2, 125)
	//blockchain.AdminSetAssessmentCriteria(3, 3, 2500)
	//
	//// AddCityPioneerAdmin 给城市合约、IntoMining设置用户（增加用户合约余额）添加管理员权限
	blockchain.AddCityPioneerAdmin()

	// 管理员设置开始考核时间,先交保证金，后考核
	//blockchain.AdminSetStartTime()
}
