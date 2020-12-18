package main

import (
	"gitlab.com/mmdaz/arvan-challenge/pkg"
	"gitlab.com/mmdaz/arvan-challenge/pkg/http"
	"gitlab.com/mmdaz/arvan-challenge/pkg/postgres"
	"gitlab.com/mmdaz/arvan-challenge/services/wallet/api"
	"gitlab.com/mmdaz/arvan-challenge/services/wallet/internal/wallet"
	"gitlab.com/mmdaz/arvan-challenge/services/wallet/pkg/repositories"
	http2 "net/http"
)

func main() {
	config := pkg.NewConfig("wallet", "/home/muhammad/go/src/gitlab.com/mmdaz/arvan-challenge/services/wallet/config.yml")
	postgresCli := postgres.NewPGXPostgres(postgres.Option{
		Host: config.Postgres.Host,
		Port: config.Postgres.Port,
		User: config.Postgres.User,
		Pass: config.Postgres.Pass,
		Db:   config.Postgres.DB,
	}, config.Postgres.ConnectionsCount)

	walletRepo := repositories.NewWalletRepository(postgresCli)
	walletCore := wallet.NewCore(walletRepo)

	httpServer := http.NewHTTPServer()
	httpHandler := api.NewHttpHandler(walletCore)
	httpServer.AddRoutes(http.Route{
		Method:       http2.MethodPost,
		Path:         "/increaseCash",
		IsAuthorized: false,
		Function:     httpHandler.IncreaseCash,
	})
	httpServer.Start(config.HttpServer.Address)
}
