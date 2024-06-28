package middlewares

import (
	"fmt"
	"slices"
	"time"

	"github.com/labstack/echo/v4"
)

func NoCacheMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			c.Response().Header().Set("Pragma", "no-cache")
			c.Response().Header().Set("Expires", "0")
			return next(c)
		}
	}
}

// 指定した秒数キャッシュするミドルウェア
func CacheMiddleware(duration int) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Cache-Control", fmt.Sprintf("max-age=%d", duration))
			return next(c)
		}
	}
}

// 指定した分にキャッシュがクリアされるように Cache-Control ヘッダを設定するミドルウェア
func TimeCacheMiddleware(minutes []int) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			now := time.Now()
			nowMinute := now.Minute()
			durations := make([]int, 0, len(minutes))
			for _, m := range minutes {
				durations = append(durations, m-nowMinute)
			}
			slices.Sort(durations)

			c.Response().Header().Set("Cache-Control", fmt.Sprintf("max-age=%d", durations[0]*60))
			return next(c)
		}
	}
}
