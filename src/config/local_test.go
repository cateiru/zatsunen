package config_test

import (
	"testing"

	"github.com/cateiru/zatsunen/src/config"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/require"
)

func TestLocalConfig(t *testing.T) {
	local := config.SetLocalConfig("/path")

	snaps.MatchSnapshot(t, local)

	require.Equal(t, local.GetMode(), "local")
	snaps.MatchSnapshot(t, local.GetConfig())
}
