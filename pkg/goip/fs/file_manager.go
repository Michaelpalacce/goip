package fs

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

// Reads a file and returns the content as well as an error if any
func ReadJsonFile(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	slog.Info("Successfully opened", "file", path)

	defer jsonFile.Close()

	return io.ReadAll(jsonFile)
}

// ReadConfigFile will check expected locations for a config.json file
func ReadConfigFile() (data []byte, err error) {
	possibleLocations := []string{"/app/config.json", "./config.json"}

	for _, location := range possibleLocations {
		data, _ = ReadJsonFile(location)

		if data != nil {
			return data, nil
		}
	}

	return nil, fmt.Errorf("could not find any configuration at expected locations %+v", possibleLocations)
}
