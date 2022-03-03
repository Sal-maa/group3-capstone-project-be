package admin

// import (
// 	_midware "capstone/be/delivery/middleware"
// 	_entity "capstone/be/entity"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	"github.com/labstack/echo/v4"
// 	"github.com/magiconair/properties/assert"
// )

// type mockRepoSuccess struct{}

// func (m mockRepoSuccess) GetAllAdmin(limit, offset int, status, category, date string) (requests []_entity.RequestResponse, total int, err error) {
// 	return []_entity.RequestResponse{
// 		{
// 			Id: 1,
// 			User: _entity.User{
// 				Id:   1,
// 				Name: "Siska Kohl",
// 			},
// 			Asset: _entity.AssetReq{
// 				Id:           1,
// 				Code:         "asset-1645748000-1",
// 				Name:         "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
// 				CategoryName: "Computer",
// 				Image:        "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
// 			},
// 			Activity:    "Borrow",
// 			RequestTime: time.Date(2022, 02, 27, 12, 23, 23, 0, time.UTC),
// 			ReturnTime:  time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC),
// 			Status:      "Approved by Admin",
// 			Description: "try to borrow",
// 		},
// 	}, 1, nil
// }

// func (m mockRepoSuccess) GetAllManager(divLogin, limit, offset int, status, category, date string) (requests []_entity.RequestResponse, total int, err error) {
// 	return []_entity.RequestResponse{
// 		{
// 			Id: 1,
// 			User: _entity.User{
// 				Id:   1,
// 				Name: "Siska Kohl",
// 			},
// 			Asset: _entity.AssetReq{
// 				Id:           1,
// 				Code:         "asset-1645748000-1",
// 				Name:         "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
// 				CategoryName: "Computer",
// 				Image:        "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
// 			},
// 			Activity:    "Borrow",
// 			RequestTime: time.Date(2022, 02, 27, 12, 23, 23, 0, time.UTC),
// 			ReturnTime:  time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC),
// 			Status:      "Approved by Admin",
// 			Description: "try to borrow",
// 		},
// 	}, 1, nil
// }

// func (m mockRepoSuccess) GetAllProcureManager(limit, offset int, status, category, date string) (requests []_entity.Procure, total int, err error) {
// 	return []_entity.Procure{
// 		{
// 			Id: 1,
// 			User: _entity.User{
// 				Id:   1,
// 				Name: "Siska Kohl",
// 			},
// 			Category:    "Computer",
// 			Image:       "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
// 			Activity:    "Borrow",
// 			RequestTime: time.Date(2022, 02, 27, 12, 23, 23, 0, time.UTC),
// 			Status:      "Waiting approval from Manager",
// 			Description: "try to procure",
// 		},
// 	}, 1, nil
// }

// func (m mockRepoSuccess) GetUserDivision(id int) (divId int, code int, err error) {
// 	return 1, http.StatusOK, nil
// }

// // success
// func TestAdminGetAll(t *testing.T) {
// 	t.Run("TestAdminGetAllSuccess", func(t *testing.T) {
// 		token, _, _ := _midware.CreateToken(1, "Administrator")

// 		request := httptest.NewRequest(http.MethodGet, "/", nil)
// 		request.Header.Set("Content-Type", "application/json")
// 		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
// 		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		response := httptest.NewRecorder()

// 		e := echo.New()

// 		context := e.NewContext(request, response)
// 		context.SetPath("/requests/admin")

// 		adminController := New(mockRepoSuccess{})
// 		_midware.JWTMiddleWare()(adminController.AdminGetAll())(context)

// 		actual := map[string]interface{}{}
// 		body := response.Body.String()
// 		json.Unmarshal([]byte(body), &actual)

