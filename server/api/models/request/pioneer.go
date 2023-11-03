package request

type GetRechargeWeightByPioneerAddress struct {
	Pioneer string `json:"pioneer"  form:"pioneer" binding:"required"`
}

type RechargeWeight struct {
	Pioneer  string `json:"pioneer" form:"pioneer" binding:"omitempty"`
	User     string `json:"user" form:"user" binding:"omitempty"`
	Start    string `json:"start" form:"start" binding:"omitempty"`
	End      string `json:"end" form:"end" binding:"omitempty"`
	Page     int    `json:"page" form:"page" binding:"omitempty"`
	PageSize int    `json:"page_size" form:"page_size" binding:"omitempty,numeric"`
}

type RewardWithdraw struct {
	Pioneer  string `json:"pioneer" form:"pioneer" binding:"omitempty"`
	Start    string `json:"start" form:"start" binding:"omitempty"`
	End      string `json:"end" form:"end" binding:"omitempty"`
	Page     int    `json:"page" form:"page" binding:"omitempty"`
	PageSize int    `json:"page_size" form:"page_size" binding:"omitempty,numeric"`
}

type Reward struct {
	Pioneer  string `json:"pioneer" form:"pioneer" binding:"omitempty"`
	Start    string `json:"start" form:"start" binding:"omitempty"`
	End      string `json:"end" form:"end" binding:"omitempty"`
	Page     int    `json:"page" form:"page" binding:"omitempty"`
	PageSize int    `json:"page_size" form:"page_size" binding:"omitempty,numeric"`
}
