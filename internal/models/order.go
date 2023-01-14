package models

import (
	"time"
)

type Order struct {
	OrderUid          string    `json:"order_uid" validate:"required,min=19,max=19" gorm:"primary_key; unique"`
	TrackNumber       string    `json:"track_number" validate:"required,min=14,max=14"`
	Entry             string    `json:"entry" validate:"required,min=4,max=4"`
	Locale            string    `json:"locale" validate:"oneof=ru en"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id" validate:"required,min=4,max=4"`
	DeliveryService   string    `json:"delivery_service" validate:"required,min=5,max=5"`
	ShardKey          string    `json:"shardkey"`
	SmId              int       `json:"sm_id" validate:"gte=0,lte=100"`
	DateCreated       time.Time `json:"date_created" format:"2006-01-02T06:22:19Z" validate:"required"`
	OofShard          string    `json:"oof_shard" validate:"required,max=2"`
	Delivery          Delivery  `json:"delivery" validate:"required" gorm:"foreignKey:OrderRefer"`
	Payment           Payment   `json:"payment" validate:"required" gorm:"foreignKey:OrderRefer"`
	Items             []Item    `json:"items" validate:"required" gorm:"foreignKey:OrderRefer"`
}
