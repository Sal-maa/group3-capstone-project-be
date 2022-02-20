package asset

import (
	_common "capstone/be/delivery/common"
	_helper "capstone/be/delivery/helper"
	_entity "capstone/be/entity"
	_assetRepo "capstone/be/repository/asset"

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

func (uc AssetController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		assetData := _entity.CreateAsset{}

		// detect failure in binding
		if err := c.Bind(&assetData); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}

		image := strings.Title(strings.ToLower(strings.TrimSpace(assetData.Image)))
		name := strings.Title(strings.ToLower(strings.TrimSpace(assetData.Name)))
		// entry_date := assetData.Entry_date
		status := strings.Title(strings.ToLower(strings.TrimSpace(assetData.Status)))
		address := strings.Title(strings.ToLower(strings.TrimSpace(assetData.Address)))
		description := strings.Title(strings.ToLower(strings.TrimSpace(assetData.Description)))
		quantity := assetData.Quantity

		check := []string{image, name, status, address, description}
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
			if address == "" {
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
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "price cannot be null or negative value"))
		}

		// prepare input to repository
		createAssetData := _entity.Asset{
			Image:       assetData.Image,
			Name:        assetData.Name,
			Entry_date:  assetData.Entry_date,
			Status:      assetData.Status,
			Address:     assetData.Address,
			Description: assetData.Description,
			Quantity:    assetData.Quantity,
		}

		// calling repository
		createAsset, code, err := uc.repository.Create(createAssetData)

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		return c.JSON(http.StatusOK, _common.CreateAssetResponse(createAsset))
	}
}
