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

		// check input string
		check := []string{shortName, description}

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
		reqData := _entity.Borrow{}
		reqData.Asset.ShortName = shortName
		reqData.Description = description

		// check role
		role := _midware.ExtractRole(c)

		switch role {
		case "Administrator":
			// if return time not set, then set to max time
			if newReq.ReturnTime == (time.Time{}) {
				newReq.ReturnTime = time.Unix(1<<63-62135596801, 999999999)
			}

			reqData.User.Id = newReq.EmployeeId
			reqData.ReturnTime = newReq.ReturnTime
			reqData.Status = "Approved by Admin"

			if code, err := rc.repository.Borrow(reqData); err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}

			return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Success create request"))
		case "Employee":
			// set return time to maximum value
			newReq.ReturnTime = time.Unix(1<<63-62135596801, 999999999)

			reqData.User.Id = _midware.ExtractId(c)
			reqData.ReturnTime = newReq.ReturnTime
			reqData.Status = "Waiting for approval"

			if code, err := rc.repository.Borrow(reqData); err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}

			return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Success create request"))
		default:
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusOK, "Invalid request"))
		}
	}
}

func (rc RequestController) Procure() echo.HandlerFunc {
	return func(c echo.Context) error {
		// check authorization
		if role := _midware.ExtractRole(c); role != "Administrator" {
			return c.JSON(http.StatusForbidden, _common.NoDataResponse(http.StatusForbidden, "Only Admin can create request"))
		}

		newReq := _entity.CreateProcure{}

		// handle failure in binding
		if err := c.Bind(&newReq); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Failed to bind data"))
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
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Failed to Upload Image"))
		}

		reqData.User.Id = _midware.ExtractId(c)
		reqData.Description = newReq.Description
		reqData.UpdatedAt = time.Now()

		code, err := rc.repository.Procure(reqData)

		if err != nil {
			log.Println(err)
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
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
			return c.JSON(http.StatusInternalServerError, _common.NoDataResponse(http.StatusInternalServerError, "Failed Get Request by Id"))
		}

		newReq := _entity.UpdateBorrow{}
		// handle failure in binding
		if err := c.Bind(&newReq); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Failed to Bind Data"))
		}

		role := _midware.ExtractRole(c)

		// check manager division and employee division
		idLogin := _midware.ExtractId(c)
		switch role {
		case "Manager":
			divLogin, err := rc.repository.GetUserDivision(idLogin)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, _common.NoDataResponse(http.StatusInternalServerError, "Failed Get Division Id User Login"))
			}

			divEmpl, err := rc.repository.GetUserDivision(request.User.Id)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, _common.NoDataResponse(http.StatusInternalServerError, "Failed Get Division Id User Request"))
			}

			if divEmpl != divLogin {
				return c.JSON(http.StatusForbidden, _common.NoDataResponse(http.StatusForbidden, "You're Not in The Same Division"))
			}

			if request.Status != "Waiting Approval from Manager" {
				return c.JSON(http.StatusForbidden, _common.NoDataResponse(http.StatusForbidden, "Update by Manager Only"))
			}
			request.Status = newReq.Status
			_, err = rc.repository.UpdateBorrow(request)
			if err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Failed Update Request"))
			}
			return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Success Update Request"))
		case "Administrator":
			switch request.Status {
			case "Waiting Approval From Admin":
				if request.Activity == "Peminjaman Aset" {
					if newReq.Status != "Waiting Approval from Manager" {
						return c.JSON(http.StatusForbidden, _common.NoDataResponse(http.StatusForbidden, "Status must be WAITING APPROVAL FROM MANAGER"))
					}
					request.Status = newReq.Status
					_, err = rc.repository.UpdateBorrowByAdmin(request)
					if err != nil {
						return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Failed Update Request"))
					}
				} else if request.Activity == "Pengembalian Aset" {
					request.Status = newReq.Status
					_, err = rc.repository.UpdateBorrowByAdmin(request)
					if err != nil {
						return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Failed Update Request"))
					}
				}
				return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Success Update Request"))
			case "Approved by Manager":
				request.Status = newReq.Status

				_, err = rc.repository.UpdateBorrowByAdmin(request)
				if err != nil {
					return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Failed Update Request"))
				}
				return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Success Update Request"))
			case "Rejected by Manager":
				return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Admin No Need Any Update"))
			default:
				return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Invalid Input request"))
			}
		default:
			return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Invalid Input request"))
		}
	}
}

func (rc RequestController) UpdateProcure() echo.HandlerFunc {
	return func(c echo.Context) error {
		idReq, err := strconv.Atoi(c.Param("id"))

		// detect invalid parameter
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Invalid Request Id"))
		}

		request, err := rc.repository.GetProcureById(idReq)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, _common.NoDataResponse(http.StatusInternalServerError, "Failed Get Request by Id"))
		}

		newReq := _entity.UpdateProcure{}
		// handle failure in binding
		if err := c.Bind(&newReq); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Failed to Bind Data"))
		}
		role := _midware.ExtractRole(c)

		// check manager division and employee division
		idLogin := _midware.ExtractId(c)
		if role != "Manager" {
			return c.JSON(http.StatusForbidden, _common.NoDataResponse(http.StatusForbidden, "Only Manager Can Do Approval"))
		}
		divLogin, err := rc.repository.GetUserDivision(idLogin)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, _common.NoDataResponse(http.StatusInternalServerError, "failed get division id user"))
		}

		divEmpl, err := rc.repository.GetUserDivision(request.Id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, _common.NoDataResponse(http.StatusInternalServerError, "failed get division id user"))
		}

		if divEmpl != divLogin {
			return c.JSON(http.StatusForbidden, _common.NoDataResponse(http.StatusForbidden, "You're Not in The Same Division"))
		}
		if request.Status != "Waiting Approval from Manager" {
			return c.JSON(http.StatusForbidden, _common.NoDataResponse(http.StatusForbidden, "Update by Manager Only"))
		}
		request.Status = newReq.Status

		_, err = rc.repository.UpdateProcure(request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Failed Update Request"))
		}
		return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Success Update Request"))
	}
}
