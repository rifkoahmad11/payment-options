package http

import (
	"net/http"
	"payment-options/internal/usecase"

	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	Usecase usecase.PaymentUsecase
}

func NewPaymentHandler(e *echo.Echo, uc usecase.PaymentUsecase) {
	h := &PaymentHandler{Usecase: uc}
	e.GET("/payment/options", h.GetPaymentOptions)
}

func (h *PaymentHandler) GetPaymentOptions(c echo.Context) error {
	data, err := h.Usecase.GetPaymentOptions()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"returnCode": "500",
			"returnDesc": "internal server error",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"returnCode": "200",
		"returnDesc": "success",
		"data":       data,
	})
}
