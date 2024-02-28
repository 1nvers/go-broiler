package controllers

import (
	// "crypto/sha256"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oneaushaf/go-broiler/database"
	"github.com/oneaushaf/go-broiler/helpers"
	"github.com/oneaushaf/go-broiler/models"
	"github.com/oneaushaf/go-broiler/resources"
	"gorm.io/gorm"
)

func UploadImage(c *gin.Context) {
	// weighingID := c.Param("weighing_id")
	// weighing := models.Weighing{}
	// check := database.DB.First(&weighing, "id=?", weighingID)
	// if check.Error != nil {
	// 	if errors.Is(check.Error, gorm.ErrRecordNotFound) {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"error": check.Error.Error(),
	// 		})
	// 		return
	// 	} else {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "weighing record not found",
	// 		})
	// 		return
	// 	}
	// }

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "failed to upload image",
		})
		return
	}

	ext := helpers.GetExtention(file.Filename)
	name := helpers.RandString(8)
	fileName := name + "." + ext

	// weighing.Image = fileName
	// check = database.DB.Save(&weighing)

	// if check.Error != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	err = c.SaveUploadedFile(file, "images/weighing/"+fileName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "failed to save image",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "image uploaded successfully",
		"image":   fileName,
	})
}

func CreateWeighing(c *gin.Context) {
	batch := models.Batch{}
	batchID := c.Param("batch_id")
	if batchID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	check := database.DB.First(&batch, "id=?", batchID)
	if check.Error != nil {
		if check.Error.Error() != "record not found" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": check.Error.Error(),
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid batch id",
			})
			return
		}
	}

	var body struct {
		Age      uint `binding:"required"`
		Deceased uint `binding:"required"`
		Image 	 string `binding:"required"`
	}

	if err := c.Bind(&body) ; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	requestData := gin.H{
		"age": body.Age,
		"count": (batch.CurrentQty - body.Deceased),
		"image": body.Image,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal JSON"})
		return
	}

	// Create an HTTP POST request to the target API
	targetURL := "http://127.0.0.1:8000/predict/weight"
	resp, err := http.Post(targetURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to make POST request"})
		return
	}
	defer resp.Body.Close()

	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response body"})
		return
	}

	weighing := models.Weighing{
		Age:      body.Age,
		Deceased: body.Deceased,
		BatchID:  batch.ID,
	}

	result := database.DB.Create(&weighing)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "weighing successfully created",
		"data": responseBody,
	})
}

func GetWeighings(c *gin.Context) {
	var weighings []models.Weighing
	var result []resources.WeighingResource
	if check := database.DB.Find(&weighings); check.Error != nil {
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
	for _, weighing := range weighings {
		result = append(result, resources.WeighingDefaultResource(weighing))
	}
	c.JSON(http.StatusOK, result)
}

func GetWeighing(c *gin.Context) {
	id := c.Param("weighing_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}
	var weighing models.Weighing
	check := database.DB.First(&weighing, "id=?", id)
	if check.Error != nil {
		if errors.Is(check.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": check.Error.Error(),
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "weighing record not found",
			})
			return
		}
	}
	c.JSON(http.StatusOK, resources.WeighingDefaultResource(weighing))
}

func GetWeighingsByFarm(c *gin.Context) {
	batch := models.Batch{}
	batchCode := c.Param("batch_code")
	if batchCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	if check := database.DB.Preload("Weighings").First(&batch, "code=?", batchCode); check.Error != nil {
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
	} else if len(batch.Weighings) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "weighing record not found",
		})
		return
	}

	var result []resources.WeighingResource

	for _, weighing := range batch.Weighings {
		result = append(result, resources.WeighingDefaultResource(weighing))
	}
	c.JSON(http.StatusOK, result)
}
