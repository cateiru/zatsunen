package middlewares

import (
	"fmt"
	"log/slog"
	"slices"
	"time"

	"github.com/cateiru/zatsunen/src/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoggerMiddleware(logger *slog.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(
		middleware.RequestLoggerConfig{
			LogURI:      true,
			LogStatus:   true,
			LogMethod:   true,
			LogHost:     true,
			LogError:    true,
			LogRemoteIP: true,

			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				ctx := c.Request().Context()
				attrs := []slog.Attr{
					slog.Time("time", time.Now()),
					slog.Duration("duration", time.Since(v.StartTime)),

					slog.String("method", v.Method),
					slog.Int("status", v.Status),
					slog.String("host", v.Host),
					slog.String("path", v.URIPath),
					slog.String("ip", v.RemoteIP),
				}
				message := fmt.Sprintf("%s %d %s", v.Method, v.Status, v.URIPath)

				logLevel := slog.LevelInfo
				if v.Error != nil {
					logLevel = slog.LevelError
				}

				logger.LogAttrs(ctx, logLevel, message, attrs...)
				return nil
			},
		},
	)
}

var GzipMiddleware = middleware.Gzip

func CsrfMiddleware(config config.MiddlewareConfig) echo.MiddlewareFunc {
	allowSecFetchSiteValues := config.AllowSecFetchSiteValues

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			secFetchSiteHeader := c.Request().Header.Get("Sec-Fetch-Site")
			if len(allowSecFetchSiteValues) == 0 || slices.Contains(allowSecFetchSiteValues, secFetchSiteHeader) {
				return next(c)
			}

			return c.String(403, "Forbidden")
		}
	}
}
