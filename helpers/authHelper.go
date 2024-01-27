package helpers

import (
	"errors"

	"github.com/oneaushaf/go-broiler/database"
	"github.com/oneaushaf/go-broiler/models"
	"golang.org/x/crypto/bcrypt"
)

func CheckCredentials(email string, password string) (models.User,error){
	var user models.User
	database.DB.First(&user, "email = ?", email)

	if user.ID == 0 {
		return user,errors.New("invalid email or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))

	if err != nil {
		return user,errors.New("invalid email or password")
	}

	return user,nil
}