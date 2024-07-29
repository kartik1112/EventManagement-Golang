package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kartik1112/EventManagement-Golang/db"
	"github.com/kartik1112/EventManagement-Golang/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
