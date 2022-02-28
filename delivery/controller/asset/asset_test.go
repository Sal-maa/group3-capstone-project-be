package asset

import (
	"bytes"
	_midware "capstone/be/delivery/middleware"
	_entity "capstone/be/entity"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockRepoSuccess struct{}

func (m mockRepoSuccess) Create(_entity.Asset) (int, error) {
	return http.StatusOK, nil
}

func (m mockRepoSuccess) GetAll() ([]_entity.AssetSimplified, int, error) {
	return []_entity.AssetSimplified{
		{
			Category:       "Computer",
			Name:           "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
			Image:          "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
			ShortName:      "asset-1645748000",
			Description:    "Processor : Intel Core i7-1165G7, RAM : 8GB DDR4, SSD : 512GB, VGA : Intel Iris Xe Graphics, Konektivitas : Wifi + Bluetooth, Ukuran Layar : 14 Inch FHD, Sistem Operasi : Windows 10 Home",
			UserCount:      4,
			StockAvailable: 20,
		},
	}, http.StatusOK, nil
}

func (m mockRepoSuccess) GetAssetsByCategory(int) ([]_entity.AssetSimplified, int, error) {
	return []_entity.AssetSimplified{
		{
			Category:       "Computer",
			Name:           "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
			Image:          "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
			ShortName:      "asset-1645748000",
			Description:    "Processor : Intel Core i7-1165G7, RAM : 8GB DDR4, SSD : 512GB, VGA : Intel Iris Xe Graphics, Konektivitas : Wifi + Bluetooth, Ukuran Layar : 14 Inch FHD, Sistem Operasi : Windows 10 Home",
			UserCount:      4,
			StockAvailable: 20,
		},
	}, http.StatusOK, nil
}

func (m mockRepoSuccess) GetByShortName(string) (int, int, _entity.AssetSimplified, int, error) {
	return 1, 0, _entity.AssetSimplified{
		Category:       "Computer",
		Name:           "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
		Image:          "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
		ShortName:      "asset-1645748000",
		Description:    "Processor : Intel Core i7-1165G7, RAM : 8GB DDR4, SSD : 512GB, VGA : Intel Iris Xe Graphics, Konektivitas : Wifi + Bluetooth, Ukuran Layar : 14 Inch FHD, Sistem Operasi : Windows 10 Home",
		UserCount:      4,
		StockAvailable: 20,
	}, http.StatusOK, nil
}

func (m mockRepoSuccess) SetMaintenance(string) (int, error) {
	return http.StatusOK, nil
}

func (m mockRepoSuccess) SetAvailable(string) (int, error) {
	return http.StatusOK, nil
}

func (m mockRepoSuccess) GetCategoryId(string) (int, int, error) {
	return 1, http.StatusOK, nil
}

func (m mockRepoSuccess) GetStats() (_entity.Statistics, int, error) {
	return _entity.Statistics{
		TotalAsset:       1,
		UnderMaintenance: 0,
		Borrowed:         0,
		Available:        1,
	}, http.StatusOK, nil
}

// success

func TestCreateSuccess(t *testing.T) {
	t.Run("TestCreateSuccess", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":              "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
			"category":          "Computer",
			"description":       "Processor : Intel Core i7-1165G7, RAM : 8GB DDR4, SSD : 512GB, VGA : Intel Iris Xe Graphics, Konektivitas : Wifi + Bluetooth, Ukuran Layar : 14 Inch FHD, Sistem Operasi : Windows 10 Home",
			"quantity":          10,
			"under_maintenance": false,
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/assets")

		assetController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(assetController.Create())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success create asset",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetAllSuccess(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/assets")

		assetController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(assetController.GetAll())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success get all assets",
			"data": []interface{}{
				map[string]interface{}{
					"category":        "Computer",
					"name":            "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
					"image":           "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
					"short_name":      "asset-1645748000",
					"description":     "Processor : Intel Core i7-1165G7, RAM : 8GB DDR4, SSD : 512GB, VGA : Intel Iris Xe Graphics, Konektivitas : Wifi + Bluetooth, Ukuran Layar : 14 Inch FHD, Sistem Operasi : Windows 10 Home",
					"user_count":      float64(4),
					"stock_available": float64(20),
				},
			},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetByShortNameSuccess(t *testing.T) {
	t.Run("TestGetByShortNameSuccess", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/assets/:short_name")
		context.SetParamNames("short_name")
		context.SetParamValues("asset-1645748000")

		assetController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(assetController.GetByShortName())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success get asset detail",
			"data": map[string]interface{}{
				"category":          "Computer",
				"name":              "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
				"image":             "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
				"description":       "Processor : Intel Core i7-1165G7, RAM : 8GB DDR4, SSD : 512GB, VGA : Intel Iris Xe Graphics, Konektivitas : Wifi + Bluetooth, Ukuran Layar : 14 Inch FHD, Sistem Operasi : Windows 10 Home",
				"total_asset":       float64(1),
				"under_maintenance": float64(0),
			},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUpdateSuccess(t *testing.T) {
	t.Run("TestUpdateSuccess", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"under_maintenance": true,
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/assets/:short_name")
		context.SetParamNames("short_name")
		context.SetParamValues("asset-1645748000")

		assetController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(assetController.Update())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success update asset status",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetStatsSuccess(t *testing.T) {
	t.Run("TestGetStatsSuccess", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/stats")

		assetController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(assetController.GetStats())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success get statistics",
			"data": map[string]interface{}{
				"total_asset":       float64(1),
				"under_maintenance": float64(0),
				"borrowed":          float64(0),
				"available":         float64(1),
			},
		}

		assert.Equal(t, expected, actual)
	})
}

// authorization failure

func TestCreateFailUnauthorized(t *testing.T) {
	t.Run("TestCreateFailUnauthorized", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Employee")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":              "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
			"category":          "Computer",
			"description":       "Processor : Intel Core i7-1165G7, RAM : 8GB DDR4, SSD : 512GB, VGA : Intel Iris Xe Graphics, Konektivitas : Wifi + Bluetooth, Ukuran Layar : 14 Inch FHD, Sistem Operasi : Windows 10 Home",
			"quantity":          10,
			"under_maintenance": false,
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/assets")

		assetController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(assetController.Create())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusUnauthorized),
			"message": "unauthorized",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUpdateFailUnauthorized(t *testing.T) {
	t.Run("TestUpdateFailUnauthorized", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Employee")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"under_maintenance": true,
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/assets/:short_name")
		context.SetParamNames("short_name")
		context.SetParamValues("asset-1645748000")

		assetController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(assetController.Update())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusUnauthorized),
			"message": "unauthorized",
		}

		assert.Equal(t, expected, actual)
	})
}

// binding failure

func TestCreateFailBinding(t *testing.T) {
	t.Run("TestCreateFailBinding", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":              "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
			"category":          "Computer",
			"description":       "Processor : Intel Core i7-1165G7, RAM : 8GB DDR4, SSD : 512GB, VGA : Intel Iris Xe Graphics, Konektivitas : Wifi + Bluetooth, Ukuran Layar : 14 Inch FHD, Sistem Operasi : Windows 10 Home",
			"quantity":          10,
			"under_maintenance": "false",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/assets")

		assetController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(assetController.Create())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusBadRequest),
			"message": "failed to bind data",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUpdateFailBinding(t *testing.T) {
	t.Run("TestUpdateFailBinding", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"under_maintenance": "true",
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/assets/:short_name")
		context.SetParamNames("short_name")
		context.SetParamValues("asset-1645748000")

		assetController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(assetController.Update())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusBadRequest),
			"message": "failed to bind data",
		}

		assert.Equal(t, expected, actual)
	})
}

