package request

import (
	"bytes"
	_midware "capstone/be/delivery/middleware"
	_entity "capstone/be/entity"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/magiconair/properties/assert"
)

type mockRepoSuccess1 struct{}
type mockRepoSuccess2 struct{}
type mockRepoSuccess3 struct{}

func (m mockRepoSuccess1) Borrow(reqData _entity.Borrow) (code int, err error) {
	return http.StatusOK, nil
}

func (m mockRepoSuccess2) Borrow(reqData _entity.Borrow) (code int, err error) {
	return http.StatusOK, nil
}

func (m mockRepoSuccess3) Borrow(reqData _entity.Borrow) (code int, err error) {
	return http.StatusOK, nil
}

func (m mockRepoSuccess1) Procure(reqData _entity.Procure) (code int, err error) {
	return http.StatusOK, nil
}

func (m mockRepoSuccess2) Procure(reqData _entity.Procure) (code int, err error) {
	return http.StatusOK, nil
}

func (m mockRepoSuccess3) Procure(reqData _entity.Procure) (code int, err error) {
	return http.StatusOK, nil
}

func (m mockRepoSuccess1) GetBorrowById(id int) (req _entity.Borrow, code int, err error) {
	return _entity.Borrow{
		Id: 1,
		User: _entity.User{
			Id:   1,
			Name: "Siska Kohl",
		},
		Asset: _entity.Asset{
			Id:        1,
			Name:      "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
			Image:     "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
			ShortName: "asset-1645748000",
		},
		Activity:    "Borrow",
		RequestTime: time.Date(2022, 02, 27, 12, 23, 23, 0, time.UTC),
		Status:      "Waiting approval from Admin",
		Description: "trying to borrow",
	}, http.StatusOK, nil
}

func (m mockRepoSuccess2) GetBorrowById(id int) (req _entity.Borrow, code int, err error) {
	return _entity.Borrow{
		Id: 1,
		User: _entity.User{
			Id:   1,
			Name: "Siska Kohl",
		},
		Asset: _entity.Asset{
			Id:        1,
			Name:      "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
			Image:     "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
			ShortName: "asset-1645748000",
		},
		Activity:    "Borrow",
		RequestTime: time.Date(2022, 02, 27, 12, 23, 23, 0, time.UTC),
		Status:      "Waiting approval from Manager",
		// Status:      "Approved by Admin",
		Description: "trying to borrow",
	}, http.StatusOK, nil
}

func (m mockRepoSuccess3) GetBorrowById(id int) (req _entity.Borrow, code int, err error) {
	return _entity.Borrow{
		Id: 1,
		User: _entity.User{
			Id:   1,
			Name: "Siska Kohl",
		},
		Asset: _entity.Asset{
			Id:        1,
			Name:      "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
			Image:     "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
			ShortName: "asset-1645748000",
		},
		Activity:    "Borrow",
		Status:      "Approved by Admin",
		RequestTime: time.Date(2022, 02, 27, 12, 23, 23, 0, time.UTC),
		Description: "trying to borrow",
	}, http.StatusOK, nil
}

func (m mockRepoSuccess1) GetUserDivision(id int) (divId int, code int, err error) {
	return 1, http.StatusOK, nil
}

func (m mockRepoSuccess2) GetUserDivision(id int) (divId int, code int, err error) {
	return 1, http.StatusOK, nil
}

func (m mockRepoSuccess3) GetUserDivision(id int) (divId int, code int, err error) {
	return 1, http.StatusOK, nil
}

func (m mockRepoSuccess1) UpdateBorrow(reqData _entity.Borrow) (req _entity.Borrow, code int, err error) {
	return _entity.Borrow{
		Id: 1,
		User: _entity.User{
			Id:   1,
			Name: "Siska Kohl",
		},
		Asset: _entity.Asset{
			Id:        1,
			Name:      "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
			Image:     "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
			ShortName: "asset-1645748000",
		},
		Activity:    "Borrow",
		RequestTime: time.Date(2022, 02, 27, 12, 23, 23, 0, time.UTC),
		// Status:      "Waiting approval from Admin",
		Status:      "Waiting approval from Manager",
		Description: "trying to borrow",
	}, http.StatusOK, nil
}
func (m mockRepoSuccess2) UpdateBorrow(reqData _entity.Borrow) (req _entity.Borrow, code int, err error) {
	return _entity.Borrow{
		Id: 1,
		User: _entity.User{
			Id:   1,
			Name: "Siska Kohl",
		},
		Asset: _entity.Asset{
			Id:        1,
			Name:      "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
			Image:     "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
			ShortName: "asset-1645748000",
		},
		Activity:    "Borrow",
		RequestTime: time.Date(2022, 02, 27, 12, 23, 23, 0, time.UTC),
		Status:      "Waiting approval from Manager",
		Description: "trying to borrow",
	}, http.StatusOK, nil
}

