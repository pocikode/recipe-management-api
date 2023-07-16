package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"recipe-management/app/config"
	_categoryHttpDelivery "recipe-management/category/delivery/http"
	_categoryHttpMiddleware "recipe-management/category/delivery/http/middleware"
	_categoryRepo "recipe-management/category/repository"
	_categoryUseCase "recipe-management/category/usecase"
	"recipe-management/infrastructure/database"
	"recipe-management/pkg"
	_recipeHttpDelivery "recipe-management/recipe/delivery/http"
	_recipeRepo "recipe-management/recipe/repository"
	_recipeUseCase "recipe-management/recipe/usecase"
)

func init() {
	config.ReadConfig()
}

func main() {
	db := database.NewDB()

	e := echo.New()
	middleware := _categoryHttpMiddleware.InitMiddleware()
	e.Use(middleware.CORS)
	e.Validator = pkg.NewCustomValidator()
	e.HTTPErrorHandler = pkg.NewHttpErrorHandler

	categoryRepo := _categoryRepo.NewMysqlCategoryRepository(db)
	categoryUCase := _categoryUseCase.NewCategoryUseCase(categoryRepo)
	_categoryHttpDelivery.NewCategoryHandler(e, categoryUCase)

	recipeRepo := _recipeRepo.NewGormRecipeRepository(db)
	recipeUCase := _recipeUseCase.NewRecipeUseCase(recipeRepo, categoryRepo)
	_recipeHttpDelivery.NewRecipeHandler(e, recipeUCase)

	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
