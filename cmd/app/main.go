package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/jmoiron/sqlx"

	"github.com/kyosu-1/ims/gen/api"
	"github.com/kyosu-1/ims/internal/config"
	"github.com/kyosu-1/ims/internal/handler"
	"github.com/kyosu-1/ims/internal/repository/rdb"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	db, err := sqlx.Connect("mysql", cfg.DB.DSN())
	if err != nil {
		panic(err)
	}


	e := echo.New()
	h := handler.NewHandlers(
		rdb.NewUserRepository(db),
		rdb.NewCategoryRepository(db),
		rdb.NewItemRepository(db),
	)

	api.RegisterHandlers(e, h)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Server.Port)))
}
