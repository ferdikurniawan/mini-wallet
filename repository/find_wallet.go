package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func FindWalletByAccountID(db *sql.DB, redis *redis.Client, accountID int) (*Wallets, error) {
	ctx := context.Background()
	var wallet Wallets

	//check cache first
	redisKey := fmt.Sprintf(RedisWalletInfoKey, accountID)
	resultBytes, err := redis.Get(ctx, redisKey).Bytes()
	if err != nil {
		fmt.Printf("[FindWalletByAccountID] error read from redis: %s", err.Error())
	}

	if len(resultBytes) > 0 {
		err = json.Unmarshal(resultBytes, &wallet)
		if err != nil {
			fmt.Printf("[FindWalletByAccountID] error unmarshall value from redis :%s", err.Error())
		}
		if wallet.ID != 0 {
			return &wallet, nil
		}
	}

	//check DB if cache is not found
	query := `SELECT mww.id, mww.wallet_id, mww.account_id, mww.status, mww.balance, mww.enabled_at, mwa.customer_xid
				FROM mw_wallets mww 
				INNER JOIN mw_accounts mwa
				ON (mww.account_id = mwa.id)
				WHERE account_id = $1`
	row := db.QueryRow(query, accountID)
	err = row.Scan(&wallet.ID, &wallet.WalletID, &wallet.AccountID, &wallet.Status, &wallet.Balance, &wallet.EnabledAt, &wallet.OwnedBy)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	//store to cache
	walletBytes, err := json.Marshal(wallet)
	if err != nil {
		fmt.Printf("error marshall wallet %s", err.Error())
		return &wallet, nil
	}

	err = redis.Set(ctx, redisKey, string(walletBytes), RedisTimeout).Err()
	if err != nil {
		fmt.Printf("error store to redis %s", err.Error())
	}

	return &wallet, nil
}
