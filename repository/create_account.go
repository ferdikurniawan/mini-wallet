package repository

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

func CreateAccount(db *sql.DB, data *Accounts) error {
	query := `INSERT INTO mw_accounts (customer_xid, token, salt)
				VALUES ($1,$2,$3)`
	_, err := db.Exec(query, data.CustomerID, data.Token, data.Salt)

	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				return errors.New("account already exists")
			}
		}
	}

	return err

}
