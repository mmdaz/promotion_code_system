package repositories

import (
	"gitlab.com/mmdaz/arvan-challenge/pkg/postgres"
	"gitlab.com/mmdaz/arvan-challenge/services/wallet/pkg/models"
	"time"
)

type WalletRepository interface {
	Create(phoneNumber int, cache int) error
	UpdateCache(phoneNumber int, amount int) error
	GetWalletByPhone(phoneNumber int) (*models.Wallet, error)
}

type WalletRepositoryImp struct {
	db *postgres.PGXDatabase
}

func NewWalletRepository(db *postgres.PGXDatabase) WalletRepository {
	return &WalletRepositoryImp{db: db}
}

func (w *WalletRepositoryImp) Create(phoneNumber int, cache int) error {
	_, err := w.db.Exec(`INSERT INTO wallet (phone_number, cache, updated_at) values ($1, $2, $3)`, phoneNumber, cache, time.Now())
	return err
}

func (w *WalletRepositoryImp) GetWalletByPhone(phoneNumber int) (*models.Wallet, error) {
	var wallet models.Wallet
	err := w.db.QueryRow(`SELECT phone_number, cache, updated_at FROM wallet where phone_number=$1`, phoneNumber).Scan(&wallet.PhoneNumber, &wallet.Cache, &wallet.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (w *WalletRepositoryImp) UpdateCache(phoneNumber int, amount int) error {
	_, err := w.db.Exec(`UPDATE wallet SET cache=$1, updated_at=$2 WHERE phone_number=$3`, amount, time.Now(), phoneNumber)
	return err
}
