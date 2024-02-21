package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oneaushaf/go-broiler/database"
	"github.com/oneaushaf/go-broiler/models"
	"github.com/oneaushaf/go-broiler/resources"
	"gorm.io/gorm"
)

func CreateRanch(c *gin.Context) {
	farm := models.Farm{}
	farmID := c.Param("farm_code")
	if farmID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	check := database.DB.First(&farm, "id=?", farmID)
	if check.Error != nil {
		if check.Error.Error() != "record not found" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": check.Error.Error(),
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid farm id",
			})
			return
		}
	}

	var body struct {
		Capacity uint   `binding:"required"`
		Code     string `binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "validation error",
		})
		return
	}

	Ranch := models.Ranch{
		Code:     body.Code,
		Capacity: body.Capacity,
		FarmID:   farm.ID,
	}

	result := database.DB.Create(&Ranch)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "ranch successfully created",
	})
}

func GetRanches(c *gin.Context) {
	var ranches []models.Ranch
	var result []resources.RanchResource
	if check := database.DB.Find(&ranches); check.Error != nil {
		if errors.Is(check.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "ranch record not found",
			})
			return
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error": check.Error.Error(),
			})
			return
		}
	}
	for _, Ranch := range ranches {
		result = append(result, resources.RanchDefaultResource(Ranch))
	}
	c.JSON(http.StatusOK, result)
}

func GetRanch(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}
	var Ranch models.Ranch
	check := database.DB.First(&Ranch, "code=?", code)
	if check.Error != nil {
		if errors.Is(check.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": check.Error.Error(),
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "ranch record not found",
			})
			return
		}
	}
	c.JSON(http.StatusOK, resources.RanchDefaultResource(Ranch))
}

func GetRanchesByFarm(c *gin.Context) {
	farm := models.Farm{}
	farmCode := c.Param("farm_code")
	if farmCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	if check := database.DB.Preload("Ranches").First(&farm, "code=?", farmCode); check.Error != nil {
		if errors.Is(check.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "invalid farm id",
			})
			return
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error": check.Error.Error(),
			})
			return
		}
	} else if len(farm.Ranches) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "ranch record not found",
		})
		return
	}

	var result []resources.RanchResource

	for _, Ranch := range farm.Ranches {
		result = append(result, resources.RanchDefaultResource(Ranch))
	}
	c.JSON(http.StatusOK, result)
}
