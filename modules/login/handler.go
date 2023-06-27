package login

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	Usecase Usecase
}

func (handler Handler) Login(w http.ResponseWriter, r *http.Request) {
	var userInput Admin
	
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		messageErr, _ := json.Marshal(map[string]string{"Message": "Failed decode to JSON"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	defer r.Body.Close()

	signedToken, err := handler.Usecase.Login(userInput.Username, userInput.Password) 
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"Message": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}
	
	fmt.Println(signedToken)
	w.Write([]byte(signedToken))

}