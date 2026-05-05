package service

import (
	"08-mock/entity"
	"08-mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryService_GetByID(t *testing.T) {
	repoMock := repository.NewCategoryRepositoryMock()

	// WAJIB: define expected call
	repoMock.Mock.On("FindByID", 1).
		Return(&entity.Category{
			ID:   1,
			Name: "Elektronik",
		}, nil)

	result, err := repoMock.FindByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Elektronik", result.Name)

	repoMock.Mock.AssertExpectations(t)
}

func TestCategoryRepositoryMock_GetAll(t *testing.T) {
	repoMock := repository.NewCategoryRepositoryMock()

	mockData := []entity.Category{
		{ID: 1, Name: "Elektronik"},
		{ID: 2, Name: "Fashion"},
	}

	repoMock.Mock.On("GetAll").
		Return(mockData, nil)

	result, err := repoMock.GetAll()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Elektronik", result[0].Name)

	repoMock.Mock.AssertExpectations(t)
}
