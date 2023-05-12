package products

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	Usecase UseCase
}

func (handler Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = handler.Usecase.CreateProduct(&product)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	

	response := map[string]interface{}{
		"Message": "Success",
		"Data": product,
	}

	json.NewEncoder(w).Encode(response)
}


func (handler Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)


	products, err := handler.Usecase.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"Message": "Success",
		"Data": products,
	}

	json.NewEncoder(w).Encode(response)
}

func (handler Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	product, err := handler.Usecase.GetProductById(id)
	if err != nil {
		http.Error(w, "Product Not Found", http.StatusNotFound)
		return
	}


	response := map[string]interface{}{
		"Message": "Success",
		"Data": product,
	}

	json.NewEncoder(w).Encode(response)
}

func (handler Handler) UpdateProductById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	

	var product Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	product.Id = id

	err = handler.Usecase.UpdateProductById(id, &product)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	response := map[string]interface{}{
		"Message": "Data Has Been Updated",
		"Data": product,
	}

	json.NewEncoder(w).Encode(response)

}


func (handler Handler) DeleteProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil{
		http.Error(w, "Invalid ID", http.StatusBadRequest)
	}

	err = handler.Usecase.DeleteProductById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"Message": "Data Has Been Deleted",
	}

	json.NewEncoder(w).Encode(response)
}