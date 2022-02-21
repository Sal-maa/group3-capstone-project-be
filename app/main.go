package main

import (
	_config "capstone/be/config"
	_requestController "capstone/be/delivery/controller/request"
	_userController "capstone/be/delivery/controller/user"
	_midware "capstone/be/delivery/middleware"
	_router "capstone/be/delivery/router"
	_requestRepo "capstone/be/repository/request"
	_userRepo "capstone/be/repository/user"
	_util "capstone/be/util"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

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
	requestRepo := _requestRepo.New(db)

	userController := _userController.New(userRepo)
	requestController := _requestController.New(requestRepo)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash(), middleware.CORS(), _midware.CustomLogger())

	_router.RegisterPath(e, userController, requestController)

	address := fmt.Sprintf(":%d", config.Port)
	e.Logger.Fatal(e.Start(address))
}
