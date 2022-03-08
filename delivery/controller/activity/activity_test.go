package activity

import (
	"bytes"
	_midware "capstone/be/delivery/middleware"
	_entity "capstone/be/entity"
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"net/http"
	"net/http/httptest"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockRepoSuccess struct{}

func (m mockRepoSuccess) GetAllActivityOfUser(int) ([]_entity.ActivitySimplified, int, error) {
	return []_entity.ActivitySimplified{
		{
			Id:           1,
			AssetImage:   "test",
			AssetName:    "test",
			ActivityType: "test",
			Status:       "Waiting approval",
			RequestDate:  time.Time{},
		},
	}, http.StatusOK, nil
}
func (m mockRepoSuccess) GetDetailActivityByRequestId(int) (_entity.Activity, int, error) {
	return _entity.Activity{
		Id:           1,
		Category:     "test",
		AssetImage:   "test",
		AssetName:    "test",
		StockLeft:    1,
		UserName:     "test",
		ActivityType: "test",
		Status:       "Waiting approval",
		Description:  "test",
		RequestDate:  time.Time{},
		ReturnDate:   time.Time{},
	}, http.StatusOK, nil
}

func (m mockRepoSuccess) CancelRequest(int) (int, error) {
	return http.StatusOK, nil
}

func (m mockRepoSuccess) ReturnRequest(int) (int, error) {
	return http.StatusOK, nil
}

func TestGetAllActivityOfUser(t *testing.T) {
	t.Run("TestGetAllActivityOfUser", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/activities/:user_id")
		context.SetParamNames("user_id")
		context.SetParamValues("1")
		activityController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(activityController.GetAllActivityOfUser())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success get all activities",
			"data": []interface{}{map[string]interface{}{
				"activity_type": "test",
				"asset_image":   "test",
				"asset_name":    "test",
				"id":            float64(1),
				"request_date":  "0001-01-01T00:00:00Z",
				"status":        "Waiting approval"}},
		}
		assert.Equal(t, expected, actual)
	})
}

func TestGetDetailActivityByRequestId(t *testing.T) {
	t.Run("TestGetDetailActivityByRequestId", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/activities/:user_id/:request_id")
		context.SetParamNames("user_id", "request_id")
		context.SetParamValues("1", "1")
		activityController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(activityController.GetDetailActivityByRequestId())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success get detail activity",
			"data": map[string]interface{}{
				"activity":     "test",
				"asset_image":  "test",
				"asset_name":   "test",
				"category":     "test",
				"description":  "test",
				"id":           float64(1),
				"request_date": "0001-01-01T00:00:00Z",
				"return_date":  "0001-01-01T00:00:00Z",
				"status":       "Waiting approval",
				"stock_left":   float64(1),
				"user_name":    "test",
			},
		}
		assert.Equal(t, expected, actual)
	})
}

func TestUpdateRrquestStatus(t *testing.T) {
	t.Run("TestUpdateRrquestStatus", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"status": "cancel",
		})
		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/activities/:user_id/:request_id")
		context.SetParamNames("user_id", "request_id")
		context.SetParamValues("1", "1")
		activityController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(activityController.UpdateRequestStatus())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success cancel request",
			"data": map[string]interface{}{
				"activity_type": "test",
				"asset_image":   "test",
				"asset_name":    "test",
				"category":      "test",
				"id":            float64(1),
				"note":          "test",
				"request_date":  "0001-01-01T00:00:00Z",
				"return_date":   "0001-01-01T00:00:00Z",
				"status":        "Cancelled",
				"stock_left":    float64(1),
				"user_name":     "test"},
		}
		assert.Equal(t, expected, actual)
	})
}

