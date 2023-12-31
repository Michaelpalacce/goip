# goip

A multi provider dns ip address updater.

## Inspiration

My ISP has a fee for static IPs. I don't want to pay that fee

## What Is It?

This tool will query `https://icanhazip.com` and listen for changes. If any changes are detected, it will set the IP in the chosen provider.
The configuration file for the providers must be stored either in `/app/config.json`, or project root `./config.json`

## Roadmap

- [ ] More providers?
- [x] Container Image
- [x] More details how to run this
- [x] Notifications
- [x] Fallback to `https://ifconfig.me/ip`

## Getting Started

### Docker

```bash
docker run -e CLOUDFLARE_API_TOKEN={{TOKEN}} -v ./config.json:/app/config.json stefangenov/goip
```

### From Source

```bash
go run main.go
```
> This will start `goip` with the default provider `cloudflare` and check for updates every 15 minutes.

## Notifications

### Webhook

Currently only webhook notifications are supported.

| Name | Value | Description |
|---|---|---|
| WEBHOOK_URL | - | The webhook to which to post when an update happens |

## Providers

Providers are what tells goip how to handle the change in IP address.

### Cloudflare

Cloudflare provider gives you the ability to work with multiple zones at the same time as well as multiple records in each zone. 
The provider also gives you the ability to partially configure parameters of the created record.

#### Environment

The following env variables must be present:

- `CLOUDFLARE_API_TOKEN`: This must be a token that has `Zone.DNS` permissions

#### Configuration

```json
{
    "cloudflare": {
        "zones": [
            {
                "name": "mywebsite.com",
                "records": [
                    {
                        "name": "mywebsite.com",
                        "proxied": true
                    },
                    {
                        "name": "subdomain.mywebsite.com",
                        "proxied": true
                    }
                ]
            }
        ]
    }
}

```
