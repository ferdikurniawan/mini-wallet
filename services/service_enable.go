package services

import (
	"mini-wallet/api/model"
	"mini-wallet/repository"

	"github.com/google/uuid"
)

func (s *service) EnableWallet(accountID int, customerID string) (*model.Wallets, error) {

	uuidRandom := uuid.New()
	walletID := uuidRandom.String()

	walletFromDB, err := repository.CreateWallet(s.DB, walletID, accountID)
	if err != nil {
		s.Logger.Errorf("[EnableWallet] error enabling account : %s", err.Error())
		return &model.Wallets{}, err
	}

	wallet := model.Wallets{
		WalletID:  walletFromDB.WalletID,
		OwnedBy:   customerID,
		Status:    walletFromDB.Status,
		EnabledAt: walletFromDB.EnabledAt.Format("2006-01-02T15:04:05-0700"),
		Balance:   walletFromDB.Balance,
	}
	return &wallet, nil
}

func (s *service) DisableWallet(accountID int, customerID string) (*model.Wallets, error) {

	uuidRandom := uuid.New()
	walletID := uuidRandom.String()

	walletFromDB, err := repository.DisableWallet(s.DB, walletID, accountID)
	if err != nil {
		s.Logger.Errorf("[DisableWallet] error disable account : %s", err.Error())
		return &model.Wallets{}, err
	}

	wallet := model.Wallets{
		WalletID:   walletFromDB.WalletID,
		OwnedBy:    customerID,
		Status:     walletFromDB.Status,
		DisabledAt: walletFromDB.DisabledAt.Format("2006-01-02T15:04:05-0700"),
		Balance:    walletFromDB.Balance,
	}
	return &wallet, nil
}
