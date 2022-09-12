package repository

import (
	"database/sql"
	"time"
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
}
