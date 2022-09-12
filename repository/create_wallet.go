package repository

import (
	"context"
	"database/sql"
	"time"
)

func CreateWallet(db *sql.DB, walletID string, accountID int) (Wallets, error) {
	ctx := context.Background()
	wallet := Wallets{}

	//start transaction mode
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return wallet, err
	}

	now := time.Now()

	//update the account status
	query := `INSERT INTO mw_wallets (wallet_id, account_id, status, balance, enabled_at, created_at) 
				VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = tx.ExecContext(ctx, query, walletID, accountID, "enabled", 0, now, now)
	if err != nil {
		tx.Rollback()
		return wallet, err
	}

	//insert account_activity_logs
	query = `INSERT INTO mw_activity_logs (account_id, activity, activity_time)
				VALUES ($1, $2, $3)`
	_, err = tx.ExecContext(ctx, query, accountID, "enable", now)
	if err != nil {
		tx.Rollback()
		return Wallets{}, err
	}

	err = tx.Commit()
	if err != nil {
		return Wallets{}, err
	}

	wallet.WalletID = walletID
	wallet.AccountID = accountID
	wallet.Status = "enabled"
	wallet.Balance = 0
	wallet.EnabledAt = now

	return wallet, nil
}
