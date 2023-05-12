package transactions

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	Usecase UseCase
}

func (handler Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request CreateTransactionRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "Failed to decode json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	transaction, err := handler.Usecase.CreateTransaction(&request)
	if err != nil {
		http.Error(w, "Data Can't Be Added", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"Message": "Create Success",
		"Data" : transaction,
	})
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
		}

		response = append(response, t)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"Message": "Success",
		"Data" : response,
	})
}

func (handler Handler) GetTransactionById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	transaction, err := handler.Usecase.GetTransactionById(id)
	if err != nil {
		http.Error(w, "Transaction Not Found", http.StatusNotFound)
		return
	}

	_, err = json.Marshal(transaction)
	if err != nil {
		http.Error(w, "Data Cannot Be Converted To JSON", http.StatusBadRequest)
		return
	}


	json.NewEncoder(w).Encode(map[string]interface{}{
		"Message": "Data Found",
		"Data": transaction,
	})
}

