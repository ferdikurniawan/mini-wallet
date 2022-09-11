package services

import "mini-wallet/repository"

func (s *service) FindAccountByToken(token string) (int, error) {
	account, err := repository.FindAccountByToken(s.DB, token)
	if err != nil {
		s.Logger.Errorf("[FindAccountByToken] error finding account: %s", err.Error())
		return account.ID, err
	}
	return account.ID, nil
}