// 		expected := map[string]interface{}{
// 			"code": float64(http.StatusOK),
// 			"data": []interface{}{
// 				map[string]interface{}{
// 					"Asset": map[string]interface{}{
// 						"category":    "Computer",
// 						"category_id": float64(0),
// 						"code":        "asset-1645748000-1",
// 						"id":          float64(1),
// 						"image":       "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
// 						"name":        "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
// 						"short_name":  "",
// 					},
// 					"User": map[string]interface{}{
// 						"address":    "",
// 						"avatar":     "",
// 						"created_at": "0001-01-01T00:00:00Z",
// 						"deleted_at": "0001-01-01T00:00:00Z",
// 						"division":   "",
// 						"email":      "",
// 						"gender":     "",
// 						"id":         float64(1),
// 						"name":       "Siska Kohl",
// 						"password":   "",
// 						"phone":      "",
// 						"role":       "",
// 						"updated_at": "0001-01-01T00:00:00Z",
// 					},
// 					"activity":     "Borrow",
// 					"deleted_at":   "0001-01-01T00:00:00Z",
// 					"description":  "try to borrow",
// 					"id":           float64(1),
// 					"request_time": "2022-02-27T12:23:23Z",
// 					"return_time":  "9999-12-31T23:59:59Z",
// 					"status":       "Approved by Admin",
// 					"updated_at":   "0001-01-01T00:00:00Z",
// 				},
// 			},
// 			"message":      "success get all requests",
// 			"total_record": float64(1),
// 		}

// 		assert.Equal(t, expected, actual)
// 	})
// }

// func TestManagerGetAllProcure(t *testing.T) {
// 	t.Run("TestManagerGetAllProcureSuccess", func(t *testing.T) {
// 		token, _, _ := _midware.CreateToken(1, "Manager")

// 		request := httptest.NewRequest(http.MethodGet, "/", nil)
// 		request.Header.Set("Content-Type", "application/json")
// 		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
// 		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		response := httptest.NewRecorder()

// 		e := echo.New()

// 		context := e.NewContext(request, response)
// 		context.SetPath("/requests/manager/procure")

// 		adminController := New(mockRepoSuccess{})
// 		_midware.JWTMiddleWare()(adminController.ManagerGetAllProcure())(context)

// 		actual := map[string]interface{}{}
// 		body := response.Body.String()
// 		json.Unmarshal([]byte(body), &actual)

// 		expected := map[string]interface{}{
// 			"code": float64(http.StatusOK),
// 			"data": []interface{}{
// 				map[string]interface{}{
// 					"User": map[string]interface{}{
// 						"address":    "",
// 						"avatar":     "",
// 						"created_at": "0001-01-01T00:00:00Z",
// 						"deleted_at": "0001-01-01T00:00:00Z",
// 						"division":   "",
// 						"email":      "",
// 						"gender":     "",
// 						"id":         float64(1),
// 						"name":       "Siska Kohl",
// 						"password":   "",
// 						"phone":      "",
// 						"role":       "",
// 						"updated_at": "0001-01-01T00:00:00Z",
// 					},
// 					"activity":     "Borrow",
// 					"category":     "Computer",
// 					"deleted_at":   "0001-01-01T00:00:00Z",
// 					"description":  "try to procure",
// 					"id":           float64(1),
// 					"image":        "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
// 					"request_time": "2022-02-27T12:23:23Z",
// 					"status":       "Waiting approval from Manager",
// 					"updated_at":   "0001-01-01T00:00:00Z",
// 				},
// 			},
// 			"message":      "success get all requests",
// 			"total_record": float64(1),
// 		}

// 		assert.Equal(t, expected, actual)
// 	})
// }

// func TestManagerGetAllBorrow(t *testing.T) {
// 	t.Run("TestManagerGetAllBorrowSuccess", func(t *testing.T) {
// 		token, _, _ := _midware.CreateToken(1, "Manager")

// 		request := httptest.NewRequest(http.MethodGet, "/", nil)
// 		request.Header.Set("Content-Type", "application/json")
// 		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
// 		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		response := httptest.NewRecorder()

// 		e := echo.New()

// 		context := e.NewContext(request, response)
// 		context.SetPath("/requests/manager/borrow")

// 		adminController := New(mockRepoSuccess{})
// 		_midware.JWTMiddleWare()(adminController.ManagerGetAllBorrow())(context)

// 		actual := map[string]interface{}{}
// 		body := response.Body.String()
// 		json.Unmarshal([]byte(body), &actual)

