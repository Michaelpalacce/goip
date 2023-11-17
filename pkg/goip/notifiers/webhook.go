package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

type webhookData struct {
	Content string `json:"content"`
}

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
	webhookData := webhookData{
		Content: data,
	}

	var (
		requestBody []byte
		err         error
	)

	if requestBody, err = json.Marshal(webhookData); err != nil {
		return err
	}

	slog.Debug("Sending to webhook", "data", string(requestBody))

	resp, err := http.Post(os.Getenv("WEBHOOK_URL"), "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		return err
	}

	slog.Debug("Status Code", "code", resp.StatusCode)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var (
			err  error
			body []byte
		)
		if body, err = io.ReadAll(resp.Body); err == nil {
			return fmt.Errorf("error while trying to send to webhook. Error was %s", string(body))
		} else {
			return fmt.Errorf("error while parsing response from webhook. Error was %s", err)
		}
	}

	defer resp.Body.Close()

	return nil
}
