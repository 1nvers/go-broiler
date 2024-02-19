package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/oneaushaf/go-broiler/helpers"
	"net/http"
)

func ReuqireAuth(c *gin.Context) {
	tokenString := c.GetHeader("token")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": tokenString})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, err := helpers.ValidateToken(tokenString)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Set("user_id", claims.UserID)
	c.Set("user_type", claims.UserType)
	c.Next()
}
