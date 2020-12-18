package mocks

import (
	"github.com/stretchr/testify/mock"
)

type WalletMock struct {
	mock.Mock
	IncreaseAmountFunc func(phoneNumber int, amount int) error
}

func (w *WalletMock) IncreaseAmount(phoneNumber int, amount int) error {
	return w.IncreaseAmountFunc(phoneNumber, amount)
}

func NewWalletMock() *WalletMock {
	return &WalletMock{}
}
