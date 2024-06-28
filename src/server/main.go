package server

import (
	"github.com/cateiru/zatsunen/src/config"
	"github.com/cateiru/zatsunen/src/log"
	"github.com/labstack/echo/v4"
)

func RunServer(mode string, path string) {
	configPerEnv := config.GetConfig(mode, path)
	c := configPerEnv.GetConfig()
	l := log.SetupLogger(c.LogConfig)

	l.Info(configPerEnv.GetMode())

	e := echo.New()
	// routes.Routes(r, &c, l)

	if err := e.Start(":8080"); err != nil {
		l.Error(err.Error())
	}
}
