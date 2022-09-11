package repository

import (
	"database/sql"
	"time"
)

type Accounts struct {
	ID         int
	CustomerID string
	Token      string
	Status     string
	Salt       string
	CreatedAt  time.Time
	UpdatedAt  sql.NullTime
}
