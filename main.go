package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/joshkiss/polyloggerclone/db"
	"github.com/joshkiss/polyloggerclone/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	routes.RegisterRoutes(server)
	server.Run(":8080")
}
