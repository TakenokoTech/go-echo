package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Test struct {
	Id     int    `json:"id"`
	Result string `json:"result"`
}

func NewTest(g *echo.Group) {
	g.GET("/test", getTest)
	g.GET("/test/:id", getTest)
	g.POST("/test/:id", postTest)
}

func getTest(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	println(c.Param("id"))
	if err != nil {
		c.Error(err)
		return err
	}
	return c.JSON(http.StatusOK, Test{
		Id:     id,
		Result: "ok",
	})
}

func postTest(c echo.Context) error {
	panic("aaa")
	return nil
}
