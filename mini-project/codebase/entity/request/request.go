package request

type CreateCategoryRequest struct {
	Type string `json:"type"`
}
type CreateProductDetailRequest struct {
	Name       string `json:"name"`
	Price      int64  `json:"price"`
	CategoryID int64  `json:"categoryID"`
	Desc       string `json:"desc"`
	Figure     []byte `json:"figure"`
}

type CreateCartRequest struct {
	Qty       int64 `json:"qty"`
	ProductID int64 `json:"productID" gorm:"type:bigint"`
}

type CreatePaymentRequest struct {
	CartID      int64  `json:"cartID" gorm:"type:bigint"`
	TotalPrice  int64  `json:"totalPrice"`
	PaymentSlip string `json:"paymentSlip" gorm:"default:payment.png"`
}
