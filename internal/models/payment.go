package models

type Payment struct {
	OrderRefer   string `json:"-"`
	Transaction  string `json:"transaction" validate:"required"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency" validate:"required"`
	Provider     string `json:"provider" validate:"required"`
	Amount       int    `json:"amount" validate:"gt=0"`
	PaymentDt    int    `json:"payment_dt" validate:"required"`
	Bank         string `json:"bank" validate:"required"`
	DeliveryCost int    `json:"delivery_cost" validate:"gt=0"`
	GoodsTotal   int    `json:"goods_total" validate:"gt=0"`
	CustomFee    int    `json:"custom_fee" validate:"gte=0"`
}
