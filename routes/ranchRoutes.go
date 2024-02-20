package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oneaushaf/go-broiler/controllers"
	"github.com/oneaushaf/go-broiler/middleware"
)

func RanchRoutes(r *gin.Engine) {
	r.GET("/ranches", middleware.ReuqireAuth, controllers.GetRanches)
	r.POST("/ranch", middleware.ReuqireAuth, controllers.CreateRanch)
	r.GET("/ranch/:code", middleware.ReuqireAuth, controllers.GetRanch)
}
