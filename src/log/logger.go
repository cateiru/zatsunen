package log

import (
	"log/slog"
	"os"

	"github.com/cateiru/zatsunen/src/config"
	"github.com/lmittmann/tint"
)

func SetupLogger(c config.LogConfig) *slog.Logger {
	if c.IsJsonLog {
		jsonHandler := slog.NewJSONHandler(os.Stdout, c.Options)
		return slog.New(jsonHandler)
	}

	tintHandler := tint.NewHandler(os.Stdout, &tint.Options{
		AddSource:   c.Options.AddSource,
		Level:       c.Options.Level,
		ReplaceAttr: c.Options.ReplaceAttr,
	})

	return slog.New(tintHandler)
}
