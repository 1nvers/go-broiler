package routes

import (
	"github.com/oneaushaf/go-broiler/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine){
	authGroup := r.Group("/user")
	authGroup.POST("/",controllers.Signup)
	authGroup.POST("/login",controllers.Login)
}