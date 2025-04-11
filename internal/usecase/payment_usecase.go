package usecase

import "payment-options/internal/models"

type PaymentUsecase interface {
	GetPaymentOptions() (map[string]models.PaymentMethod, error)
}
