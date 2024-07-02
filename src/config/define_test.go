package config_test

import (
	"testing"

	"github.com/cateiru/zatsunen/src/config"
	"github.com/stretchr/testify/require"
)

func TestGetConfig(t *testing.T) {
	t.Run("local", func(t *testing.T) {
		environment := config.GetConfig("local", "/path")

		require.Equal(t, environment.GetMode(), "local")
	})

	t.Run("production", func(t *testing.T) {
		environment := config.GetConfig("production", "/path")

		require.Equal(t, environment.GetMode(), "production")
	})
}
