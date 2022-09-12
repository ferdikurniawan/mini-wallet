package main

import (
	"mini-wallet/api/handlers"
	"mini-wallet/libs"
	"mini-wallet/services"
	"mini-wallet/store"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	config, err := libs.LoadConfig(".")
	if err != nil {
		e.Logger.Fatal("error loading config:", err)
	}

	db := store.InitPsql(config)
	redis := store.InitRedis(config)
	svc := services.NewService(&e.Logger, db, redis)
	ctrl := handlers.NewController(svc)

	router(e, ctrl)

	// Start server
	e.Logger.Fatal(e.Start(":7001"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
