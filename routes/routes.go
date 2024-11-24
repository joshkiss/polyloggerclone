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
	auth.POST("/entries", createEntry)
	auth.PUT("/entries/:id", placeholder)
	auth.DELETE("/entries/:id, placeholder")

	auth.GET("/stats", placeholder) // get stats for user
	auth.GET("/stats/:lang", placeholder)
}

func placeholder(ctx *gin.Context) {}
