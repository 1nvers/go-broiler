package routes

import (
	"github.com/oneaushaf/go-broiler/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine){
	authGroup := r.Group("")
	authGroup.POST("/signup",controllers.Signup)
	authGroup.POST("/login",controllers.Login)
}