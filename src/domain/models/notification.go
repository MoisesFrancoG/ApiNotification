package models

type Notification struct {
	Sender        string `json:"sender"`        // <- Quién envía el mensaje
	DestinationID string `json:"destinationID"` // <- A quién va dirigido
	Message       string `json:"message"`  
}
