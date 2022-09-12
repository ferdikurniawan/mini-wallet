package repository

import (
	"database/sql"
	"time"
)

const (
	//redis related
	RedisWalletInfoKey = "wallet_info_%d"
	RedisTimeout       = time.Second * 5 //as mentioned in the study case problem set "The maximum delay for updating the balance is 5 seconds"
)

type Accounts struct {
	ID         int
	CustomerID string
	Token      string
	Salt       string
	CreatedAt  time.Time
	UpdatedAt  sql.NullTime
}

type Wallets struct {
	ID         int
	WalletID   string
	AccountID  int
	CustomerID string
	Status     string
	Balance    int64
	EnabledAt  time.Time
	CreatedAt  time.Time
	UpdatedAt  sql.NullTime
	OwnedBy    string
	DisabledAt time.Time
}

type Deposit struct {
	ID          int
	DepositID   string
	DepositedBy string
	Status      string
	DepositedAt string
	Amount      int64
	ReferenceID string
}

type Withdraw struct {
	ID          int
	WithdrawID  string
	WithdrawnBy string
	Status      string
	WithdrawAt  string
	Amount      int64
	ReferenceID string
}
