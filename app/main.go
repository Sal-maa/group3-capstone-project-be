package main

import (
	_config "capstone/be/config"
	_paymentController "capstone/be/delivery/controller/payment"
	_userController "capstone/be/delivery/controller/user"
	_midware "capstone/be/delivery/middleware"
	_router "capstone/be/delivery/router"
	_paymentRepo "capstone/be/repository/payment"
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
	paymentRepo := _paymentRepo.New(db)

	userController := _userController.New(userRepo)
	paymentController := _paymentController.New(paymentRepo)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash(), middleware.CORS(), _midware.CustomLogger())

	_router.RegisterPath(e, userController, paymentController)
	// koneksi
	address := fmt.Sprintf(":%d", config.Port)
	e.Logger.Fatal(e.Start(address))
}
