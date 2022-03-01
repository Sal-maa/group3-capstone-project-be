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

func (ac AssetController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		// check authorization
		if role := middleware.ExtractRole(c); role != "Administrator" {
			return c.JSON(http.StatusUnauthorized, _common.NoDataResponse(http.StatusUnauthorized, "unauthorized"))
		}

		assetData := _entity.CreateAsset{}

		// detect failure in binding
		if err := c.Bind(&assetData); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}

		// prepare input
		name := strings.TrimSpace(assetData.Name)
		category := strings.TrimSpace(assetData.Category)
		description := strings.TrimSpace(assetData.Description)

		// check input string
		check := []string{name, category, description}

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

		// check input quantity
		if assetData.Quantity <= 0 {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "quantity must be greater than zero"))
		}

		// prepare input to repository
		createAssetData := _entity.Asset{}
		createAssetData.Name = name
		createAssetData.Description = description

		// get category id via repository
		id, code, err := ac.repository.GetCategoryId(category)

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		createAssetData.CategoryId = id

		short_name := fmt.Sprintf("asset-%d", (time.Now().Unix()))
		createAssetData.ShortName = short_name

		if assetData.UnderMaintenance {
			createAssetData.Status = "Asset Under Maintenance"
		} else {
			createAssetData.Status = "Available"
		}

		// detect asset image upload
		src, file, err := c.Request().FormFile("image")

		// detect failure in parsing file
		switch err {
		case nil:
			defer src.Close()

			// upload avatar to amazon s3
			id := middleware.ExtractId(c)
			image, code, err := _helper.UploadImage("asset", id, file, src)

			// detect failure while uploading avatar
			if err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}

			createAssetData.Image = image
		case http.ErrMissingFile:
			log.Println(err)
			createAssetData.Image = "default_image.png"
		case http.ErrNotMultipart:
			log.Println(err)
			createAssetData.Image = "default_image.png"
		default:
			log.Println(err)
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to upload image"))
		}

		for i := 1; i <= assetData.Quantity; i++ {
			createAssetData.Code = fmt.Sprintf("%s-%d", short_name, i)

			// calling repository
			code, err := ac.repository.Create(createAssetData)

			// detect failure in repository
			if err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}
		}

		return c.JSON(http.StatusOK, _common.CreateAssetResponse())
	}
}

func (ac AssetController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		// filter by category
		category := strings.Title(strings.TrimSpace(c.QueryParam("category")))

		if category != "" {
			// get category id via repository
			id, code, err := ac.repository.GetCategoryId(category)

			// detect failure in repository
			if err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}

			// calling repository
			assets, code, err := ac.repository.GetAssetsByCategory(id)

			// detect failure in repository
			if err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}

			return c.JSON(http.StatusOK, _common.GetAllAssetsResponse(assets))
		}

		// calling repository
		assets, code, err := ac.repository.GetAll()

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		return c.JSON(http.StatusOK, _common.GetAllAssetsResponse(assets))
	}
}

func (ac AssetController) GetByShortName() echo.HandlerFunc {
	return func(c echo.Context) error {
		short_name := c.Param("short_name")

		// calling repository
		total, maintenance, asset, code, err := ac.repository.GetByShortName(short_name)

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		return c.JSON(http.StatusOK, _common.GetByShortNameResponse(total, maintenance, asset))
	}
}

func (ac AssetController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		short_name := c.Param("short_name")

		if role := middleware.ExtractRole(c); role != "Administrator" {
			return c.JSON(http.StatusUnauthorized, _common.NoDataResponse(http.StatusUnauthorized, "unauthorized"))
		}

		updateData := _entity.UpdateAsset{}

		// detect failure in binding
		if err := c.Bind(&updateData); err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}

		if updateData.UnderMaintenance {
			// calling repository
			code, err := ac.repository.SetMaintenance(short_name)

			// detect failure in repository
			if err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}

			return c.JSON(http.StatusOK, _common.UpdateAssetResponse())
		}

		// calling repository
		code, err := ac.repository.SetAvailable(short_name)

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		return c.JSON(http.StatusOK, _common.UpdateAssetResponse())
	}
}

func (ac AssetController) GetStats() echo.HandlerFunc {
	return func(c echo.Context) error {
		// calling repository
		statistics, code, err := ac.repository.GetStats()

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		return c.JSON(http.StatusOK, _common.GetStatsResponse(statistics))
	}
}