// 		expected := map[string]interface{}{
// 			"code": float64(http.StatusOK),
// 			"data": []interface{}{
// 				map[string]interface{}{
// 					"Asset": map[string]interface{}{
// 						"category":    "Computer",
// 						"category_id": float64(0),
// 						"code":        "asset-1645748000-1",
// 						"id":          float64(1),
// 						"image":       "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
// 						"name":        "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
// 						"short_name":  "",
// 					},
// 					"User": map[string]interface{}{
// 						"address":    "",
// 						"avatar":     "",
// 						"created_at": "0001-01-01T00:00:00Z",
// 						"deleted_at": "0001-01-01T00:00:00Z",
// 						"division":   "",
// 						"email":      "",
// 						"gender":     "",
// 						"id":         float64(1),
// 						"name":       "Siska Kohl",
// 						"password":   "",
// 						"phone":      "",
// 						"role":       "",
// 						"updated_at": "0001-01-01T00:00:00Z",
// 					},
// 					"activity":     "Borrow",
// 					"deleted_at":   "0001-01-01T00:00:00Z",
// 					"description":  "try to borrow",
// 					"id":           float64(1),
// 					"request_time": "2022-02-27T12:23:23Z",
// 					"return_time":  "9999-12-31T23:59:59Z",
// 					"status":       "Approved by Admin",
// 					"updated_at":   "0001-01-01T00:00:00Z",
// 				},
// 			},
// 			"message":      "success get all requests",
// 			"total_record": float64(1),
// 		}

// 		assert.Equal(t, expected, actual)
// 	})
// }

// // failed because of role
// func TestRoleAdminGetAll(t *testing.T) {
// 	t.Run("TestRoleAdminGetAllFailed", func(t *testing.T) {
// 		token, _, _ := _midware.CreateToken(1, "Employee")

// 		request := httptest.NewRequest(http.MethodGet, "/", nil)
// 		request.Header.Set("Content-Type", "application/json")
// 		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
// 		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		response := httptest.NewRecorder()

// 		e := echo.New()

// 		context := e.NewContext(request, response)
// 		context.SetPath("/requests/admin")

// 		adminController := New(mockRepoSuccess{})
// 		_midware.JWTMiddleWare()(adminController.AdminGetAll())(context)

// 		actual := map[string]interface{}{}
// 		body := response.Body.String()
// 		json.Unmarshal([]byte(body), &actual)

// 		expected := map[string]interface{}{
// 			"code":    float64(http.StatusBadRequest),
// 			"message": "You don't have permission",
// 		}

// 		assert.Equal(t, expected, actual)
// 	})
// }

// func TestRoleGetAllProcure(t *testing.T) {
// 	t.Run("TestRoleGetAllProcureFailed", func(t *testing.T) {
// 		token, _, _ := _midware.CreateToken(1, "Administrator")

// 		request := httptest.NewRequest(http.MethodGet, "/", nil)
// 		request.Header.Set("Content-Type", "application/json")
// 		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
// 		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		response := httptest.NewRecorder()

// 		e := echo.New()

// 		context := e.NewContext(request, response)
// 		context.SetPath("/requests/manager/procure")

// 		adminController := New(mockRepoSuccess{})
// 		_midware.JWTMiddleWare()(adminController.ManagerGetAllProcure())(context)

// 		actual := map[string]interface{}{}
// 		body := response.Body.String()
// 		json.Unmarshal([]byte(body), &actual)

// 		expected := map[string]interface{}{
// 			"code":    float64(http.StatusBadRequest),
// 			"message": "You don't have permission",
// 		}

// 		assert.Equal(t, expected, actual)
// 	})
// }

// func TestRoleGetAllBorrow(t *testing.T) {
// 	t.Run("TestRoleGetAllBorrowFailed", func(t *testing.T) {
// 		token, _, _ := _midware.CreateToken(1, "Employee")

// 		request := httptest.NewRequest(http.MethodGet, "/", nil)
// 		request.Header.Set("Content-Type", "application/json")
// 		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
// 		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		response := httptest.NewRecorder()

// 		e := echo.New()

// 		context := e.NewContext(request, response)
// 		context.SetPath("/requests/manager/borrow")

// 		adminController := New(mockRepoSuccess{})
// 		_midware.JWTMiddleWare()(adminController.ManagerGetAllBorrow())(context)

