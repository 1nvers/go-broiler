package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TempHandler(c *gin.Context){
	c.JSON(http.StatusBadRequest,gin.H{
		"message":"routes not being handled yet",
	})
}