package env

import "os"

// GetWebhook will fetch the webhook url from the environment. This can be an empty string
func GetWebhook() string {
    return os.Getenv("WEBHOOK_URL")
}
