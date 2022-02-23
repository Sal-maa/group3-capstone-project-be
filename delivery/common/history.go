package common

import (
	_entity "capstone/be/entity"
	"net/http"
)

func GetAllRequestHistoryOfUserResponse(histories []_entity.UserRequestHistorySimplified, count int) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get all histories",
		"data": map[string]interface{}{
			"count":     count,
			"histories": histories,
		},
	}
}

func GetDetailRequestHistoryByRequestId(history _entity.UserRequestHistory) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get detail history",
		"data":    history,
	}
}

func GetAllUsageHistoryOfAsset(histories []_entity.AssetUsageHistory) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get all histories",
		"data":    histories,
	}
}
