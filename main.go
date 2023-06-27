package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Pos-Design/modules/login"
	"Pos-Design/modules/products"
	"Pos-Design/modules/register"
	"Pos-Design/modules/transactions"
)

func main() {
	// Connection To Database POS
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/pos?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("failed connect to database")
	}

	productRepo := products.Repository{DB: db}
	productUseCase := products.UseCase{Repo: productRepo}
	productHandler := products.Handler{Usecase: productUseCase}

	transactionRepo := transactions.Repository{DB: db}
	transactionUseCase := transactions.UseCase{Repo: transactionRepo, ProductRepo: productRepo}
	transactionHandler := transactions.Handler{Usecase: transactionUseCase}

	RegisterRepo := register.Repository{DB: db}
	RegisterUseCase := register.UseCase{Repo: RegisterRepo}
	RegisterHandler := register.Handler{Usecase: RegisterUseCase}

	LoginRepo := login.Repository{DB: db}
	loginUsecase := login.Usecase{Repo: LoginRepo}
	loginHandler := login.Handler{Usecase: loginUsecase}



	// New Router
	r := mux.NewRouter()

	// Handler Register
	r.HandleFunc("/register", RegisterHandler.Register).Methods("POST")

	// Handler Login
	r.HandleFunc("/login", loginHandler.Login).Methods("POST")

	// Handler Products
	r.HandleFunc("/products", JwtMiddleware(productHandler.GetAllProducts)).Methods("GET")
	r.HandleFunc("/products/{id}", JwtMiddleware(productHandler.GetProductById)).Methods("GET")
	r.HandleFunc("/products", JwtMiddleware(productHandler.CreateProduct)).Methods("POST")
	r.HandleFunc("/products/{id}", JwtMiddleware(productHandler.UpdateProductById)).Methods("PUT")
	r.HandleFunc("/products/{id}", JwtMiddleware(productHandler.DeleteProductById)).Methods("DELETE")
	r.HandleFunc("/products/{id}/status", JwtMiddleware(productHandler.SoftDelete)).Methods("PATCH")


	// Handler Transactions
	r.HandleFunc("/transactions", JwtMiddleware(transactionHandler.GetAllTransactions)).Methods("GET")
	r.HandleFunc("/transactions", JwtMiddleware(transactionHandler.CreateTransaction)).Methods("POST")
	r.HandleFunc("/transactions/{id}", JwtMiddleware(transactionHandler.GetTransactionById)).Methods("GET")
	
	// Set Port
	fmt.Println("Server started at localhost:9000")
	http.ListenAndServe(":9000", r)
}