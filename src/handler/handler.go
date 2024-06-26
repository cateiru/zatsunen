package handler

import (
	"log/slog"

	"github.com/cateiru/zatsunen/src/config"
)

type Handler struct {
	Logger *slog.Logger
	Config *config.Config
}

func NewHandler(config *config.Config, logger *slog.Logger) *Handler {
	return &Handler{
		Logger: logger,
		Config: config,
	}
}
