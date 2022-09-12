package model

const (
	APISuccess = "success"
	APIFail    = "fail"
	APIError   = "error"
)

type APIOutput struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

type User struct {
	CustomerID string `form:"customer_xid"`
}

type UserToken struct {
	Token string `json:"token"`
}

type Accounts struct {
	ID         int    `json:"id,omitempty"`
	CustomerID string `json:"customer_xid"`
	Token      string `json:"token,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
}

type Wallets struct {
	ID        int    `json:"-"`
	WalletID  string `json:"id"`
	OwnedBy   string `json:"owned_by"`
	Status    string `json:"status"`
	EnabledAt string `json:"enabled_at"`
	Balance   int64  `json:"balance"`
}
