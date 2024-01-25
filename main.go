package main

import (
	"os"

	"github.com/1nvers/go-broiler/controllers"
	"github.com/1nvers/go-broiler/initializers"
	"github.com/1nvers/go-broiler/middleware"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnv()
	initializers.ConnectDatabase()
	initializers.SyncDatabase()
}

func main(){
	r := gin.Default()
	r.POST("/signup",controllers.Signup)
	r.POST("/login",controllers.Login)
	r.GET("validate",middleware.ReuqireAuth, controllers.Validate)
	port := os.Getenv("PORT")
	r.Run(":"+port)
}