package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
	"recipe-management/domain"
	"recipe-management/infrastructure/database"
)

type GormCategoryTestSuite struct {
	database.GormTestSuite
	db         *gorm.DB
	mock       sqlmock.Sqlmock
	category   *domain.Category
	repository domain.CategoryRepository
}
