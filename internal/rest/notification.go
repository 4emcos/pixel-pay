package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type NotificationResponse struct {
	Notification bool `json:"notification"`
}

var (
	notificationUrl = os.Getenv("NOTIFICATION_HOST")
)

func Notification() error {
	if notificationUrl == "" {
		notificationUrl = "localhost"
	}
	resp, err := http.Get("http://" + notificationUrl + ":8099/notification")

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