func (m mockRepoSuccess3) UpdateBorrow(reqData _entity.Borrow) (req _entity.Borrow, code int, err error) {
	return _entity.Borrow{
		Id: 1,
		User: _entity.User{
			Id:   1,
			Name: "Siska Kohl",
		},
		Asset: _entity.Asset{
			Id:        1,
			Name:      "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
			Image:     "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
			ShortName: "asset-1645748000",
		},
		Activity:    "Borrow",
		RequestTime: time.Date(2022, 02, 27, 12, 23, 23, 0, time.UTC),
		Status:      "Waiting approval from Manager",
		Description: "trying to borrow",
	}, http.StatusOK, nil
}

func (m mockRepoSuccess1) UpdateProcure(reqData _entity.Procure) (req _entity.Procure, code int, err error) {
	return _entity.Procure{
		Id: 1,
		User: _entity.User{
			Id:   1,
			Name: "Siska Kohl",
		},
		Category:    "Computer",
		Image:       "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
		Activity:    "Procure",
		RequestTime: time.Date(2022, 02, 27, 12, 23, 23, 0, time.UTC),
		Status:      "Waiting approval from manager",
		Description: "trying to procure",
	}, http.StatusOK, nil
}

func (m mockRepoSuccess2) UpdateProcure(reqData _entity.Procure) (req _entity.Procure, code int, err error) {
	return _entity.Procure{
		Id: 1,
		User: _entity.User{
			Id:   1,
			Name: "Siska Kohl",
		},
		Category:    "Computer",
		Image:       "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
		Activity:    "Procure",
		RequestTime: time.Date(2022, 02, 27, 12, 23, 23, 0, time.UTC),
		Status:      "Waiting approval from manager",
		Description: "trying to procure",
	}, http.StatusOK, nil
}

func (m mockRepoSuccess3) UpdateProcure(reqData _entity.Procure) (req _entity.Procure, code int, err error) {
	return _entity.Procure{
		Id: 1,
		User: _entity.User{
			Id:   1,
			Name: "Siska Kohl",
		},
		Category:    "Computer",
		Image:       "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
		Activity:    "Procure",
		RequestTime: time.Date(2022, 02, 27, 12, 23, 23, 0, time.UTC),
		Status:      "Waiting approval from manager",
		Description: "trying to procure",
	}, http.StatusOK, nil
}

func (m mockRepoSuccess1) GetProcureById(id int) (req _entity.Procure, code int, err error) {
	return _entity.Procure{
		Id: 1,
		User: _entity.User{
			Id:   1,
			Name: "Siska Kohl",
		},
		Category:    "Computer",
		Image:       "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
		Activity:    "Procure",
		RequestTime: time.Date(2022, 02, 27, 12, 23, 23, 0, time.UTC),
		Status:      "Waiting approval from manager",
		Description: "trying to procure",
	}, http.StatusOK, nil
}

func (m mockRepoSuccess2) GetProcureById(id int) (req _entity.Procure, code int, err error) {
	return _entity.Procure{
		Id: 1,
		User: _entity.User{
			Id:   1,
			Name: "Siska Kohl",
		},
		Category:    "Computer",
		Image:       "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
		Activity:    "Procure",
		RequestTime: time.Date(2022, 02, 27, 12, 23, 23, 0, time.UTC),
		Status:      "Waiting approval from manager",
		Description: "trying to procure",
	}, http.StatusOK, nil
}

func (m mockRepoSuccess3) GetProcureById(id int) (req _entity.Procure, code int, err error) {
	return _entity.Procure{
		Id: 1,
		User: _entity.User{
			Id:   1,
			Name: "Siska Kohl",
		},
		Category:    "Computer",
		Image:       "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
		Activity:    "Procure",
		RequestTime: time.Date(2022, 02, 27, 12, 23, 23, 0, time.UTC),
		Status:      "Waiting approval from manager",
		Description: "trying to procure",
	}, http.StatusOK, nil
}

func (m mockRepoSuccess1) UpdateBorrowByAdmin(reqData _entity.Borrow) (_entity.Borrow, error) {
	return _entity.Borrow{
		Id: 1,
		User: _entity.User{
			Id:   1,
			Name: "Siska Kohl",
		},
		Asset: _entity.Asset{
			Id:        1,
			Name:      "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
			Image:     "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
			ShortName: "asset-1645748000",
		},
		Activity:    "Borrow",
		RequestTime: time.Date(2022, 02, 27, 12, 23, 23, 0, time.UTC),
		Status:      "Waiting approval from Admin",
		Description: "trying to borrow",
	}, nil
}

