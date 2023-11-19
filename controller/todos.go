package controller

import (
	"github.com/labstack/echo/v4"
	"go-echo/domain/todos"
	"net/http"
)

type todosController struct {
	u todos.Repository
}

func SetupTodosController(g *echo.Group, u todos.Repository) {
	tc := todosController{u: u}
	g.GET("/todos", tc.getUser)
	g.POST("/todos", tc.postUser)
}

func (tc todosController) getUser(c echo.Context) error {
	return c.JSON(http.StatusOK, tc.u.SelectAll())
}

func (tc todosController) postUser(c echo.Context) error {
	var todo todos.Todo
	if err := c.Bind(&todo); err != nil {
		return err
	}
	if err := c.Validate(&todo); err != nil {
		return err
	}
	tc.u.Insert(todo)
	return c.JSON(http.StatusOK, tc.u.SelectAll())
}
