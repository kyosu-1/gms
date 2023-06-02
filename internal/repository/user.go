package repository

import (
	"context"

	"github.com/kyosu-1/ims/internal/model"
)

type UserRepository interface {
	FindUserByID(ctx context.Context, id string) (*model.User, error)
}
