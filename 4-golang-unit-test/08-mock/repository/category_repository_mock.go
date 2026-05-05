package repository

import (
	"08-mock/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

type categoryRepositoryMock struct{}

func NewCategoryRepositoryMock() *CategoryRepositoryMock {
	return &CategoryRepositoryMock{}
}

func (r *CategoryRepositoryMock) GetAll() ([]entity.Category, error) {
	args := r.Mock.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	categories := args.Get(0).([]entity.Category)
	return categories, args.Error(1)
}

func (r *CategoryRepositoryMock) FindByID(id int) (*entity.Category, error) {
	args := r.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	categories := args.Get(0).(*entity.Category)
	return categories, args.Error(1)
}