// 		actual := map[string]interface{}{}
// 		body := response.Body.String()
// 		json.Unmarshal([]byte(body), &actual)

// 		expected := map[string]interface{}{
// 			"code":    float64(http.StatusBadRequest),
// 			"message": "You don't have permission",
// 		}

// 		assert.Equal(t, expected, actual)
// 	})
// }

// // failed because of query param
// func TestQueryParamAdminGetAll(t *testing.T) {
// 	t.Run("TestPageAdminGetAllFailed", func(t *testing.T) {
// 		token, _, _ := _midware.CreateToken(1, "Administrator")

// 		request := httptest.NewRequest(http.MethodGet, "/?p=a", nil)
// 		request.Header.Set("Content-Type", "application/json")
// 		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
// 		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		response := httptest.NewRecorder()

// 		e := echo.New()

// 		context := e.NewContext(request, response)
// 		context.SetPath("/requests/admin")
// 		context.QueryParam("p")
// 		adminController := New(mockRepoSuccess{})
// 		_midware.JWTMiddleWare()(adminController.AdminGetAll())(context)

// 		actual := map[string]interface{}{}
// 		body := response.Body.String()
// 		json.Unmarshal([]byte(body), &actual)

// 		expected := map[string]interface{}{
// 			"code":    float64(http.StatusBadRequest),
// 			"message": "Error parsing page",
// 		}

// 		assert.Equal(t, expected, actual)
// 	})
// 	t.Run("TestRecordOfPageAdminGetAllFailed", func(t *testing.T) {
// 		token, _, _ := _midware.CreateToken(1, "Administrator")

// 		request := httptest.NewRequest(http.MethodGet, "/?rp=a", nil)
// 		request.Header.Set("Content-Type", "application/json")
// 		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
// 		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		response := httptest.NewRecorder()

// 		e := echo.New()

// 		context := e.NewContext(request, response)
// 		context.SetPath("/requests/admin")
// 		context.QueryParam("rp")
// 		adminController := New(mockRepoSuccess{})
// 		_midware.JWTMiddleWare()(adminController.AdminGetAll())(context)

// 		actual := map[string]interface{}{}
// 		body := response.Body.String()
// 		json.Unmarshal([]byte(body), &actual)

// 		expected := map[string]interface{}{
// 			"code":    float64(http.StatusBadRequest),
// 			"message": "Error parsing record of page",
// 		}

// 		assert.Equal(t, expected, actual)
// 	})
// 	t.Run("TestStatusAdminGetAllFailed", func(t *testing.T) {
// 		token, _, _ := _midware.CreateToken(1, "Administrator")

// 		request := httptest.NewRequest(http.MethodGet, "/?s=uhuy", nil)
// 		request.Header.Set("Content-Type", "application/json")
// 		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
// 		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		response := httptest.NewRecorder()

// 		e := echo.New()

// 		context := e.NewContext(request, response)
// 		context.SetPath("/requests/admin")
// 		context.QueryParam("s")
// 		adminController := New(mockRepoSuccess{})
// 		_midware.JWTMiddleWare()(adminController.AdminGetAll())(context)

// 		actual := map[string]interface{}{}
// 		body := response.Body.String()
// 		json.Unmarshal([]byte(body), &actual)

// 		expected := map[string]interface{}{
// 			"code":    float64(http.StatusBadRequest),
// 			"message": "Bad request",
// 		}

// 		assert.Equal(t, expected, actual)
// 	})
// 	t.Run("TestCategoryAdminGetAllFailed", func(t *testing.T) {
// 		token, _, _ := _midware.CreateToken(1, "Administrator")

// 		request := httptest.NewRequest(http.MethodGet, "/?c=uhuy", nil)
// 		request.Header.Set("Content-Type", "application/json")
// 		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
// 		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		response := httptest.NewRecorder()

// 		e := echo.New()

// 		context := e.NewContext(request, response)
// 		context.SetPath("/requests/admin")
// 		context.QueryParam("c")
// 		adminController := New(mockRepoSuccess{})
// 		_midware.JWTMiddleWare()(adminController.AdminGetAll())(context)

// 		actual := map[string]interface{}{}
// 		body := response.Body.String()
// 		json.Unmarshal([]byte(body), &actual)

