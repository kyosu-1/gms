package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handlers) GetHealth(ctx echo.Context) error {
	return ctx.JSON(200, map[string]string{"status": "OK"})
}
