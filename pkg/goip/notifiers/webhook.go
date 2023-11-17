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
	err := w.sendToWebhook(fmt.Sprintf("IP change detected. New IP set: %s", ip))

	if err != nil {
		return err
	}

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

// Auth sends a welcome message
func (w *WebhookNotifier) Auth() error {
	err := w.sendToWebhook("`goip` is starting its watch")

	if err != nil {
		return err
	}

	return nil
}

// sendToWebhook will send the given data to the webhook
func (w *WebhookNotifier) sendToWebhook(data string) error {
	webhookURL := os.Getenv("WEBHOOK_URL")

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBufferString(data))

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
