package main

import (
	"example.com/restapi-dev/db"
	"example.com/restapi-dev/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
