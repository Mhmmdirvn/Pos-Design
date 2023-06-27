package transactions

import (
	"Pos-Design/modules/products"
	"errors"
	"time"
)

type Transaction struct {
	Id        int                `gorm:"primarykey" json:"id"`
	TimeStamp time.Time          `json:"timestamp"`
	Total     int                `json:"total"`
	AdminID   int                `json:"admin_id"`
	Admin     *Admin             `json:"admin"`
	Items     []TransactionItems `json:"items"`
}

type Admin struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TransactionItems struct {
	Id            int               `gorm:"primarykey"`
	TransactionID int               `gorm:"foreignkey:TransactionID"`
	ProductID     int               `gorm:"foreignkey:ProductID"`
	Quantity      int               `json:"quantity"`
	Price         int               `json:"price"`
	Product       *products.Product `json:"product"`
}

type CreateItemRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CreateTransactionRequest struct {
	Items []CreateItemRequest `json:"items"`
}

type ResponseWithMap struct {
	Message string
	Data    []map[string]interface{}
}

type ResponseCreateProduct struct {
	Message string
	Data Transaction
}

type ResponseGetProductByID struct {
	Message string
	Data Transaction
}

var (
	ErrProductIdNotFound     = errors.New("product id not found")
	ErrStockNotEnough        = errors.New("stock not enough")
	ErrProductHasBeenRemoved = errors.New("product has been removed")
)
