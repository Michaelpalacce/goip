# goip

A multi provider dns ip address updater.

## Inspiration

My ISP has a fee for static IPs. I don't want to pay that fee

## What Is It?

This tool will query `https://icanhazip.com` and listen for changes. If any changes are detected, it will set the IP in the chosen provider.
The configuration file for the providers must be stored either in `/app/config.json`, or project root `./config.json`

## Roadmap

- [ ] More providers?
- [ ] Container Image
- [ ] More details how to run this

## Getting Started

```bash
go run main.go
```
> This will start `goip` with the default provider `cloudflare` and check for updates every 15 minutes.

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
                "name": "stefangenov.site",
                "records": [
                    {
                        "name": "stefangenov.site",
                        "proxied": true
                    },
                    {
                        "name": "*-public.stefangenov.site",
                        "proxied": true
                    }
                ]
            }
        ]
    }
}
```
