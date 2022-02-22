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
	"strconv"
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

		// handle get asset id
		assetId, err := rc.repository.GetAssetId(newReq)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to check asset id"))
		}
		if assetId == 0 {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "asset not found"))
		}
		reqData.Asset.Id = assetId

		// handle maintenance status
		asset, err := rc.repository.CheckMaintenance(reqData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to check asset status"))
		}
		if asset.Status == "Asset Under Maintenance" {
			return c.JSON(http.StatusForbidden, _common.NoDataResponse(http.StatusForbidden, "Sorry, asset is under maintenace"))
		}
		reqData.Activity = newReq.Activity
		reqData.RequestTime = newReq.RequestTime
		reqData.ReturnTime = newReq.ReturnTime
		reqData.Status = "Waiting Approval"
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

func (rc RequestController) UpdateBorrow() echo.HandlerFunc {
	return func(c echo.Context) error {
		idReq, err := strconv.Atoi(c.Param("id"))
		// detect invalid parameter
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid request id"))
		}

		role := middleware.ExtractRole(c)
		if role != "Manager" {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "you don't have permission"))
		}

		// check manager division and employee division
		idLogin := middleware.ExtractId(c)
		divLogin, err := rc.repository.GetUserDivision(idLogin)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed get division id user"))
		}

		request, err := rc.repository.GetBorrowById(idReq)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed get request by id"))
		}

		divEmpl, err := rc.repository.GetUserDivision(request.Id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed get division id user"))
		}

		if divEmpl != divLogin {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "You're not in the same division"))
		}

		newReq := _entity.UpdateBorrow{}
		// handle failure in binding
		if err := c.Bind(&newReq); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}
		reqData := _entity.Borrow{}
		reqData.Id = idReq
		reqData.Status = newReq.Status

		_, err = rc.repository.UpdateBorrow(reqData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed update request"))
		}

		return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Success update request"))
	}
}

func (rc RequestController) UpdateProcure() echo.HandlerFunc {
	return func(c echo.Context) error {
		idReq, err := strconv.Atoi(c.Param("id"))

		// detect invalid parameter
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid request id"))
		}
		role := middleware.ExtractRole(c)
		if role != "Manager" {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "you don't have permission"))
		}

		// check manager division and employee division
		idLogin := middleware.ExtractId(c)
		divLogin, err := rc.repository.GetUserDivision(idLogin)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed get division id user"))
		}

		request, err := rc.repository.GetBorrowById(idReq)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed get request by id"))
		}

		divEmpl, err := rc.repository.GetUserDivision(request.Id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed get division id user"))
		}

		if divEmpl != divLogin {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "You're not in the same division"))
		}

		newReq := _entity.UpdateProcure{}
		// handle failure in binding
		if err := c.Bind(&newReq); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}
		reqData := _entity.Procure{}
		reqData.Id = idReq
		reqData.Status = newReq.Status

		_, err = rc.repository.UpdateProcure(reqData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed update request"))
		}

		return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Success update request"))
	}
}

func (rc RequestController) UpdateBorrowByAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		idReq, err := strconv.Atoi(c.Param("id"))
		// detect invalid parameter
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid request id"))
		}

		role := middleware.ExtractRole(c)
		if role != "Administrator" {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "you don't have permission"))
		}

		// check manager division and employee division
		idLogin := middleware.ExtractId(c)
		divLogin, err := rc.repository.GetUserDivision(idLogin)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed get division id user"))
		}

		request, err := rc.repository.GetBorrowById(idReq)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed get request by id"))
		}

		divEmpl, err := rc.repository.GetUserDivision(request.Id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed get division id user"))
		}

		if divEmpl != divLogin {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "You're not in the same division"))
		}

		newReq := _entity.UpdateBorrow{}
		// handle failure in binding
		if err := c.Bind(&newReq); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}
		reqData := _entity.Borrow{}
		reqData.Id = idReq
		reqData.Status = newReq.Status

		_, err = rc.repository.UpdateBorrowByAdmin(reqData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed update request"))
		}

		return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Success update request"))
	}
}

func (rc RequestController) UpdateProcureByAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		idReq, err := strconv.Atoi(c.Param("id"))
		// detect invalid parameter
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid request id"))
		}

		role := middleware.ExtractRole(c)
		if role != "Administrator" {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "you don't have permission"))
		}

		// check manager division and employee division
		idLogin := middleware.ExtractId(c)
		divLogin, err := rc.repository.GetUserDivision(idLogin)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed get division id user"))
		}

		request, err := rc.repository.GetBorrowById(idReq)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed get request by id"))
		}

		divEmpl, err := rc.repository.GetUserDivision(request.Id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed get division id user"))
		}

		if divEmpl != divLogin {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "You're not in the same division"))
		}

		newReq := _entity.UpdateProcure{}
		// handle failure in binding
		if err := c.Bind(&newReq); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}
		reqData := _entity.Procure{}
		reqData.Id = idReq
		reqData.Status = newReq.Status

		_, err = rc.repository.UpdateProcureByAdmin(reqData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed update request"))
		}

		return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Success update request"))
	}
}
