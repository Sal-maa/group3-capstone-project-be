package user

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

func (m mockRepoSuccess) Create(_entity.User) (_entity.UserSimplified, int, error) {
	return _entity.UserSimplified{
		Id:     1,
		Name:   "Salmaa",
		Email:  "salma@sirclo.com",
		Phone:  "08123456789",
		Avatar: "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/default_avatar.png",
	}, http.StatusOK, nil
}

func (m mockRepoSuccess) LoginByEmail(string) (_entity.User, int, error) {
	return _entity.User{
		Id:       1,
		Name:     "Salmaa",
		Password: "$2a$04$mE8tA7CWbuouRX5Sj5THgOW2SylADQ1H.wzjWcaL/H2KPUGbScXAm",
		Avatar:   "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/default_avatar.png",
	}, http.StatusOK, nil
}

func (m mockRepoSuccess) LoginByPhone(string) (_entity.User, int, error) {
	return _entity.User{
		Id:       1,
		Name:     "Salmaa",
		Password: "$2a$04$mE8tA7CWbuouRX5Sj5THgOW2SylADQ1H.wzjWcaL/H2KPUGbScXAm",
		Avatar:   "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/default_avatar.png",
	}, http.StatusOK, nil
}

func (m mockRepoSuccess) GetById(int) (_entity.User, int, error) {
	return _entity.User{
		Id:       1,
		Name:     "Salmaa",
		Email:    "salma@sirclo.com",
		Phone:    "08123456789",
		Password: "$2a$04$mE8tA7CWbuouRX5Sj5THgOW2SylADQ1H.wzjWcaL/H2KPUGbScXAm",
		Avatar:   "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/default_avatar.png",
	}, http.StatusOK, nil
}

func (m mockRepoSuccess) GetAll() ([]_entity.UserSimplified, int, error) {
	return []_entity.UserSimplified{
		{
			Id:     1,
			Name:   "Salmaa",
			Email:  "salma@sirclo.com",
			Phone:  "08123456789",
			Avatar: "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/default_avatar.png",
		},
	}, http.StatusOK, nil
}

func (m mockRepoSuccess) Update(_entity.User) (_entity.UserSimplified, int, error) {
	return _entity.UserSimplified{
		Id:     1,
		Name:   "Salmaa",
		Email:  "salma@sirclo.com",
		Phone:  "08123456789",
		Avatar: "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/default_avatar.png",
	}, http.StatusOK, nil
}

func (m mockRepoSuccess) Delete(int) (int, error) {
	return http.StatusOK, nil
}

// success

