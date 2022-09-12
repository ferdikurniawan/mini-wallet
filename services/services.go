package services

import (
	"database/sql"
	"mini-wallet/api/model"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

type (
	service struct {
		Logger echo.Logger
		DB     *sql.DB
		Redis  *redis.Client
	}

	Service interface {
		CreateAccount(req model.User) (model.UserToken, error)
		FindAccountByToken(token string) (*model.Accounts, error)
		EnableWallet(accountID int, customerID string) (*model.Wallets, error)
		FindWalletByAccountID(accountID int) (*model.Wallets, error)
		DepositWalletByWalletID(walletID, customerID int, amount, balance int64, customerXID, referenceID string) (*model.TransactionResponse, error)
		WithdrawWalletByWalletID(walletID, customerID int, amount, balance int64, customerXID, referenceID string) (*model.TransactionResponse, error)
		DisableWallet(accountID int, customerID string) (*model.Wallets, error)
	}
)

func NewService(logger *echo.Logger, db *sql.DB, redis *redis.Client) Service {
	svc := service{}
	svc.Logger = *logger
	svc.DB = db
	svc.Redis = redis
	return &svc
}
