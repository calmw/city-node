package request

type GetLocationByUserAddress struct {
	User string `json:"user" binding:"required"`
}

type Location struct {
	User     string `json:"user" form:"user" binding:"omitempty"`
	Page     int    `json:"page" form:"page" binding:"omitempty"`
	PageSize int    `json:"page_size" form:"page_size" binding:"omitempty,numeric"`
}
