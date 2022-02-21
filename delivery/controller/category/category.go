package category

import (
	_common "capstone/be/delivery/common"
	_categoryRepo "capstone/be/repository/category"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	repository _categoryRepo.Category
}

func New(category _categoryRepo.Category) *CategoryController {
	return &CategoryController{repository: category}
}

func (uc CategoryController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		// calling repository
		categories, code, err := uc.repository.GetAll()

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		return c.JSON(http.StatusOK, _common.GetAllCategoryResponse(categories))
	}
}
