package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oneaushaf/go-broiler/database"
	"github.com/oneaushaf/go-broiler/models"
	"github.com/oneaushaf/go-broiler/resources"
)

func CreateBatch(c *gin.Context){
	var body struct {
		InitialQty	  uint 	
		RanchCode	  string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Failed to read body",
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

	c.JSON(http.StatusOK, gin.H{})
}

func GetBatches(c *gin.Context){
	var batches []models.Batch
	var result []resources.BatchResource
	check := database.DB.Find(&batches)
	if check.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else if len(batches) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	}
	for _, batch := range batches {
		result = append(result, resources.BatchDefaultResource(batch))
	}
	c.JSON(http.StatusOK, gin.H{
		"batches": result,
	})
}


func GetBatch(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"invalid request",
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
	} 
	c.JSON(http.StatusOK, gin.H{
		"Batch": resources.BatchDefaultResource(batch),
	})
}