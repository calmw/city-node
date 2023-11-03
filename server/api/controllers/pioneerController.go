package controllers

import (
	"city-node-server/api/common/statecode"
	"city-node-server/api/models/request"
	"city-node-server/api/models/response"
	"city-node-server/api/services"
	"github.com/gin-gonic/gin"
)

type PioneerController struct {
}

//func (c *PioneerController) GetRechargeWeightByPioneerAddress(ctx *gin.Context) {
//	res := response.Gin{Res: ctx}
//	req := request.GetRechargeWeightByPioneerAddress{}
//
//	errCode := validate.NewPioneer().GetRechargeWeightByPioneerAddress(ctx, &req)
//	if errCode != statecode.CommonSuccess {
//		res.Response(ctx, errCode, nil)
//		return
//	}
//
//	errCode, data := services.NewPioneer().GetRechargeWeightByPioneerAddress(&req)
//	if errCode != statecode.CommonSuccess {
//		res.Response(ctx, errCode, nil)
//		return
//	}
//
//	res.Response(ctx, statecode.CommonSuccess, data)
//	return
//}

func (c *PioneerController) RechargeWeight(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.RechargeWeight{}
	ctx.ShouldBindQuery(&req)

	errCode, total, data := services.NewPioneer().RechargeWeight(&req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}

	res.Response(ctx, statecode.CommonSuccess, map[string]interface{}{
		"total": total,
		"list":  data,
	})
	return
}

func (c *PioneerController) RewardWithdraw(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.RewardWithdraw{}
	ctx.ShouldBindQuery(&req)

	errCode, total, data := services.NewPioneer().RewardWithdraw(&req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}

	res.Response(ctx, statecode.CommonSuccess, map[string]interface{}{
		"total": total,
		"list":  data,
	})
	return
}

func (c *PioneerController) Reward(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	var req request.Reward
	ctx.ShouldBindQuery(&req)

	errCode, total, data := services.NewPioneer().Reward(&req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}

	res.Response(ctx, statecode.CommonSuccess, map[string]interface{}{
		"total": total,
		"list":  data,
	})
	return
}
