package common

import (
	_entity "capstone/be/entity"
	"net/http"
)

func CreateAssetResponse(product _entity.Asset) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success create asset",
		"data": map[string]interface{}{
			"id":    product.Id,
			"image": product.Image,
			"name":  product.Name,
			// "entry_date":  product.Entry_date,
			"status": product.Status,
			// "address":     product.Address,
			"description": product.Description,
			"quantity":    product.Quantity,
		},
	}
}
func GetAllCategoryResponse(categories []_entity.Category) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success create categories",
		"data":    categories,
	}
}

func GetAllAssetsResponse(assets []_entity.Asset) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get all assets",
		"data":    assets,
	}
}
func GetAssetByIdResponse(asset _entity.Asset) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get asset by id",
		"data": map[string]interface{}{
			"id":          asset.Id,
			"image":       asset.Image,
			"name":        asset.Name,
			"status":      asset.Status,
			"category_id": asset.CategoryId,
			"description": asset.Description,
			"quantity":    asset.Quantity,
		},
	}
}

func UpdateAssetResponse(product _entity.Asset) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success create asset",
		"data": map[string]interface{}{
			"id":    product.Id,
			"image": product.Image,
			"name":  product.Name,
			// "entry_date":  product.Entry_date,
			"status": product.Status,
			// "address":     product.Address,
			"description": product.Description,
			"quantity":    product.Quantity,
		},
	}
}
