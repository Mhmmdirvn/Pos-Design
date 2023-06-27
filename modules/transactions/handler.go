package transactions

import (
	"encoding/json"
	"net/http"

)

type Handler struct {
	Usecase UseCase
}

func (handler Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var request CreateTransactionRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	transaction, err := handler.Usecase.CreateTransaction(r.Context(), &request)
	if err != nil {
		if err == ErrProductIdNotFound {
			json.NewEncoder(w).Encode(map[string]string{
				"Message": err.Error(),
			})
			return
		} else if err == ErrProductHasBeenRemoved {
			json.NewEncoder(w).Encode(map[string]string{
				"Message": err.Error(),
			})
			return
		} else if err == ErrStockNotEnough {
			json.NewEncoder(w).Encode(map[string]string{
				"Message": err.Error(),
			})
		}
	}

	

	// json.NewEncoder(w).Encode(map[string]interface{}{
	// 	
	// 	"Data" : transaction,
	// })

	response := &ResponseCreateProduct{
		Message: "Create Success",
		Data: *transaction,
	}

	json.NewEncoder(w).Encode(response)
}

func (handler Handler) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	transactions, err := handler.Usecase.GetAllTransactions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := []map[string]interface{}{}

	for _, data := range transactions {
		t := map[string]interface{}{
			"Id": data.Id,
			"TimeStamp": data.TimeStamp,
			"Total": data.Total,
			"admin_id": data.AdminID,
		}

		response = append(response, t)
	}

	// json.NewEncoder(w).Encode(map[string]interface{}{
	// 	"Message": "Success",
	// 	"Data" : response,
	// })
	responseMessage := &ResponseWithMap{
		Message: "Success",
		Data: response,
	}

	json.NewEncoder(w).Encode(responseMessage)
}

func (handler Handler) GetTransactionById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transaction, err := handler.Usecase.GetTransactionById(r.Context())
	if err != nil {
		http.Error(w, "Transaction Not Found", http.StatusNotFound)
		return
	}

	_, err = json.Marshal(transaction)
	if err != nil {
		http.Error(w, "Data Cannot Be Converted To JSON", http.StatusBadRequest)
		return
	}


	// json.NewEncoder(w).Encode(map[string]interface{}{
	// 	"Message": "Data Found",
	// 	"Data": transaction,
	// })


	response := &ResponseGetProductByID{
		Message: "Data Found",
		Data: *transaction,
	}

	json.NewEncoder(w).Encode(response)
}

