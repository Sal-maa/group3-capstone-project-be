package request

import (
	_common "capstone/be/delivery/common"
	_helper "capstone/be/delivery/helper"
	_midware "capstone/be/delivery/middleware"
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
		newReq := _entity.CreateBorrow{}

		// handle failure in binding
		if err := c.Bind(&newReq); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}

		// prepare input string
		newReq.ShortName = strings.TrimSpace(newReq.ShortName)
		newReq.Description = strings.TrimSpace(newReq.Description)

		// handle borrow request based on role
		role := _midware.ExtractRole(c)

		reqData := _entity.Borrow{}

		switch role {
		case "Administrator":
			// if return time not set, then set to max time
			if newReq.ReturnTime == (time.Time{}) {
				newReq.ReturnTime = time.Unix(1<<63-62135596801, 999999999)
			}

			// prepare input to repository
			reqData.User.Id = newReq.EmployeeId
			reqData.Asset.ShortName = newReq.ShortName
			reqData.Description = newReq.Description
			reqData.ReturnTime = newReq.ReturnTime
			reqData.Status = "Approved by Admin"

			if code, err := rc.repository.Borrow(reqData); err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}

			return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Success create request"))
		case "Employee":
			// set return time to maximum value
			newReq.ReturnTime = time.Unix(1<<63-62135596801, 999999999)

			// prepare input to repository
			reqData.User.Id = _midware.ExtractId(c)
			reqData.Asset.ShortName = newReq.ShortName
			reqData.Description = newReq.Description
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
		role := _midware.ExtractRole(c)
		if role != "Administrator" {
			return c.JSON(http.StatusForbidden, _common.NoDataResponse(http.StatusForbidden, "Only Admin Can Create Request"))
		}
		idLogin := _midware.ExtractId(c)
		newReq := _entity.CreateProcure{}
		// handle failure in binding
		if err := c.Bind(&newReq); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Failed to Bind Data"))
		}

		reqData := _entity.Procure{}
		// check category id
		categoryId, err := rc.repository.GetCategoryId(newReq.Category)
		if categoryId == 0 {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Category Not Found"))
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, _common.NoDataResponse(http.StatusInternalServerError, "Failed to Get Category Id"))
		}

		// detect image upload
		src, file, err := c.Request().FormFile("image")

		// detect failure in parsing file
		switch err {
		case nil:
			defer src.Close()

			// upload avatar to amazon s3
			image, code, err := _helper.UploadImage("procure", idLogin, file, src)

			// detect failure while uploading avatar
			if err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}

			newReq.Image = image
		case http.ErrMissingFile:
			image := newReq.Image[strings.LastIndex(newReq.Image, "/")+1:]
			newReq.Image = image
		case http.ErrNotMultipart:
			image := newReq.Image[strings.LastIndex(newReq.Image, "/")+1:]
			newReq.Image = image
		default:
			log.Println(err)
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Failed to Upload Image"))
		}
		reqData.User.Id = idLogin
		reqData.Category = categoryId
		reqData.Image = fmt.Sprintf("https://capstone-group3.s3.ap-southeast-1.amazonaws.com/%s", newReq.Image)
		reqData.Activity = newReq.Activity
		reqData.RequestTime = time.Now()
		reqData.Status = "Waiting Approval from Manager"
		reqData.Description = newReq.Description
		reqData.UpdatedAt = time.Now()

		_, err = rc.repository.Procure(reqData)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Failed Create Request"))
		}

		return c.JSON(http.StatusOK, _common.NoDataResponse(http.StatusOK, "Success Create Request"))
	}
}

func (rc RequestController) UpdateBorrow() echo.HandlerFunc {
	return func(c echo.Context) error {
		idReq, err := strconv.Atoi(c.Param("id"))
		// detect invalid parameter
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Invalid Request Id"))
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
