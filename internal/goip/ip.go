package goip

import "github.com/Michaelpalacce/goip/internal/goip/http"

func GetPublicIp() ([]byte, error) {
	return http.GetBody("https://icanhazip.com")
}
