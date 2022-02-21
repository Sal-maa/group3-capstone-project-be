package router

import (
	_history "capstone/be/delivery/controller/history"
	_user "capstone/be/delivery/controller/user"
	_midware "capstone/be/delivery/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

const swagger string = "<a href=\"https://app.swaggerhub.com/apis-docs/bagusbpg6/group3-capstone-API/1.0.0\">Visit API documentation</a>"

func RegisterPath(
	e *echo.Echo,
	userController *_user.UserController,
	historyController *_history.HistoryController,
) {
	// Root
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, swagger)
	})

	// Login
	e.POST("/login", userController.Login())

	// User
	e.GET("/users", userController.GetAll(), _midware.JWTMiddleWare())
	e.GET("/users/:id", userController.GetById(), _midware.JWTMiddleWare())
	e.PUT("/users/:id", userController.Update(), _midware.JWTMiddleWare())

	// History
	e.GET("/histories/users/:user_id", historyController.GetAllUsageHistoryOfUser(), _midware.JWTMiddleWare())
	e.GET("/histories/users/:user_id/:request_id", historyController.GetDetailUsageHistoryByRequestId(), _midware.JWTMiddleWare())
	e.GET("/histories/assets/:asset_id", historyController.GetAllUsageHistoryOfAsset(), _midware.JWTMiddleWare())
}
