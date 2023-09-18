package request

type GetLocationByUserAddress struct {
	User string `json:"user" binding:"required"`
}
