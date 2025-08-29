package repository

import "golang-unit-testing/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
