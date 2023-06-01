package main

import (
	"fmt"
	"github.com/labstack/echo/v4"

	"github.com/kyosu-1/ims/gen/api"
	"github.com/kyosu-1/ims/internal/config"
	"github.com/kyosu-1/ims/internal/handler"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	h := handler.NewHandlers()

	api.RegisterHandlers(e, h)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Server.Port)))
}
