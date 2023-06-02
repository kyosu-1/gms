package repository

import (
	"github.com/kyosu-1/ims/internal/model"
)

type CategoryRepository interface {
	FindAllCategories() ([]*model.Category, error)
	FindCategoryByID(id string) (*model.Category, error)
	SaveCategory(category *model.Category) error
	DeleteCategory(id string) error
}
