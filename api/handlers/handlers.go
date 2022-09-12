package handlers

import (
	"mini-wallet/api/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *controller) InitHandler(echoCtx echo.Context) error {
	var user model.User
	err := echoCtx.Bind(&user)
	if err != nil || user.CustomerID == "" {
		msg := map[string]string{
			"customer_xid": "customer_xid must exist",
		}
		apitOutput := buildAPIOutput(model.APIFail, msg)
		return echoCtx.JSON(http.StatusBadRequest, apitOutput)
	}

	userToken, err := c.Service.CreateAccount(user)
	if err != nil && err.Error() == "account already exists" {
		apiOutput := buildAPIOutput(model.APIFail, err.Error())
		return echoCtx.JSON(http.StatusBadRequest, apiOutput)
	} else if err != nil && err.Error() != "account already exists" {
		apiOutput := buildAPIOutput(model.APIError, "error when doing execution to database")
		return echoCtx.JSON(http.StatusInternalServerError, apiOutput)
	}

	apiOutput := buildAPIOutput(model.APISuccess, userToken)
	return echoCtx.JSON(http.StatusCreated, apiOutput)
}

func (c *controller) EnableHandler(echoCtx echo.Context) error {
	msg, token, err := c.checkAuth(echoCtx)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIFail, msg)
		return echoCtx.JSON(http.StatusUnauthorized, apiOutput)
	}

	account, err := c.Service.FindAccountByToken(token)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIError, err.Error())
		return echoCtx.JSON(http.StatusInternalServerError, apiOutput)
	}

	walletData, err := c.Service.FindWalletByAccountID(account.ID)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIError, err.Error())
		return echoCtx.JSON(http.StatusInternalServerError, apiOutput)
	}

	if walletData.Status == "enabled" {
		msg := map[string]string{
			"Enable": "wallet already enabled",
		}
		apiOutput := buildAPIOutput(model.APIFail, msg)
		return echoCtx.JSON(http.StatusBadRequest, apiOutput)
	}

	wallet, err := c.Service.EnableWallet(account.ID, account.CustomerID)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIError, err.Error())
		return echoCtx.JSON(http.StatusInternalServerError, apiOutput)
	}

	dataWallet := map[string]model.Wallets{
		"wallet": *wallet,
	}

	apiOutput := buildAPIOutput(model.APISuccess, dataWallet)
	return echoCtx.JSON(http.StatusCreated, apiOutput)
}

func (c *controller) ViewWalletHandler(echoCtx echo.Context) error {
	msg, token, err := c.checkAuth(echoCtx)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIFail, msg)
		return echoCtx.JSON(http.StatusUnauthorized, apiOutput)
	}

	account, err := c.Service.FindAccountByToken(token)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIError, err.Error())
		return echoCtx.JSON(http.StatusInternalServerError, apiOutput)
	}

	walletData, err := c.Service.FindWalletByAccountID(account.ID)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIError, err.Error())
		return echoCtx.JSON(http.StatusInternalServerError, apiOutput)
	}

	if walletData.Status != "enabled" {
		msg := map[string]string{
			"Enable": "wallet is not enabled yet",
		}
		apiOutput := buildAPIOutput(model.APIFail, msg)
		return echoCtx.JSON(http.StatusForbidden, apiOutput)
	}

	dataWallet := map[string]model.Wallets{
		"wallet": *walletData,
	}

	apiOutput := buildAPIOutput(model.APISuccess, dataWallet)

	return echoCtx.JSON(http.StatusOK, apiOutput)
}

func (c *controller) DepositWalletHandler(echoCtx echo.Context) error {

	msg, token, err := c.checkAuth(echoCtx)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIFail, msg)
		return echoCtx.JSON(http.StatusUnauthorized, apiOutput)
	}

	account, err := c.Service.FindAccountByToken(token)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIError, err.Error())
		return echoCtx.JSON(http.StatusInternalServerError, apiOutput)
	}

	walletData, err := c.Service.FindWalletByAccountID(account.ID)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIError, err.Error())
		return echoCtx.JSON(http.StatusInternalServerError, apiOutput)
	}

	if walletData.Status != "enabled" {
		msg := map[string]string{
			"Enable": "wallet is not enabled yet",
		}
		apiOutput := buildAPIOutput(model.APIFail, msg)
		return echoCtx.JSON(http.StatusForbidden, apiOutput)
	}

	var transaction model.Transaction

	err = echoCtx.Bind(&transaction)
	if err != nil && transaction.Amount < 1 {
		msg := map[string]string{
			"amount": "amount must be greater than zero",
		}
		apitOutput := buildAPIOutput(model.APIFail, msg)
		return echoCtx.JSON(http.StatusBadRequest, apitOutput)
	}

	resp, err := c.Service.DepositWalletByWalletID(walletData.ID, account.ID, transaction.Amount, walletData.Balance, account.CustomerID, transaction.ReferenceID)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIError, err.Error())
		return echoCtx.JSON(http.StatusInternalServerError, apiOutput)
	}

	dataDeposit := map[string]model.TransactionResponse{
		"deposit": *resp,
	}

	apiOutput := buildAPIOutput(model.APISuccess, dataDeposit)
	return echoCtx.JSON(http.StatusOK, apiOutput)
}

