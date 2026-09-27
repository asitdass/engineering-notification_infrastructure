package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type CreateNotificationRequest struct {
	Recipient string `json:"recipient"`
	Channel string `json:"channel"`
	Template string `json:"template"`
}

func CreateNotification(w http.ResponseWriter, r *http.Request) {
	var request CreateNotificationRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Reciepient: %s\nChannel: %s\nTemplate: %s\n", 
	request.Recipient, 
	request.Channel, 
	request.Template,
)
}