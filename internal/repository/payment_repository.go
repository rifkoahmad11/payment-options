package repository

import "payment-options/internal/models"

type PaymentRepository interface {
	CallOVO() models.PaymentMethod
	CallDANA() models.PaymentMethod
	CallGoPay() models.PaymentMethod
	CallShopee() models.PaymentMethod
	CallOneKlik() models.PaymentMethod
	CallBRIDD() models.PaymentMethod
}
