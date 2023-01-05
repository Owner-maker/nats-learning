package models

type Payment struct {
	Transaction  string `json:"transaction" binding:"required"`
	RequestId    string `json:"request_id" binding:"required"`
	Currency     string `json:"currency" binding:"required"`
	Provider     string `json:"provider" binding:"required"`
	Amount       int    `json:"amount" binding:"required"`
	PaymentDt    int    `json:"payment_dt" binding:"required"`
	Bank         string `json:"bank" binding:"required"`
	DeliveryCost int    `json:"delivery_cost" binding:"required"`
	GoodsTotal   int    `json:"goods_total" binding:"required"`
	CustomFee    int    `json:"custom_fee" binding:"required"`
}
