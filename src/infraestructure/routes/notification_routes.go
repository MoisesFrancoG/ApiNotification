package routes

import (
	"main/src/application/services"
	"main/src/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func NotificationRoutes(router *gin.Engine) {
	service := services.NewNotificationService()
	controller := controllers.NewNotificationController(service)

	router.POST("/notifications", controller.ReceiveNotification)
}
