package main

import (
	_config "capstone/be/config"
	_activityController "capstone/be/delivery/controller/activity"
	_adminController "capstone/be/delivery/controller/admin"
	_assetController "capstone/be/delivery/controller/asset"
	_historyController "capstone/be/delivery/controller/history"
	_requestController "capstone/be/delivery/controller/request"
	_userController "capstone/be/delivery/controller/user"
	_midware "capstone/be/delivery/middleware"
	_router "capstone/be/delivery/router"
	_activityRepo "capstone/be/repository/activity"
	_adminRepo "capstone/be/repository/admin"
	_assetRepo "capstone/be/repository/asset"
	_historyRepo "capstone/be/repository/history"
	_requestRepo "capstone/be/repository/request"
	_userRepo "capstone/be/repository/user"

	_util "capstone/be/util"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
}

func main() {
	config := _config.GetConfig()

	db, err := _util.GetDBInstance(config)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	userRepo := _userRepo.New(db)
	activityRepo := _activityRepo.New(db)
	assetRepo := _assetRepo.New(db)
	historyRepo := _historyRepo.New(db)
	requestRepo := _requestRepo.New(db)
	adminRepo := _adminRepo.New(db)

	userController := _userController.New(userRepo)
	activityController := _activityController.New(activityRepo)
	assetController := _assetController.New(assetRepo)
	historyController := _historyController.New(historyRepo)
	requestController := _requestController.New(requestRepo)
	adminController := _adminController.New(adminRepo)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash(), middleware.CORS(), _midware.CustomLogger())

	_router.RegisterPath(e,
		userController,
		assetController,
		historyController,
		requestController,
		activityController,
		adminController,
	)

	address := fmt.Sprintf(":%d", config.Port)
	e.Logger.Fatal(e.Start(address))
}
