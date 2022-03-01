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

type mockRepoSuccess struct{}

func (m mockRepoSuccess) Borrow(reqData _entity.Borrow) (code int, err error) {
	return http.StatusOK, nil
}

func (m mockRepoSuccess) Procure(reqData _entity.Procure) (code int, err error) {
	return http.StatusOK, nil
}

func (m mockRepoSuccess) GetBorrowById(id int) (req _entity.Borrow, code int, err error) {
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

func (m mockRepoSuccess) GetUserDivision(id int) (divId int, code int, err error) {
	return 1, http.StatusOK, nil
}

func (m mockRepoSuccess) UpdateBorrow(reqData _entity.Borrow) (req _entity.Borrow, code int, err error) {
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

func (m mockRepoSuccess) UpdateProcure(reqData _entity.Procure) (req _entity.Procure, code int, err error) {
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
		Status:      "Waiting approval from Admin",
		Description: "trying to procure",
	}, http.StatusOK, nil
}

func (m mockRepoSuccess) GetProcureById(id int) (req _entity.Procure, code int, err error) {
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
		Status:      "Waiting approval from Admin",
		Description: "trying to procure",
	}, http.StatusOK, nil
}

func (m mockRepoSuccess) UpdateBorrowByAdmin(reqData _entity.Borrow) (_entity.Borrow, error) {
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

func (m mockRepoSuccess) ReturnAdmin(reqData _entity.Borrow) (updatedReq _entity.Borrow, code int, err error) {
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

		requestController := New(mockRepoSuccess{})
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

		requestController := New(mockRepoSuccess{})
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

		requestController := New(mockRepoSuccess{})
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

		assetController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(assetController.UpdateBorrow())(context)

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

		assetController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(assetController.UpdateBorrow())(context)

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
