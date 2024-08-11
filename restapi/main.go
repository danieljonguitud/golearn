package main

import (
	"danieljonguitud.com/restapi/db"
	"danieljonguitud.com/restapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	sdb.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
