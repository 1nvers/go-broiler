package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oneaushaf/go-broiler/database"
	"github.com/oneaushaf/go-broiler/models"
	"github.com/oneaushaf/go-broiler/resources"
)

func CreateFarm(c *gin.Context){
	var body struct{
		Code    string  `binding:"required"`
		Adress  string  `binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "validation error",
			"details" : err.Error(),
		})
		return
	}

	farm := models.Farm{
		Code: body.Code,
		Adress:  body.Adress,
	}

	result := database.DB.Create(&farm)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error" : result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success" : "Farm successfully created",
	})
}

func GetFarms(c *gin.Context){
	var farms []models.Farm
	var result []resources.FarmResource
	check := database.DB.Find(&farms)
	if check.Error != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error" : check.Error.Error(),
		})
	} 

	for _, farm := range farms {
		result = append(result, resources.FarmDefaultResource(farm))
	}

	c.JSON(http.StatusOK,result)
}


func GetFarm(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusBadRequest,gin.H{
			"error" : "Invalid request",
		})
		return
	}

	var Farm models.Farm

	check := database.DB.First(&Farm, "code=?", code)
	if check.Error != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error" : check.Error.Error(),
		})
		return
	} else if check.RowsAffected == 0 {
		c.JSON(http.StatusNotFound,gin.H{
			"error":check.Error.Error(),
		})
	}
	c.JSON(http.StatusOK,resources.FarmDefaultResource(Farm))
}