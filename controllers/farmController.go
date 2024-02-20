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
		Code    string
		Adress  string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Failed to read body",
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
			"error" : "server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func GetFarms(c *gin.Context){
	var farms []models.Farm
	var result []resources.FarmResource
	check := database.DB.Find(&farms)
	if check.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else if len(farms) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	}
	for _, farm := range farms {
		result = append(result, resources.FarmDefaultResource(farm))
	}
	c.JSON(http.StatusOK, gin.H{
		"farms": result,
	})
}


func GetFarm(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	var Farm models.Farm
	check := database.DB.First(&Farm, "code=?", code)
	if check.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else if check.RowsAffected == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, gin.H{
		"farm": resources.FarmDefaultResource(Farm),
	})
}