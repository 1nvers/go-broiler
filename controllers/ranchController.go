package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oneaushaf/go-broiler/database"
	"github.com/oneaushaf/go-broiler/models"
	"github.com/oneaushaf/go-broiler/resources"
)

func CreateRanch(c *gin.Context){
	var body struct{
		Capacity uint
		Code     string
		FarmID	 uint
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Failed to read body",
		})
		return
	}

	Ranch := models.Ranch{
		Code: body.Code,
		Capacity:  body.Capacity,
		FarmID: body.FarmID,
	}

	farm := models.Farm{} 

	check := database.DB.First(&farm,"id=?",body.FarmID)
	if check.Error != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error" : "server error",
		})
		return
	} else if check.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest,gin.H{
			"error" : "server error",
		})
		return
	} 

	result := database.DB.Create(&Ranch)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error" : result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func GetRanches(c *gin.Context){
	var ranches []models.Ranch
	var result []resources.RanchResource
	check := database.DB.Find(&ranches)
	if check.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else if len(ranches) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	}
	for _, Ranch := range ranches {
		result = append(result, resources.RanchDefaultResource(Ranch))
	}
	c.JSON(http.StatusOK, gin.H{
		"ranches": result,
	})
}


func GetRanch(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"invalid request",
		})
		return
	}
	var Ranch models.Ranch
	check := database.DB.First(&Ranch, "code=?", code)
	if check.Error != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":check.Error.Error(),
		})
		return
	} 
	c.JSON(http.StatusOK, gin.H{
		"ranch": resources.RanchDefaultResource(Ranch),
	})
}