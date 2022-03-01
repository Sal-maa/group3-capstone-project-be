package history

import (
	_midware "capstone/be/delivery/middleware"
	_entity "capstone/be/entity"
	"errors"

	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockRepoSuccess struct{}

func (m mockRepoSuccess) GetAllRequestHistoryOfUser(int, int) (int, []_entity.UserRequestHistorySimplified, int, error) {
	return 1, []_entity.UserRequestHistorySimplified{
		{
			Id:           1,
			Category:     "test",
			AssetName:    "test",
			AssetImage:   "test",
			RequestDate:  time.Time{},
			ActivityType: "test",
		},
	}, http.StatusOK, nil
}
func (m mockRepoSuccess) GetDetailRequestHistoryByRequestId(int) (_entity.UserRequestHistory, int, error) {
	return _entity.UserRequestHistory{
		Id:          1,
		Category:    "test",
		AssetName:   "test",
		AssetImage:  "test",
		StockLeft:   1,
		UserName:    "test",
		RequestDate: time.Time{},
		ReturnDate:  time.Time{},
		Status:      "test",
		Description: "test",
	}, http.StatusOK, nil
}
func (m mockRepoSuccess) GetAllUsageHistoryOfAsset(string) (_entity.AssetInfo, []_entity.AssetUser, int, error) {
	return _entity.AssetInfo{
			Category:   "",
			AssetName:  "",
			AssetImage: "",
		}, []_entity.AssetUser{
			{
				Name:        "",
				RequestDate: time.Time{},
				Status:      "",
			},
		}, http.StatusOK, nil
}

func TestGetAllRequestHistoryOfUser(t *testing.T) {
	t.Run("TestGetAllRequestHistoryOfUser", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/histories/users/:user_id")
		context.SetParamNames("user_id")
		context.SetParamValues("1")
		historyController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(historyController.GetAllRequestHistoryOfUser())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success get all histories",
			"data": map[string]interface{}{
				"count": float64(1),
				"histories": []interface{}{map[string]interface{}{
					"activity_type": "test",
					"asset_image":   "test",
					"asset_name":    "test",
					"category":      "test",
					"id":            float64(1),
					"request_date":  "0001-01-01T00:00:00Z"}},
			},
		}
		assert.Equal(t, expected, actual)
	})
}

func TestGetDetailRequestHistoryByRequestId(t *testing.T) {
	t.Run("TestGetDetailRequestHistoryByRequestId", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/histories/users/:user_id/:request_id")
		context.SetParamNames("user_id", "request_id")
		context.SetParamValues("1", "1")
		historyController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(historyController.GetDetailRequestHistoryByRequestId())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success get detail history",
			"data": map[string]interface{}{
				"asset_image":  "test",
				"asset_name":   "test",
				"category":     "test",
				"description":  "test",
				"id":           float64(1),
				"request_date": "0001-01-01T00:00:00Z",
				"return_date":  "0001-01-01T00:00:00Z",
				"status":       "test",
				"stock_left":   float64(1),
				"user_name":    "test",
			},
		}
		assert.Equal(t, expected, actual)
	})
}

func TestGetAllUsageHistoryOfAsset(t *testing.T) {
	t.Run("TestGetAllUsageHistoryOfAsset", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/histories/users/:short_name")
		context.SetParamNames("short_name")
		context.SetParamValues("dell")
		historyController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(historyController.GetAllUsageHistoryOfAsset())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success get all histories",
			"data": map[string]interface{}{
				"asset_image": "",
				"asset_name":  "",
				"category":    "",
				"users": []interface{}{map[string]interface{}{
					"request_date": "0001-01-01T00:00:00Z",
					"status":       "",
					"user_name":    ""}}},
		}
		assert.Equal(t, expected, actual)
	})
}

