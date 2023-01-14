package models

type Delivery struct {
	OrderRefer string `json:"-"`
	Name       string `json:"name" validate:"required,max=30"`
	Phone      string `json:"phone" validate:"required"`
	Zip        string `json:"zip" validate:"required,max=10"`
	City       string `json:"city" validate:"required,max=30"`
	Address    string `json:"address" validate:"required,max=30"`
	Region     string `json:"region" validate:"required,max=30"`
	Email      string `json:"email" validate:"required,max=30"`
}