// invalid user id
func TestGetAllActivityOfUserInvalidUSerid(t *testing.T) {
	t.Run("TestGetAllActivityOfUserInvalidUSerid", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/activities/:user_id")
		context.SetParamNames("user_id")
		context.SetParamValues("test")
		activityController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(activityController.GetAllActivityOfUser())(context)

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
func TestGetDetailActivityByRequestIdInvalidUserid(t *testing.T) {
	t.Run("TestGetDetailActivityByRequestIdInvalidUserid", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/activities/:user_id/:request_id")
		context.SetParamNames("user_id", "request_id")
		context.SetParamValues("test", "1")
		activityController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(activityController.GetDetailActivityByRequestId())(context)

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

// invalid request id
func TestGetDetailActivityByRequestIdInvalidReqid(t *testing.T) {
	t.Run("TestGetDetailActivityByRequestIdInvalidReqid", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/activities/:user_id/:request_id")
		context.SetParamNames("user_id", "request_id")
		context.SetParamValues("1", "test")
		activityController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(activityController.GetDetailActivityByRequestId())(context)

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

//unautorized
func TestGetDetailActivityByRequestIdUnauth(t *testing.T) {
	t.Run("TestGetDetailActivityByRequestIdInvalidReqid", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(3, "")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/activities/:user_id/:request_id")
		context.SetParamNames("user_id", "request_id")
		context.SetParamValues("1", "1")
		activityController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(activityController.GetDetailActivityByRequestId())(context)

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

//unauthorized

func TestGetAllActivityOfUserUnauthorized(t *testing.T) {
	t.Run("TestGetAllActivityOfUserUnauthorized", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Employee")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/activities/:user_id")
		context.SetParamNames("user_id")
		context.SetParamValues("2")
		activityController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(activityController.GetAllActivityOfUser())(context)

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

func TestGetDetailActivityByRequestIdUnauthorized(t *testing.T) {
	t.Run("TestGetDetailActivityByRequestIdUnauthorized", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Employee")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/activities/:user_id/:request_id")
		context.SetParamNames("user_id", "request_id")
		context.SetParamValues("2", "1")
		activityController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(activityController.GetAllActivityOfUser())(context)

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

func TestUpdateRrquestStatusRetrunAssset(t *testing.T) {
	t.Run("TestUpdateRrquestStatusFailRepo", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(2, "Employee")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"status": "cancel",
		})
		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/activities/:user_id/:request_id")
		context.SetParamNames("user_id", "request_id")
		context.SetParamValues("1", "1")
		activityController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(activityController.UpdateRequestStatus())(context)

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

//cannot return asset
func TestUpdateRrquestStatusRetrunAsset(t *testing.T) {
	t.Run("TestUpdateRrquestStatusFailRepo", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"status": "return",
		})
		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/activities/:user_id/:request_id")
		context.SetParamNames("user_id", "request_id")
		context.SetParamValues("1", "1")
		activityController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(activityController.UpdateRequestStatus())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusBadRequest),
			"message": "cannot return asset",
		}
		assert.Equal(t, expected, actual)
	})
}

//invalid input status
func TestUpdateRrquestStatusFailRepo(t *testing.T) {
	t.Run("TestUpdateRrquestStatusFailRepo", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodPut, "/", nil)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/activities/:user_id/:request_id")
		context.SetParamNames("user_id", "request_id")
		context.SetParamValues("1", "1")
		activityController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(activityController.UpdateRequestStatus())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusBadRequest),
			"message": "invalid input status",
		}
		assert.Equal(t, expected, actual)
	})
}

// invalid user id
func TestUpdateRrquestStatusFailRepoInvalidUSerid(t *testing.T) {
	t.Run("TestUpdateRrquestStatusFailRepoInvalidUSerid", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodPut, "/", nil)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/activities/:user_id/:request_id")
		context.SetParamNames("user_id", "request_id")
		context.SetParamValues("s", "1")
		activityController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(activityController.UpdateRequestStatus())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code": float64(400), "message": "invalid user id",
		}
		assert.Equal(t, expected, actual)
	})
}

