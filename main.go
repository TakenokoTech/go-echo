package main

import (
	"github.com/labstack/echo/v4"
	"go-echo/controller"
	"go-echo/utils"
	"net/http"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(utils.LoggerMiddleware())
	e.Use(utils.RecoverMiddleware())
	e.GET("/", getIndex)

	v1 := e.Group("/v1")
	controller.NewTest(v1)

	go utils.RoutingList(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func getIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
