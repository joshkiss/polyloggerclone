package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(srv *gin.Engine) {
	srv.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the PolyLoggerClone API"})
	})
	srv.POST("/login", login)
	srv.POST("/register", signup)
}
