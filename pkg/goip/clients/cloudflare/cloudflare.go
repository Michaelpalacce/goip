package cloudflare

import (
	"fmt"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

type Cloudflare struct {
	api *cloudflare.API
}

// Auth to Cloudflare with the given token
func (c *Cloudflare) Auth() error {
	token := os.Getenv("CLOUDFLARE_API_TOKEN")
	api, err := cloudflare.NewWithAPIToken(token)

	if err != nil {
		return fmt.Errorf("could not authenticate to Cloudflare with the given token, error was: %s", err)
	}

	c.api = api
	return nil
}