// authorization failure
func TestGetAllRequestHistoryOfUserUnauthorized(t *testing.T) {
	t.Run("TestGetAllRequestHistoryOfUserUnauthorized", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Employee")

		request := httptest.NewRequest(http.MethodPost, "/", nil)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/histories/users/:user_id")
		context.SetParamNames("user_id")
		context.SetParamValues("2")

		assetController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(assetController.GetAllRequestHistoryOfUser())(context)

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
func TestGetDetailRequestHistoryByRequestIdUnauthorized(t *testing.T) {
	t.Run("TestGetDetailRequestHistoryByRequestIdUnauthorized", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Employee")

		request := httptest.NewRequest(http.MethodPost, "/", nil)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/histories/users/:user_id/:request_id")
		context.SetParamNames("user_id", "request_id")
		context.SetParamValues("2", "1")

		assetController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(assetController.GetDetailRequestHistoryByRequestId())(context)

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

// invalid user id
func TestGetAllRequestHistoryOfUserInvalidUserid(t *testing.T) {
	t.Run("TestGetAllRequestHistoryOfUserInvalidUserid", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Employee")

		request := httptest.NewRequest(http.MethodPost, "/", nil)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/histories/users/:user_id")
		context.SetParamNames("user_id")
		context.SetParamValues("test")

		assetController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(assetController.GetDetailRequestHistoryByRequestId())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusBadRequest),
			"message": "invalid user id",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetDetailRequestHistoryByRequestIdInvalidUserid(t *testing.T) {
	t.Run("TestGetAllRequestHistoryOfUserUnauthorized", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Employee")

		request := httptest.NewRequest(http.MethodPost, "/", nil)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/histories/users/:user_id/:request_id")
		context.SetParamNames("user_id", "request_id")
		context.SetParamValues("test", "1")

		assetController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(assetController.GetDetailRequestHistoryByRequestId())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusBadRequest),
			"message": "invalid user id",
		}

		assert.Equal(t, expected, actual)
	})
}

//invalid request id
func TestGetDetailRequestHistoryByRequestIdInvalidRequestId(t *testing.T) {
	t.Run("TestGetDetailRequestHistoryByRequestIdInvalidRequestId", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Employee")

		request := httptest.NewRequest(http.MethodPost, "/", nil)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/histories/users/:user_id/:request_id")
		context.SetParamNames("user_id", "request_id")
		context.SetParamValues("1", "test")

		assetController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(assetController.GetDetailRequestHistoryByRequestId())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusBadRequest),
			"message": "invalid request id",
		}

		assert.Equal(t, expected, actual)
	})
}

type mockRepoFail struct{}

func (m mockRepoFail) GetAllRequestHistoryOfUser(int, int) (int, []_entity.UserRequestHistorySimplified, int, error) {
	return 0, []_entity.UserRequestHistorySimplified{}, http.StatusInternalServerError, errors.New("internal server error")
}
func (m mockRepoFail) GetDetailRequestHistoryByRequestId(int) (_entity.UserRequestHistory, int, error) {
	return _entity.UserRequestHistory{}, http.StatusInternalServerError, errors.New("internal server error")
}
func (m mockRepoFail) GetAllUsageHistoryOfAsset(string) (_entity.AssetInfo, []_entity.AssetUser, int, error) {
	return _entity.AssetInfo{}, []_entity.AssetUser{}, http.StatusInternalServerError, errors.New("internal server error")
}

func TestGetAllRequestHistoryOfUserFaildRepo(t *testing.T) {
	t.Run("TestGetAllRequestHistoryOfUserFaildRepo", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodPost, "/", nil)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/histories/users/:user_id")
		context.SetParamNames("user_id")
		context.SetParamValues("1")

		assetController := New(mockRepoFail{})
		_midware.JWTMiddleWare()(assetController.GetAllRequestHistoryOfUser())(context)

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

func TestGetDetailRequestHistoryByRequestIdFaildRepo(t *testing.T) {
	t.Run("TestGetDetailRequestHistoryByRequestIdFaildRepo", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodPost, "/", nil)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/histories/users/:user_id/:request_id")
		context.SetParamNames("user_id", "request_id")
		context.SetParamValues("1", "1")

		assetController := New(mockRepoFail{})
		_midware.JWTMiddleWare()(assetController.GetDetailRequestHistoryByRequestId())(context)

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

func TestGetAllUsageHistoryOfAssetFaildRepo(t *testing.T) {
	t.Run("TestGetAllUsageHistoryOfAssetFaildRepo", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodPost, "/", nil)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/histories/users/:short_name")
		context.SetParamNames("short_name")
		context.SetParamValues("dell")

		assetController := New(mockRepoFail{})
		_midware.JWTMiddleWare()(assetController.GetAllUsageHistoryOfAsset())(context)

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
