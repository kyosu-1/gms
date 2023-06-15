package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/gorilla/sessions"

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
	store := sessions.NewCookieStore([]byte(cfg.Session.Secret))

	e := echo.New()
	h := handler.NewHandlers(
		rdb.NewUserRepository(db),
		rdb.NewCategoryRepository(db),
		rdb.NewItemRepository(db),
		store,
	)

	api.RegisterHandlers(e, h)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Server.Port)))
}