// invalid req id
func TestUpdateRrquestStatusFsailRepoInvalidReqid(t *testing.T) {
	t.Run("TestUpdateRrquestStatusFsailRepoInvalidReqid", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodPut, "/", nil)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/activities/:user_id/:request_id")
		context.SetParamNames("user_id", "request_id")
		context.SetParamValues("1", "s")
		activityController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(activityController.UpdateRequestStatus())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(400),
			"message": "invalid request id",
		}
		assert.Equal(t, expected, actual)
	})
}

type mockRepoFail struct{}

func (m mockRepoFail) GetAllActivityOfUser(int) ([]_entity.ActivitySimplified, int, error) {
	return []_entity.ActivitySimplified{}, http.StatusInternalServerError, errors.New("internal server error")
}
func (m mockRepoFail) GetDetailActivityByRequestId(int) (_entity.Activity, int, error) {
	return _entity.Activity{}, http.StatusInternalServerError, errors.New("internal server error")
}

func (m mockRepoFail) CancelRequest(int) (int, error) {
	return http.StatusInternalServerError, errors.New("internal server error")
}

func (m mockRepoFail) ReturnRequest(int) (int, error) {
	return http.StatusInternalServerError, errors.New("internal server error")
}

func (m mockRepoFail) UpdateRequestStatus(int) (int, error) {
	return http.StatusInternalServerError, errors.New("internal server error")
}
func TestGetAllActivityOfUserFaildRepo(t *testing.T) {
	t.Run("TestGetAllActivityOfUserFaildRepo", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/activities/:user_id")
		context.SetParamNames("user_id")
		context.SetParamValues("1")
		activityController := New(mockRepoFail{})
		_midware.JWTMiddleWare()(activityController.GetAllActivityOfUser())(context)

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

func TestGetDetailActivityByRequestIdFaildRepo(t *testing.T) {
	t.Run("TestGetDetailActivityByRequestIdFaildRepo", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/activities/:user_id/:request_id")
		context.SetParamNames("user_id", "request_id")
		context.SetParamValues("1", "1")
		activityController := New(mockRepoFail{})
		_midware.JWTMiddleWare()(activityController.GetDetailActivityByRequestId())(context)

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

func TestUpdateRrquestStastus(t *testing.T) {
	t.Run("TestUpdateRrquestStatus", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"status": "cancel",
		})
		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/activities/:user_id/:request_id")
		context.SetParamNames("user_id", "request_id")
		context.SetParamValues("1", "1")
		activityController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(activityController.UpdateRequestStatus())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success cancel request",
			"data": map[string]interface{}{
				"activity_type": "test",
				"asset_image":   "test",
				"asset_name":    "test",
				"category":      "test",
				"id":            float64(1),
				"note":          "test",
				"request_date":  "0001-01-01T00:00:00Z",
				"return_date":   "0001-01-01T00:00:00Z",
				"status":        "Cancelled",
				"stock_left":    float64(1),
				"user_name":     "test"},
		}
		assert.Equal(t, expected, actual)
	})
}

//failed bind data
func TestUpdateRrquestStastusBindfail(t *testing.T) {
	t.Run("TestUpdateRrquestStatus", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"status": 1,
		})
		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/activities/:user_id/:request_id")
		context.SetParamNames("user_id", "request_id")
		context.SetParamValues("1", "1")
		activityController := New(mockRepoFail{})
		_midware.JWTMiddleWare()(activityController.UpdateRequestStatus())(context)

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

func TestUpdateRrquestStastusFaild(t *testing.T) {
	t.Run("TestUpdateRrquestStatus", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"status": "cancel",
		})
		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/activities/:user_id/:request_id")
		context.SetParamNames("user_id", "request_id")
		context.SetParamValues("1", "1")
		activityController := New(mockRepoFail{})
		_midware.JWTMiddleWare()(activityController.UpdateRequestStatus())(context)

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
