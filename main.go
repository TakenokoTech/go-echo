package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go-echo/di"
	"go-echo/utils"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(utils.LoggerMiddleware())
	e.Use(utils.RecoverMiddleware())
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	di.Setup(e, e.Group("/v1"))
	go utils.RoutingList(e)
	e.Logger.Fatal(e.Start(":8080"))
}
