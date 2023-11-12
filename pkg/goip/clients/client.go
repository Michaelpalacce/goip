package clients

import (
	"fmt"
)

// Client is a general interface implemented by all clients
type Client interface {
	Auth() error
	SetIp(ip string) error
	CheckEnv() error
}

// CreateClientBasedOnProvider will return an authenticated, fully loaded client
func CreateClientBasedOnProvider(provider string) (Client, error) {
	var client Client
	switch provider {
	case "cloudflare":
		client = &Cloudflare{}
	default:
		return nil, fmt.Errorf("could not create a provider of type: %s", provider)
	}

	if err := client.CheckEnv(); err != nil {
		return nil, fmt.Errorf("error while validating provider (%s) environment: %s", provider, err)
	}

	if err := client.Auth(); err != nil {
		return nil, fmt.Errorf("error while trying to auth: %s", err)
	}

	return client, nil
}
