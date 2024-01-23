package model

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/vinialeixo/crud-golang/src/configuration/rest_err"
)

const (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	//claims é padronizado em toda linguagem de programação que for utilizado o jwt
	//objeto que vai especificar quais campos que terá no dentro do jwt

	claims := jwt.MapClaims{
		"id":    ud.id,
		"email": ud.email,
		"name":  ud.name,
		"age":   ud.age,
		"exp":   time.Now().Add(time.Hour * 1).Unix(), //1 hora de expiramento
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_err.NewInternalServerError(
			fmt.Sprintf("error trying to generate jwt token, err=%s", err.Error()))
	}

	return tokenString, nil
}
