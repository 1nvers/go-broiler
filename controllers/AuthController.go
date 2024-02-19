package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oneaushaf/go-broiler/helpers"
)

func Signup(c *gin.Context){
	var body struct{
		FirstName    string
		LastName     string
		Phone		 string
		Email 	     string
		UserType 	 string
		Password     string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Failed to read body",
		})
		return
	}

	err := CreateUser(body.FirstName,body.LastName,body.Phone,body.Email,body.UserType,body.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context){
	var body struct {
		Email 	 string
		Password string
	}
	
	if c.Bind(&body)!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to read body",
		})
		return
	}

	user, err := helpers.CheckCredentials(body.Email,body.Password)

	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid email or password",
		})
		return
	}

	tokenString, err := helpers.GenerateTokens(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Failed to create token",
		})
		return
	}

	// c.SetSameSite(http.SameSiteLaxMode)
	// c.SetCookie("Authorization",tokenString, 3600*24 ,"","",true,true)
	c.JSON(http.StatusOK,gin.H{
		"token":tokenString,
	})
}