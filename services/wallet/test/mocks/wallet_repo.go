package mocks

import (
	"github.com/stretchr/testify/mock"
	"gitlab.com/mmdaz/arvan-challenge/services/wallet/pkg/models"
)

type WalletRepoMock struct {
	mock.Mock
	CreateFunc           func(phoneNumber int, cache int) error
	UpdateCacheFunc      func(phoneNumber int, amount int) error
	GetWalletByPhoneFunc func(phoneNumber int) (*models.Wallet, error)
}

func (w *WalletRepoMock) Create(phoneNumber int, cache int) error {
	return w.CreateFunc(phoneNumber, cache)
}

func (w *WalletRepoMock) UpdateCache(phoneNumber int, amount int) error {
	return w.UpdateCacheFunc(phoneNumber, amount)
}

func (w *WalletRepoMock) GetWalletByPhone(phoneNumber int) (*models.Wallet, error) {
	return w.GetWalletByPhoneFunc(phoneNumber)
}

func NewWalletRepoMock() *WalletRepoMock {
	return &WalletRepoMock{}
}
