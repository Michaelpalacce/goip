package env

import "os"

func GetWebhook() string {
    return os.Getenv("WEBHOOK_URL")
}
