package handlers

import (
	"errors"
	"fmt"
	"mini-wallet/api/model"
	"strings"
)

func buildAPIOutput(status string, data interface{}) model.APIOutput {
	if status == model.APIError {
		msgString := fmt.Sprintf("%x", data)
		return model.APIOutput{
			Status:  status,
			Message: msgString,
		}
	}
	return model.APIOutput{
		Data:   data,
		Status: status,
	}
}

func buildErrorMessage(key, value string) map[string]string {
	msg := map[string]string{
		key: value,
	}
	return msg
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
