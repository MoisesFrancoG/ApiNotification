package main

import (
	"log"
	"main/src/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Definir rutas
	routes.NotificationRoutes(router)

	// Iniciar servidor en el puerto 8081

	log.Println("Server running on port 8081")
	router.Run(":8081")
}
