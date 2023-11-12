package main

import (
	"flag"
	"log"
	"log/slog"

	"github.com/Michaelpalacce/goip/internal/goip"
	"github.com/Michaelpalacce/goip/pkg/goip/clients"
	"github.com/Michaelpalacce/goip/pkg/goip/clients/cloudflare"
)

func main() {
	goip.ConfigureLogging()

	var (
		provider string
	)

	flag.StringVar(&provider, "provider", "cloudflare", "Which provider to use? Available: cloudflare,")
	flag.Parse()

	slog.Info("Provider chosen", "provider", provider)

	var client clients.Client

	switch provider {
	case "cloudflare":
		client = &cloudflare.Cloudflare{}
	default:
		log.Fatalf("could not create a provider of type: %s", provider)
	}

	if err := client.CheckEnv(); err != nil {
		log.Fatalf("error while validating provider (%s) environment: %s", provider, err)
	}

	if err := client.Auth(); err != nil {
		log.Fatalf("error while trying to auth: %s", err)
	}

	var (
		publicIp []byte
		err      error
	)

	if publicIp, err = goip.GetPublicIp(); err != nil {
		log.Fatalf("Error while trying to fetch public IP: %s", err)
	}

	slog.Info("Fetched public IP from https://icanhazip.com", "publicIp", string(publicIp))

	if err := client.SetIp(string(publicIp)); err != nil {
		log.Fatalf("Error while setting ip: %s", err)
	}
}
