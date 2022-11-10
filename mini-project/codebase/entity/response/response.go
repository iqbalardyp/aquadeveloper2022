package response

type BaseResponse struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type GetCategoryResponse struct {
	Type string `json:"type"`
}

type BaseProductResponse struct {
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Category string `json:"category"`
	Figure   []byte `json:"figure"`
}

type GetProductDetailResponse struct {
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Category string `json:"category"`
	Desc     string `json:"desc"`
	Figure   []byte `json:"figure"`
}

type GetCartResponse struct {
	Qty     int64 `json:"qty"`
	Product BaseProductResponse
}
