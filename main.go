package main

import (
	"github.com/gin-gonic/gin"
	"github.com/1nvers/go-broiler/initializers"
	"os"
)

func init(){
	initializers.LoadEnv()
	initializers.ConnectDatabase()
	initializers.SyncDatabase()
}

func main(){
	r := gin.Default()
	r.GET("/test",func (c *gin.Context)  {
		c.JSON(200,gin.H{
			"test":"success",
		})
	})
	port := os.Getenv("PORT")
	r.Run(":"+port)
}