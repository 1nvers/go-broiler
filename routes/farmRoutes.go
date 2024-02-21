package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oneaushaf/go-broiler/controllers"
	"github.com/oneaushaf/go-broiler/middleware"
)

func FarmRoutes(r *gin.Engine) {
	r.POST("/farms", middleware.ReuqireAuth ,controllers.CreateFarm)
	r.GET("/farms", middleware.ReuqireAuth ,controllers.GetFarms)
	
	r.GET("/farms/:farm_code", middleware.ReuqireAuth ,controllers.GetFarm)
	r.PUT("/farms/:farm_code", middleware.ReuqireAuth ,controllers.TempHandler) //
	r.DELETE("/farms/:farm_code", middleware.ReuqireAuth ,controllers.TempHandler) // 
}
