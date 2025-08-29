// package repository

// import (
// 	"golang-unit-testing/entity"

// 	"github.com/stretchr/testify/mock"
// )

// type CategoryRepositoryMock struct {
// 	Mock mock.Mock
// }

// func (repositoryMock CategoryRepositoryMock) FindById(id string) *entity.Category {
// 	arguments := repositoryMock.Mock.Called(id)
// 	if arguments.Get(0) == nil {
// 		return nil
// 	} else {
// 		category := arguments.Get(0).(entity.Category)
// 		return &category
// 	}
// }
