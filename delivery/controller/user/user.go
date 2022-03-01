package user

import (
	_common "capstone/be/delivery/common"
	_helper "capstone/be/delivery/helper"
	_midware "capstone/be/delivery/middleware"
	_entity "capstone/be/entity"
	_userRepo "capstone/be/repository/user"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	repository _userRepo.User
}

func New(user _userRepo.User) *UserController {
	return &UserController{repository: user}
}

func (uc UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		loginData := _entity.Login{}

		// handle failure in binding
		if err := c.Bind(&loginData); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}

		// prepare input string
		input := strings.TrimSpace(loginData.Input)
		password := strings.TrimSpace(loginData.Password)

		// check input string
		check := []string{input, password}

		for _, s := range check {
			// check empty string in required input
			if s == "" {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "input cannot be empty"))
			}

			// check malicious character in input
			if err := _helper.CheckStringInput(s); err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, s+": "+err.Error()))
			}
		}

		loginUser := _entity.User{}

		// check input pattern
		if err := _helper.CheckEmailPattern(input); err == nil {
			// if login by email, calling repository
			login, code, err := uc.repository.LoginByEmail(input)

			// detect failure in repository
			if err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}

			loginUser = login
		} else if err = _helper.CheckPhonePattern(input); err == nil {
			// if login by phone, calling repository
			login, code, err := uc.repository.LoginByPhone(input)

			// detect failure in repository
			if err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}

			loginUser = login
		} else {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, input+": invalid email or phone"))
		}

		// detect password mismatch
		if err := bcrypt.CompareHashAndPassword([]byte(loginUser.Password), []byte(password)); err != nil {
			return c.JSON(http.StatusUnauthorized, _common.NoDataResponse(http.StatusUnauthorized, "password does not match"))
		}

		// create token based on user id
		token, expire, err := _midware.CreateToken(loginUser.Id, loginUser.Role)

		// detect failure in creating token
		if err != nil {
			return c.JSON(http.StatusInternalServerError, _common.NoDataResponse(http.StatusInternalServerError, "failed to create token"))
		}

		return c.JSON(http.StatusOK, _common.LoginResponse(loginUser, token, expire))
	}
}

func (uc UserController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		// calling repository
		users, code, err := uc.repository.GetAll()

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		return c.JSON(http.StatusOK, _common.GetAllUsersResponse(users))
	}
}

func (uc UserController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		// detect invalid parameter
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid user id"))
		}

		// calling repository
		user, code, err := uc.repository.GetById(id)

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		return c.JSON(http.StatusOK, _common.GetUserByIdResponse(user))
	}
}

func (uc UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		// detect invalid parameter
		if err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "invalid user id"))
		}

		// check authorization
		if id != _midware.ExtractId(c) {
			return c.JSON(http.StatusUnauthorized, _common.NoDataResponse(http.StatusUnauthorized, "unauthorized"))
		}

		userData := _entity.UpdateUser{}

		// detect failure in binding
		if err := c.Bind(&userData); err != nil {
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to bind data"))
		}

		// prepare input string
		division := strings.Title(strings.ToLower(strings.TrimSpace(userData.Division)))
		name := strings.Title(strings.ToLower(strings.TrimSpace(userData.Name)))
		email := strings.TrimSpace(userData.Email)
		phone := strings.TrimSpace(userData.Phone)
		password := strings.TrimSpace(userData.Password)
		gender := strings.Title(strings.ToLower(strings.TrimSpace(userData.Gender)))
		address := strings.Title(strings.ToLower(strings.TrimSpace(userData.Address)))

		// calling repository to get existing user data
		updateUserData, code, err := uc.repository.GetById(id)

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		// detect change in user division
		if division != "" {
			// check malicious character in input
			if err := _helper.CheckStringInput(division); err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, division+": "+err.Error()))
			}

			updateUserData.Division = division
		}

		// detect change in user name
		if name != "" {
			// check malicious character in input
			if err := _helper.CheckStringInput(name); err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, name+": "+err.Error()))
			}

			updateUserData.Name = name
		}

		// detect change in user email
		if email != "" {
			// check malicious character in input
			if err := _helper.CheckStringInput(email); err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, email+": "+err.Error()))
			}

			// check email pattern
			if err := _helper.CheckEmailPattern(email); err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, email+": "+err.Error()))
			}

			updateUserData.Email = email
		}

		// detect change in user phone
		if phone != "" {
			// check malicious character in input
			if err := _helper.CheckStringInput(phone); err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, phone+": "+err.Error()))
			}

			// check phone pattern
			if err := _helper.CheckPhonePattern(phone); err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, phone+": "+err.Error()))
			}

			updateUserData.Phone = phone
		}

		// detect change in user password
		if password != "" {
			// check malicious character in input
			if err := _helper.CheckStringInput(password); err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, password+": "+err.Error()))
			}

			// check password pattern
			if err := _helper.CheckPasswordPattern(password); err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, password+": "+err.Error()))
			}

			// hashing password
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

			// detect failure in hashing password
			if err != nil {
				return c.JSON(http.StatusInternalServerError, _common.NoDataResponse(http.StatusInternalServerError, "failed to hash password"))
			}

			updateUserData.Password = string(hashedPassword)
		}

		// detect change in user gender
		if gender != "" {
			// check malicious character in input
			if err := _helper.CheckStringInput(gender); err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, gender+": "+err.Error()))
			}

			updateUserData.Gender = gender
		}

		// detect change in user address
		if address != "" {
			// check malicious character in input
			if err := _helper.CheckStringInput(address); err != nil {
				return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, address+": "+err.Error()))
			}

			updateUserData.Address = address
		}

		// detect avatar image upload
		src, file, err := c.Request().FormFile("avatar")

		// detect failure in parsing file
		switch err {
		case nil:
			defer src.Close()

			// upload avatar to amazon s3
			avatar, code, err := _helper.UploadImage("user", id, file, src)

			// detect failure while uploading avatar
			if err != nil {
				return c.JSON(code, _common.NoDataResponse(code, err.Error()))
			}

			// delete previous avatar from amazon s3
			filename := updateUserData.Avatar[strings.LastIndex(updateUserData.Avatar, "/")+1:]

			if filename != "default_avatar.png" {

				if err = _helper.DeleteImage(filename); err != nil {
					log.Println(err)
				}
			}

			updateUserData.Avatar = avatar
		case http.ErrMissingFile:
			avatar := updateUserData.Avatar[strings.LastIndex(updateUserData.Avatar, "/")+1:]
			updateUserData.Avatar = avatar
		case http.ErrNotMultipart:
			avatar := updateUserData.Avatar[strings.LastIndex(updateUserData.Avatar, "/")+1:]
			updateUserData.Avatar = avatar
		default:
			log.Println(err)
			return c.JSON(http.StatusBadRequest, _common.NoDataResponse(http.StatusBadRequest, "failed to upload avatar"))
		}

		updateUserData.Id = id

		// calling repository
		updatedUser, code, err := uc.repository.Update(updateUserData)

		// detect failure in repository
		if err != nil {
			return c.JSON(code, _common.NoDataResponse(code, err.Error()))
		}

		return c.JSON(http.StatusOK, _common.UpdateUserResponse(updatedUser))
	}
}
