package config_test

import (
	"testing"

	"github.com/cateiru/zatsunen/src/config"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/require"
)

func TestGetMode(t *testing.T) {
	local := config.SetLocalConfig("/path")

	require.Equal(t, local.GetMode(), "local")
}

func TestGetConfig(t *testing.T) {
	local := config.SetLocalConfig("/path")

	snaps.MatchSnapshot(t, local.GetConfig())
}

func TestSetLocalConfig(t *testing.T) {
	local := config.SetLocalConfig("/path")

	snaps.MatchSnapshot(t, local)
}
