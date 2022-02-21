package request

import (
	_common "capstone/be/delivery/common"
	"capstone/be/delivery/middleware"
	_entity "capstone/be/entity"
	_requestRepo "capstone/be/repository/request"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RequestController struct {
	repository _requestRepo.Request
}

func New(request _requestRepo.Request) *RequestController {
	return &RequestController{repository: request}
}

func (rc RequestController) Borrow() echo.HandlerFunc {
	return func(c echo.Context) error {
		idLogin := middleware.ExtractId(c)
		newReq := _entity.CreateBorrow{}
		fmt.Println("newReq", newReq)
		// handle failure in binding
		if err := c.Bind(&newReq); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}
		reqData := _entity.Borrow{}
		// prepare input string
		reqData.User.Id = idLogin
		reqData.Asset.Id = newReq.AssetId
		reqData.Activity = newReq.Activity
		reqData.RequestTime = newReq.RequestTime
		reqData.ReturnTime = newReq.ReturnTime
		reqData.Status = "Menunggu Persetujuan Admin"
		reqData.Description = newReq.Description

		_, err := rc.repository.Borrow(reqData)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed create request"))
		}

		return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Success create request"))
	}
}
