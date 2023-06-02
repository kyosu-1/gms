package repository

import (
	"context"

	"github.com/kyosu-1/ims/internal/model"
)

type ItemRepository interface {
	FindAllItems(ctx context.Context) ([]*model.Item, error)
	FindItemByID(ctx context.Context, id int) (*model.Item, error)
	SaveItem(ctx context.Context, item *model.Item) error
	DeleteItem(ctx context.Context, id string) error
}
