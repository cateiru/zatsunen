package log

import (
	"log/slog"
	"os"

	"github.com/cateiru/zatsunen/src/config"
)

func SetupLogger(c config.LogConfig) *slog.Logger {
	jsonHandler := slog.NewJSONHandler(os.Stdout, c.Options)
	logger := slog.New(jsonHandler)

	return logger
}
