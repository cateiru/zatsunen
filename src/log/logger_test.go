package log_test

import (
	"log/slog"
	"testing"

	"github.com/cateiru/zatsunen/src/config"
	"github.com/cateiru/zatsunen/src/log"
	"github.com/stretchr/testify/require"
)

func TestSetLogger(t *testing.T) {
	t.Run("json log", func(t *testing.T) {
		logConfig := config.LogConfig{
			Options: &slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelInfo,
				ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
					return a
				},
			},
			IsJsonLog: true,
		}

		logger := log.SetupLogger(logConfig)
		require.NotNil(t, logger)

		_, isJsonLog := logger.Handler().(*slog.JSONHandler)
		require.True(t, isJsonLog)
	})

	t.Run("text log", func(t *testing.T) {
		logConfig := config.LogConfig{
			Options: &slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelInfo,
				ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
					return a
				},
			},
			IsJsonLog: false,
		}

		logger := log.SetupLogger(logConfig)
		require.NotNil(t, logger)

		_, textLog := logger.Handler().(*slog.JSONHandler)
		require.False(t, textLog)
	})
}
