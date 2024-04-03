package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type NotificationResponse struct {
	Notification bool `json:"notification"`
}

func Notification() error {
	resp, err := http.Get("http://localhost:8099/notification")

	if err != nil {
		return fmt.Errorf("err call notification")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("err call notification")
	}

	var notificationResponse NotificationResponse
	err = json.Unmarshal(body, &notificationResponse)

	if notificationResponse.Notification {
		return nil
	}

	return fmt.Errorf("err call notification")

}
