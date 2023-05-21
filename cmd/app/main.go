package main

import (
	"github.com/labstack/echo/v4"

	"github.com/kyosu-1/ims/gen/api"
	"github.com/kyosu-1/ims/internal/handler"
)

func main() {
	e := echo.New()
	h := handler.NewHandlers()

	api.RegisterHandlers(e, h)

	e.Logger.Fatal(e.Start(":8080"))
}