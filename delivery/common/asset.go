package common

import (
	_entity "capstone/be/entity"
	"net/http"
)

func CreateAssetResponse() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success create asset",
	}
}

func GetAllAssetsResponse(assets []_entity.AssetSimplified) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get all assets",
		"data":    assets,
	}
}

func GetByShortNameResponse(total int, asset _entity.AssetSimplified) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get asset detail",
		"data": map[string]interface{}{
			"category":    asset.Category,
			"name":        asset.Name,
			"image":       asset.Image,
			"description": asset.Description,
			"total_asset": total,
		},
	}
}

func UpdateAssetResponse() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success update asset status",
	}
}

func GetStatsResponse(statistics _entity.Statistics) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get statistics",
		"data":    statistics,
	}
}
