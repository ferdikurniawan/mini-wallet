package handlers

import (
	"errors"
	"mini-wallet/api/model"
	"strings"
)

func buildAPIOutput(status string, data interface{}) model.APIOutput {
	return model.APIOutput{
		Data:   data,
		Status: status,
	}
}

func getToken(auth string) (string, error) {
	authArr := strings.Split(auth, " ")
	if len(authArr) < 2 {
		return "", errors.New("invalid token")
	}

	if authArr[0] != "Token" {
		return "", errors.New("invalid token prefix")
	}

	return authArr[1], nil
}
