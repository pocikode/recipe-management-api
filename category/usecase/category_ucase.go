package usecase

import (
	"recipe-management/domain"
	"recipe-management/pkg"
	customError "recipe-management/pkg/error"
)

type categoryUseCase struct {
	categoryRepository domain.CategoryRepository
}

func NewCategoryUseCase(cr domain.CategoryRepository) domain.CategoryUseCase {
	return &categoryUseCase{cr}
}

func (cu categoryUseCase) Fetch(pagination *pkg.Pagination) (*pkg.Pagination, error) {
	pagination, err := cu.categoryRepository.Fetch(*pagination)
	if err != nil {
		return nil, err
	}

	return pagination, nil
}

func (cu categoryUseCase) Create(c *domain.Category) (*domain.Category, error) {
	if c.Name == "" {
		return nil, customError.NewError(400, "Name is required")
	}
	c, err := cu.categoryRepository.Create(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
