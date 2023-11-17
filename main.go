package main

import (
	"flag"
	"log"
	"log/slog"

	"github.com/Michaelpalacce/goip/pkg/goip/clients"
	"github.com/Michaelpalacce/goip/pkg/goip/logger"
	notifier "github.com/Michaelpalacce/goip/pkg/goip/notifiers"
	"github.com/Michaelpalacce/goip/pkg/goip/watcher"
)

func main() {
	// Configure Logging
	logger.ConfigureLogging()

	// Parse arguments
	var (
		providerInput string
		notifierInput string
		interval      int
		client        clients.Client
		notif         notifier.Notifier
		err           error
	)

	flag.StringVar(&providerInput, "provider", "cloudflare", "Which provider to use? Available: cloudflare,")
	flag.StringVar(&notifierInput, "notifier", "", "Which notifier to use? Available: webhook,")
	flag.IntVar(&interval, "interval", 15, "How often to check for updates?")
	flag.Parse()

	// Create provider
	slog.Info("Provider chosen", "provider", providerInput)

	if client, err = clients.CreateClientBasedOnInput(providerInput); err != nil {
		log.Fatalf("error creating the client for %s. Error was: %s", providerInput, err)
	}

	if notif, err = notifier.CreateNotifierBasedOnInput(notifierInput); err != nil {
		log.Fatalf("error creating the notifier for %s. Error was: %s", notifierInput, err)
	}

	watcher := watcher.Watcher{
		Client:   client,
		Notifier: notif,
	}

	watcher.Watch(interval)
}
