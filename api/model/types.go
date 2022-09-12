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

type Transaction struct {
	Amount      int64  `form:"amount"`
	ReferenceID string `form:"reference_id"`
}

type IsDisable struct {
	IsDisabled string `form:"is_disabled"`
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
	ID         int    `json:"-"`
	WalletID   string `json:"id"`
	OwnedBy    string `json:"owned_by"`
	Status     string `json:"status"`
	EnabledAt  string `json:"enabled_at,omitempty"`
	DisabledAt string `json:"disabled_at,omitempty"`
	Balance    int64  `json:"balance"`
}

type TransactionResponse struct {
	ID            int    `json:"-"`
	TransactionID string `json:"id"`
	DepositedBy   string `json:"deposited_by,omitempty"`
	WithdrawnBy   string `json:"withdrawn_by,omitempty"`
	Status        string `json:"status"`
	DepositedAt   string `json:"deposited_at,omitempty"`
	WithdrawnAt   string `json:"withdraw_at,omitempty"`
	Amount        int64  `json:"amount"`
	ReferenceID   string `json:"reference_id"`
}
