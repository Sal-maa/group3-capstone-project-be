package activity

import (
	_common "capstone/be/delivery/common"
	_midware "capstone/be/delivery/middleware"
	_entity "capstone/be/entity"
	_activityRepo "capstone/be/repository/activity"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ActivityController struct {
	repository _activityRepo.Activity
}

func New(activity _activityRepo.Activity) *ActivityController {
	return &ActivityController{repository: activity}
}

func (ac ActivityController) GetAllActivityOfUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id, err := strconv.Atoi(c.Param("user_id"))

		// detect invalid parameter
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid user id"))
		}

		// check authorization
		if user_id != _midware.ExtractId(c) && _midware.ExtractRole(c) != "Administrator" {
			return c.JSON(http.StatusUnauthorized, _common.NoDataResponse(http.StatusUnauthorized, "unauthorized"))
		}

		// calling repository
		activities, code, err := ac.repository.GetAllActivityOfUser(user_id)

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		return c.JSON(http.StatusOK, _common.GetAllActivityOfUserResponse(activities))
	}
}

func (ac ActivityController) GetDetailActivityByRequestId() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id, err := strconv.Atoi(c.Param("user_id"))

		// detect invalid parameter
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid user id"))
		}

		// check authorization
		if user_id != _midware.ExtractId(c) && _midware.ExtractRole(c) != "Administrator" {
			return c.JSON(http.StatusUnauthorized, _common.NoDataResponse(http.StatusUnauthorized, "unauthorized"))
		}

		request_id, err := strconv.Atoi(c.Param("request_id"))

		// detect invalid parameter
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid request id"))
		}

		// calling repository
		activity, code, err := ac.repository.GetDetailActivityByRequestId(request_id)

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		return c.JSON(http.StatusOK, _common.GetDetailActivityByRequestId(activity))
	}
}

func (ac ActivityController) UpdateRequestStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id, err := strconv.Atoi(c.Param("user_id"))

		// detect invalid parameter
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid user id"))
		}

		// check authorization
		if user_id != _midware.ExtractId(c) && _midware.ExtractRole(c) != "Administrator" {
			return c.JSON(http.StatusUnauthorized, _common.NoDataResponse(http.StatusUnauthorized, "unauthorized"))
		}

		request_id, err := strconv.Atoi(c.Param("request_id"))

		// detect invalid parameter
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid request id"))
		}

		activityData := _entity.UpdateActivityStatus{}

		// detect failure in binding
		if err := c.Bind(&activityData); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}

		// prepare input string
		status := strings.ToLower(strings.TrimSpace(activityData.Status))

		// calling repository to get existing activity data
		updateActivityData, code, err := ac.repository.GetDetailActivityByRequestId(request_id)

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		updatedActivity := _entity.Activity{}

		switch status {
		case "cancel":
			if updateActivityData.Status == "Cancelled" {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "request already cancelled"))
			} else if updateActivityData.ActivityType == "Return" {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "asset already in return process"))
			} else if updateActivityData.Status != "Waiting approval" {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "request cannot be cancelled"))
			}

			code, err := ac.repository.CancelRequest(request_id)

			if err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}

			updatedActivity = updateActivityData
			updatedActivity.Status = "Cancelled"

			return c.JSON(http.StatusOK, _common.UpdateRequestStatusResponse(updatedActivity, "cancel"))
		case "return":
			if updateActivityData.Status == "Cancelled" {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "request already cancelled"))
			} else if updateActivityData.ActivityType == "Return" {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "asset already in return process"))
			} else if updateActivityData.Status != "Approved by Admin" {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "cannot return asset"))
			}

			code, err := ac.repository.ReturnRequest(request_id)

			if err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}

			updatedActivity = updateActivityData
			updatedActivity.Status = "Waiting approval"
			updatedActivity.ActivityType = "Return"

			return c.JSON(http.StatusOK, _common.UpdateRequestStatusResponse(updatedActivity, "return"))
		default:
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid input status"))
		}
	}
}