func TestCreateSuccess(t *testing.T) {
	t.Run("TestCreateSuccess", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "Salmaa",
			"email":    "salma@sirclo.com",
			"phone":    "08123456789",
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users")

		userController := New(mockRepoSuccess{})
		userController.Create()(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success create user",
			"data": map[string]interface{}{
				"id":   float64(1),
				"name": "Salmaa",
			},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestLoginByEmailSuccess(t *testing.T) {
	t.Run("TestLoginByEmailSuccess", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"input":    "salma@sirclo.com",
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/login")

		userController := New(mockRepoSuccess{})
		userController.Login()(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		data, _ := actual["data"].(map[string]interface{})
		token := data["token"]

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success login",
			"data": map[string]interface{}{
				"id":     float64(1),
				"name":   "Salmaa",
				"token":  token,
				"avatar": "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/default_avatar.png",
			},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestLoginByPhoneSuccess(t *testing.T) {
	t.Run("TestLoginByPhoneSuccess", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"input":    "08123456789",
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/login")

		userController := New(mockRepoSuccess{})
		userController.Login()(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		data, _ := actual["data"].(map[string]interface{})
		token := data["token"]

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success login",
			"data": map[string]interface{}{
				"id":     float64(1),
				"name":   "Salmaa",
				"token":  token,
				"avatar": "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/default_avatar.png",
			},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetByIdSuccess(t *testing.T) {
	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockRepoSuccess{})
		userController.GetById()(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success get user by id",
			"data": map[string]interface{}{
				"id":     float64(1),
				"name":   "Salmaa",
				"email":  "salma@sirclo.com",
				"phone":  "08123456789",
				"avatar": "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/default_avatar.png",
			},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestGetAllSuccess(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users")

		userController := New(mockRepoSuccess{})
		userController.GetAll()(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success get all users",
			"data": []interface{}{
				map[string]interface{}{
					"id":     float64(1),
					"name":   "Salmaa",
					"email":  "salma@sirclo.com",
					"phone":  "08123456789",
					"avatar": "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/default_avatar.png",
				},
			},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUpdateSuccess(t *testing.T) {
	t.Run("TestUpdateSuccess", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1)

		requestBody, _ := json.Marshal(map[string]string{
			"name":     "Salmaa",
			"email":    "salma@sirclo.com",
			"phone":    "08123456789",
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(userController.Update())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success update user",
			"data": map[string]interface{}{
				"id":     float64(1),
				"name":   "Salmaa",
				"email":  "salma@sirclo.com",
				"phone":  "08123456789",
				"avatar": "https://capstone-group3.s3.ap-southeast-1.amazonaws.com/default_avatar.png",
			},
		}

		assert.Equal(t, expected, actual)
	})
}

func TestDeleteSuccess(t *testing.T) {
	t.Run("TestDeleteSuccess", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1)

		request := httptest.NewRequest(http.MethodDelete, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(userController.Delete())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusOK),
			"message": "success delete user",
		}

		assert.Equal(t, expected, actual)
	})
}

// binding failure

func TestCreateFailBinding(t *testing.T) {
	t.Run("TestCreateFailBinding", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "Salmaa",
			"email":    "salma@sirclo.com",
			"phone":    8123456789,
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users")

		userController := New(mockRepoSuccess{})
		userController.Create()(context)

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

func TestLoginByEmailFailBinding(t *testing.T) {
	t.Run("TestLoginByEmailFailBinding", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]interface{}{
			"input":    "salma@sirclo.com",
			"password": 12345,
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/login")

		userController := New(mockRepoSuccess{})
		userController.Login()(context)

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

func TestLoginByPhoneFailBinding(t *testing.T) {
	t.Run("TestLoginByPhoneFailBinding", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]interface{}{
			"input":    "08123456789",
			"password": 1234,
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/login")

		userController := New(mockRepoSuccess{})
		userController.Login()(context)

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
		token, _, _ := _midware.CreateToken(1)

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "Salmaa",
			"email":    "salma@sirclo.com",
			"phone":    8123456789,
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(userController.Update())(context)

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
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "Salmaa",
			"email":    "salma@sirclo.com",
			"phone":    "",
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users")

		userController := New(mockRepoSuccess{})
		userController.Create()(context)

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

func TestLoginByEmailFailEmptyInput(t *testing.T) {
	t.Run("TestLoginByEmailFailEmptyInput", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"input":    "salma@sirclo.com",
			"password": "",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/login")

		userController := New(mockRepoSuccess{})
		userController.Login()(context)

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

func TestLoginByPhoneFailEmptyInput(t *testing.T) {
	t.Run("TestLoginByPhoneFailEmptyInput", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"input":    "08123456789",
			"password": "",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/login")

		userController := New(mockRepoSuccess{})
		userController.Login()(context)

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

// string pattern

func TestCreateFailMaliciousCharacter(t *testing.T) {
	t.Run("TestCreateFailMaliciousCharacter", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "Salmaa",
			"email":    "salma@sirclo.com",
			"phone":    "; --",
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users")

		userController := New(mockRepoSuccess{})
		userController.Create()(context)

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

func TestLoginByEmailFailMaliciousCharacter(t *testing.T) {
	t.Run("TestLoginByEmailFailMaliciousCharacter", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"input":    "salma@sirclo.com",
			"password": "; --",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/login")

		userController := New(mockRepoSuccess{})
		userController.Login()(context)

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

func TestLoginByPhoneFailMaliciousCharacter(t *testing.T) {
	t.Run("TestLoginByEmailFailMaliciousCharacter", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"input":    "08123456789",
			"password": "; --",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/login")

		userController := New(mockRepoSuccess{})
		userController.Login()(context)

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

func TestUpdateFailMaliciousCharacter1(t *testing.T) {
	t.Run("TestUpdateFailMaliciousCharacter1", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1)

		requestBody, _ := json.Marshal(map[string]string{
			"name":     "; --",
			"email":    "salma@sirclo.com",
			"phone":    "08123456789",
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(userController.Update())(context)

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

func TestUpdateFailMaliciousCharacter2(t *testing.T) {
	t.Run("TestUpdateFailMaliciousCharacter2", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1)

		requestBody, _ := json.Marshal(map[string]string{
			"name":     "Salmaa",
			"email":    "; --",
			"phone":    "08123456789",
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(userController.Update())(context)

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

func TestUpdateFailMaliciousCharacter3(t *testing.T) {
	t.Run("TestUpdateFailMaliciousCharacter3", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1)

		requestBody, _ := json.Marshal(map[string]string{
			"name":     "Salmaa",
			"email":    "salma@sirclo.com",
			"phone":    "; --",
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(userController.Update())(context)

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

func TestUpdateFailMaliciousCharacter4(t *testing.T) {
	t.Run("TestUpdateFailMaliciousCharacter4", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1)

		requestBody, _ := json.Marshal(map[string]string{
			"name":     "Salmaa",
			"email":    "salma@sirclo.com",
			"phone":    "08123456789",
			"password": "; --",
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(userController.Update())(context)

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

func TestCreateFailInvalidEmail(t *testing.T) {
	t.Run("TestCreateFailInvalidEmail", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "Salmaa",
			"email":    "salma@sir@clo.com",
			"phone":    "08123456789",
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users")

		userController := New(mockRepoSuccess{})
		userController.Create()(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusBadRequest),
			"message": "salma@sir@clo.com: email must contain exactly one local and domain name",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestLoginByEmailFailInvalidEmail(t *testing.T) {
	t.Run("TestLoginByEmailFailInvalidEmail", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"input":    "salma@sir@clo.com",
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/login")

		userController := New(mockRepoSuccess{})
		userController.Login()(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusBadRequest),
			"message": "salma@sir@clo.com: invalid email or phone",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUpdateFailInvalidEmail(t *testing.T) {
	t.Run("TestUpdateFailInvalidEmail", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1)

		requestBody, _ := json.Marshal(map[string]string{
			"name":     "Salmaa",
			"email":    "salma@sir@clo.com",
			"phone":    "08123456789",
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(userController.Update())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusBadRequest),
			"message": "salma@sir@clo.com: email must contain exactly one local and domain name",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestCreateFailInvalidPhone(t *testing.T) {
	t.Run("TestCreateFailInvalidPhone", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "Salmaa",
			"email":    "salma@sirclo.com",
			"phone":    "+628123456789",
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users")

		userController := New(mockRepoSuccess{})
		userController.Create()(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusBadRequest),
			"message": "+628123456789: phone number must contain numbers only",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestLoginByPhoneFailInvalidPhone(t *testing.T) {
	t.Run("TestLoginByEmailFailInvalidPhone", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"input":    "+628123456789",
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/login")

		userController := New(mockRepoSuccess{})
		userController.Login()(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusBadRequest),
			"message": "+628123456789: invalid email or phone",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUpdateFailInvalidPhone(t *testing.T) {
	t.Run("TestUpdateFailInvalidPhone", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1)

		requestBody, _ := json.Marshal(map[string]string{
			"name":     "Salmaa",
			"email":    "salma@sirclo.com",
			"phone":    "+628123456789",
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(userController.Update())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusBadRequest),
			"message": "+628123456789: phone number must contain numbers only",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestCreateFailInvalidPassword(t *testing.T) {
	t.Run("TestCreateFailInvalidPassword", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "Salmaa",
			"email":    "salma@sirclo.com",
			"phone":    "08123456789",
			"password": "aaa",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users")

		userController := New(mockRepoSuccess{})
		userController.Create()(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusBadRequest),
			"message": "aaa: password must be minimum 6 characters long",
		}

		assert.Equal(t, expected, actual)
	})
}

func TestUpdateFailInvalidPassword(t *testing.T) {
	t.Run("TestUpdateFailInvalidPassword", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1)

		requestBody, _ := json.Marshal(map[string]string{
			"name":     "Salmaa",
			"email":    "salma@sirclo.com",
			"phone":    "08123456789",
			"password": "aaa",
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(userController.Update())(context)

		actual := map[string]interface{}{}
		body := response.Body.String()
		json.Unmarshal([]byte(body), &actual)

		expected := map[string]interface{}{
			"code":    float64(http.StatusBadRequest),
			"message": "aaa: password must be minimum 6 characters long",
		}

		assert.Equal(t, expected, actual)
	})
}

// invalid parameter

func TestGetByIdFailInvalidParameter(t *testing.T) {
	t.Run("TestGetByIdFailInvalidParameter", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("a")

		userController := New(mockRepoSuccess{})
		userController.GetById()(context)

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

func TestUpdateFailInvalidParameter(t *testing.T) {
	t.Run("TestUpdateFailInvalidParameter", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1)

		requestBody, _ := json.Marshal(map[string]string{
			"name":     "Salmaa",
			"email":    "salma@sirclo.com",
			"phone":    "08123456789",
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("a")

		userController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(userController.Update())(context)

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

func TestDeleteFailInvalidParameter(t *testing.T) {
	t.Run("TestDeleteInvalidParameter", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1)

		request := httptest.NewRequest(http.MethodDelete, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("a")

		userController := New(mockRepoSuccess{})
		_midware.JWTMiddleWare()(userController.Delete())(context)

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

type mockRepoFail struct{}

func (m mockRepoFail) Create(_entity.User) (_entity.UserSimplified, int, error) {
	return _entity.UserSimplified{}, http.StatusInternalServerError, errors.New("internal server error")
}

func (m mockRepoFail) LoginByEmail(string) (_entity.User, int, error) {
	return _entity.User{}, http.StatusInternalServerError, errors.New("internal server error")
}

func (m mockRepoFail) LoginByPhone(string) (_entity.User, int, error) {
	return _entity.User{}, http.StatusInternalServerError, errors.New("internal server error")
}

func (m mockRepoFail) GetById(int) (_entity.User, int, error) {
	return _entity.User{}, http.StatusInternalServerError, errors.New("internal server error")
}

func (m mockRepoFail) GetAll() ([]_entity.UserSimplified, int, error) {
	return nil, http.StatusInternalServerError, errors.New("internal server error")
}

func (m mockRepoFail) Update(_entity.User) (_entity.UserSimplified, int, error) {
	return _entity.UserSimplified{}, http.StatusInternalServerError, errors.New("internal server error")
}

func (m mockRepoFail) Delete(int) (int, error) {
	return http.StatusInternalServerError, errors.New("internal server error")
}

func TestCreateFailRepo(t *testing.T) {
	t.Run("TestCreateFailRepo", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "Salmaa",
			"email":    "salma@sirclo.com",
			"phone":    "08123456789",
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users")

		userController := New(mockRepoFail{})
		userController.Create()(context)

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

func TestLoginByEmailFailRepo(t *testing.T) {
	t.Run("TestLoginByEmailFailRepo", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"input":    "salma@sirclo.com",
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/login")

		userController := New(mockRepoFail{})
		userController.Login()(context)

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

func TestLoginByPhoneFailRepo(t *testing.T) {
	t.Run("TestLoginByPhoneFailRepo", func(t *testing.T) {
		requestBody, _ := json.Marshal(map[string]string{
			"input":    "08123456789",
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/login")

		userController := New(mockRepoFail{})
		userController.Login()(context)

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

func TestGetByIdFailRepo(t *testing.T) {
	t.Run("TestGetByIdFailRepo", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockRepoFail{})
		userController.GetById()(context)

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
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users")

		userController := New(mockRepoFail{})
		userController.GetAll()(context)

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
		token, _, _ := _midware.CreateToken(1)

		requestBody, _ := json.Marshal(map[string]string{
			"name":     "Salmaa",
			"email":    "salma@sirclo.com",
			"phone":    "08123456789",
			"password": "74nSA&ge%#fwJ",
		})

		request := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockRepoFail{})
		_midware.JWTMiddleWare()(userController.Update())(context)

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

func TestDeleteFailRepo(t *testing.T) {
	t.Run("TestDeleteFailRepo", func(t *testing.T) {
		token, _, _ := _midware.CreateToken(1)

		request := httptest.NewRequest(http.MethodDelete, "/", nil)
		request.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		response := httptest.NewRecorder()

		e := echo.New()

		context := e.NewContext(request, response)
		context.SetPath("/users/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := New(mockRepoFail{})
		_midware.JWTMiddleWare()(userController.Delete())(context)

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
