package login

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Admin struct {
	ID       int    `gorm:"primarykey"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type MyClaims struct {
	jwt.StandardClaims
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type AdminResponse struct {
	Message string
	Data jwt.MapClaims
}

var APPLICATION_NAME = "Muhammad Irvan"
var LOGIN_EXPIRATION_DURATION = time.Duration(12) * time.Hour
var JWT_SIGNATURE_KEY = []byte("jaksdjhhwhjhds")
