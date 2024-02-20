package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oneaushaf/go-broiler/controllers"
	"github.com/oneaushaf/go-broiler/middleware"
)

func BatchRoutes(r *gin.Engine) {
	r.GET("/batches", middleware.ReuqireAuth, controllers.GetBatches)
	r.POST("/batch", middleware.ReuqireAuth, controllers.CreateBatch)
	r.GET("/batch/:id", middleware.ReuqireAuth, controllers.GetBatch)
}
