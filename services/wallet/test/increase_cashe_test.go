package test

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/stretchr/testify/assert"
	"gitlab.com/mmdaz/arvan-challenge/services/wallet/pkg/models"
	"testing"
)

func TestIncreaseCache(t *testing.T) {
	testPhoneNumber := 9147569102
	testCache := 1000
	walletRepo.CreateFunc = func(phoneNumber int, cache int) error {
		assert.Equal(t, phoneNumber, testPhoneNumber)
		assert.Equal(t, cache, testCache)
		return nil
	}

	walletRepo.GetWalletByPhoneFunc = func(phoneNumber int) (*models.Wallet, error) {
		assert.Equal(t, phoneNumber, testPhoneNumber)
		return nil, pgx.ErrNoRows
	}

	walletRepo.UpdateCacheFunc = func(phoneNumber int, amount int) error {
		assert.Equal(t, phoneNumber, testPhoneNumber)
		assert.Equal(t, amount, testCache)
		return nil
	}

	err := walletCore.Increase(fmt.Sprint(testPhoneNumber), testCache)
	assert.Nil(t, err)

}
