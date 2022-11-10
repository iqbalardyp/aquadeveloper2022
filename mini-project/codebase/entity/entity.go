package entity

type Category struct {
	ID   int64  `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Type string `json:"type"`
}

type Product struct {
	ID         int64    `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Name       string   `json:"name"`
	Price      int64    `json:"price"`
	CategoryID int64    `json:"categoryID" gorm:"type:bigint"`
	Category   Category `gorm:"foreignKey:CategoryID"`
	Desc       string   `json:"desc" gorm:"type:text;default:Lorem ipsum"`
	Figure     []byte   `json:"figure"`
}

type Cart struct {
	ID        int64   `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Qty       int64   `json:"qty" gorm:"default:1"`
	ProductID int64   `json:"productID" gorm:"type:bigint"`
	Product   Product `gorm:"foreignKey:ProductID"`
}

type Payment struct {
	ID          int64  `json:"id" gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Timestamp   int64  `json:"timestamp" gorm:"autoCreateTime"`
	PaymentSlip []byte `json:"paymentSlip"`
}
