package api

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"gitlab.com/mmdaz/arvan-challenge/services/wallet/internal/wallet"
	"net/http"
)

type HttpHandler struct {
	walletCore *wallet.Core
}

type IncreaseCacheRequest struct {
	Amount int `json:"amount"`
}

func NewHttpHandler(walletCore *wallet.Core) *HttpHandler {
	return &HttpHandler{walletCore: walletCore}
}

func (h *HttpHandler) IncreaseCash(ctx *gin.Context) {
	phoneNumber := ctx.Request.Header.Get("PhoneNumber")

	var requestBody IncreaseCacheRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	err := h.walletCore.Increase(phoneNumber, requestBody.Amount)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
