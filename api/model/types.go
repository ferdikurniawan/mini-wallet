package model

const (
	APISuccess = "success"
	APIFail    = "fail"
	APIError   = "error"
)

type APIOutput struct {
	Data    interface{} `json:"data,omitempty"`
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
}

type User struct {
	CustomerID string `form:"customer_xid"`
}

type UserToken struct {
	Token string `json:"token"`
}
