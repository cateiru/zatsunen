package config_test

import (
	"testing"

	"github.com/cateiru/zatsunen/src/config"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/require"
)

func TestProductionConfig(t *testing.T) {
	production := config.SetProductionConfig("/path")

	snaps.MatchSnapshot(t, production)

	require.Equal(t, production.GetMode(), "production")
	snaps.MatchSnapshot(t, production.GetConfig())
}
