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

// CreateClientBasedOnInput will return an authenticated, fully loaded client
func CreateClientBasedOnInput(input string) (Client, error) {
	var client Client
	switch input {
	case "cloudflare":
		client = &Cloudflare{}
	default:
		return nil, fmt.Errorf("could not create a provider of type: %s", input)
	}

	if err := client.CheckEnv(); err != nil {
		return nil, fmt.Errorf("error while validating provider (%s) environment: %s", input, err)
	}

	if err := client.Auth(); err != nil {
		return nil, fmt.Errorf("error while trying to auth: %s", err)
	}

	return client, nil
}
