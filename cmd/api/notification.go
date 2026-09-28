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
	if(r.Method != http.MethodPost){
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var request CreateNotificationRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	err = validateCreateNotificationRequest(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Reciepient: %s\nChannel: %s\nTemplate: %s\n", 
	request.Recipient, 
	request.Channel, 
	request.Template,
)
}