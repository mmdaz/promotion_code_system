package promotion_code

import (
	"context"
	"errors"
	"github.com/go-redsync/redsync/v4"
	"github.com/labstack/gommon/log"
	"gitlab.com/mmdaz/arvan-challenge/pkg"
	"gitlab.com/mmdaz/arvan-challenge/pkg/wallet"
	"gitlab.com/mmdaz/arvan-challenge/services/promotion_code/internal"
	"gitlab.com/mmdaz/arvan-challenge/services/promotion_code/pkg/repositories"
	"sync"
)

type Core struct {
	config            *pkg.Config
	redisMutex        *redsync.Mutex
	ctx               context.Context
	walletCli         *wallet.HttpClient
	promotionCodeRepo repositories.PromotionCodeRepo
	mu                sync.Mutex
	updatePublisher   *internal.Publisher
}

func NewCore(config *pkg.Config, redis *redsync.Mutex, ctx context.Context, promotionCodeRepo repositories.PromotionCodeRepo, walletCli *wallet.HttpClient, updatePublisher *internal.Publisher) *Core {
	return &Core{
		config:            config,
		redisMutex:        redis,
		ctx:               ctx,
		promotionCodeRepo: promotionCodeRepo,
		walletCli:         walletCli,
		updatePublisher:   updatePublisher,
	}
}

func (c *Core) ApplyPromotionCode(phoneNumber string) error {
	phone, err := pkg.NormalizePhone(phoneNumber, "IR")
	if err != nil {
		return err
	}

	err = c.addCode(phone.Number)
	if err != nil {
		return err
	}

	c.updatePublisher.ReceiveUpdate(phone.Number)

	err = c.walletCli.IncreaseAmount(phone.Number, 1000)
	if err != nil {
		return err
	}

	return nil
}

func (c *Core) addCode(phoneNumber int) error {
	if err := c.redisMutex.Lock(); err != nil {
		return err
	}
	//c.mu.Lock()
	defer func() {
		//c.mu.Unlock()
		if _, err := c.redisMutex.Unlock(); err != nil {
			log.Error(err)
		}
	}()

	codeCounts, err := c.promotionCodeRepo.GetAppliedCodeCounts(c.config.PromotionCode.StartTime, c.config.PromotionCode.EndTime)
	if err != nil {
		return err
	}

	if codeCounts >= c.config.PromotionCode.MaxCodes {
		return errors.New("codes were finished")
	}

	err = c.promotionCodeRepo.Create(phoneNumber, c.config.PromotionCode.CodeValue)
	if err != nil {
		return err
	}

	return nil
}
