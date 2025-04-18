package usecase

import (
	"payment-options/internal/models"
	"payment-options/internal/repository"
)

type paymentUsecase struct {
	repo repository.PaymentRepository
}

func NewPaymentUsecase(r repository.PaymentRepository) PaymentUsecase {
	return &paymentUsecase{repo: r}
}

func (u *paymentUsecase) GetPaymentOptions() (map[string]models.PaymentMethod, error) {
	result := make(map[string]models.PaymentMethod)

	// OVO
	ovo, err := u.repo.CallOVO()
	if err != nil {
		return nil, err
	}
	result["ovo"] = ovo

	// DANA
	dana, err := u.repo.CallDANA()
	if err != nil {
		return nil, err
	}
	result["dana"] = dana

	// GoPay
	gopay, err := u.repo.CallGoPay()
	if err != nil {
		return nil, err
	}
	result["gopay"] = gopay

	// ShopeePay
	shopee, err := u.repo.CallShopee()
	if err != nil {
		return nil, err
	}
	result["shopee"] = shopee

	// OneKlik
	oneklik, err := u.repo.CallOneKlik()
	if err != nil {
		return nil, err
	}
	result["oneklik"] = oneklik

	// BRIDD
	bridd, err := u.repo.CallBRIDD()
	if err != nil {
		return nil, err
	}
	result["bridd"] = bridd

	return result, nil
}
