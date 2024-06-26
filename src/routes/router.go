package routes

import (
	"log/slog"

	"github.com/cateiru/zatsunen/src/config"
	"github.com/cateiru/zatsunen/src/handler"
	"github.com/cateiru/zatsunen/src/middlewares"
	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux, c *config.Config, l *slog.Logger) {
	h := handler.NewHandler(c, l)

	r.Use(middlewares.Logger(h.Logger))

	r.Get("/", h.RootHandler)
}
