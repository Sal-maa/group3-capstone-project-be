package common

import (
	_entity "capstone/be/entity"
	"net/http"
)

func NoDataResponse(code int, message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    code,
		"message": message,
	}
}

func CreateUserResponse(user _entity.UserSimplified) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success create user",
		"data": map[string]interface{}{
			"id":   user.Id,
			"name": user.Name,
		},
	}
}

func LoginResponse(user _entity.User, token string, expire int64) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success login",
		"data": map[string]interface{}{
			"id":     user.Id,
			"name":   user.Name,
			"token":  token,
			"expire": expire,
			"avatar": user.Avatar,
		},
	}
}

func GetUserByIdResponse(user _entity.User) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get user by id",
		"data": map[string]interface{}{
			"id":     user.Id,
			"name":   user.Name,
			"email":  user.Email,
			"phone":  user.Phone,
			"avatar": user.Avatar,
		},
	}
}

func GetAllUsersResponse(users []_entity.UserSimplified) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success get all users",
		"data":    users,
	}
}

func UpdateUserResponse(user _entity.UserSimplified) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success update user",
		"data":    user,
	}
}
