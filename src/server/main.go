package server

import (
	"net/http"

	"github.com/cateiru/zatsunen/src/config"
	"github.com/cateiru/zatsunen/src/log"
	"github.com/cateiru/zatsunen/src/routes"
	"github.com/go-chi/chi/v5"
)

func RunServer(mode string, path string) {
	configPerEnv := config.GetConfig(mode, path)
	c := configPerEnv.GetConfig()
	l := log.SetupLogger(c.LogConfig)

	l.Info(configPerEnv.GetMode())

	r := chi.NewRouter()
	routes.Routes(r, &c, l)

	http.ListenAndServe(":8080", r)
}
