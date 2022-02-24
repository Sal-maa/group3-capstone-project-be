package admin

import (
	_common "capstone/be/delivery/common"
	"log"
	"strconv"
	"strings"

	// _helper "capstone/be/delivery/helper"
	_midware "capstone/be/delivery/middleware"
	"net/http"

	_adminRepo "capstone/be/repository/admin"

	"github.com/labstack/echo/v4"
)

// "github.com/labstack/echo/v4"

type AdminController struct {
	repository _adminRepo.Admin
}

func New(admin _adminRepo.Admin) *AdminController {
	return &AdminController{repository: admin}
}

func (ac AdminController) HomePageGetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := _midware.ExtractRole(c)
		if role != "Administrator" {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "You don't have permission"))
		}

		// filter by page number
		p := c.QueryParam("p")
		// default value for page
		if p == "" {
			p = "1"
		}

		limit, err := strconv.Atoi(p)

		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Error parsing page"))
		}

		// filter by records per page
		rp := c.QueryParam("rp")

		offset, err := strconv.Atoi(rp)

		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Error parsing record of page"))
		}

		// filter by status
		status := c.QueryParam("s")

		// default value for status
		if status == "" {
			status = "all"
		}

		// to prevent sql injection
		if strings.ContainsAny(status, ";") {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Bad request"))
		}

		allstatus := map[string]int{"all": 1, "Waiting Approval": 1, "Approved": 1, "Rejected": 1, "Request to Return": 1}

		if _, exist := allstatus[status]; !exist {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Bad request"))
		}
		// filter by date
		date := c.QueryParam("d")

		// filter by category
		category := c.QueryParam("category")

		// to prevent sql injection
		if strings.ContainsAny(category, ";") {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Bad request"))
		}

		// default value for category
		if category == "" {
			category = "all"
		}

		categories := map[string]int{"all": 1, "Computer": 1, "Computer Accessories": 1, "Networking": 1, "UPS": 1, "Printer and Scanner": 1, "Electronics": 1, "Others": 1}

		if _, exist := categories[category]; !exist {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Bad request"))
		}

		requests, err := ac.repository.GetAllNewRequest(limit, offset, status, date)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, _common.NoDataResponse(http.StatusInternalServerError, "Failed to read data"))
		}

		return c.JSON(http.StatusOK, _common.GetAllRequestResponse(requests))
	}
}