func (m mockRepoSuccess2) UpdateBorrowByAdmin(reqData _entity.Borrow) (_entity.Borrow, error) {
	return _entity.Borrow{
		Id: 1,
		User: _entity.User{
			Id:   1,
			Name: "Siska Kohl",
		},
		Asset: _entity.Asset{
			Id:        1,
			Name:      "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
			Image:     "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
			ShortName: "asset-1645748000",
		},
		Activity:    "Borrow",
		RequestTime: time.Date(2022, 02, 27, 12, 23, 23, 0, time.UTC),
		Status:      "Waiting approval from Admin",
		Description: "trying to borrow",
	}, nil
}

func (m mockRepoSuccess3) UpdateBorrowByAdmin(reqData _entity.Borrow) (_entity.Borrow, error) {
	return _entity.Borrow{
		Id: 1,
		User: _entity.User{
			Id:   1,
			Name: "Siska Kohl",
		},
		Asset: _entity.Asset{
			Id:        1,
			Name:      "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
			Image:     "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
			ShortName: "asset-1645748000",
		},
		Activity:    "Borrow",
		RequestTime: time.Date(2022, 02, 27, 12, 23, 23, 0, time.UTC),
		Status:      "Waiting approval from Admin",
		Description: "trying to borrow",
	}, nil
}

func (m mockRepoSuccess1) ReturnAdmin(reqData _entity.Borrow) (updatedReq _entity.Borrow, code int, err error) {
	return _entity.Borrow{
		Id: 1,
		User: _entity.User{
			Id:   1,
			Name: "Siska Kohl",
		},
		Asset: _entity.Asset{
			Id:        1,
			Name:      "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
			Image:     "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
			ShortName: "asset-1645748000",
		},
		RequestTime: time.Date(2022, 02, 27, 12, 23, 23, 0, time.UTC),
		Activity:    "Request to Return",
		Description: "trying to borrow",
		Status:      "Approved by Admin",
	}, http.StatusOK, nil
}

func (m mockRepoSuccess2) ReturnAdmin(reqData _entity.Borrow) (updatedReq _entity.Borrow, code int, err error) {
	return _entity.Borrow{
		Id: 1,
		User: _entity.User{
			Id:   1,
			Name: "Siska Kohl",
		},
		Asset: _entity.Asset{
			Id:        1,
			Name:      "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
			Image:     "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
			ShortName: "asset-1645748000",
		},
		RequestTime: time.Date(2022, 02, 27, 12, 23, 23, 0, time.UTC),
		Activity:    "Request to Return",
		Description: "trying to borrow",
		Status:      "Approved by Admin",
	}, http.StatusOK, nil
}

func (m mockRepoSuccess3) ReturnAdmin(reqData _entity.Borrow) (updatedReq _entity.Borrow, code int, err error) {
	return _entity.Borrow{
		Id: 1,
		User: _entity.User{
			Id:   1,
			Name: "Siska Kohl",
		},
		Asset: _entity.Asset{
			Id:        1,
			Name:      "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
			Image:     "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
			ShortName: "asset-1645748000",
		},
		RequestTime: time.Date(2022, 02, 27, 12, 23, 23, 0, time.UTC),
		Activity:    "Request to Return",
		Description: "trying to borrow",
		Status:      "Approved by Admin",
	}, http.StatusOK, nil
}

// success

