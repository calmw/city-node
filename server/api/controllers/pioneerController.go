package controllers

import (
	"city-node-server/api/common/statecode"
	"city-node-server/api/models/request"
	"city-node-server/api/models/response"
	"city-node-server/api/services"
	"city-node-server/api/validate"
	"github.com/gin-gonic/gin"
)

type PioneerController struct {
}

func (c *PioneerController) GetRechargeWeightByPioneerAddress(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.GetRechargeWeightByPioneerAddress{}

	errCode := validate.NewPioneer().GetRechargeWeightByPioneerAddress(ctx, &req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}

	errCode, data := services.NewPioneer().GetRechargeWeightByPioneerAddress(&req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}

	res.Response(ctx, statecode.CommonSuccess, data)
	return
}
