package common

import (
	_entity "capstone/be/entity"
	"net/http"
)

func GetAllRequestResponse(requests []_entity.RequestResponse) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get all requests",
		"data":    requests,
	}
}
