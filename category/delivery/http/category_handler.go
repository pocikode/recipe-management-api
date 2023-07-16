package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"recipe-management/domain"
	"recipe-management/pkg"
	"strconv"
)

type CategoryHandler struct {
	categoryUCase domain.CategoryUseCase
}

func NewCategoryHandler(e *echo.Echo, cu domain.CategoryUseCase) {
	handler := &CategoryHandler{
		categoryUCase: cu,
	}

	e.GET("/categories", handler.FetchCategory)
	e.POST("/categories", handler.CreateCategory)
}

func (ch *CategoryHandler) FetchCategory(c echo.Context) error {
	pagination := pkg.Pagination{}
	pagination.Page, _ = strconv.Atoi(c.QueryParam("page"))
	pagination.Limit, _ = strconv.Atoi(c.QueryParam("limit"))

	fetch, err := ch.categoryUCase.Fetch(&pagination)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, fetch.ToResponse())
}

func (ch *CategoryHandler) CreateCategory(c echo.Context) error {
	var category domain.Category

	if err := c.Bind(&category); err != nil {
		return err
	}

	if err := c.Validate(&category); err != nil {
		return err
	}

	cat, err := ch.categoryUCase.Create(&category)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, cat)
}
