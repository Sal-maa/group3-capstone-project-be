package history

import (
	_common "capstone/be/delivery/common"
	_midware "capstone/be/delivery/middleware"
	_historyRepo "capstone/be/repository/history"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type HistoryController struct {
	repository _historyRepo.History
}

func New(history _historyRepo.History) *HistoryController {
	return &HistoryController{repository: history}
}

func (hc HistoryController) GetAllUsageHistoryOfUser() echo.HandlerFunc {
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

		// pagination
		p := strings.TrimSpace(c.QueryParam("page"))
		if p == "" {
			p = "1"
		}

		page, err := strconv.Atoi(p)

		// detect invalid page number
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid page number"))
		}

		// calling repository
		histories, code, err := hc.repository.GetAllUsageHistoryOfUser(user_id, page)

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		return c.JSON(http.StatusOK, _common.GetAllUsageHistoryOfUserResponse(histories))
	}
}

func (hc HistoryController) GetDetailUsageHistoryByRequestId() echo.HandlerFunc {
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
		history, code, err := hc.repository.GetDetailUsageHistoryByRequestId(request_id)

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		return c.JSON(http.StatusOK, _common.GetDetailUsageHistoryByRequestId(history))
	}
}

func (hc HistoryController) GetAllUsageHistoryOfAsset() echo.HandlerFunc {
	return func(c echo.Context) error {
		short_name := c.Param("short_name")

		// calling repository
		histories, code, err := hc.repository.GetAllUsageHistoryOfAsset(short_name)

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		return c.JSON(http.StatusOK, _common.GetAllUsageHistoryOfAsset(histories))
	}
}
