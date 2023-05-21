package handler

import (
	"github.com/labstack/echo/v4"
	
	"github.com/kyosu-1/ims/gen/api"
)

type Handlers struct{}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) GetHealth(ctx echo.Context) error {
	return ctx.JSON(200, map[string]string{"status": "OK"})
}

func (h *Handlers) GetCategories(ctx echo.Context) error {
	// TODO: implement
	return nil
}

func (h *Handlers) PostCategories(ctx echo.Context) error {
	// TODO: implement
	return nil
}

func (h *Handlers) DeleteCategoriesCategoryID(ctx echo.Context, categoryID api.CategoryID) error {
	// TODO: implement
	return nil
}

func (h *Handlers) GetCategoriesCategoryID(ctx echo.Context, categoryID api.CategoryID) error {
	// TODO: implement
	return nil
}

func (h *Handlers) GetItems(ctx echo.Context) error {
	// TODO: implement
	return nil
}

func (h *Handlers) PostItems(ctx echo.Context) error {
	// TODO: implement
	return nil
}

func (h *Handlers) DeleteItemsItemID(ctx echo.Context, itemID api.ItemID) error {
	// TODO: implement
	return nil
}

func (h *Handlers) GetItemsItemID(ctx echo.Context, itemID api.ItemID) error {
	// TODO: implement
	return nil
}

func (h *Handlers) PostItemsItemIDBorrow(ctx echo.Context, itemID api.ItemID) error {
	// TODO: implement
	return nil
}

func (h *Handlers) PostItemsItemIDReturn(ctx echo.Context, itemID api.ItemID) error {
	// TODO: implement
	return nil
}

func (h *Handlers) GetMe(ctx echo.Context) error {
	// TODO: implement
	return nil
}

func (h *Handlers) PostSignin(ctx echo.Context) error {
	// TODO: implement
	return nil
}

func (h *Handlers) PostSignout(ctx echo.Context) error {
	// TODO: implement
	return nil
}
