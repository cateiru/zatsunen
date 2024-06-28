package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) RootHandler(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("Hello, %s", h.Config.Mode))
}
