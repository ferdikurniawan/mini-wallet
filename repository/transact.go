package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

func DepositWallet(db *sql.DB, referenceID, customerXID, transactionID string, walletID, customerID int, amount, balance int64) (*Deposit, error) {

	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	//insert into transaction_logs
	query := `INSERT INTO mw_transaction_logs 
				(wallet_id, balance_before, deposit_amt, balance_after, created_at, reference_id, transaction_id, created_by, status)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err = tx.ExecContext(ctx, query, walletID, balance, amount, balance+amount, now, referenceID, transactionID, customerID, "success")
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				tx.Rollback()
				return nil, errors.New("reference id must be unique")
			}
		}
		tx.Rollback()
		return nil, err
	}

	//update wallet balance
	query = `UPDATE mw_wallets
				SET balance = balance + $1, updated_at = $2
				WHERE id = $3`

	_, err = tx.ExecContext(ctx, query, amount, now, walletID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	//commit transaction
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	deposit := Deposit{
		DepositID:   transactionID,
		Status:      "success",
		DepositedBy: customerXID,
		DepositedAt: now.Format("2006-01-02T15:04:05-0700"),
		Amount:      amount,
		ReferenceID: referenceID,
	}

	return &deposit, nil
}

func WithdrawWallet(db *sql.DB, referenceID, customerXID, transactionID string, walletID, customerID int, amount, balance int64) (*Withdraw, error) {

	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	//insert into transaction_logs
	query := `INSERT INTO mw_transaction_logs 
				(wallet_id, balance_before, withdraw_amt, balance_after, created_at, reference_id, transaction_id, created_by, status)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err = tx.ExecContext(ctx, query, walletID, balance, amount, balance-amount, now, referenceID, transactionID, customerID, "success")
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				tx.Rollback()
				return nil, errors.New("reference id must be unique")
			}
		}
		tx.Rollback()
		return nil, err
	}

	//update wallet balance
	query = `UPDATE mw_wallets
				SET balance = balance - $1
				WHERE id = $2`

	_, err = tx.ExecContext(ctx, query, amount, walletID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	//commit transaction
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	withdraw := Withdraw{
		WithdrawID:  transactionID,
		Status:      "success",
		WithdrawnBy: customerXID,
		WithdrawAt:  now.Format("2006-01-02T15:04:05-0700"),
		Amount:      amount,
		ReferenceID: referenceID,
	}

	return &withdraw, nil
}
