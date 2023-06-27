package register

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Usecase UseCase
}

func (handler Handler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var register Admin

	err := json.NewDecoder(r.Body).Decode(&register)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = handler.Usecase.Register(&register)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	response := &Response{
		Message: "Register Success",
		Data: register,
	}

	json.NewEncoder(w).Encode(response)
}