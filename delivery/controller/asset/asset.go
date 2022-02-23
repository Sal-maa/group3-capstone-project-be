package asset

import (
	_common "capstone/be/delivery/common"
	_helper "capstone/be/delivery/helper"
	"capstone/be/delivery/middleware"
	"fmt"
	"time"

	_entity "capstone/be/entity"
	_assetRepo "capstone/be/repository/asset"

	// _userRepo "capstone/be/repository/user"

	"log"
	"strconv"

	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type AssetController struct {
	repository _assetRepo.Asset
}

func New(asset _assetRepo.Asset) *AssetController {
	return &AssetController{repository: asset}
}

func (uc AssetController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		p := strings.TrimSpace(c.QueryParam("page"))
		log.Println(p)
		if p == "" {
			p = "1"
		}

		page, err := strconv.Atoi(p)
		log.Println(page)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid page number"))
		}
		// limit := 8
		assets, code, err := uc.repository.GetAll(page)
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}
		return c.JSON(http.StatusOK, _common.GetAllAssetsResponse(assets))
	}
}
func (uc AssetController) GetAssetByCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		p := strings.TrimSpace(c.QueryParam("page"))
		log.Println(p)
		if p == "" {
			p = "1"
		}
		page, err := strconv.Atoi(p)
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid page number"))
		}
		category := c.QueryParam("category")

		asset, code, err := uc.repository.GetAssetByCategory(category, page)
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}
		return c.JSON(http.StatusOK, _common.GetAssetByCategoryResponse(asset))
	}
}
func (uc AssetController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		// detect invalid parameter
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid asset id"))
		}
		// calling repository
		asset, code, err := uc.repository.GetById(id)
		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}
		return c.JSON(http.StatusOK, _common.GetAssetByIdResponse(asset))
	}
}

func (uc AssetController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := middleware.ExtractRole(c)
		if role != "Administrator" {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "you don't have permission"))
		}

		assetData := _entity.CreateAsset{}
		if err := c.Bind(&assetData); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}

		image := strings.Title(strings.ToLower(strings.TrimSpace(assetData.Image)))
		name := strings.Title(strings.ToLower(strings.TrimSpace(assetData.Name)))
		status := strings.Title(strings.ToLower(strings.TrimSpace(assetData.Status)))
		description := strings.Title(strings.ToLower(strings.TrimSpace(assetData.Description)))
		quantity := assetData.Quantity
		check := []string{image, name, status, description}

		for _, s := range check {
			// check empty string in required input
			if image == "" {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "input cannot be empty"))
			}
			if name == "" {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "input cannot be empty"))
			}
			if status == "" {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "input cannot be empty"))
			}
			if description == "" {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "input cannot be empty"))
			}
			// check malicious character in input
			if err := _helper.CheckStringInput(s); err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, s+": "+err.Error()))
			}
		}
		if quantity == 0 || quantity < 0 {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "quantity cannot be null or negative value"))
		}
		// prepare input to repository
		createAssetData := _entity.Asset{
			CodeAsset:   assetData.CodeAsset,
			Image:       assetData.Image,
			Name:        assetData.Name,
			Short_Name:  assetData.Short_Name,
			Status:      assetData.Status,
			Description: assetData.Description,
			Quantity:    assetData.Quantity,
		}
		// calling repository
		createAsset, code, err := uc.repository.Create(createAssetData)

		// var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

		// s := make([]rune, 5)
		// for i := range s {
		// 	s[i] = letters[rand.Intn(len(letters))]
		// }
		short_name := fmt.Sprintf("asset-%d", (time.Now().Unix()))

		for i := 0; i < createAssetData.Quantity; i++ {
			createAssetData.CodeAsset = fmt.Sprintf("%s-%d", short_name, i)
			createAsset, _, _ = uc.repository.Create(createAssetData)
		}

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}
		return c.JSON(http.StatusOK, _common.CreateAssetResponse(createAsset))
	}
}

func (uc AssetController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid asset id"))
		}

		role := middleware.ExtractRole(c)
		if role != "Administrator" {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "you don't have permission"))
		}
		assetData := _entity.UpdateAsset{}

		// detect failure in binding
		if err := c.Bind(&assetData); err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}
		// prepare input string
		image := strings.Title(strings.ToLower(strings.TrimSpace(assetData.Image)))
		name := strings.Title(strings.ToLower(strings.TrimSpace(assetData.Name)))
		status := strings.Title(strings.ToLower(strings.TrimSpace(assetData.Status)))
		description := strings.Title(strings.ToLower(strings.TrimSpace(assetData.Description)))
		quantity := assetData.Quantity
		// calling repository to get existing user data
		updateAssetData, code, err := uc.repository.GetById(id)

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		check := []string{image, name, status, description}

		for _, s := range check {
			// check empty string in required input
			if image == "" {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "input cannot be empty"))
			}
			if name == "" {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "input cannot be empty"))
			}
			if status == "" {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "input cannot be empty"))
			}
			if description == "" {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "input cannot be empty"))
			}
			// check malicious character in input
			if err := _helper.CheckStringInput(s); err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, s+": "+err.Error()))
			}
		}
		if quantity == 0 || quantity < 0 {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "quantity cannot be null or negative value"))
		}
		// calling repository
		UpdateAsset, code, err := uc.repository.Update(updateAssetData)

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}
		return c.JSON(http.StatusOK, _common.UpdateAssetResponse(UpdateAsset))
	}
}
