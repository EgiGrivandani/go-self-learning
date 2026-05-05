package service

import (
	"08-mock/entity"
	"08-mock/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service *CategoryService) GetAll() ([]entity.Category, error) {
	category, err := service.Repository.FindAll()
	if err == nil {
		return nil, errors.New("category not found")
	}
	return category, nil
}
