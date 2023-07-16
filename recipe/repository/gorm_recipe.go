package repository

import (
	"errors"
	"gorm.io/gorm"
	"recipe-management/domain"
	"recipe-management/infrastructure/database"
	"recipe-management/pkg"
	"strconv"
)

type gormRecipeRepository struct {
	db *gorm.DB
}

func (repo gormRecipeRepository) Fetch(pagination *pkg.Pagination) error {
	var recipes []domain.Recipe

	query := repo.db.Debug()

	if pagination.Keyword != "" {
		query = query.Where("title like ?", "%"+pagination.Keyword+"%")
	}

	if filterCategory, ok := pagination.Filters["category_id"]; ok {
		if filterCategory != "" {
			var categoryID int
			categoryID, err := strconv.Atoi(filterCategory)
			if err != nil {
				categoryID = 0
			}

			query = query.Where("category_id = ?", categoryID)
		}
	}

	err := query.Scopes(database.Paginate(recipes, pagination, query)).Find(&recipes).Error
	if err != nil {
		return err
	}

	pagination.Rows = recipes

	return nil
}

func (repo gormRecipeRepository) GetByID(id uint64) (*domain.Recipe, error) {
	var recipe domain.Recipe

	if err := repo.db.Joins("Category").Preload("Ingredients").Preload("Steps").
		First(&recipe, "recipes.id=?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrRecipeNotFound
		}

		return nil, err
	}

	return &recipe, nil
}

func (repo gormRecipeRepository) Create(recipe *domain.Recipe) (*domain.Recipe, error) {
	if err := repo.db.Create(&recipe).Error; err != nil {
		return nil, err
	}

	return recipe, nil
}

func (repo gormRecipeRepository) Update(recipe *domain.Recipe, updateData *domain.Recipe) (*domain.Recipe, error) {
	recipe.CategoryID = updateData.CategoryID
	recipe.Title = updateData.Title
	recipe.Description = updateData.Description
	recipe.CookingTime = updateData.CookingTime
	recipe.ThumbnailURL = updateData.ThumbnailURL
	recipe.VideoURL = updateData.VideoURL

	err := repo.db.Transaction(func(tx *gorm.DB) error {
		if err := repo.db.Save(&recipe).Error; err != nil {
			return err
		}

		if err := repo.db.Model(&recipe).Association("Ingredients").Replace(updateData.Ingredients); err != nil {
			return err
		}

		if err := repo.db.Model(&recipe).Association("Steps").Replace(updateData.Steps); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (repo gormRecipeRepository) Delete(recipe *domain.Recipe) error {
	if err := repo.db.Delete(&recipe).Error; err != nil {
		return err
	}

	return nil
}

func NewGormRecipeRepository(db *gorm.DB) domain.RecipeRepository {
	return &gormRecipeRepository{db: db}
}
