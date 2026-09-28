package main

import (
	"fmt"
)

func validateCreateNotificationRequest(request CreateNotificationRequest) error {
	if request.Recipient == "" {
		return fmt.Errorf("recipient is required")
	}
	if request.Channel == "" {
		return fmt.Errorf("channel is required")
	}
	if request.Channel != "email" && request.Channel != "sms" && request.Channel != "push" {
		return fmt.Errorf("invalid channel")
	}
	if request.Template == "" {
		return fmt.Errorf("template is required")
	}
	return nil
}