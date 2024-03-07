package controllers

import (
	"city-node-server/api/common/statecode"
	"city-node-server/api/models/response"
	"github.com/gin-gonic/gin"
)

type TestController struct{}

func (c *TestController) Account(ctx *gin.Context) {
	res := response.Gin{Res: ctx}

	res.Response(ctx, statecode.CommonSuccess, "0x17E56C5f4E271a2Ce0920580784C6397e247C9d9")

	return
}
