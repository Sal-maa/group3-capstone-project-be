package common

import (
	_entity "capstone/be/entity"
	"net/http"
)

func LoginResponse(user _entity.User, token string, expire int64) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success login",
		"data": map[string]interface{}{
			"id":     user.Id,
			"role":   user.Role,
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
			"id":       user.Id,
			"division": user.Division,
			"role":     user.Role,
			"name":     user.Name,
			"email":    user.Email,
			"phone":    user.Phone,
			"gender":   user.Gender,
			"address":  user.Address,
			"avatar":   user.Avatar,
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
