package models

type Item struct {
	OrderRefer  string `json:"-"`
	ChrtId      int    `json:"chrt_id" validate:"required"`
	TrackNumber string `json:"track_number" validate:"required,min=14,max=14"`
	Price       int    `json:"price" validate:"gt=0"`
	Rid         string `json:"rid" validate:"required,min=21,max=21"`
	Name        string `json:"name" validate:"required"`
	Sale        int    `json:"sale" validate:"gt=0"`
	Size        string `json:"size" validate:"required"`
	TotalPrice  int    `json:"total_price" validate:"gt=0"`
	NmId        int    `json:"nm_id" validate:"required"`
	Brand       string `json:"brand" validate:"required"`
	Status      int    `json:"status" validate:"required,min=0, max=999"`
}
