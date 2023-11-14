package controllers

import (
	"city-node-server/api/common/statecode"
	"city-node-server/api/models/request"
	"city-node-server/api/models/response"
	"city-node-server/api/services"
	"github.com/gin-gonic/gin"
)

type MiningController struct{}

// LedgerDetails 合约充值提现团队汇总
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