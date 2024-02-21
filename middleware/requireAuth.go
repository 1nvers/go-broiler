package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/oneaushaf/go-broiler/helpers"
	"net/http"
)

func ReuqireAuth(c *gin.Context) {
	tokenString := c.GetHeader("token")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no authorization token detected",
		})
		return
	}

	claims, err := helpers.ValidateToken(tokenString)

	if err != nil {
		c.JSON(http.StatusUnauthorized,gin.H{
			"error" : "invalid token",
		})
		return
	}

	c.Set("user_id", claims.UserID)
	c.Set("user_type", claims.UserType)
	c.Next()
}