// 		expected := map[string]interface{}{
// 			"code":    float64(http.StatusBadRequest),
// 			"message": "Bad request",
// 		}

// 		assert.Equal(t, expected, actual)
// 	})
// }

// func TestQueryParamGetAllProcure(t *testing.T) {
// 	t.Run("TestPageGetAllProcureFailed", func(t *testing.T) {
// 		token, _, _ := _midware.CreateToken(1, "Manager")

// 		request := httptest.NewRequest(http.MethodGet, "/?p=a", nil)
// 		request.Header.Set("Content-Type", "application/json")
// 		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
// 		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		response := httptest.NewRecorder()

// 		e := echo.New()

// 		context := e.NewContext(request, response)
// 		context.SetPath("/requests/manager/procure")
// 		context.QueryParam("p")

// 		adminController := New(mockRepoSuccess{})
// 		_midware.JWTMiddleWare()(adminController.ManagerGetAllProcure())(context)

// 		actual := map[string]interface{}{}
// 		body := response.Body.String()
// 		json.Unmarshal([]byte(body), &actual)

// 		expected := map[string]interface{}{
// 			"code":    float64(http.StatusBadRequest),
// 			"message": "Error parsing page",
// 		}

// 		assert.Equal(t, expected, actual)
// 	})
// 	t.Run("TestRecordOfPageGetAllProcureFailed", func(t *testing.T) {
// 		token, _, _ := _midware.CreateToken(1, "Manager")

// 		request := httptest.NewRequest(http.MethodGet, "/?rp=a", nil)
// 		request.Header.Set("Content-Type", "application/json")
// 		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
// 		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		response := httptest.NewRecorder()

// 		e := echo.New()

// 		context := e.NewContext(request, response)
// 		context.SetPath("/requests/manager/procure")
// 		context.QueryParam("rp")

// 		adminController := New(mockRepoSuccess{})
// 		_midware.JWTMiddleWare()(adminController.ManagerGetAllProcure())(context)

// 		actual := map[string]interface{}{}
// 		body := response.Body.String()
// 		json.Unmarshal([]byte(body), &actual)

// 		expected := map[string]interface{}{
// 			"code":    float64(http.StatusBadRequest),
// 			"message": "Error parsing record of page",
// 		}

// 		assert.Equal(t, expected, actual)
// 	})
// 	t.Run("TestStatusGetAllProcureFailed", func(t *testing.T) {
// 		token, _, _ := _midware.CreateToken(1, "Manager")

// 		request := httptest.NewRequest(http.MethodGet, "/?s=uhuy", nil)
// 		request.Header.Set("Content-Type", "application/json")
// 		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
// 		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		response := httptest.NewRecorder()

// 		e := echo.New()

// 		context := e.NewContext(request, response)
// 		context.SetPath("/requests/manager/procure")
// 		context.QueryParam("s")

// 		adminController := New(mockRepoSuccess{})
// 		_midware.JWTMiddleWare()(adminController.ManagerGetAllProcure())(context)

// 		actual := map[string]interface{}{}
// 		body := response.Body.String()
// 		json.Unmarshal([]byte(body), &actual)

// 		expected := map[string]interface{}{
// 			"code":    float64(http.StatusBadRequest),
// 			"message": "Bad request",
// 		}

// 		assert.Equal(t, expected, actual)
// 	})
// 	t.Run("TestCategoryGetAllProcureFailed", func(t *testing.T) {
// 		token, _, _ := _midware.CreateToken(1, "Manager")

// 		request := httptest.NewRequest(http.MethodGet, "/?c=uhuy", nil)
// 		request.Header.Set("Content-Type", "application/json")
// 		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
// 		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		response := httptest.NewRecorder()

// 		e := echo.New()

// 		context := e.NewContext(request, response)
// 		context.SetPath("/requests/manager/procure")
// 		context.QueryParam("c")

// 		adminController := New(mockRepoSuccess{})
// 		_midware.JWTMiddleWare()(adminController.ManagerGetAllProcure())(context)

// 		actual := map[string]interface{}{}
// 		body := response.Body.String()
// 		json.Unmarshal([]byte(body), &actual)

