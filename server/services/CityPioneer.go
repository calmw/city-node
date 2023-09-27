package services

import "city-node-server/blockchain"

func InitCityPioneer() {

	// 管理员设置TOX代币地址
	//blockchain.AdminSetTOXAddress()
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
	blockchain.AddCityPioneerAdmin()

	// 管理员设置开始考核时间,先交保证金，后考核
	//blockchain.AdminSetStartTime(time.Now().Unix())
}
