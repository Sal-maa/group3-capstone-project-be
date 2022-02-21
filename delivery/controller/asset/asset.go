package asset

import (
	_common "capstone/be/delivery/common"
	_helper "capstone/be/delivery/helper"
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
	// userRepository _userRepo.User
}

func New(asset _assetRepo.Asset) *AssetController {
	return &AssetController{repository: asset}
}

func (uc AssetController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		// calling repository
		assets, code, err := uc.repository.GetAll()

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		return c.JSON(http.StatusOK, _common.GetAllAssetsResponse(assets))
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
		// role := "Administrator"
		// println(role)
		// if role != _midware.ExtractRole(c) {
		// 	fmt.Println(role)

		// 	log.Println(role)

		// 	return c.JSON(http.StatusUnauthorized, _common.NoDataResponse(http.StatusUnauthorized, "unauthorized role"))
		// }
		assetData := _entity.CreateAsset{}
		userData := _entity.UserSimplified{}
		// detect failure in binding
		if err := c.Bind(&assetData); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}
		if err := c.Bind(&userData); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}

		image := strings.Title(strings.ToLower(strings.TrimSpace(assetData.Image)))
		name := strings.Title(strings.ToLower(strings.TrimSpace(assetData.Name)))
		// entry_date := strings.Title(strings.ToLower(strings.TrimSpace(assetData.Entry_date)))

		// entry_date := assetData.Entry_date
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
			// if entry_date == "" {
			// 	return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "input cannot be empty"))
			// }
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
			Image: assetData.Image,
			Name:  assetData.Name,
			// Entry_date:  assetData.Entry_date,
			Status:      assetData.Status,
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

func (uc AssetController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid asset id"))
		}

		// // check authorization
		// if id != _midware.ExtractId(c) {
		// 	return c.JSON(http.StatusUnauthorized, _common.NoDataResponse(http.StatusUnauthorized, "unauthorized"))
		// }

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
			// if entry_date == "" {
			// 	return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "input cannot be empty"))
			// }
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
		// detect avatar image upload

		// calling repository
		UpdateAsset, code, err := uc.repository.Update(updateAssetData)

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		return c.JSON(http.StatusOK, _common.UpdateAssetResponse(UpdateAsset))
	}
}