// 		expected := map[string]interface{}{
// 			"code":    float64(http.StatusBadRequest),
// 			"message": "Bad request",
// 		}

// 		assert.Equal(t, expected, actual)
// 	})
// }

// func TestQueryParamGetAllBorrow(t *testing.T) {
// 	t.Run("TestPageGetAllBorrowFailed", func(t *testing.T) {
// 		token, _, _ := _midware.CreateToken(1, "Manager")

// 		request := httptest.NewRequest(http.MethodGet, "/?p=a", nil)
// 		request.Header.Set("Content-Type", "application/json")
// 		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
// 		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		response := httptest.NewRecorder()

// 		e := echo.New()

// 		context := e.NewContext(request, response)
// 		context.SetPath("/requests/manager/borrow")
// 		context.QueryParam("p")
// 		adminController := New(mockRepoSuccess{})
// 		_midware.JWTMiddleWare()(adminController.ManagerGetAllBorrow())(context)

// 		actual := map[string]interface{}{}
// 		body := response.Body.String()
// 		json.Unmarshal([]byte(body), &actual)

// 		expected := map[string]interface{}{
// 			"code":    float64(http.StatusBadRequest),
// 			"message": "Error parsing page",
// 		}

// 		assert.Equal(t, expected, actual)
// 	})
// 	t.Run("TestRecordOfPageGetAllBorrowFailed", func(t *testing.T) {
// 		token, _, _ := _midware.CreateToken(1, "Manager")

// 		request := httptest.NewRequest(http.MethodGet, "/?rp=a", nil)
// 		request.Header.Set("Content-Type", "application/json")
// 		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
// 		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		response := httptest.NewRecorder()

// 		e := echo.New()

// 		context := e.NewContext(request, response)
// 		context.SetPath("/requests/manager/borrow")
// 		context.QueryParam("rp")
// 		adminController := New(mockRepoSuccess{})
// 		_midware.JWTMiddleWare()(adminController.ManagerGetAllBorrow())(context)

// 		actual := map[string]interface{}{}
// 		body := response.Body.String()
// 		json.Unmarshal([]byte(body), &actual)

// 		expected := map[string]interface{}{
// 			"code":    float64(http.StatusBadRequest),
// 			"message": "Error parsing record of page",
// 		}

// 		assert.Equal(t, expected, actual)
// 	})
// 	t.Run("TestStatusGetAllBorrowFailed", func(t *testing.T) {
// 		token, _, _ := _midware.CreateToken(1, "Manager")

// 		request := httptest.NewRequest(http.MethodGet, "/?s=uhuy", nil)
// 		request.Header.Set("Content-Type", "application/json")
// 		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
// 		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		response := httptest.NewRecorder()

// 		e := echo.New()

// 		context := e.NewContext(request, response)
// 		context.SetPath("/requests/manager/borrow")
// 		context.QueryParam("s")
// 		adminController := New(mockRepoSuccess{})
// 		_midware.JWTMiddleWare()(adminController.ManagerGetAllBorrow())(context)

// 		actual := map[string]interface{}{}
// 		body := response.Body.String()
// 		json.Unmarshal([]byte(body), &actual)

// 		expected := map[string]interface{}{
// 			"code":    float64(http.StatusBadRequest),
// 			"message": "Bad request",
// 		}

// 		assert.Equal(t, expected, actual)
// 	})
// 	t.Run("TestCategoryGetAllBorrowFailed", func(t *testing.T) {
// 		token, _, _ := _midware.CreateToken(1, "Manager")

// 		request := httptest.NewRequest(http.MethodGet, "/?c=uhuy", nil)
// 		request.Header.Set("Content-Type", "application/json")
// 		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
// 		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		response := httptest.NewRecorder()

// 		e := echo.New()

// 		context := e.NewContext(request, response)
// 		context.SetPath("/requests/manager/borrow")
// 		context.QueryParam("c")
// 		adminController := New(mockRepoSuccess{})
// 		_midware.JWTMiddleWare()(adminController.ManagerGetAllBorrow())(context)

// 		actual := map[string]interface{}{}
// 		body := response.Body.String()
// 		json.Unmarshal([]byte(body), &actual)