func TestBorrowSuccess(t *testing.T) {
	t.Run("TestBorrowEmployeeSuccess", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Employee")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"short_name":  "asset-1645748000",
			"description": "pinjam lenovo",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/requests/borrow")

		requestController := New(mockRepoSuccess1{})
		_midware.JWTMiddleWare()(requestController.Borrow())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success create request",
		}

		assert.Equal(t, expected, actual)
	})
	t.Run("TestBorrowAdministratorSuccess", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"short_name":  "asset-1645748000",
			"employee_id": 1,
			"description": "pinjam kipas",
			"return_time": "2022-02-21T21:30:05+07:00",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/requests/borrow")

		requestController := New(mockRepoSuccess1{})
		_midware.JWTMiddleWare()(requestController.Borrow())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success create request",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestProcureSuccess(t *testing.T) {
	t.Run("TestProcureAdministratorSuccess", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"image":       "image.jpg",
			"category":    "category name",
			"description": "alasan pengajuan",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/requests/procure")

		requestController := New(mockRepoSuccess1{})
		_midware.JWTMiddleWare()(requestController.Procure())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success create request",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetBorrowById(t *testing.T) {
	t.Run("TestGetBorrowByIdSuccess", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		request := httptest.NewRequest(http.MethodPut, "/", nil)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/requests/borrow/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		requestController := New(mockRepoSuccess1{})
		_midware.JWTMiddleWare()(requestController.GetBorrowById())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code": float64(http.StatusOK),
			"data": map[string]interface{}{
				"Asset": map[string]interface{}{
					"category_id": float64(0),
					"code":        "",
					"created_at":  "0001-01-01T00:00:00Z",
					"deleted_at":  "0001-01-01T00:00:00Z",
					"description": "",
					"id":          float64(1),
					"image":       "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
					"name":        "Dell Latitude 3420 (i7-1165G7, 8GB, 512GB)",
					"short_name":  "asset-1645748000",
					"status":      "",
					"updated_at":  "0001-01-01T00:00:00Z",
				},
				"User": map[string]interface{}{
					"address":    "",
					"avatar":     "",
					"created_at": "0001-01-01T00:00:00Z",
					"deleted_at": "0001-01-01T00:00:00Z",
					"division":   "",
					"email":      "",
					"gender":     "",
					"id":         float64(1),
					"name":       "Siska Kohl",
					"password":   "",
					"phone":      "",
					"role":       "",
					"updated_at": "0001-01-01T00:00:00Z",
				},
				"activity":     "Borrow",
				"deleted_at":   "0001-01-01T00:00:00Z",
				"description":  "trying to borrow",
				"id":           float64(1),
				"request_time": "2022-02-27T12:23:23Z",
				"return_time":  "0001-01-01T00:00:00Z",
				"status":       "Waiting approval from Admin",
				"updated_at":   "0001-01-01T00:00:00Z",
			},
			"message": "success get request",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetProcureById(t *testing.T) {
	t.Run("TestGetProcureByIdSuccess", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Manager")

		request := httptest.NewRequest(http.MethodPut, "/", nil)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/requests/procure/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		requestController := New(mockRepoSuccess1{})
		_midware.JWTMiddleWare()(requestController.GetProcureById())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code": float64(http.StatusOK),
			"data": map[string]interface{}{
				"User": map[string]interface{}{
					"address":    "",
					"avatar":     "",
					"created_at": "0001-01-01T00:00:00Z",
					"deleted_at": "0001-01-01T00:00:00Z",
					"division":   "",
					"email":      "",
					"gender":     "",
					"id":         float64(1),
					"name":       "Siska Kohl",
					"password":   "",
					"phone":      "",
					"role":       "",
					"updated_at": "0001-01-01T00:00:00Z",
				},
				"activity":     "Procure",
				"category":     "Computer",
				"deleted_at":   "0001-01-01T00:00:00Z",
				"description":  "trying to procure",
				"id":           float64(1),
				"image":        "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/asset-6-1645748000.png",
				"request_time": "2022-02-27T12:23:23Z",
				"status":       "Waiting approval from manager",
				"updated_at":   "0001-01-01T00:00:00Z",
			},
			"message": "success get request",
		}

		assert.Equal(t, expected, actual)
	})
}
func TestUpdateBorrowSuccess(t *testing.T) {
	t.Run("TestUpdateBorrowEmployeeAdminSuccess", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"approved": true,
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/requests/borrow/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		requestController := New(mockRepoSuccess1{})
		_midware.JWTMiddleWare()(requestController.UpdateBorrow())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success update request",
		}

		assert.Equal(t, expected, actual)
	})
	t.Run("TestUpdateBorrowAdminManagerSuccess", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Manager")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"approved": true,
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/requests/borrow/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		requestController := New(mockRepoSuccess2{})
		_midware.JWTMiddleWare()(requestController.UpdateBorrow())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success update request",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUpdateProcureSuccess(t *testing.T) {
	t.Run("TestUpdateProcureSuccess", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Manager")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"approved": true,
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/requests/procure/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		requestController := New(mockRepoSuccess1{})
		_midware.JWTMiddleWare()(requestController.UpdateProcure())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success update request",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestReturnAdminSuccess(t *testing.T) {
	t.Run("TestReturnAdminSuccess", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1, "Administrator")

		requestBody, _ := json.Marshal(map[string]interface{}{
			"askingreturn": true,
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/requests/admin/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		requestController := New(mockRepoSuccess3{})
		_midware.JWTMiddleWare()(requestController.AdminReturn())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success asking return",
		}

		assert.Equal(t, expected, actual)
	})
}

//
