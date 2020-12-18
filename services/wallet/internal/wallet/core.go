package wallet

import (
	"errors"
	"github.com/jackc/pgx"
	"gitlab.com/mmdaz/arvan-challenge/services/wallet/pkg/repositories"
	"strconv"
)

type Core struct {
	walletRepo repositories.WalletRepository
}

func NewCore(walletRepo repositories.WalletRepository) *Core {
	return &Core{walletRepo: walletRepo}
}

func (c *Core) Increase(phoneNumber string, amount int) error {

	phone, err := strconv.Atoi(phoneNumber)
	if err != nil {
		return err
	}

	if amount <= 0 {
		return errors.New("amount invalid")
	}

	wallet, err := c.walletRepo.GetWalletByPhone(phone)

	switch err {
	case pgx.ErrNoRows:
		err = c.walletRepo.Create(phone, amount)
		if err != nil {
			return err
		}
	case nil:
		newAmount := wallet.Cache + amount
		err = c.walletRepo.UpdateCache(phone, newAmount)
		if err != nil {
			return err
		}
	default:
		return err
	}
	return nil
}
