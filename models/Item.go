package models

type Item struct {
	ChrtId      int    `json:"chrt_id" binding:"required"`
	TrackNumber string `json:"track_number" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Rid         string `json:"rid" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Sale        int    `json:"sale" binding:"required"`
	Size        string `json:"size" binding:"required"`
	TotalPrice  int    `json:"total_price" binding:"required"`
	NmId        int    `json:"nm_id" binding:"required"`
	Brand       string `json:"brand" binding:"required"`
	Status      int    `json:"status" binding:"required"`
}
