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

func CreateAssetResponse(asset _entity.Asset) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "success create asset",
		"data": map[string]interface{}{
			"id":          asset.Id,
			"image":       asset.Image,
			"name":        asset.Name,
			"status":      asset.Status,
			"description": asset.Description,
			"quantity":    asset.Quantity,
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

func GetAllAssetsResponse(assets []_entity.AssetSimplified) map[string]interface{} {
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

func UpdateAssetResponse(asset _entity.Asset) map[string]interface{} {
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
	}
}
