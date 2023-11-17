package watcher

import (
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/Michaelpalacce/goip/pkg/goip/clients"
	"github.com/Michaelpalacce/goip/pkg/goip/network"
	notifier "github.com/Michaelpalacce/goip/pkg/goip/notifiers"
)

type Watcher struct {
	ip      string
	Client  clients.Client
    Notifier notifier.Notifier
}

// Watch will watch every interval minutes for a change in the in memory ip address and if a change is found, it will update the cloudflare record
func (w *Watcher) Watch(interval int) {
	if w.Client == nil {
		log.Fatal("watcher started, but Client not passed")
	}

	for {
		// Fetch public IP
		var (
			publicIp []byte
			err      error
		)

		if publicIp, err = network.GetPublicIp(); err != nil {
			log.Fatalf("error while trying to fetch public IP: %s", err)
		}

		slog.Info("Fetched public IP from https://icanhazip.com", "publicIp", string(publicIp))

		ipToSet := string(publicIp)

		// Checks if an update is needed
		if w.ip != ipToSet {
			slog.Debug("IP change detected. Setting new IP", "newIP", ipToSet)
			w.ip = ipToSet

			go func() {
				if err := w.Client.SetIp(w.ip); err != nil {
					slog.Error("error while trying to set IP", "IP", ipToSet)
				}
			}()

			go func() {
				webhook := os.Getenv("WEBHOOK_URL")
				if webhook != "" {
					// Post to the webhook
				}
			}()
		} else {
			slog.Debug("No IP change detected")
		}

		time.Sleep(time.Duration(interval) * time.Minute)
	}
}
