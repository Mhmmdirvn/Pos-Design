package products

import (
	"encoding/json"
	"net/http"

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

	

	response := &ResponseAddAndEditData{
		Message: "Data Added Successfully",
		Data: product,
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
	response := &ProductsResponse{
		Message: "Data Found",
		Data: products,
	}

	json.NewEncoder(w).Encode(response)
}

func (handler Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	product, err := handler.Usecase.GetProductById(r.Context())
	if err != nil {
		http.Error(w, "Product Not Found", http.StatusNotFound)
		return
	}


	response := &ProductResponse{
		Message: "Data Found",
		Data: product,
	}

	json.NewEncoder(w).Encode(response)
}

func (handler Handler) UpdateProductById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)


	

	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		messageErr, _ := json.Marshal(map[string]string{"Message": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	if err := handler.Usecase.UpdateProductById(r.Context(), &product); err != nil {
		if err != nil {
			if err == ErrPoductHasBeenRemoved {
				json.NewEncoder(w).Encode(map[string]string{
					"Message": err.Error(),
				})
				return
			} else if err == ErrProductIdNotFound {
				json.NewEncoder(w).Encode(map[string]string{
					"Message": err.Error(),
				})
				return
			}
		}
	}
	



	response := &ResponseAddAndEditData{
		Message: "Data Has Been Updated",
		Data: product,
	}

	json.NewEncoder(w).Encode(response)

}



func (handler Handler) SoftDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var requestBody RequestBodyStatus
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		messageErr, _ := json.Marshal(map[string]string{"Message": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	product, err := handler.Usecase.SoftDelete(r.Context(), requestBody.Status)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"Message": err.Error(),
		})
		return
	}

	response := &ProductResponse{
		Message: "Success",
		Data: product,
	}

	json.NewEncoder(w).Encode(response)
}



func (handler Handler) DeleteProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	

	err := handler.Usecase.DeleteProductById(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"Message": "Data Has Been Deleted",
	}

	json.NewEncoder(w).Encode(response)
}