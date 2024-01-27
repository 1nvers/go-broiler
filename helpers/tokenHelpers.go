package helpers

import (
	"os"
	"time"

	"github.com/oneaushaf/go-broiler/models"
	"github.com/dgrijalva/jwt-go"
)

func GenerateTokens(user models.User)(string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"sub":user.ID,
		"exp":time.Now().Add(time.Hour*24*30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	return tokenString,err
}