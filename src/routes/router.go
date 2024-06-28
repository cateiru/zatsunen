package routes

import (
	"log/slog"

	"github.com/cateiru/zatsunen/src/config"
	"github.com/cateiru/zatsunen/src/handler"
	"github.com/cateiru/zatsunen/src/middlewares"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, c *config.Config, l *slog.Logger) {
	h := handler.NewHandler(c, l)

	e.Use(middlewares.LoggerMiddleware(h.Logger))
	e.Use(middlewares.CsrfMiddleware(c.MiddlewareConfig))
	e.Use(middlewares.GzipMiddleware())

	e.GET("/", h.RootHandler)
}
