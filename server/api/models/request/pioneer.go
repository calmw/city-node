package request

type GetRechargeWeightByPioneerAddress struct {
	Pioneer string `json:"pioneer" binding:"required"`
}
