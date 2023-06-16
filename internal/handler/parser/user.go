package parser

import (
	"github.com/kyosu-1/ims/gen/api"
	"github.com/kyosu-1/ims/internal/model"
)

func (m *Model) User(user *model.User) *api.User {
	return &api.User{
		Id:   user.ID,
		Name: user.Name,
	}
}
