package main

import (
	"fmt"
	"log"

	"github.com/Michaelpalacce/goip/pkg/utils/http"
)

func main() {
	url := "https://google.com"

	if body, err := http.GetBody(url); err == nil {
		fmt.Printf("Body: %s\n", body)
	} else {
		log.Fatalf("Error while trying to fetch url: %s, error was: %s", url, err)
	}
}
