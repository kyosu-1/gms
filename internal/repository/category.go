package repository

import (
	"context"

	"github.com/kyosu-1/ims/internal/model"
)

type CategoryRepository interface {
	FindAllCategories(ctx context.Context) ([]*model.Category, error)
	FindCategoryByID(ctx context.Context, id string) (*model.Category, error)
	SaveCategory(ctx context.Context, category *model.Category) error
	DeleteCategory(ctx context.Context, id string) error
}
