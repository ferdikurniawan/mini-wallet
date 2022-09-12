package handlers

import (
	"mini-wallet/services"

	"github.com/labstack/echo/v4"
)

type (
	controller struct {
		Service services.Service
	}

	Controller interface {
		InitHandler(echoCtx echo.Context) error
		EnableHandler(echoCtx echo.Context) error
		ViewWalletHandler(echoCtx echo.Context) error
		DepositWalletHandler(echoCtx echo.Context) error
		WithdrawWalletHandler(echoCtx echo.Context) error
		DisableHandler(echoCtx echo.Context) error
	}
)

func NewController(service services.Service) Controller {
	ctrl := controller{}
	ctrl.Service = service
	return &ctrl
}
