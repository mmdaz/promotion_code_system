package main

import (
	"gitlab.com/mmdaz/arvan-challenge/pkg"
	"gitlab.com/mmdaz/arvan-challenge/pkg/http"
	"gitlab.com/mmdaz/arvan-challenge/pkg/redis"
	"gitlab.com/mmdaz/arvan-challenge/services/promotion_code/api"
	"gitlab.com/mmdaz/arvan-challenge/services/promotion_code/internal/promotion_code"
	http2 "net/http"
)

func main() {
	config := pkg.NewConfig("promotion_code", "")

	//postgresCli := postgres.NewPGXPostgres(postgres.Option{
	//	Host: config.Postgres.Host,
	//	Port: config.Postgres.Port,
	//	User: config.Postgres.User,
	//	Pass: config.Postgres.Pass,
	//	Db:   config.Postgres.DB,
	//}, config.Postgres.ConnectionsCount)

	redisCli := redis.NewRedisWithOption(config.Redis.Enable, redis.Option{
		Host:       config.Redis.Host,
		Port:       config.Redis.Port,
		PoolSize:   config.Redis.PoolSize,
		DB:         config.Redis.DB,
		Pass:       config.Redis.Pass,
		MaxRetries: config.Redis.MaxRetries,
	})

	promotionCodeCore := promotion_code.NewCore(redisCli)

	httpServer := http.NewHTTPServer()
	httpHandler := api.NewHttpHandler(promotionCodeCore)
	httpServer.Start(config.HttpServer.Address)
	httpServer.AddRoutes(http.Route{
		Method:       http2.MethodPost,
		Path:         "/applyCode",
		IsAuthorized: false,
		Function:     httpHandler.ApplyCode,
	})

}
