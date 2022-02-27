package request

import (
	_common "capstone/be/delivery/common"
	_helper "capstone/be/delivery/helper"
	_midware "capstone/be/delivery/middleware"
	_entity "capstone/be/entity"
	_requestRepo "capstone/be/repository/request"
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
		newReq := _entity.CreateBorrow{}

		// handle failure in binding
		if err := c.Bind(&newReq); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}

		// prepare input string
		shortName := strings.TrimSpace(newReq.ShortName)
		description := strings.TrimSpace(newReq.Description)

		// check empty string in required input
		// description is not required, such as when admin assigns asset to employee
		if shortName == "" {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "input cannot be empty"))
		}

		// check input string
		check := []string{shortName, description}

		for _, s := range check {
			// check malicious character in input
			if err := _helper.CheckStringInput(s); err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, s+": "+err.Error()))
			}
		}

		// prepare input to repository
		reqData := _entity.Borrow{}
		reqData.Asset.ShortName = shortName
		reqData.Description = description

		// check role
		role := _midware.ExtractRole(c)

		switch role {
		case "Administrator":
			// if return time not set, then set to max time
			if newReq.ReturnTime == (time.Time{}) {
				newReq.ReturnTime = time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC)
			}

			reqData.User.Id = newReq.EmployeeId
			reqData.ReturnTime = newReq.ReturnTime
			reqData.Status = "Approved by Admin"

			// calling repository
			code, err := rc.repository.Borrow(reqData)

			// detect failure in repository
			if err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}

			return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "success create request"))
		case "Employee":
			// set return time to maximum value
			newReq.ReturnTime = time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC)

			reqData.User.Id = _midware.ExtractId(c)
			reqData.ReturnTime = newReq.ReturnTime
			reqData.Status = "Waiting approval"

			if code, err := rc.repository.Borrow(reqData); err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}

			return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "success create request"))
		default:
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusOK, "only admin/employee can make request"))
		}
	}
}

func (rc RequestController) Procure() echo.HandlerFunc {
	return func(c echo.Context) error {
		// check authorization
		if role := _midware.ExtractRole(c); role != "Administrator" {
			return c.JSON(http.StatusForbidden, _common.NoDataResponse(http.StatusForbidden, "only admin can create request"))
		}

		newReq := _entity.CreateProcure{}

		// handle failure in binding
		if err := c.Bind(&newReq); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}

		// prepare input string
		category := strings.TrimSpace(newReq.Category)
		description := strings.TrimSpace(newReq.Description)

		// check input string
		check := []string{category, description}

		for _, s := range check {
			// check empty string in required input
			if s == "" {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "input cannot be empty"))
			}

			// check malicious character in input
			if err := _helper.CheckStringInput(s); err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, s+": "+err.Error()))
			}
		}

		// prepare input to repository
		reqData := _entity.Procure{}
		reqData.User.Id = _midware.ExtractId(c)
		reqData.Category = category
		reqData.Description = description

		// detect image upload
		src, file, err := c.Request().FormFile("image")

		// detect failure in parsing file
		switch err {
		case nil:
			defer src.Close()

			// upload avatar to amazon s3
			image, code, err := _helper.UploadImage("procure", reqData.User.Id, file, src)

			// detect failure while uploading avatar
			if err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}

			reqData.Image = image
		case http.ErrMissingFile:
			reqData.Image = "default_image.png"
		case http.ErrNotMultipart:
			reqData.Image = "default_image.png"
		default:
			log.Println(err)
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to upload image"))
		}

		reqData.User.Id = _midware.ExtractId(c)
		reqData.Description = newReq.Description
		reqData.Status = "Waiting approval"

		// calling repository
		code, err := rc.repository.Procure(reqData)

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "success create request"))
	}
}

func (rc RequestController) UpdateBorrow() echo.HandlerFunc {
	return func(c echo.Context) error {
		// get request id to be updated
		idReq, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid request id"))
		}

		// get existing borrow request by id
		request, code, err := rc.repository.GetBorrowById(idReq)

		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		newStatus := _entity.UpdateBorrow{}

		// handle failure in binding
		if err := c.Bind(&newStatus); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}

		// check role and division of currently logged in user
		role := _midware.ExtractRole(c)
		idLogin := _midware.ExtractId(c)

		switch role {
		case "Manager":
			// this is the case where the logged in user is a manager
			// and he/she wants to approve or reject request ONLY IF
			// request status is waiting approval

			// check manager division
			divLogin, code, err := rc.repository.GetUserDivision(idLogin)

			if err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}

			// check user division
			divEmpl, code, err := rc.repository.GetUserDivision(request.User.Id)

			if err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}

			// check authorization
			if divEmpl != divLogin {
				return c.JSON(http.StatusForbidden, _common.NoDataResponse(http.StatusForbidden, "you are not in the same division"))
			}

			// check request status
			if request.Status != "Waiting approval" {
				return c.JSON(http.StatusForbidden, _common.NoDataResponse(http.StatusForbidden, "cannot approve/reject this request"))
			}

			// set request status
			if newStatus.Approved {
				request.Status = "Approved by Manager"
			} else {
				request.Status = "Rejected by Manager"
			}

			// calling repository
			_, code, err = rc.repository.UpdateBorrow(request)

			// detect failure in repository
			if err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}

			return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "success update request"))
		case "Administrator":
			// this is the case where the logged in user is an administrator
			// and he/she wants to approve or reject request ONLY AFTER the request
			// has been approved by manager

			// check request status
			if request.Status != "Approved by Manager" {
				return c.JSON(http.StatusForbidden, _common.NoDataResponse(http.StatusForbidden, "cannot approve/reject this request"))
			}

			// set request status
			if newStatus.Approved {
				request.Status = "Approved by Admin"
			} else {
				request.Status = "Rejected by Admin"
			}

			// calling repository
			_, code, err = rc.repository.UpdateBorrow(request)

			// detect failure in repository
			if err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}

			return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "success update request"))
		default:
			// this is the case where the logged in user is ordinary employee
			return c.JSON(http.StatusForbidden, _common.NoDataResponse(http.StatusForbidden, "not allowed to update request status"))
		}
	}
}

func (rc RequestController) UpdateProcure() echo.HandlerFunc {
	return func(c echo.Context) error {
		// get request id to be updated
		idReq, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid request id"))
		}

		// check role of currently logged in user
		if role := _midware.ExtractRole(c); role != "Manager" {
			return c.JSON(http.StatusForbidden, _common.NoDataResponse(http.StatusForbidden, "not allowed to update request status"))
		}

		newStatus := _entity.UpdateProcure{}

		// handle failure in binding
		if err := c.Bind(&newStatus); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}

		// get existing procure request by id
		request, code, err := rc.repository.GetProcureById(idReq)

		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		// check request status
		if request.Status != "Waiting approval" {
			return c.JSON(http.StatusForbidden, _common.NoDataResponse(http.StatusForbidden, "cannot approve/reject this request"))
		}

		// set request status
		if newStatus.Approved {
			request.Status = "Approved by Manager"
		} else {
			request.Status = "Rejected by Manager"
		}

		// calling repository
		_, code, err = rc.repository.UpdateProcure(request)

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "success update request"))
	}
}
