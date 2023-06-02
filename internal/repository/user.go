package repository

import (
	"github.com/kyosu-1/ims/internal/model"
)

type UserRepository interface {
	FindUserByID(id string) (*model.User, error)
}
