package repository

import "database/sql"

func CreateAccount(db *sql.DB, data *Accounts) error {
	query := `INSERT INTO mw_accounts (customer_xid, token, status)
				VALUES ($1,$2,$3)`
	_, err := db.Exec(query, data.CustomerID, data.Token, data.Status)
	return err
}