package services

import (
	"mini-wallet/api/model"
	"mini-wallet/libs/token"
	"mini-wallet/repository"
)

func (s *service) CreateAccount(req model.User) (model.UserToken, error) {

	randSalt := token.GenSalt(10)

	withSalt := req.CustomerID + randSalt //append customer id with salt for security purpose
	token := token.GenToken(withSalt)

	accountData := repository.Accounts{
		CustomerID: req.CustomerID,
		Token:      token,
		Status:     "disabled",
		Salt:       randSalt,
	}

	err := repository.CreateAccount(s.DB, &accountData)
	if err != nil {
		s.Logger.Errorf("[CreateAccount] error insert to db %s", err.Error())
		return model.UserToken{}, err
	}

	return model.UserToken{Token: token}, nil

}
