package main

import (
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

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

	e.Use(session.MiddlewareWithConfig(session.Config{
		Store: store,
		Skipper: func(ctx echo.Context) bool {
			return ctx.Path() == "/health" || ctx.Path() == "/signin"
		},
	}))

	api.RegisterHandlers(e, h)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Server.Port)))
}
