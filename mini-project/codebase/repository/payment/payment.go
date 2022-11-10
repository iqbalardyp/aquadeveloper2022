package payment

import (
	"ecommerce/codebase/entity"

	"gorm.io/gorm"
)

type PaymentRepo struct {
	db *gorm.DB
}

type IPaymentRepo interface {
	Create(payment entity.Payment) (entity.Payment, error)
}

func NewPaymentRepo(db *gorm.DB) *PaymentRepo {
	return &PaymentRepo{db: db}
}

func (p PaymentRepo) Create(payment entity.Payment) (entity.Payment, error) {
	if err := p.db.Debug().Create(&payment).Error; err != nil {
		return entity.Payment{}, err
	}
	return payment, nil
}
