package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oneaushaf/go-broiler/helpers"
)

func Signup(c *gin.Context) {
	var body struct {
		FirstName string `binding:"required"`
		LastName  string `binding:"required"`
		Phone     string `binding:"required"`
		Email     string `binding:"required"`
		UserType  string `binding:"required"`
		Password  string `binding:"required"`
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Validation error",
		})
		return
	}

	err := CreateUser(body.FirstName, body.LastName, body.Phone, body.Email, body.UserType, body.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success" : "Signup completed successfuly",
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string  `binding:"required"`
		Password string  `binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation error",
		})
		return
	}

	user, err := helpers.CheckCredentials(body.Email, body.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tokenString, err := helpers.GenerateTokens(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
