package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Michaelpalacce/goip/pkg/goip/clients"
	"github.com/Michaelpalacce/goip/pkg/goip/clients/cloudflare"
)

func main() {
	var (
		provider string
	)

	flag.StringVar(&provider, "provider", "cloudflare", "Which provider to use? Available: cloudflare,")

	flag.Parse()

	fmt.Printf("Provider chosen to be used: %s", provider)

	var client clients.Client

	switch provider {
	case "cloudflare":
		client = &cloudflare.Cloudflare{}
    default:
        log.Fatalf("could not create a provider of type: %s", provider)
	}

	if err := client.Auth(); err != nil {
		log.Fatalf("error while trying to auth: %s", err)
	}
}
