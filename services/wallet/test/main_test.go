package test

import (
	"gitlab.com/mmdaz/arvan-challenge/services/wallet/internal/wallet"
	"gitlab.com/mmdaz/arvan-challenge/services/wallet/test/mocks"
	"os"
	"testing"
)

var (
	walletRepo = mocks.NewWalletRepoMock()
	walletCore *wallet.Core
)

func TestMain(m *testing.M) {
	//config := pkg.NewConfig("wallet", "/home/muhammad/go/src/gitlab.com/mmdaz/arvan-challenge/services/wallet/config.yml")
	walletCore = wallet.NewCore(walletRepo)
	r := m.Run()
	os.Exit(r)
}
