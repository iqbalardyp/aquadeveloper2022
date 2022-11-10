package payment

import (
	"ecommerce/codebase/entity"
	"ecommerce/codebase/repository/payment"

	"github.com/jinzhu/copier"
)

type IPaymentUseCase interface{}

type PaymentUseCase struct {
	paymentRepository payment.IPaymentRepo
}

func NewPaymentUseCase(paymentRepository payment.IPaymentRepo) *PaymentUseCase {
	return &PaymentUseCase{paymentRepository: paymentRepository}
}

func (p PaymentUseCase) CreatePayment(req entity.Payment) (entity.Payment, error) {
	var payment entity.Payment
	copier.Copy(&payment, &req)

	paymentResponse, err := p.paymentRepository.Create(payment)
	if err != nil {
		return entity.Payment{}, err
	}
	return paymentResponse, nil
}