// 		expected := map[string]interface{}{
// 			"code":    float64(http.StatusBadRequest),
// 			"message": "Bad request",
// 		}

// 		assert.Equal(t, expected, actual)
// 	})
// }

// // failed because of Repo
// type mockRepoFail struct{}

// func (m mockRepoFail) GetAllAdmin(limit, offset int, status, category, date string) (requests []_entity.RequestResponse, total int, err error) {
// 	return []_entity.RequestResponse{}, 0, errors.New("Failed to read data")
// }

// func (m mockRepoFail) GetAllManager(divLogin, limit, offset int, status, category, date string) (requests []_entity.RequestResponse, total int, err error) {
// 	return []_entity.RequestResponse{}, 0, errors.New("Failed to read data")
// }

// func (m mockRepoFail) GetAllProcureManager(limit, offset int, status, category, date string) (requests []_entity.Procure, total int, err error) {
// 	return []_entity.Procure{}, 0, errors.New("Failed to read data")
// }

// func (m mockRepoFail) GetUserDivision(id int) (divId int, code int, err error) {
// 	return 1, http.StatusOK, nil
// }
// func TestRepoAdminGetAll(t *testing.T) {
// 	t.Run("TestRepoAdminGetAllFailed", func(t *testing.T) {
// 		token, _, _ := _midware.CreateToken(1, "Administrator")

// 		request := httptest.NewRequest(http.MethodGet, "/", nil)
// 		request.Header.Set("Content-Type", "application/json")
// 		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
// 		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		response := httptest.NewRecorder()

// 		e := echo.New()

// 		context := e.NewContext(request, response)
// 		context.SetPath("/requests/admin")

// 		adminController := New(mockRepoFail{})
// 		_midware.JWTMiddleWare()(adminController.AdminGetAll())(context)

// 		actual := map[string]interface{}{}
// 		body := response.Body.String()
// 		json.Unmarshal([]byte(body), &actual)

// 		expected := map[string]interface{}{
// 			"code":    float64(http.StatusInternalServerError),
// 			"message": "Failed to read data",
// 		}

// 		assert.Equal(t, expected, actual)
// 	})
// }

// func TestRepoGetAllProcure(t *testing.T) {
// 	t.Run("TestRepoGetAllProcureFailed", func(t *testing.T) {
// 		token, _, _ := _midware.CreateToken(1, "Manager")

// 		request := httptest.NewRequest(http.MethodGet, "/", nil)
// 		request.Header.Set("Content-Type", "application/json")
// 		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
// 		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		response := httptest.NewRecorder()

// 		e := echo.New()

// 		context := e.NewContext(request, response)
// 		context.SetPath("/requests/manager/procure")

// 		adminController := New(mockRepoFail{})
// 		_midware.JWTMiddleWare()(adminController.ManagerGetAllProcure())(context)

// 		actual := map[string]interface{}{}
// 		body := response.Body.String()
// 		json.Unmarshal([]byte(body), &actual)

// 		expected := map[string]interface{}{
// 			"code":    float64(http.StatusInternalServerError),
// 			"message": "Failed to read data",
// 		}

// 		assert.Equal(t, expected, actual)
// 	})
// }

// func TestRepoGetAllBorrow(t *testing.T) {
// 	t.Run("TestRepoGetAllBorrowFailed", func(t *testing.T) {
// 		token, _, _ := _midware.CreateToken(1, "Manager")

// 		request := httptest.NewRequest(http.MethodGet, "/", nil)
// 		request.Header.Set("Content-Type", "application/json")
// 		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
// 		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 		response := httptest.NewRecorder()

// 		e := echo.New()

// 		context := e.NewContext(request, response)
// 		context.SetPath("/requests/manager/borrow")

// 		adminController := New(mockRepoFail{})
// 		_midware.JWTMiddleWare()(adminController.ManagerGetAllBorrow())(context)

// 		actual := map[string]interface{}{}
// 		body := response.Body.String()
// 		json.Unmarshal([]byte(body), &actual)

// 		expected := map[string]interface{}{
// 			"code":    float64(http.StatusInternalServerError),
// 			"message": "Failed to read data",
// 		}

// 		assert.Equal(t, expected, actual)
// 	})
// }
