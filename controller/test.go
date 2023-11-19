package controller

import (
	"github.com/labstack/echo/v4"
	"go-echo/domain/test"
	"go-echo/domain/users"
	"net/http"
	"strconv"
)

type testController struct {
	u users.Repository
}

func SetupTestController(g *echo.Group, u users.Repository) {
	tc := testController{u: u}
	g.GET("/test", tc.getTest)
	g.GET("/test/:id", tc.getTest)
	g.POST("/test/:id", tc.postTest)
}

func (tc testController) getTest(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
		return err
	}
	return c.JSON(http.StatusOK, test.NewTest(id, "ok"))
}

func (tc testController) postTest(c echo.Context) error {
	panic("aaa")
	return nil
}
