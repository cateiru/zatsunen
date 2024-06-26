package middlewares

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func Logger(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fh := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			start := time.Now()

			defer func() {
				ctx := r.Context()
				end := time.Now()
				status := ww.Status()

				attrs := []slog.Attr{
					slog.Time("time", time.Now()),
					slog.Duration("duration", end.Sub(start)),

					slog.String("method", r.Method),
					slog.Int("status", ww.Status()),
					slog.String("host", r.Host),
					slog.String("path", r.URL.Path),
					slog.String("query", r.URL.RawQuery),
					slog.String("ip", r.RemoteAddr),
				}

				logLevel := slog.LevelInfo
				if status >= 500 && status < 500 {
					logLevel = slog.LevelError
				}

				logger.LogAttrs(ctx, logLevel, "request", attrs...)
			}()

			next.ServeHTTP(w, r)

		}

		return http.HandlerFunc(fh)
	}
}
