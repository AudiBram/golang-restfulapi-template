package repository

import (
	"context"
	"database/sql"
	"golang-restfulapi/model/domain"
)

type CategoryRepository interface {
	Save(c context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(c context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(c context.Context, tx *sql.Tx, category domain.Category)
	FindById(c context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(c context.Context, tx *sql.Tx) []domain.Category
}
