package transactions

import (
	"Pos-Design/modules/products"
	"time"
)

type Transaction struct {
	Id        int                `gorm:"primarykey" json:"id"`
	TimeStamp time.Time          `json:"timestamp"`
	Total     int                `json:"total"`
	Items     []TransactionItems `json:"items"`
}

type TransactionItems struct {
	Id            int `gorm:"primarykey"`
	TransactionID int `gorm:"foreignkey:TransactionID"`
	ProductID     int `gorm:"foreignkey:ProductID"`
	Quantity      int `json:"quantity"`
	Price         int `json:"price"`
	Product       *products.Product `json:"product"`
}

type CreateItemRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CreateTransactionRequest struct {
	Items []CreateItemRequest
}
