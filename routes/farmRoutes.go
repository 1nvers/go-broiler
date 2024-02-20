package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oneaushaf/go-broiler/controllers"
	"github.com/oneaushaf/go-broiler/middleware"
)

func FarmRoutes(r *gin.Engine) {
	r.GET("/farms", middleware.ReuqireAuth ,controllers.GetFarms)
	r.POST("/farm", middleware.ReuqireAuth ,controllers.CreateFarm)
	r.GET("/farm/:code", middleware.ReuqireAuth ,controllers.GetFarm)
}
