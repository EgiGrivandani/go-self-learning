package repository

import "08-mock/entity"

type CategoryRepository interface {
	FindAll() ([]entity.Category, error)
	FindByID(id int) (*entity.Category, error)
}

type categoryRepository struct{}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

func (r *categoryRepository) FindAll() ([]entity.Category, error) {
	var categories []entity.Category

	categories = append(categories, entity.Category{ID: 1, Name: "Laptop"})
	return categories, nil
}

func (r *categoryRepository) FindByID(id int) (*entity.Category, error) {
	category := &entity.Category{ID: id, Name: "Laptop"}
	return category, nil
}
