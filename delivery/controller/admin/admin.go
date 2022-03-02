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
	adminRepository _adminRepo.Admin
}

func New(admin _adminRepo.Admin) *AdminController {
	return &AdminController{adminRepository: admin}
}

func (ac AdminController) AdminGetAll() echo.HandlerFunc {
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

		offset, err := strconv.Atoi(p)

		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Error parsing page"))
		}

		// filter by records per page
		rp := c.QueryParam("rp")
		// default value for page
		if rp == "" {
			rp = "5"
		}

		limit, err := strconv.Atoi(rp)

		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Error parsing record of page"))
		}

		// filter by status

		status := strings.ToUpper(c.QueryParam("s"))
		// default value for status
		if status == "" {
			status = "ALL"
		}

		// to prevent sql injection
		if strings.ContainsAny(status, ";") {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Bad request"))
		}

		allstatus := map[string]string{"ALL": "all", "WAITING-APPROVAL": "Waiting Approval", "APPROVED": "Approved", "REJECTED": "Rejected", "REQUEST-TO-RETURN": "Request to Return"}

		if _, exist := allstatus[status]; !exist {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Bad request"))
		}

		status = allstatus[status]
		// filter by date
		date := c.QueryParam("d")

		// filter by category
		category := strings.ToUpper(c.QueryParam("c"))

		// to prevent sql injection
		if strings.ContainsAny(category, ";") {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Bad request"))
		}

		// default value for category
		if category == "" {
			category = "ALL"
		}

		categories := map[string]string{"ALL": "all", "COMPUTER": "Computer", "COMPUTER-ACCESSORIES": "Computer Accessories", "NETWORKING": "Networking", "UPS": "UPS", "PRINTER-AND-SCANNER": "Printer and Scanner", "ELECTRONICS": "Electronics", "OTHERS": "Others"}

		if _, exist := categories[category]; !exist {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Bad request"))
		}
		category = categories[category]
		requests, total, err := ac.adminRepository.GetAllAdmin(limit, offset, status, category, date)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, _common.NoDataResponse(http.StatusInternalServerError, "Failed to read data"))
		}

		return c.JSON(http.StatusOK, _common.GetAllRequestResponse(requests, total))
	}
}

func (ac AdminController) ManagerGetAllBorrow() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := _midware.ExtractRole(c)
		if role != "Manager" {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "You don't have permission"))
		}

		idLogin := _midware.ExtractId(c)
		divLogin, _, err := ac.adminRepository.GetUserDivision(idLogin)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed get division id user"))
		}

		// filter by page number
		p := c.QueryParam("p")
		// default value for page
		if p == "" {
			p = "1"
		}

		offset, err := strconv.Atoi(p)

		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Error parsing page"))
		}

		// filter by records per page
		rp := c.QueryParam("rp")
		// default value for page
		if rp == "" {
			rp = "5"
		}

		limit, err := strconv.Atoi(rp)

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

		allstatus := map[string]string{"ALL": "all", "WAITING-APPROVAL": "Waiting Approval", "APPROVED": "Approved", "REJECTED": "Rejected", "REQUEST-TO-RETURN": "Request to Return"}

		if _, exist := allstatus[status]; !exist {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Bad request"))
		}

		status = allstatus[status]
		// filter by date
		date := c.QueryParam("d")

		// filter by category
		category := strings.ToUpper(c.QueryParam("c"))

		// to prevent sql injection
		if strings.ContainsAny(category, ";") {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Bad request"))
		}

		// default value for category
		if category == "" {
			category = "ALL"
		}

		categories := map[string]string{"ALL": "all", "COMPUTER": "Computer", "COMPUTER-ACCESSORIES": "Computer Accessories", "NETWORKING": "Networking", "UPS": "UPS", "PRINTER-AND-SCANNER": "Printer and Scanner", "ELECTRONICS": "Electronics", "OTHERS": "Others"}

		if _, exist := categories[category]; !exist {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Bad request"))
		}
		category = categories[category]

		requests, total, err := ac.adminRepository.GetAllManager(divLogin, limit, offset, status, category, date)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, _common.NoDataResponse(http.StatusInternalServerError, "Failed to read data"))
		}

		return c.JSON(http.StatusOK, _common.GetAllRequestResponse(requests, total))
	}
}

func (ac AdminController) ManagerGetAllProcure() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := _midware.ExtractRole(c)
		if role != "Manager" {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "You don't have permission"))
		}

		// filter by page number
		p := c.QueryParam("p")
		// default value for page
		if p == "" {
			p = "1"
		}

		offset, err := strconv.Atoi(p)

		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Error parsing page"))
		}

		// filter by records per page
		rp := c.QueryParam("rp")
		// default value for page
		if rp == "" {
			rp = "5"
		}

		limit, err := strconv.Atoi(rp)

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

		allstatus := map[string]string{"ALL": "all", "WAITING-APPROVAL": "Waiting Approval", "APPROVED": "Approved", "REJECTED": "Rejected", "REQUEST-TO-RETURN": "Request to Return"}

		if _, exist := allstatus[status]; !exist {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Bad request"))
		}

		status = allstatus[status]
		// filter by date
		date := c.QueryParam("d")

		// filter by category
		category := strings.ToUpper(c.QueryParam("c"))

		// to prevent sql injection
		if strings.ContainsAny(category, ";") {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Bad request"))
		}

		// default value for category
		if category == "" {
			category = "ALL"
		}

		categories := map[string]string{"ALL": "all", "COMPUTER": "Computer", "COMPUTER-ACCESSORIES": "Computer Accessories", "NETWORKING": "Networking", "UPS": "UPS", "PRINTER-AND-SCANNER": "Printer and Scanner", "ELECTRONICS": "Electronics", "OTHERS": "Others"}

		if _, exist := categories[category]; !exist {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "Bad request"))
		}
		category = categories[category]

		requests, total, err := ac.adminRepository.GetAllProcureManager(limit, offset, status, category, date)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, _common.NoDataResponse(http.StatusInternalServerError, "Failed to read data"))
		}

		return c.JSON(http.StatusOK, _common.GetAllProcureRequestResponse(requests, total))
	}
}
