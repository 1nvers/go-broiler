package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oneaushaf/go-broiler/controllers"
	"github.com/oneaushaf/go-broiler/middleware"
)

func BatchRoutes(r *gin.Engine) {
	r.POST("/ranches/:ranch_code/batches", middleware.ReuqireAuth, controllers.CreateBatch)
	r.GET("/ranches/:ranch_code/batches", middleware.ReuqireAuth, controllers.GetBatchesByRanch)
	r.GET("/batches", middleware.ReuqireAuth, controllers.GetBatches)

	r.GET("/batches/:batches_id", middleware.ReuqireAuth, controllers.GetBatch)
	r.PUT("/batches/:batches_id", middleware.ReuqireAuth, controllers.TempHandler)
	r.DELETE("/batches/:batches_id", middleware.ReuqireAuth, controllers.TempHandler)
}
