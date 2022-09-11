package model

const (
	APISuccess = "success"
	APIFail    = "fail"
	APIError   = "error"
)

type APIOutput struct {
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
}

type User struct {
	CustomerID string `form:"customer_xid"`
}

type UserToken struct {
	Token string `json:"token"`
}
