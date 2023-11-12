package network

import (
	"fmt"
	"io"
	"net/http"
)

func GetBody(url string) ([]byte, error) {
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()

		if body, err := io.ReadAll(resp.Body); err == nil {
			return body, nil
		}
	}

	return nil, fmt.Errorf("http: Error while trying to fetch url: %s", url)
}
