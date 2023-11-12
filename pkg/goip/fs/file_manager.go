package fs

import (
	"io"
	"log/slog"
	"os"
)

func ReadJsonFile(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)

	if err != nil {
        return nil, err
	}

	slog.Info("Successfully opened", "file", path)

	defer jsonFile.Close()

	return io.ReadAll(jsonFile)
}
