package promotion_code

import (
	"context"
	"errors"
	redsync "github.com/go-redsync/redsync/v4"
	"gitlab.com/mmdaz/arvan-challenge/pkg"
	"gitlab.com/mmdaz/arvan-challenge/services/promotion_code/pkg/repositories"
)

type Core struct {
	config            *pkg.Config
	redisMutex        *redsync.Mutex
	ctx               context.Context
	promotionCodeRepo repositories.PromotionCodeRepo
}

func NewCore(config *pkg.Config, redis *redsync.Mutex, ctx context.Context, promotionCodeRepo repositories.PromotionCodeRepo) *Core {
	return &Core{
		config:            config,
		redisMutex:        redis,
		ctx:               ctx,
		promotionCodeRepo: promotionCodeRepo,
	}
}

func (c *Core) ApplyPromotionCode(phoneNumber string) error {

	if pkg.ValidatePhoneNumbers(phoneNumber) {
		return errors.New("phone number invalid")
	}

	if err := c.redisMutex.LockContext(c.ctx); err != nil {
		return err
	}

	codeCounts, err := c.promotionCodeRepo.GetAppliedCodeCounts(c.config.PromotionCode.StartTime, c.config.PromotionCode.EndTime)
	if err != nil {
		return err
	}

	if codeCounts > c.config.PromotionCode.MaxCodes {
		return errors.New("codes were finished")
	}

	err = c.promotionCodeRepo.Create(phoneNumber, c.config.PromotionCode.CodeValue)
	if err != nil {
		return err
	}

	// add to wallet

	if _, err := c.redisMutex.UnlockContext(c.ctx); err != nil {
		return err
	}

	return nil
}
