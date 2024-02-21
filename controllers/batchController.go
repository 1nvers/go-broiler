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

func CreateBatch(c *gin.Context){
	var body struct {
		InitialQty	  uint 	 `binding:"required"`
		RanchCode	  string `binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Validation error",
		})
		return
	}

	ranch := models.Ranch{}
	check := database.DB.First(&ranch,"code=?",body.RanchCode)
	if check.Error != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error" : check.Error.Error(),
		})
		return
	} 

	batch := models.Batch{
		InitialQty: body.InitialQty,
		CurrentQty: body.InitialQty,
		Finished: false,
		RanchID:  ranch.ID,
	}

	result := database.DB.Create(&batch)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error" : result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success" : "Batch successfully created",
	})
}

func GetBatches(c *gin.Context){
	var batches []models.Batch
	var result []resources.BatchResource
	check := database.DB.Find(&batches)
	if check.Error != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error" : check.Error.Error(),
		})
	}
	for _, batch := range batches {
		result = append(result, resources.BatchDefaultResource(batch))
	}
	c.JSON(http.StatusOK, result)
}


func GetBatch(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid request",
		})
		return
	}
	var batch models.Batch
	check := database.DB.First(&batch, "id=?", id)
	if check.Error != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":check.Error.Error(),
		})
		return
	} else if check.RowsAffected == 0 {
		c.JSON(http.StatusNotFound,gin.H{
			"error":check.Error.Error(),
		})
	}
	c.JSON(http.StatusOK, resources.BatchDefaultResource(batch))
}

func GetBatchesByRanch(c *gin.Context) {
	ranch := models.Ranch{}
	ranchCode := c.Param("ranch_code")
	if ranchCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	if check := database.DB.Preload("Batches").First(&ranch, "code=?", ranchCode); check.Error != nil {
		if errors.Is(check.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "invalid ranch id",
			})
			return
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error": check.Error.Error(),
			})
			return
		}
	} else if len(ranch.Batches) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "ranch record not found",
		})
		return
	}

	var result []resources.BatchResource

	for _, batch := range ranch.Batches {
		result = append(result, resources.BatchDefaultResource(batch))
	}
	c.JSON(http.StatusOK, result)
}