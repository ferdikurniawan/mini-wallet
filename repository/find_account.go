package repository

import "database/sql"

func FindAccountByToken(db *sql.DB, token string) (*Accounts, error) {
	var account Accounts
	query := `SELECT id, customer_xid, token, created_at, updated_at, salt
				FROM mw_accounts
				WHERE token = $1`
	row := db.QueryRow(query, token)
	err := row.Scan(&account.ID, &account.CustomerID, &account.Token, &account.CreatedAt, &account.UpdatedAt, &account.Salt)
	if err != nil {
		return &account, err
	}

	return &account, nil
}
