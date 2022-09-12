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

func (s *service) FindWalletByAccountID(accountID int) (*model.Wallets, error) {
	walletFromDB, err := repository.FindWalletByAccountID(s.DB, s.Redis, accountID)
	if err != nil {
		s.Logger.Errorf("[FindWalletByAccountID] error finding wallet: %s", err.Error())
		return nil, err
	}
	wallet := model.Wallets{
		ID:        walletFromDB.ID,
		WalletID:  walletFromDB.WalletID,
		OwnedBy:   walletFromDB.OwnedBy,
		Status:    walletFromDB.Status,
		EnabledAt: walletFromDB.EnabledAt.Format("2006-01-02T15:04:05-0700"),
		Balance:   walletFromDB.Balance,
	}
	return &wallet, nil
}
