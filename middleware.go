package main

import (
	"Pos-Design/modules/login"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

func JwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			messageErr, _ := json.Marshal(map[string]string{"Message": "Invalid Token"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(messageErr)
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		token, err := jwt.ParseWithClaims(tokenString, &login.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("signing method invalid")
			}

			return login.JWT_SIGNATURE_KEY, nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		claims, ok := token.Claims.(*login.MyClaims) 
		if !ok || !token.Valid {
			fmt.Println(&ok)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.Background()
		ctx = context.WithValue(ctx, "idPrms", id)
		ctx = context.WithValue(ctx, "id_admin", claims.Id)
		ctx = context.WithValue(ctx, "name", claims.Username)

		next(w, r.WithContext(ctx))
	})
}