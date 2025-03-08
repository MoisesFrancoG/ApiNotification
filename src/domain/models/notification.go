package models

type Notification struct {
	Sender        string `json:"sender"`        
	DestinationID string `json:"destinationID"` 
	Message       string `json:"message"`  
}
