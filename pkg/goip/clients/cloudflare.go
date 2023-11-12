package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/Michaelpalacce/goip/pkg/goip/fs"
	"github.com/cloudflare/cloudflare-go"
)

// Record represents one Zone Record
type Record struct {
	Name    string `json:"name"`
	Proxied bool   `json:"proxied"`
}

// Zone (s) are how Cloudflare separates different DNS endpoints
type Zone struct {
	Name    string   `json:"name"`
	Records []Record `json:"records"`
}

// CloudflareConfig is the structure of the json config that is expected
type CloudflareConfig struct {
	Cloudflare struct {
		Zones []Zone `json:"zones"`
	} `json:"cloudflare"`
}

// Cloudflare is the Cloudflare client that will support Authentication and setting records
type Cloudflare struct {
	api *cloudflare.API
}

// CheckEnv validates that everything needed is present
func (c Cloudflare) CheckEnv() error {
	tokens := []string{"CLOUDFLARE_API_TOKEN"}

	for _, token := range tokens {
		_, exists := os.LookupEnv(token)

		if !exists {
			return fmt.Errorf("%s not set", token)
		}
	}

	return nil
}

// Auth to Cloudflare with the given token
func (c *Cloudflare) Auth() error {
	api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))

	if err != nil {
		return fmt.Errorf("could not authenticate to Cloudflare with the given token, error was: %s", err)
	}

	c.api = api
	return nil
}

// SetIp sets the IP for the given zones based on the configuration
func (c Cloudflare) SetIp(ip string) error {
	var (
		data []byte
		err  error
	)

	// @TODO Where to get this from?
	if data, err = fs.ReadJsonFile("./config.json"); err != nil {
		return err
	}

	var config CloudflareConfig

	if err := json.Unmarshal(data, &config); err != nil {
		return err
	}

	zones := config.Cloudflare.Zones
	for _, zone := range zones {
		slog.Debug("Setting IP for zone", "zone", zone.Name)
		if err := c.setIpForZone(ip, zone); err != nil {
			return err
		}
	}

	return nil
}

// setIpForZone sets the public ip for a specific zone
func (c Cloudflare) setIpForZone(ip string, zone Zone) error {
	zoneID, err := c.api.ZoneIDByName(zone.Name)
	if err != nil {
		return err
	}
	slog.Debug("Found zone", "zoneId", zoneID, "zoneName", zone.Name)

	for _, r := range zone.Records {
		slog.Debug("Setting IP for record", "record", r)
		c.setIpForRecord(ip, zoneID, r)
	}

	return nil
}

// setIpForRecord will update the specific record
func (c Cloudflare) setIpForRecord(ip string, zoneID string, record Record) error {
	records, _, err := c.api.ListDNSRecords(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.ListDNSRecordsParams{})
	if err != nil {
		return err
	}

	for _, r := range records {
		if r.Name == record.Name {
			slog.Info("Updating record", "recordName", record.Name)

			_, err := c.api.UpdateDNSRecord(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.UpdateDNSRecordParams{
				ID:      r.ID,
				Content: ip,
				Proxied: cloudflare.BoolPtr(record.Proxied),
			})

			if err != nil {
				return err
			}
		}
	}

	return nil
}
