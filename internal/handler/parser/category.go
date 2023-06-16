package parser

import (
	"github.com/kyosu-1/ims/gen/api"
	"github.com/kyosu-1/ims/internal/model"
)

type Model struct{}

type Api struct{}

func (m *Model) Category(category *model.Category) *api.Category {
	return &api.Category{
		Id:   category.ID,
		Name: category.Name,
	}
}

func (a *Api) Category(category *api.Category) *model.Category {
	return &model.Category{
		ID:   category.Id,
		Name: category.Name,
	}
}
