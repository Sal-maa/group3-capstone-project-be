package router

import (
	_user "capstone/be/delivery/controller/user"
	_midware "capstone/be/delivery/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

const swagger string = "<a href=\"https://app.swaggerhub.com/apis-docs/bagusbpg6/group3-capstone-API/1.0.0\">Visit API documentation</a>"

func RegisterPath(
	e *echo.Echo,
	userController *_user.UserController,
) {
	// Root
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, swagger)
	})

	// Login
	e.POST("/login", userController.Login())

	// User
	e.POST("/users", userController.Create())
	e.GET("/users", userController.GetAll())
	e.GET("/users/:id", userController.GetById())
	e.PUT("/users/:id", userController.Update(), _midware.JWTMiddleWare())
	e.DELETE("/users/:id", userController.Delete(), _midware.JWTMiddleWare())
}
