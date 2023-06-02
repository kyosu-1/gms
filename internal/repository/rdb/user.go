package rdb

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/kyosu-1/ims/internal/model"
)

const (
	userTable = "users"
)

type User struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	user := &User{}
	query, args, err := squirrel.Select("*").From(userTable).Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return nil, err
	}
	if err := r.db.Get(user, query, args...); err != nil {
		return nil, err
	}
	return &model.User{
		ID:       user.ID,
		Name:     user.Name,
		Password: user.Password,
	}, nil
}
