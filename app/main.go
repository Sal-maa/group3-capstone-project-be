package main

import (
	_config "capstone/be/config"
	_assetController "capstone/be/delivery/controller/asset"
	_historyController "capstone/be/delivery/controller/history"
	_requestController "capstone/be/delivery/controller/request"
	_userController "capstone/be/delivery/controller/user"
	_midware "capstone/be/delivery/middleware"
	_router "capstone/be/delivery/router"
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
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	config := _config.GetConfig()

	db, err := _util.GetDBInstance(config)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	userRepo := _userRepo.New(db)
	assetRepo := _assetRepo.New(db)
	historyRepo := _historyRepo.New(db)
	requestRepo := _requestRepo.New(db)

	userController := _userController.New(userRepo)
	assetController := _assetController.New(assetRepo)
	historyController := _historyController.New(historyRepo)
	requestController := _requestController.New(requestRepo)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash(), middleware.CORS(), _midware.CustomLogger())

	_router.RegisterPath(e, userController, assetController, historyController, requestController)

	address := fmt.Sprintf(":%d", config.Port)
	e.Logger.Fatal(e.Start(address))
}
