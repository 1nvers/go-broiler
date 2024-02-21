package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oneaushaf/go-broiler/controllers"
	"github.com/oneaushaf/go-broiler/middleware"
)

func RanchRoutes(r *gin.Engine) {
	r.POST("/farms/:farm_code/ranches", middleware.ReuqireAuth, controllers.CreateRanch)
	r.GET("/farms/:farm_code/ranches", middleware.ReuqireAuth, controllers.GetRanchesByFarm)
	r.GET("/ranches", middleware.ReuqireAuth, controllers.GetRanches)

	r.GET("/ranches/:ranch_code", middleware.ReuqireAuth, controllers.GetRanch)
	r.PUT("/ranches/:ranch_code", middleware.ReuqireAuth, controllers.TempHandler)
	r.DELETE("/ranches/:ranch_code", middleware.ReuqireAuth, controllers.TempHandler)
}
