package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type indexController struct {
}

func SetupIndexController(e *echo.Echo) {
	ic := indexController{}
	e.GET("/", ic.getIndex)
}

func (ic indexController) getIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
