package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oneaushaf/go-broiler/middleware"
)

func Ranchroutes(r *gin.Engine) {
	r.GET("/ranches", middleware.ReuqireAuth, func(c *gin.Context) {})
	r.GET("/ranches/:id", middleware.ReuqireAuth, func(c *gin.Context) {})
}
