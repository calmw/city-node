package controllers

import (
	"city-node-server/api/common/statecode"
	"city-node-server/api/models/request"
	"city-node-server/api/models/response"
	"city-node-server/api/services"
	"city-node-server/api/validate"
	"github.com/gin-gonic/gin"
)

type UserLocationController struct {
}

func (c *UserLocationController) GetLocationByUserAddress(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.GetLocationByUserAddress{}

	errCode := validate.NewUserLocation().GetLocationByUserAddress(ctx, &req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}

	errCode, data := services.NewUserLocation().GetLocationByUserAddress(&req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}

	res.Response(ctx, statecode.CommonSuccess, data)
	return
}

func (c *UserLocationController) UserLocation(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.Location{}
	ctx.ShouldBindQuery(&req)

	errCode, total, data := services.NewUserLocation().UserLocation(&req)
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

func (c *UserLocationController) GetCityIdByCityName(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := request.CityName{}
	ctx.ShouldBindQuery(&req)

	errCode, data := services.NewUserLocation().CityId(&req)
	if errCode != statecode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}

	res.Response(ctx, statecode.CommonSuccess, data)
	return
}
