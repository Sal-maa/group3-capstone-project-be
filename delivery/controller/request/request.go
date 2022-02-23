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

		role := middleware.ExtractRole(c)
		switch role {
		case "Administrator":
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
			statAsset, err := rc.repository.CheckMaintenance(assetId)
			if err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to check asset status"))
			}
			if statAsset == "Asset Under Maintenance" {
				return c.JSON(http.StatusForbidden, _common.NoDataResponse(http.StatusForbidden, "Sorry, asset is under maintenace"))
			}
			reqData.Activity = newReq.Activity
			reqData.RequestTime = time.Now()
			reqData.ReturnTime = newReq.ReturnTime
			reqData.Status = "Waiting Approval"
			reqData.Description = newReq.Description
			reqData.UpdatedAt = time.Now()

			_, err = rc.repository.Borrow(reqData)
			if err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed create request"))
			}
			return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Success create request"))
		case "Employee":
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
			statAsset, err := rc.repository.CheckMaintenance(assetId)
			if err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to check asset status"))
			}
			if statAsset == "Asset Under Maintenance" {
				return c.JSON(http.StatusForbidden, _common.NoDataResponse(http.StatusForbidden, "Sorry, asset is under maintenace"))
			}
			reqData.Activity = newReq.Activity
			reqData.RequestTime = time.Now()
			reqData.Status = "Waiting Approval"
			reqData.Description = newReq.Description
			reqData.UpdatedAt = time.Now()

			_, err = rc.repository.Borrow(reqData)
			if err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed create request"))
			}
			return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Success create request"))
		default:
			return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Invalid input request"))
		}
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
		categoryId, err := rc.repository.GetCategoryId(newReq)
		if categoryId == 0 {
			// add new category if category isn't exist
			_, err := rc.repository.AddCategory(newReq.Category)
			if err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to add new category"))
			}
			categoryId, err = rc.repository.GetCategoryId(newReq)
			if err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to get category_id"))
			}
		}
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to get category_id"))
		}
		reqData.Category = categoryId

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
		reqData.RequestTime = time.Now()
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

		request, err := rc.repository.GetBorrowById(idReq)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed get request by id"))
		}

		newReq := _entity.UpdateBorrow{}
		// handle failure in binding
		if err := c.Bind(&newReq); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}

		role := middleware.ExtractRole(c)

		// check manager division and employee division
		idLogin := middleware.ExtractId(c)
		switch role {
		case "Manager":
			divLogin, err := rc.repository.GetUserDivision(idLogin)
			if err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed get division id user"))
			}

			divEmpl, err := rc.repository.GetUserDivision(request.User.Id)
			if err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed get division id user"))
			}

			if divEmpl != divLogin {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "You're not in the same division"))
			}

			if request.Status != "Waiting Approval" {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Approval by Manager Only"))
			}
			request.Status = newReq.Status
			_, err = rc.repository.UpdateBorrow(request)
			if err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed update request"))
			}
			return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Success update request"))
		case "Administrator":
			if request.Status != "Approve by Manager" {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Rejected by Manager/Administrator Only"))
			}
			request.Status = newReq.Status

			_, err = rc.repository.UpdateBorrowByAdmin(request)
			if err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed update request"))
			}
			return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Success update request"))
		default:
			return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Invalid input request"))
		}
	}
}

func (rc RequestController) UpdateProcure() echo.HandlerFunc {
	return func(c echo.Context) error {
		idReq, err := strconv.Atoi(c.Param("id"))

		// detect invalid parameter
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid request id"))
		}

		request, err := rc.repository.GetProcureById(idReq)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed get request by id"))
		}

		newReq := _entity.UpdateProcure{}
		// handle failure in binding
		if err := c.Bind(&newReq); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}
		role := middleware.ExtractRole(c)

		// check manager division and employee division
		idLogin := middleware.ExtractId(c)
		switch role {
		case "Manager":
			divLogin, err := rc.repository.GetUserDivision(idLogin)
			if err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed get division id user"))
			}

			divEmpl, err := rc.repository.GetUserDivision(request.Id)
			if err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed get division id user"))
			}

			if divEmpl != divLogin {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "You're not in the same division"))
			}
			if request.Status != "Waiting Approval" {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Approval by Manager Only"))
			}
			request.Status = newReq.Status

			_, err = rc.repository.UpdateProcure(request)
			if err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed update request"))
			}
			return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Success update request"))
		case "Administrator":
			if request.Status != "Approve by Manager" {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Rejected by Manager/Administrator Only"))
			}
			request.Status = newReq.Status

			_, err = rc.repository.UpdateProcureByAdmin(request)
			if err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed update request"))
			}
			return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Success update request"))
		default:
			return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Invalid input request"))
		}
	}
}
