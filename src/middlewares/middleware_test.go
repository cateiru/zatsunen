package middlewares_test

import (
	"net/http"
	"testing"

	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/cateiru/zatsunen/src/config"
	"github.com/cateiru/zatsunen/src/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestCsrfMiddleware(t *testing.T) {
	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	}

	t.Run("成功: Same-Origin", func(t *testing.T) {
		config := config.MiddlewareConfig{
			AllowSecFetchSiteValues: []string{"Same-Origin"},
		}

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.R.Header.Set("Sec-Fetch-Site", "Same-Origin")

		c := m.Echo()

		middlewareFunc := middlewares.CsrfMiddleware(config)
		middlewareFunc(handler)(c)

		require.Equal(t, http.StatusOK, c.Response().Status)
	})

	t.Run("失敗: Same-Origin", func(t *testing.T) {
		config := config.MiddlewareConfig{
			AllowSecFetchSiteValues: []string{"Same-Origin"},
		}

		mock, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		c := mock.Echo()

		middlewareFunc := middlewares.CsrfMiddleware(config)
		middlewareFunc(handler)(c)

		require.Equal(t, http.StatusForbidden, c.Response().Status)
	})

	t.Run("AllowSecFetchSiteValues が空の場合はすべての値を許可", func(t *testing.T) {
		values := []string{
			"Cross-Site",
			"Same-Site",
			"Same-Origin",
			"None",
		}
		config := config.MiddlewareConfig{
			AllowSecFetchSiteValues: []string{},
		}

		for _, value := range values {
			t.Run(value, func(t *testing.T) {

				m, err := easy.NewMock("/", http.MethodGet, "")
				require.NoError(t, err)
				m.R.Header.Set("Sec-Fetch-Site", value)

				c := m.Echo()

				middlewareFunc := middlewares.CsrfMiddleware(config)
				middlewareFunc(handler)(c)

				require.Equal(t, http.StatusOK, c.Response().Status)
			})
		}
	})
}
