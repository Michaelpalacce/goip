package main

import (
	"flag"
	"log"
	"log/slog"

	"github.com/Michaelpalacce/goip/pkg/goip/clients"
	"github.com/Michaelpalacce/goip/pkg/goip/logger"
	"github.com/Michaelpalacce/goip/pkg/goip/watcher"
)

func main() {
	// Configure Logging
	logger.ConfigureLogging()

	// Parse arguments
	var (
		provider string
		interval int
		client   clients.Client
		err      error
	)

	flag.StringVar(&provider, "provider", "cloudflare", "Which provider to use? Available: cloudflare,")
	flag.IntVar(&interval, "interval", 15, "How often to check for updates?")
	flag.Parse()

	// Create provider
	slog.Info("Provider chosen", "provider", provider)

	if client, err = clients.CreateClientBasedOnProvider(provider); err != nil {
		log.Fatalf("error creating the client for %s. Error was: %s", provider, err)
	}

	watcher := watcher.Watcher{
		Client: client,
	}

	watcher.Watch(interval)
}
