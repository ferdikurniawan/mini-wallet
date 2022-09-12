package main

import (
	"mini-wallet/api/handlers"

	"github.com/labstack/echo/v4"
)

const (
	APIPrefix              = "/api/v1"
	InitEndpoint           = APIPrefix + "/init"
	WalletEndpoint         = APIPrefix + "/wallet"
	DepositWalletEndpoint  = APIPrefix + "/wallet/deposits"
	WithdrawWalletEndpoint = APIPrefix + "/wallet/withdrawals"
)

func router(webServer *echo.Echo, ctrl handlers.Controller) {
	webServer.POST(InitEndpoint, ctrl.InitHandler)
	webServer.POST(WalletEndpoint, ctrl.EnableHandler)
	webServer.GET(WalletEndpoint, ctrl.ViewWalletHandler)
	webServer.PATCH(WalletEndpoint, ctrl.DisableHandler)
	webServer.POST(DepositWalletEndpoint, ctrl.DepositWalletHandler)
	webServer.POST(WithdrawWalletEndpoint, ctrl.WithdrawWalletHandler)
}
