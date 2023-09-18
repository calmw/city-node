package validate

import (
	"city-node-server/api/common/statecode"
	"city-node-server/api/models/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
)

type Pioneer struct{}

func NewPioneer() *Pioneer {
	return &Pioneer{}
}

func (v *Pioneer) GetRechargeWeightByPioneerAddress(c *gin.Context, req *request.GetRechargeWeightByPioneerAddress) int {

	err := c.BindQuery(req)
	fmt.Println(err)
	if err == io.EOF {
		return statecode.ParameterEmptyErr
	} else if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			if e.Field() == "pioneer" && e.Tag() == "required" {
				return statecode.PioneerEmpty
			}
		}
		return statecode.CommonErrServerErr
	}

	return statecode.CommonSuccess
}
