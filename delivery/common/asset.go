package common

import (
	_entity "capstone/be/entity"
	"net/http"
)

func CreateAssetResponse(asset _entity.AssetSimplified) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success create asset",
		"data": map[string]interface{}{
			"id":            asset.Id,
			"code_asset":    asset.CodeAsset,
			"image":         asset.Image,
			"name":          asset.Name,
			"short_name":    asset.Short_Name,
			"status":        asset.Status,
			"description":   asset.Description,
			"category_name": asset.CategoryName,
		},
	}
}

func GetAllAssetsResponse(assets []_entity.AssetSimplified) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get all assets",
		"data":    assets,
	}
}

func GetAssetByCategoryResponse(asset _entity.AssetSimplified) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get asset by category",
		"data": map[string]interface{}{
			"id":            asset.Id,
			"image":         asset.Image,
			"name":          asset.Name,
			"status":        asset.Status,
			"category_name": asset.CategoryName,
			"description":   asset.Description,
			"quantity":      asset.Quantity,
		},
	}
}

func GetAssetByIdResponse(asset _entity.Asset) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get asset by id",
		"data": map[string]interface{}{
			"id":            asset.Id,
			"image":         asset.Image,
			"name":          asset.Name,
			"status":        asset.Status,
			"category_name": asset.Category.Name,
			"description":   asset.Description,
			"quantity":      asset.Quantity,
		},
	}
}

func UpdateAssetResponse(asset _entity.AssetSimplified) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success edit asset",
		"data": map[string]interface{}{
			"id":          asset.Id,
			"image":       asset.Image,
			"name":        asset.Name,
			"status":      asset.Status,
			"description": asset.Description,
			"quantity":    asset.Quantity,
		},
		// "totalPage": asset.TotalPage,
	}
}

func GetStatsResponse(statistics _entity.Statistics) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get statistics",
		"data":    statistics,
	}
}
