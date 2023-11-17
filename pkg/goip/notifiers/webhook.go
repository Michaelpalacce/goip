package notifier

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

type WebhookNotifier struct {
}

// Sends a notification about the IP change
func (w *WebhookNotifier) Notify(ip string) error {
	webhookURL := os.Getenv("WEBHOOK_URL")

	data := fmt.Sprintf("IP change detected. New IP set: %s", ip)

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBufferString(data))

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// CheckEnv validates that everything needed is present
func (w *WebhookNotifier) CheckEnv() error {
	tokens := []string{"WEBHOOK_URL"}

	for _, token := range tokens {
		_, exists := os.LookupEnv(token)

		if !exists {
			return fmt.Errorf("%s not set", token)
		}
	}

	return nil
}

func (w *WebhookNotifier) Auth() error {
	return nil
}
