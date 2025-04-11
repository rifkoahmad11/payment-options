package main

import (
	httpDelivery "payment-options/internal/http"
	"payment-options/internal/repository"
	"payment-options/internal/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	repo := repository.NewPaymentRepo()
	uc := usecase.NewPaymentUsecase(repo)
	httpDelivery.NewPaymentHandler(e, uc)

	e.Logger.Fatal(e.Start(":8081"))
}
