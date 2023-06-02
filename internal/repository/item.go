package repository

import (
	"github.com/kyosu-1/ims/internal/model"
)

type ItemRepository interface {
	FindAllItems() ([]*model.Item, error)
	FindItemByID(id int) (*model.Item, error)
	SaveItem(item *model.Item) error
	DeleteItem(id string) error
}