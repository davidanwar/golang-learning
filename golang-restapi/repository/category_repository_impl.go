package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restapi/helper"
	"golang-restapi/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	script := "INSERT INTO category(name) VALUES(?)"
	result, err := tx.ExecContext(ctx, script, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	script := "UPDATE category SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, category.Name, category.Id)
	helper.PanicIfError(err)
	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	script := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error) {
	script := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, script, id)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, err
	} else {
		return category, errors.New("Category Not Found")
	}

}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	script := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, script)
	helper.PanicIfError(err)
	defer rows.Close()

	categories := []domain.Category{}
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
