package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"recipe-management/domain"
	"recipe-management/pkg"
	"strconv"
)

type recipeHandler struct {
	recipeUCase domain.RecipeUseCase
}

func NewRecipeHandler(e *echo.Echo, ru domain.RecipeUseCase) {
	handler := recipeHandler{recipeUCase: ru}

	e.GET("/recipes", handler.FetchRecipe)
	e.GET("/recipes/:id", handler.GetRecipeByID)
	e.POST("/recipes", handler.CreateRecipe)
	e.PUT("/recipes/:id", handler.UpdateRecipe)
	e.DELETE("/recipes/:id", handler.DeleteRecipe)
}

func (rh *recipeHandler) FetchRecipe(ctx echo.Context) error {
	pagination := pkg.Pagination{}
	pagination.Page, _ = strconv.Atoi(ctx.QueryParam("page"))
	pagination.Limit, _ = strconv.Atoi(ctx.QueryParam("limit"))
	pagination.Keyword = ctx.QueryParam("keyword")
	pagination.Filters = map[string]string{"category_id": ctx.QueryParam("category_id")}

	if err := rh.recipeUCase.Fetch(&pagination); err != nil {
		return err
	}

	var rows []RecipeListSerializer
	castedRows, _ := pagination.Rows.([]domain.Recipe)

	for _, recipe := range castedRows {
		rows = append(rows, RecipeListSerializer{
			ID:           recipe.ID,
			CategoryID:   recipe.CategoryID,
			Title:        recipe.Title,
			Description:  recipe.Description,
			CookingTime:  recipe.CookingTime,
			ThumbnailURL: recipe.ThumbnailURL,
			VideoURL:     recipe.VideoURL,
			CreatedAt:    recipe.CreatedAt,
			UpdatedAt:    recipe.UpdatedAt,
		})
	}

	pagination.Rows = rows

	return ctx.JSON(http.StatusOK, pagination.ToResponse())
}

func (rh *recipeHandler) GetRecipeByID(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return domain.ErrRecipeNotFound
	}

	recipe, err := rh.recipeUCase.GetByID(id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, recipe)
}

func (rh *recipeHandler) CreateRecipe(ctx echo.Context) error {
	var recipe domain.Recipe

	if err := ctx.Bind(&recipe); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(&recipe); err != nil {
		return err
	}

	r, err := rh.recipeUCase.Create(&recipe)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, r)
}

func (rh *recipeHandler) UpdateRecipe(ctx echo.Context) error {
	var updateData domain.Recipe

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return domain.ErrRecipeNotFound
	}

	if err = ctx.Bind(&updateData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = ctx.Validate(&updateData); err != nil {
		return err
	}

	r, err := rh.recipeUCase.Update(id, &updateData)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, r)
}

func (rh *recipeHandler) DeleteRecipe(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return domain.ErrRecipeNotFound
	}

	if err = rh.recipeUCase.Delete(id); err != nil {
		return err
	}

	return ctx.JSON(http.StatusNoContent, "Success")
}
