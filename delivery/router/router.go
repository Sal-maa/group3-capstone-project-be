package router

import (
	_activity "capstone/be/delivery/controller/activity"
	_admin "capstone/be/delivery/controller/admin"
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
	activityController *_activity.ActivityController,
	adminController *_admin.AdminController,
) {
	// Root
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, swagger)
	})

	// Login
	e.POST("/login", userController.Login())

	// User
	e.GET("/users/:id", userController.GetById(), _midware.JWTMiddleWare())
	e.PUT("/users/:id", userController.Update(), _midware.JWTMiddleWare())

	// Asset
	e.POST("/assets", assetController.Create(), _midware.JWTMiddleWare())
	e.GET("/assets", assetController.GetAll(), _midware.JWTMiddleWare())
	e.GET("/assets/:short_name", assetController.GetByShortName(), _midware.JWTMiddleWare())
	e.PUT("/assets/:short_name", assetController.Update(), _midware.JWTMiddleWare())
	e.GET("/stats", assetController.GetStats(), _midware.JWTMiddleWare())

	// History
	e.GET("/histories/users/:user_id", historyController.GetAllRequestHistoryOfUser(), _midware.JWTMiddleWare())
	e.GET("/histories/users/:user_id/:request_id", historyController.GetDetailRequestHistoryByRequestId(), _midware.JWTMiddleWare())
	e.GET("/histories/assets/:short_name", historyController.GetAllUsageHistoryOfAsset(), _midware.JWTMiddleWare())

	// Request by Employee
	e.POST("/requests/borrow", requestController.Borrow(), _midware.JWTMiddleWare())
	e.POST("/requests/procure", requestController.Procure(), _midware.JWTMiddleWare())

	// Update by Manager and Admin
	e.PUT("/requests/borrow/:id", requestController.UpdateBorrow(), _midware.JWTMiddleWare())
	e.PUT("/requests/procure/:id", requestController.UpdateProcure(), _midware.JWTMiddleWare())

	// Admin & Manager Page
	e.GET("/requests/admin", adminController.AdminGetAll(), _midware.JWTMiddleWare())
	e.GET("/requests/manager", adminController.ManagerGetAll(), _midware.JWTMiddleWare())

	// Activity
	e.GET("/activities/:user_id", activityController.GetAllActivityOfUser(), _midware.JWTMiddleWare())
	e.GET("/activities/:user_id/:request_id", activityController.GetDetailActivityByRequestId(), _midware.JWTMiddleWare())
	e.PUT("/activities/:user_id/:request_id", activityController.UpdateRequestStatus(), _midware.JWTMiddleWare())
}
