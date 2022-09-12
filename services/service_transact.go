package services

import (
	"mini-wallet/api/model"
	"mini-wallet/repository"

	"github.com/google/uuid"
)

func (s *service) DepositWalletByWalletID(walletID, customerID int, amount, balance int64, customerXID, referenceID string) (*model.TransactionResponse, error) {
	transIDUUID := uuid.New()
	transID := transIDUUID.String()

	res, err := repository.DepositWallet(s.DB, referenceID, customerXID, transID, walletID, customerID, amount, balance)
	if err != nil {
		s.Logger.Errorf("[DepositWalletByWalletID] error deposit to wallet: %s", err.Error())
		return nil, err
	}

	transResponse := model.TransactionResponse{
		TransactionID: transID,
		DepositedBy:   customerXID,
		Status:        res.Status,
		DepositedAt:   res.DepositedAt,
		Amount:        res.Amount,
		ReferenceID:   referenceID,
	}

	return &transResponse, nil
}

func (s *service) WithdrawWalletByWalletID(walletID, customerID int, amount, balance int64, customerXID, referenceID string) (*model.TransactionResponse, error) {
	transIDUUID := uuid.New()
	transID := transIDUUID.String()

	res, err := repository.WithdrawWallet(s.DB, referenceID, customerXID, transID, walletID, customerID, amount, balance)
	if err != nil {
		s.Logger.Errorf("[WithdrawWalletByWalletID] error withdraw to wallet: %s", err.Error())
		return nil, err
	}

	transResponse := model.TransactionResponse{
		TransactionID: transID,
		WithdrawnBy:   customerXID,
		Status:        res.Status,
		WithdrawnAt:   res.WithdrawAt,
		Amount:        res.Amount,
		ReferenceID:   referenceID,
	}

	return &transResponse, nil
}
