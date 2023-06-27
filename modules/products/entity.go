package products

import (
	"errors"
	"time"
)

type Product struct {
	Id        int    `gorm:"primarykey" json:"id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Stock     int    `json:"stock"`
	DeletedAt *time.Time
}

type ProductResponse struct {
	Message string
	Data    *Product
}

type ProductsResponse struct {
	Message string
	Data    []Product
}

type ResponseAddAndEditData struct {
	Message string
	Data    Product
}

type RequestBodyStatus struct {
	Status string `json:"status"`
}

var (
	ErrProductAlreadyDeleted = errors.New("product already deleted")
	ErrProductNotDeleted     = errors.New("product is not deleted yet")
	ErrInvalidStatus         = errors.New("invalid status")
	ErrChangedStatus         = errors.New("status data cannot changed")
	ErrPoductHasBeenRemoved  = errors.New("product has been removed")
	ErrProductIdNotFound     = errors.New("data has been deleted")
)
