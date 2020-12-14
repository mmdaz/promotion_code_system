package api

import (
	"github.com/gin-gonic/gin"
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
	err := h.promotionCodeCore.ApplyPromotionCode()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "")
	}
}
