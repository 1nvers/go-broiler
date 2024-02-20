package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oneaushaf/go-broiler/database"
	"github.com/oneaushaf/go-broiler/models"
	"github.com/oneaushaf/go-broiler/resources"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(firstName string, lastName string, phone string, email string, userType string, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return errors.New("failed to hash password")
	}

	user := models.User{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Email:     email,
		UserType:  userType,
		Password:  string(hash)}

	result := database.DB.Create(&user)

	if result.Error != nil {
		return errors.New("failed to create user")
	}
	return nil
}

func GetAuth(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	var user models.User
	check := database.DB.First(&user, "id=?", userID)
	if check.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else if check.RowsAffected == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, gin.H{
		"user": resources.UserDefaultResource(user),
	})
}
func GetUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	var user models.User
	check := database.DB.First(&user, "id=?", userID)
	if check.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else if check.RowsAffected == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, gin.H{
		"user": resources.UserDefaultResource(user),
	})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	var result []resources.UserResource

	check := database.DB.Find(&users)
	if check.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else if len(users) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	}
	for _, user := range users {
		result = append(result, resources.UserDefaultResource(user))
	}
	c.JSON(http.StatusOK, gin.H{
		"users": result,
	})
}
