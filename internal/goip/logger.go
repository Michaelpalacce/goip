package goip

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

// ConfigureLogging will set a default logger
func ConfigureLogging() {
	w := os.Stderr

	// set global logger with custom options
	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))

}
