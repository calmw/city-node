package controllers

import (
	"city-node-server/api/common/statecode"
	"city-node-server/api/models/request"
	"city-node-server/api/models/response"
	"city-node-server/api/services"
	"github.com/gin-gonic/gin"
)

type MiningController struct{}

// LedgerDetails 合约充值提现团队汇总详情
func (c *MiningController) LedgerDetails(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.LedgerDetails{}
	ctx.ShouldBindQuery(&req)

	errCode, result, bscIn, bscOut, matchIn, matchOut := services.NewMining().LedgerDetails(&req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}

	res.Response(ctx, statecode.CommonSuccess, map[string]interface{}{
		"bscIn":    bscIn,
		"bscOut":   bscOut,
		"matchIn":  matchIn,
		"matchOut": matchOut,
		"list":     result,
	})
	return
}

// Ledger 合约充值提现团队汇总
func (c *MiningController) Ledger(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.LedgerDetails{}
	ctx.ShouldBindQuery(&req)

	errCode, result, bscIn, bscOut, matchIn, matchOut := services.NewMining().Ledger(&req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}

	res.Response(ctx, statecode.CommonSuccess, map[string]interface{}{
		"bscIn":    bscIn,
		"bscOut":   bscOut,
		"matchIn":  matchIn,
		"matchOut": matchOut,
		"list":     result,
	})
	return
}

// Ledger 合约充值提现团队汇总
func (c *MiningController) RechargeSum(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.RechargeSum{}
	ctx.ShouldBindQuery(&req)

	errCode, result := services.NewMining().RechargeSum(&req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}

	res.Response(ctx, statecode.CommonSuccess, map[string]interface{}{
		"list": result,
	})
	return
}
