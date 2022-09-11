package services

import (
	"database/sql"
	"mini-wallet/api/model"

	"github.com/labstack/echo/v4"
)

type (
	service struct {
		Logger echo.Logger
		DB     *sql.DB
	}

	Service interface {
		CreateAccount(req model.User) (model.UserToken, error)
	}
)

func NewService(logger *echo.Logger, db *sql.DB) Service {
	svc := service{}
	svc.Logger = *logger
	svc.DB = db
	return &svc
}
