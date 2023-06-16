package rdb

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/kyosu-1/ims/internal/model"
)

const (
	categoryTable = "categories"
)

type Category struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) FindAllCategories(ctx context.Context) ([]*model.Category, error) {
	categories := []*model.Category{}
	query, args, err := squirrel.Select("*").From(categoryTable).ToSql()
	if err != nil {
		return nil, err
	}
	if err := r.db.Select(&categories, query, args...); err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) FindCategoryByID(ctx context.Context, id string) (*model.Category, error) {
	category := &Category{}
	query, args, err := squirrel.Select("*").From(categoryTable).Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return nil, err
	}

	if err := r.db.Get(category, query, args...); err != nil {
		return nil, err
	}

	return &model.Category{
		ID:   category.ID,
		Name: category.Name,
	}, nil
}

func (r *CategoryRepository) SaveCategory(ctx context.Context, category *model.Category) error {
	txn, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	query, args, err := squirrel.Insert(categoryTable).Columns("id", "name").Values(category.ID, category.Name).ToSql()
	if err != nil {
		return err
	}

	stmt, err := txn.Preparex(query)
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(args...); err != nil {
		return err
	}

	return txn.Commit()
}

func (r *CategoryRepository) DeleteCategory(ctx context.Context, id string) error {
	txn, err := r.db.Beginx()
	if err != nil {
		return err
	}

	query, args, err := squirrel.Delete(categoryTable).Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}

	stmt, err := txn.Preparex(query)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(args...); err != nil {
		return err
	}

	return txn.Commit()
}
