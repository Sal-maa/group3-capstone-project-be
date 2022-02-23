package router

import (
	_asset "capstone/be/delivery/controller/asset"
	_history "capstone/be/delivery/controller/history"
	_request "capstone/be/delivery/controller/request"
	_user "capstone/be/delivery/controller/user"
	_midware "capstone/be/delivery/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

const swagger string = "<a href=\"https://app.swaggerhub.com/apis-docs/bagusbpg6/group3-capstone-API/1.0.0\">Visit API documentation</a>"

func RegisterPath(
	e *echo.Echo,
	userController *_user.UserController,
	assetController *_asset.AssetController,
	historyController *_history.HistoryController,
	requestController *_request.RequestController,
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

	// Asset
	e.POST("/assets", assetController.Create())

	e.GET("/assets", assetController.GetAll())
	e.GET("/assets/:id", assetController.GetById())
	e.PUT("/assets/:id", assetController.Update())

	// History
	e.GET("/histories/users/:user_id", historyController.GetAllUsageHistoryOfUser(), _midware.JWTMiddleWare())
	e.GET("/histories/users/:user_id/:request_id", historyController.GetDetailUsageHistoryByRequestId(), _midware.JWTMiddleWare())
	e.GET("/histories/assets/:asset_id", historyController.GetAllUsageHistoryOfAsset(), _midware.JWTMiddleWare())

	// Request by Employee
	e.POST("/requests/borrow", requestController.Borrow(), _midware.JWTMiddleWare())
	e.PUT("/requests/borrow", requestController.CancelBorrow(), _midware.JWTMiddleWare())
	e.POST("/requests/procure", requestController.Procure(), _midware.JWTMiddleWare())
	e.PUT("/requests/procure", requestController.CancelProcure(), _midware.JWTMiddleWare())

	// Update by Manager
	e.PUT("/requests/borrow/:id", requestController.UpdateBorrow(), _midware.JWTMiddleWare())
	e.PUT("/requests/procure/:id", requestController.UpdateProcure(), _midware.JWTMiddleWare())

	// Update by Administrator
	e.PUT("/requests/borrow/:id", requestController.UpdateBorrowByAdmin(), _midware.JWTMiddleWare())
	e.PUT("/requests/procure/:id", requestController.UpdateProcureByAdmin(), _midware.JWTMiddleWare())
}
