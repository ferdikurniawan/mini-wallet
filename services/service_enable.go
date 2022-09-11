package services

import "mini-wallet/repository"

func (s *service) EnableAccount(accountID int) error {
	err := repository.EnableAccountByUserID(s.DB, accountID)
	if err != nil {
		s.Logger.Errorf("[EnableAccount] error enabling account : %s", err.Error())
		return err
	}
	return nil
}
