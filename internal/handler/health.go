package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handlers) GetHealth(ec echo.Context) error {
	return ec.JSON(200, map[string]string{"status": "OK"})
}
