package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"main/src/domain/models"
)

type NotificationService struct{}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (s *NotificationService) ProcessNotification(notification models.Notification) {
	// Validar que el mensaje tiene destino
	if notification.DestinationID == "" {
		log.Println("Error: DestinationID está vacío, no se enviará el mensaje")
		return
	}

	log.Printf("Processing notification for %s: %s", notification.DestinationID, notification.Message)

	// Construir el payload con la estructura esperada por el WebSocket
	payload := map[string]interface{}{
		"sender":        "API_REST",
		"DestinationID": notification.DestinationID,
		"content":       notification.Message,  // Aquí aseguramos que se envíe el mensaje
		"time":          time.Now().Format(time.RFC3339), // Agregamos la hora del mensaje
	}

	// Convertir a JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error encoding JSON: %v", err)
		return
	}

	// Enviar la notificación al WebSocket vía HTTP
	resp, err := http.Post("http://localhost:4000/send-message", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error sending notification to WebSocket: %v", err)
		return
	}
	defer resp.Body.Close()

	log.Println("Notification sent to WebSocket successfully")
}
