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
	//TODO handle error pas init

	apiOutput := buildAPIOutput(model.APISuccess, userToken)
	return echoCtx.JSON(http.StatusCreated, apiOutput)
}

func (c *controller) EnableHandler(echoCtx echo.Context) error {
	msg, token, err := c.checkAuth(echoCtx)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIFail, msg)
		return echoCtx.JSON(http.StatusUnauthorized, apiOutput)
	}

	accountID, err := c.Service.FindAccountByToken(token)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIError, err.Error())
		return echoCtx.JSON(http.StatusInternalServerError, apiOutput)
	}

	err = c.Service.EnableAccount(accountID)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIError, err.Error())
		return echoCtx.JSON(http.StatusInternalServerError, apiOutput)
	}

	return nil
}

func (c *controller) checkAuth(echoCtx echo.Context) (map[string]string, string, error) {
	reqHeader := echoCtx.Request().Header
	reqToken := reqHeader.Get("Authorization")
	token, err := getToken(reqToken)
	if err != nil {
		msg := map[string]string{
			"Token": err.Error(),
		}
		return msg, token, err
	}
	return map[string]string{}, token, nil
}
