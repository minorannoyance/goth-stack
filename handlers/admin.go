package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func PrivateHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "Secret")
}
