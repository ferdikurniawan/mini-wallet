package handlers

import (
	"mini-wallet/api/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *controller) InitHandler(echoCtx echo.Context) error {
	var user model.User
	err := echoCtx.Bind(&user)
	if err != nil {
		//TODO use jsend format
		return echoCtx.JSON(http.StatusBadRequest, "bad request")
	}

	userToken, err := c.Service.CreateAccount(user)

	apiOutput := buildAPIOutput(model.APISuccess, userToken)
	return echoCtx.JSON(http.StatusCreated, apiOutput)
}

func (c *controller) EnableHandler(echoCtx echo.Context) error {
	reqHeader := echoCtx.Request().Header
	reqToken := reqHeader.Get("Authorization")
	_, err := getToken(reqToken)
	if err != nil {
		msg := map[string]string{
			"Token": err.Error(),
		}
		apiOutput := buildAPIOutput(model.APIFail, msg)
		return echoCtx.JSON(http.StatusUnauthorized, apiOutput)
	}
	return nil
}
