package services

import (
	"mini-wallet/api/model"
	"mini-wallet/repository"
)

func (s *service) FindAccountByToken(token string) (*model.Accounts, error) {
	account := model.Accounts{}
	accountFromDB, err := repository.FindAccountByToken(s.DB, token)
	if err != nil {
		s.Logger.Errorf("[FindAccountByToken] error finding account: %s", err.Error())
		return &account, err
	}

	account.ID = accountFromDB.ID
	account.CustomerID = accountFromDB.CustomerID
	return &account, nil
}