func (c *controller) WithdrawWalletHandler(echoCtx echo.Context) error {

	msg, token, err := c.checkAuth(echoCtx)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIFail, msg)
		return echoCtx.JSON(http.StatusUnauthorized, apiOutput)
	}

	account, err := c.Service.FindAccountByToken(token)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIError, err.Error())
		return echoCtx.JSON(http.StatusInternalServerError, apiOutput)
	}

	walletData, err := c.Service.FindWalletByAccountID(account.ID)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIError, err.Error())
		return echoCtx.JSON(http.StatusInternalServerError, apiOutput)
	}

	if walletData.Status != "enabled" {
		msg := map[string]string{
			"Enable": "wallet is not enabled yet",
		}
		apiOutput := buildAPIOutput(model.APIFail, msg)
		return echoCtx.JSON(http.StatusForbidden, apiOutput)
	}

	var transaction model.Transaction

	err = echoCtx.Bind(&transaction)
	if err != nil {
		msg := map[string]string{
			"request": "bad request, please check again",
		}
		apitOutput := buildAPIOutput(model.APIFail, msg)
		return echoCtx.JSON(http.StatusBadRequest, apitOutput)
	}

	if transaction.Amount < 1 {
		msg := map[string]string{
			"amount": "amount should be greater than zero",
		}
		apitOutput := buildAPIOutput(model.APIFail, msg)
		return echoCtx.JSON(http.StatusBadRequest, apitOutput)
	}

	if transaction.Amount > walletData.Balance {
		msg := map[string]string{
			"balance": "insufficient balance",
		}
		apitOutput := buildAPIOutput(model.APIFail, msg)
		return echoCtx.JSON(http.StatusBadRequest, apitOutput)
	}

	resp, err := c.Service.WithdrawWalletByWalletID(walletData.ID, account.ID, transaction.Amount, walletData.Balance, account.CustomerID, transaction.ReferenceID)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIError, err.Error())
		return echoCtx.JSON(http.StatusInternalServerError, apiOutput)
	}

	dataDeposit := map[string]model.TransactionResponse{
		"withdraw": *resp,
	}

	apiOutput := buildAPIOutput(model.APISuccess, dataDeposit)
	return echoCtx.JSON(http.StatusOK, apiOutput)
}

func (c *controller) DisableHandler(echoCtx echo.Context) error {
	msg, token, err := c.checkAuth(echoCtx)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIFail, msg)
		return echoCtx.JSON(http.StatusUnauthorized, apiOutput)
	}

	account, err := c.Service.FindAccountByToken(token)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIError, err.Error())
		return echoCtx.JSON(http.StatusInternalServerError, apiOutput)
	}

	walletData, err := c.Service.FindWalletByAccountID(account.ID)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIError, err.Error())
		return echoCtx.JSON(http.StatusInternalServerError, apiOutput)
	}

	if walletData.Status == "disabled" {
		msg := map[string]string{
			"Disabled": "wallet already disabled",
		}
		apiOutput := buildAPIOutput(model.APIFail, msg)
		return echoCtx.JSON(http.StatusBadRequest, apiOutput)
	}

	var isDisable model.IsDisable

	err = echoCtx.Bind(&isDisable)
	if err != nil {
		msg := map[string]string{
			"is_disabled": "invalid form value",
		}
		apitOutput := buildAPIOutput(model.APIFail, msg)
		return echoCtx.JSON(http.StatusBadRequest, apitOutput)
	}

	wallet, err := c.Service.DisableWallet(account.ID, account.CustomerID)
	if err != nil {
		apiOutput := buildAPIOutput(model.APIError, err.Error())
		return echoCtx.JSON(http.StatusInternalServerError, apiOutput)
	}

	dataWallet := map[string]model.Wallets{
		"wallet": *wallet,
	}

	apiOutput := buildAPIOutput(model.APISuccess, dataWallet)
	return echoCtx.JSON(http.StatusCreated, apiOutput)
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