// empty input

func TestCreateFailEmptyInput(t *testing.T) {
	t.Run("TestCreateFailEmptyInput", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":              "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
			"category":          "Computer",
			"description":       "",
			"quantity":          10,
			"under_maintenance": false,
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/assets")

		assetController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(assetController.Create())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusBadRequest),
			"message": "input cannot be empty",
		}

		assert.Equal(t, expected, actual)
	})
}

// wrong quantity

func TestCreateFailInvalidQuantity(t *testing.T) {
	t.Run("TestCreateFailInvalidQuantity", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":              "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
			"category":          "Computer",
			"description":       "Processor : Intel Core i7-1165G7, RAM : 8GB DDR4, SSD : 512GB, VGA : Intel Iris Xe Graphics, Konektivitas : Wifi + Bluetooth, Ukuran Layar : 14 Inch FHD, Sistem Operasi : Windows 10 Home",
			"quantity":          0,
			"under_maintenance": false,
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/assets")

		assetController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(assetController.Create())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusBadRequest),
			"message": "quantity must be greater than zero",
		}

		assert.Equal(t, expected, actual)
	})
}

// string pattern

func TestCreateFailMaliciousCharacter(t *testing.T) {
	t.Run("TestCreateFailMaliciousCharacter", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":              "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
			"category":          "Computer",
			"description":       "; --",
			"quantity":          10,
			"under_maintenance": false,
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/assets")

		assetController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(assetController.Create())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusBadRequest),
			"message": "; --: input cannot contain forbidden character",
		}

		assert.Equal(t, expected, actual)
	})
}

