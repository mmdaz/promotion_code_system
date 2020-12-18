package test

import (
	"github.com/go-redsync/redsync/v4"
	"gitlab.com/mmdaz/arvan-challenge/pkg"
	"gitlab.com/mmdaz/arvan-challenge/pkg/redis"
	"gitlab.com/mmdaz/arvan-challenge/services/promotion_code/internal"
	"gitlab.com/mmdaz/arvan-challenge/services/promotion_code/internal/promotion_code"
	"gitlab.com/mmdaz/arvan-challenge/services/promotion_code/test/mocks"
	"os"
	"testing"
	"time"
)

var (
	promotionCodeCore *promotion_code.Core
	promotionCodeRepo = mocks.NewMockPromotionCodeRepo()
	kafkaMock         = mocks.NewKafkaMock()
	walletMock = mocks.NewWalletMock()
)

func TestMain(m *testing.M) {
	config := pkg.NewConfig("promotion_code", "/home/muhammad/go/src/gitlab.com/mmdaz/arvan-challenge/services/promotion_code/test/config.yml")

	redisCli := redis.NewRedisWithOption(redis.Option{
		Host:     config.Redis.Host,
		Port:     config.Redis.Port,
		PoolSize: config.Redis.PoolSize,
		DB:       config.Redis.DB,
		Pass:     config.Redis.Pass,
	})

	publisher := internal.NewPublisher(kafkaMock, config)
	redisCli.Client.FlushAll()

	rs := redsync.New(redisCli.SyncPool)
	mutex := rs.NewMutex(config.PromotionCode.LockKey, redsync.WithExpiry(time.Hour))

	promotionCodeCore = promotion_code.NewCore(config, mutex, promotionCodeRepo, walletMock, publisher)
	r := m.Run()
	os.Exit(r)
}
