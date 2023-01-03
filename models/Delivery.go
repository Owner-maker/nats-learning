package models

type Delivery struct {
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Zip     string `json:"zip" binding:"required"`
	City    string `json:"city" binding:"required"`
	Address string `json:"address" binding:"required"`
	Region  string `json:"region" binding:"required"`
	Email   string `json:"email" binding:"required"`
}
