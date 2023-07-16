package usecase

import (
	"recipe-management/domain"
	"recipe-management/pkg"
)

type recipeUseCase struct {
	recipeRepo   domain.RecipeRepository
	categoryRepo domain.CategoryRepository
}

func (r recipeUseCase) Fetch(pagination *pkg.Pagination) error {
	if err := r.recipeRepo.Fetch(pagination); err != nil {
		return err
	}

	return nil
}

func (r recipeUseCase) GetByID(id uint64) (*domain.Recipe, error) {
	recipe, err := r.recipeRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (r recipeUseCase) Create(recipe *domain.Recipe) (*domain.Recipe, error) {
	_, err := r.categoryRepo.GetByID(recipe.CategoryID)
	if err != nil {
		return nil, err
	}

	for i := range recipe.Ingredients {
		recipe.Ingredients[i].Sort = uint8(i)
	}

	for i := range recipe.Steps {
		recipe.Steps[i].Sort = uint8(i)
	}

	recipe, err = r.recipeRepo.Create(recipe)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (r recipeUseCase) Update(id uint64, updateData *domain.Recipe) (*domain.Recipe, error) {
	recipe, err := r.recipeRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if recipe.CategoryID != updateData.CategoryID {
		newCategory, err := r.categoryRepo.GetByID(updateData.CategoryID)
		if err != nil {
			return nil, err
		}
		recipe.Category = *newCategory
	}

	for i := range updateData.Ingredients {
		updateData.Ingredients[i].Sort = uint8(i)
	}

	for i := range updateData.Steps {
		updateData.Steps[i].Sort = uint8(i)
	}

	recipe, err = r.recipeRepo.Update(recipe, updateData)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (r recipeUseCase) Delete(id uint64) error {
	recipe, err := r.recipeRepo.GetByID(id)
	if err != nil {
		return err
	}

	if err = r.recipeRepo.Delete(recipe); err != nil {
		return err
	}

	return nil
}

func NewRecipeUseCase(rr domain.RecipeRepository, cr domain.CategoryRepository) domain.RecipeUseCase {
	return &recipeUseCase{recipeRepo: rr, categoryRepo: cr}
}
