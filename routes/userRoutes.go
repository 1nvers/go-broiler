package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oneaushaf/go-broiler/controllers"
	"github.com/oneaushaf/go-broiler/middleware"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/users/auth",middleware.ReuqireAuth, controllers.GetAuth)
	r.GET("/users/:id",middleware.ReuqireAuth, controllers.GetUser)
	r.GET("/users",middleware.ReuqireAuth, controllers.GetUsers)
}
