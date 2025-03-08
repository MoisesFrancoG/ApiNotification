package controllers

import (
	"log"
	"main/src/application/services"
	"main/src/domain/models"
	"main/src/infraestructure/rabbit"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	service *services.NotificationService
}

func NewNotificationController(service *services.NotificationService) *NotificationController {
	return &NotificationController{service: service}
}

func (c *NotificationController) ReceiveNotification(ctx *gin.Context) {
	var rawMessage map[string]string

	if err := ctx.ShouldBindJSON(&rawMessage); err != nil {
		log.Println("Invalid request payload")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	notification := models.Notification{
		Sender:        "API_REST",
		DestinationID: "123", 
		Message:       rawMessage["message"],
	}

	if sender, exists := rawMessage["sender"]; exists {
		notification.Sender = sender
	}
	if destination, exists := rawMessage["destinationID"]; exists {
		notification.DestinationID = destination
	}

	log.Printf("Processed Notification: %+v", notification)

	c.service.ProcessNotification(notification)
	rabbit.SendMessagetoQueue(notification.Message)
	ctx.JSON(http.StatusOK, gin.H{"status": "Notification processed"})
}
