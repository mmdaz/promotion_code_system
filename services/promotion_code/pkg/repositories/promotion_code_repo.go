package repositories

import (
	"gitlab.com/mmdaz/arvan-challenge/pkg/postgres"
	"time"
)

type PromotionCodeRepo interface {
	Create(phoneNumber string, code string) error
	GetAppliedCodeCounts(startTime time.Time, endTime time.Time) (int32, error)
}

type PromotionCodeImp struct {
	db *postgres.PGXDatabase
}

func NewPromotionCodeRepo(db *postgres.PGXDatabase) PromotionCodeRepo {
	return &PromotionCodeImp{db: db}
}

func (p PromotionCodeImp) Create(phoneNumber string, code string) error {
	_, err := p.db.Exec(`INSERT INTO promotion_code (phone_number, code, created_at) values ($1, $2, $3)`, phoneNumber, code, time.Now())
	return err
}

func (p PromotionCodeImp) GetAppliedCodeCounts(startTime time.Time, endTime time.Time) (int32, error) {
	var codeCounts int32
	err := p.db.QueryRow(`SELECT COUNT(*) FROM promotion_code WHERE created_at >= $1 AND created_at <= $2;`, startTime, endTime).Scan(&codeCounts)
	if err != nil {
		return 0, err
	}
	return codeCounts, nil
}
