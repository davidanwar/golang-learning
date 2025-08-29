package service

import (
	"errors"
	"golang-unit-testing/entity"
	"golang-unit-testing/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (categoryService CategoryService) Get(id string) (*entity.Category, error) {
	category := categoryService.Repository.FindById(id)

	if category == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}
