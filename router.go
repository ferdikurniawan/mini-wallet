package main

import (
	"mini-wallet/api/handlers"

	"github.com/labstack/echo/v4"
)

const (
	APIPrefix      = "/api/v1"
	InitEndpoint   = APIPrefix + "/init"
	EnableEndpoint = APIPrefix + "/enable"
)

func router(webServer *echo.Echo, ctrl handlers.Controller) {
	webServer.GET("/", hello)
	webServer.POST(InitEndpoint, ctrl.InitHandler)
	webServer.POST(EnableEndpoint, ctrl.EnableHandler)
}
