package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/oneaushaf/go-broiler/models"
)

type Claims struct{
	UserID 	 uint `json:"userID"`
	UserType string `json:"userType"`
	jwt.StandardClaims
}

func GenerateTokens(user models.User)(string,error){
	expirationTime := time.Now().Add(24 * time.Hour)
			claims := &Claims{
				UserID: user.ID,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	return tokenString,err
}

func ValidateToken(tokenString string) (claims *Claims, er error) {
	token, err := jwt.ParseWithClaims(tokenString,&Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return nil , err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok{
		err = errors.New("failed to get claims")
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix(){
		err = errors.New("token is expired")
		return nil, err
	}
	
	return claims,nil
}