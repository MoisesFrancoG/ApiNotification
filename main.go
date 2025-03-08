package main

import (
	"log"
	"main/src/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.NotificationRoutes(router)


	log.Println("Server running on port 8081")
	router.Run(":8081")
}
