package main

import (
	"context"
	"github.com/go-redsync/redsync/v4"
	"github.com/labstack/gommon/log"
	"gitlab.com/mmdaz/arvan-challenge/pkg"
	"gitlab.com/mmdaz/arvan-challenge/pkg/http"
	"gitlab.com/mmdaz/arvan-challenge/pkg/postgres"
	"gitlab.com/mmdaz/arvan-challenge/pkg/redis"
	"gitlab.com/mmdaz/arvan-challenge/services/promotion_code/api"
	"gitlab.com/mmdaz/arvan-challenge/services/promotion_code/internal/promotion_code"
	"gitlab.com/mmdaz/arvan-challenge/services/promotion_code/pkg/repositories"
	http2 "net/http"
)

func main() {
	config := pkg.NewConfig("promotion_code", "/home/muhammad/go/src/gitlab.com/mmdaz/arvan-challenge/services/promotion_code/config.yml")

	log.Info(config.PromotionCode.StartTime)
	postgresCli := postgres.NewPGXPostgres(postgres.Option{
		Host: config.Postgres.Host,
		Port: config.Postgres.Port,
		User: config.Postgres.User,
		Pass: config.Postgres.Pass,
		Db:   config.Postgres.DB,
	}, config.Postgres.ConnectionsCount)

	redisCli := redis.NewRedisWithOption(redis.Option{
		Host:       config.Redis.Host,
		Port:       config.Redis.Port,
		PoolSize:   config.Redis.PoolSize,
		DB:         config.Redis.DB,
		Pass:       config.Redis.Pass,
		MaxRetries: config.Redis.MaxRetries,
	})

	promotionCodeRepo := repositories.NewPromotionCodeRepo(postgresCli)

	rs := redsync.New(redisCli.SyncPool)
	mutex := rs.NewMutex(config.PromotionCode.LockKey)
	ctx := context.Background()

	promotionCodeCore := promotion_code.NewCore(config, mutex, ctx, promotionCodeRepo)

	httpServer := http.NewHTTPServer()
	httpHandler := api.NewHttpHandler(promotionCodeCore)
	httpServer.AddRoutes(http.Route{
		Method:       http2.MethodPost,
		Path:         "/applyCode",
		IsAuthorized: false,
		Function:     httpHandler.ApplyCode,
	})
	httpServer.Start(config.HttpServer.Address)
}
