package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Pos-Design/modules/products"
	"Pos-Design/modules/transactions"
)

func main() {
	// Connection To Database POS
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3307)/pos?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("failed connect to database")
	}

	productRepo := products.Repository{DB: db}
	productUseCase := products.UseCase{Repo: productRepo}
	productHandler := products.Handler{Usecase: productUseCase}

	transactionRepo := transactions.Repository{DB: db}
	transactionUseCase := transactions.UseCase{Repo: transactionRepo, ProductRepo: productRepo}
	transactionHandler := transactions.Handler{Usecase: transactionUseCase}

	r := mux.NewRouter()

	// Handler Products
	r.HandleFunc("/products", productHandler.GetAllProducts).Methods("GET")
	r.HandleFunc("/products/{id}", productHandler.GetProductById).Methods("GET")
	r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", productHandler.UpdateProductById).Methods("PUT")
	r.HandleFunc("/products/{id}", productHandler.DeleteProductById).Methods("DELETE")

	// Handler Transactions
	r.HandleFunc("/transactions", transactionHandler.GetAllTransactions).Methods("GET")
	r.HandleFunc("/transactions", transactionHandler.CreateTransaction).Methods("POST")
	r.HandleFunc("/transactions/{id}", transactionHandler.GetTransactionById).Methods("GET")
	
	// Set Port
	fmt.Println("Server started at localhost:9000")
	http.ListenAndServe(":9000", r)
}