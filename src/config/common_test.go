package config_test

import (
	"testing"

	"github.com/cateiru/zatsunen/src/config"
	"github.com/gkampitakis/go-snaps/snaps"
)

func TestGetCommonConfig(t *testing.T) {
	commonConfig := config.GetCommonConfig()

	snaps.MatchSnapshot(t, commonConfig)
}
