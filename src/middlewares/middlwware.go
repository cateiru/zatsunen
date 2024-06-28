package middlewares

import (
	"fmt"
	"log/slog"
	"time"

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
				message := fmt.Sprintf("%s %d(%s) %s", v.Method, v.Status, v.URIPath)

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
