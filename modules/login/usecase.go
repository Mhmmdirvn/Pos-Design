package login

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Usecase struct {
	Repo Repository
}

func (usecase Usecase) Login(username string, password string) (string, error) {
	admin, err := usecase.Repo.Login(username, password)
	if err != nil {
		return "", err
	}

	claims := MyClaims{
		Id: admin.ID,
		Username: admin.Username,
		StandardClaims: jwt.StandardClaims{
			Issuer: APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}