type mockRepoFail struct{}

func (m mockRepoFail) Create(_entity.Asset) (int, error) {
	return http.StatusInternalServerError, errors.New("internal server error")
}

func (m mockRepoFail) GetAll() ([]_entity.AssetSimplified, int, error) {
	return []_entity.AssetSimplified{}, http.StatusInternalServerError, errors.New("internal server error")
}

func (m mockRepoFail) GetAssetsByCategory(int) ([]_entity.AssetSimplified, int, error) {
	return []_entity.AssetSimplified{}, http.StatusInternalServerError, errors.New("internal server error")
}

func (m mockRepoFail) GetByShortName(string) (int, int, _entity.AssetSimplified, int, error) {
	return 0, 0, _entity.AssetSimplified{}, http.StatusInternalServerError, errors.New("internal server error")
}

func (m mockRepoFail) SetMaintenance(string) (int, error) {
	return http.StatusInternalServerError, errors.New("internal server error")
}

func (m mockRepoFail) SetAvailable(string) (int, error) {
	return http.StatusInternalServerError, errors.New("internal server error")
}

func (m mockRepoFail) GetCategoryId(string) (int, int, error) {
	return 0, http.StatusInternalServerError, errors.New("internal server error")
}

func (m mockRepoFail) GetStats() (_entity.Statistics, int, error) {
	return _entity.Statistics{}, http.StatusInternalServerError, errors.New("internal server error")
}

func TestCreateFailRepo(t *testing.T) {
	t.Run("TestCreateFailRepo", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":              "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
			"category":          "Computer",
			"description":       "Processor : Intel Core i7-1165G7, RAM : 8GB DDR4, SSD : 512GB, VGA : Intel Iris Xe Graphics, Konektivitas : Wifi + Bluetooth, Ukuran Layar : 14 Inch FHD, Sistem Operasi : Windows 10 Home",
			"quantity":          10,
			"under_maintenance": false,
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/assets")

		assetController := New(mockRepoFail{})
		_midware.JWTMiddleWare()(assetController.Create())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusInternalServerError),
			"message": "internal server error",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetAllFailRepo(t *testing.T) {
	t.Run("TestGetAllFailRepo", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/assets")

		assetController := New(mockRepoFail{})
		_midware.JWTMiddleWare()(assetController.GetAll())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusInternalServerError),
			"message": "internal server error",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetByShortNameFailRepo(t *testing.T) {
	t.Run("TestGetByShortNameFailRepo", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/assets/:short_name")
		context.SetParamNames("short_name")
		context.SetParamValues("asset-1645748000")

		assetController := New(mockRepoFail{})
		_midware.JWTMiddleWare()(assetController.GetByShortName())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusInternalServerError),
			"message": "internal server error",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUpdateFailRepo(t *testing.T) {
	t.Run("TestUpdateFailRepo", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"under_maintenance": true,
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/assets/:short_name")
		context.SetParamNames("short_name")
		context.SetParamValues("asset-1645748000")

		assetController := New(mockRepoFail{})
		_midware.JWTMiddleWare()(assetController.Update())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusInternalServerError),
			"message": "internal server error",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetStatsFailRepo(t *testing.T) {
	t.Run("TestGetStatsFailRepo", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/stats")

		assetController := New(mockRepoFail{})
		_midware.JWTMiddleWare()(assetController.GetStats())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusInternalServerError),
			"message": "internal server error",
		}

		assert.Equal(t, expected, actual)
	})
}
