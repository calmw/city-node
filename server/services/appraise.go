package services

import "city-node-server/pkg/blockchain"

func InitAppraise() {
	appraise := blockchain.NewAppraise()
	// AdminSetStar 管理员设置Star合约地址
	appraise.AdminSetStar()
	// AdminSetCity 管理员设置城市合约
	appraise.AdminSetCity()
	// AdminSetPioneer 管理员设置先锋合约
	appraise.AdminSetPioneer()
	// AddAdmin 添加管理权限
	appraise.AddAdmin()
	// AdminSetPioneerBatch 管理员设置先锋批次
	appraise.AdminSetPioneerBatch("", 3)

}
