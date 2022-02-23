package common

import (
	_entity "capstone/be/entity"
	"net/http"
)

func GetAllActivityOfUserResponse(activities []_entity.ActivitySimplified) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get all activities",
		"data":    activities,
	}
}

func GetDetailActivityByRequestId(activity _entity.Activity) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get detail activity",
		"data": map[string]interface{}{
			"id":           activity.Id,
			"category":     activity.Category,
			"asset_name":   activity.AssetName,
			"asset_image":  activity.AssetImage,
			"user_name":    activity.UserName,
			"request_date": activity.RequestDate,
			"return_date":  activity.ReturnDate,
			"status":       activity.Status,
			"description":  activity.Description,
			"stock_left":   activity.StockLeft,
		},
	}
}

func UpdateRequestStatusResponse(activity _entity.Activity, info string) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success " + info + " request",
		"data":    activity,
	}
}
