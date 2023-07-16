package repository

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"recipe-management/domain"
	"recipe-management/infrastructure/database"
	"recipe-management/pkg"
)

type gormCategoryRepository struct {
	db *gorm.DB
}

func NewMysqlCategoryRepository(db *gorm.DB) domain.CategoryRepository {
	return &gormCategoryRepository{db}
}

func (repo gormCategoryRepository) Fetch(pagination pkg.Pagination) (*pkg.Pagination, error) {
	var categories []domain.Category

	repo.db.Scopes(database.Paginate(categories, &pagination, repo.db)).Find(&categories)
	pagination.Rows = categories

	return &pagination, nil
}

func (repo gormCategoryRepository) Create(c *domain.Category) (*domain.Category, error) {
	if err := repo.db.Create(c).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil, domain.ErrCategoryConflict
		}

		return nil, err
	}

	return c, nil
}

func (repo gormCategoryRepository) GetByName(name string) (*domain.Category, error) {
	var category domain.Category
	if err := repo.db.Where("name = ?", name).First(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (repo gormCategoryRepository) GetByID(id uint64) (*domain.Category, error) {
	var category domain.Category

	result := repo.db.Where("id = ?", id).First(&category)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrCategoryNotFound
		}

		return nil, err
	}

	return &category, nil
}
