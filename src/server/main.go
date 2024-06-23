package server

import (
	"github.com/cateiru/zatsunen/src/config"
	"github.com/cateiru/zatsunen/src/log"
)

func RunServer(mode string, path string) {
	configPerEnv := config.GetConfig(mode, path)
	logger := log.SetupLogger(configPerEnv.GetConfig().LogConfig)

	logger.Info(configPerEnv.GetMode())
}
