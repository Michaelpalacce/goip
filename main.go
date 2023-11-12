package main

import (
	"flag"
	"log"
	"log/slog"

	"github.com/Michaelpalacce/goip/pkg/goip/clients"
	"github.com/Michaelpalacce/goip/pkg/goip/logger"
	"github.com/Michaelpalacce/goip/pkg/goip/network"
)

func main() {
	// Configure Logging
	logger.ConfigureLogging()

	// Parse arguments
	var (
		provider string
	)

	flag.StringVar(&provider, "provider", "cloudflare", "Which provider to use? Available: cloudflare,")
	flag.Parse()

	// Create provider
	slog.Info("Provider chosen", "provider", provider)

	var client clients.Client

	switch provider {
	case "cloudflare":
		client = &clients.Cloudflare{}
	default:
		log.Fatalf("could not create a provider of type: %s", provider)
	}

	if err := client.CheckEnv(); err != nil {
		log.Fatalf("error while validating provider (%s) environment: %s", provider, err)
	}

	if err := client.Auth(); err != nil {
		log.Fatalf("error while trying to auth: %s", err)
	}

	// Fetch public IP
	var (
		publicIp []byte
		err      error
	)

	if publicIp, err = network.GetPublicIp(); err != nil {
		log.Fatalf("Error while trying to fetch public IP: %s", err)
	}

	slog.Info("Fetched public IP from https://icanhazip.com", "publicIp", string(publicIp))

	// Set the fetched ip to the records
	if err := client.SetIp(string(publicIp)); err != nil {
		log.Fatalf("Error while setting ip: %s", err)
	}
}
