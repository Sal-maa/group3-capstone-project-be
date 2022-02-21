package request

import (
	_common "capstone/be/delivery/common"
	_helper "capstone/be/delivery/helper"
	"capstone/be/delivery/middleware"
	_entity "capstone/be/entity"
	_requestRepo "capstone/be/repository/request"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

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
		// handle failure in binding
		if err := c.Bind(&newReq); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}
		reqData := _entity.Borrow{}
		// prepare input string
		reqData.User.Id = idLogin
		reqData.Asset.Id = newReq.AssetId
		// handle maintenance status
		asset, err := rc.repository.CheckMaintenance(reqData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to check asset asset"))
		}
		if asset.Status == "Asset Under Maintenance" {
			return c.JSON(http.StatusForbidden, _common.NoDataResponse(http.StatusForbidden, "Sorry, asset is under maintenace"))
		}
		reqData.Activity = newReq.Activity
		reqData.RequestTime = newReq.RequestTime
		reqData.ReturnTime = newReq.ReturnTime
		reqData.Status = "Waiting Approval from Admin"
		reqData.Description = newReq.Description
		reqData.UpdatedAt = time.Now()

		_, err = rc.repository.Borrow(reqData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed create request"))
		}

		return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Success create request"))
	}
}

func (rc RequestController) Procure() echo.HandlerFunc {
	return func(c echo.Context) error {
		idLogin := middleware.ExtractId(c)
		newReq := _entity.CreateProcure{}
		// handle failure in binding
		if err := c.Bind(&newReq); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}

		reqData := _entity.Procure{}
		reqData.User.Id = idLogin

		// check category id
		category, err := rc.repository.GetCategoryId(newReq)
		if category == 0 {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "category not found"))
		}
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to get category_id"))
		}
		reqData.Category = category

		// detect image upload
		src, file, err := c.Request().FormFile("image")

		// detect failure in parsing file
		switch err {
		case nil:
			defer src.Close()

			// upload avatar to amazon s3
			image, code, err := _helper.UploadImage("asset", reqData.User.Id, file, src)

			// detect failure while uploading avatar
			if err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}

			reqData.Image = image
		case http.ErrMissingFile:
			image := reqData.Image[strings.LastIndex(reqData.Image, "/")+1:]
			reqData.Image = image
		case http.ErrNotMultipart:
			image := reqData.Image[strings.LastIndex(reqData.Image, "/")+1:]
			reqData.Image = image
		default:
			log.Println(err)
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to upload avatar"))
		}

		reqData.Activity = newReq.Activity
		reqData.RequestTime = newReq.RequestTime
		reqData.Status = "Waiting Approval from Admin"
		reqData.Description = newReq.Description
		reqData.UpdatedAt = time.Now()

		_, err = rc.repository.Procure(reqData)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed create request"))
		}

		return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Success create request"))
	}
}
