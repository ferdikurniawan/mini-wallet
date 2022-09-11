package repository

import (
	"context"
	"database/sql"
	"time"
)

func EnableAccountByUserID(db *sql.DB, accountID int) error {
	//TODO use transaction
	ctx := context.Background()

	//start transaction mode
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	now := time.Now()

	//update the account status
	query := `UPDATE mw_accounts 
				SET status = 'enabled', updated_at = $2 
				WHERE id = $1`
	_, err = tx.ExecContext(ctx, query, accountID, now)
	if err != nil {
		tx.Rollback()
		return err
	}

	//insert account_activity_logs
	query = `INSERT INTO mw_activity_logs (account_id, activity, activity_time)
				VALUES ($1, $2, $3)`
	_, err = tx.ExecContext(ctx, query, accountID, "enable", now)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
