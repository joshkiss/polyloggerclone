package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshkiss/polyloggerclone/middlewares"
)

func RegisterRoutes(srv *gin.Engine) {
	srv.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the PolyLoggerClone API"})
	})
	srv.POST("/login", login)
	srv.POST("/register", signup)

	auth := srv.Group("/api")
	auth.Use(middlewares.Authenticate)
	auth.POST("/entries")
	auth.PUT("/entries/:id")
	auth.DELETE("/entries/:id")

	auth.GET("/stats") // get stats for user
	auth.GET("/stats/:lang")
}
