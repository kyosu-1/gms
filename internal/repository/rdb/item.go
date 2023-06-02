package rdb

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/kyosu-1/ims/internal/model"
)

type ItemRepository struct {
	db *sqlx.DB
}

func NewItemRepository(db *sqlx.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) FindAllItems(ctx context.Context) ([]*model.Item, error) {
	return nil, nil
}

func (r *ItemRepository) FindItemByID(ctx context.Context, id string) (*model.Item, error) {
	return nil, nil
}

func (r *ItemRepository) SaveItem(ctx context.Context, item *model.Item) error {
	return nil
}

func (r *ItemRepository) DeleteItem(ctx context.Context, id string) error {
	return nil
}
