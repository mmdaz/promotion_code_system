package api

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"gitlab.com/mmdaz/arvan-challenge/services/promotion_code/internal/promotion_code"
	"net/http"
)

type HttpHandler struct {
	promotionCodeCore *promotion_code.Core
}

func NewHttpHandler(promotionCodeCore *promotion_code.Core) *HttpHandler {
	return &HttpHandler{promotionCodeCore: promotionCodeCore}
}

func (h HttpHandler) ApplyCode(ctx *gin.Context) {
	phoneNumber := ctx.Request.Header.Get("PhoneNumber")
	err := h.promotionCodeCore.ApplyPromotionCode(phoneNumber)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
	}
}
