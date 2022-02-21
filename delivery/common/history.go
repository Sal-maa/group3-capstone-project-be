package common

import (
	_entity "capstone/be/entity"
	"net/http"
)

func GetAllUsageHistoryOfUserResponse(histories []_entity.UserUsageHistorySimplified) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get all histories",
		"data":    histories,
	}
}

func GetDetailUsageHistoryByRequestId(history _entity.UserUsageHistory) map[string]interface{} {
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
