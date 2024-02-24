package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oneaushaf/go-broiler/controllers"
	"github.com/oneaushaf/go-broiler/middleware"
)

func WeighingRoutes(r *gin.Engine) {
	r.POST("/batch/:batch_id/weighings",middleware.ReuqireAuth, controllers.CreateWeighing)
	r.GET("/batch/:batch_id/weighings",middleware.ReuqireAuth, controllers.GetWeighingsByFarm)
	r.GET("/weighings",middleware.ReuqireAuth, controllers.GetWeighings)
	r.POST("/weighings/:weighing_id/image",middleware.ReuqireAuth, controllers.UploadImage)
	r.GET("/weighings/:weighing_id/image",middleware.ReuqireAuth, controllers.UploadImage)
	
	r.GET("/weighings/:weighing_id",middleware.ReuqireAuth, controllers.GetWeighing)
	r.PUT("/weighings/:weighing_id",middleware.ReuqireAuth, controllers.TempHandler)
	r.DELETE("/weighings/:weighing_id",middleware.ReuqireAuth, controllers.TempHandler)
}