package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-echo/domain/users"
	"net/http"
)

type usersController struct {
	u users.Repository
}

func SetupUsersController(g *echo.Group, u users.Repository) {
	uc := usersController{
		u: u,
	}
	g.GET("/users", uc.getUser)
	g.GET("/users/:id", uc.getUserById)
	g.POST("/users", uc.postUser)
}

func (uc usersController) getUser(c echo.Context) error {
	return c.JSON(http.StatusOK, uc.u.SelectAll())
}

func (uc usersController) getUserById(c echo.Context) error {
	return c.JSON(http.StatusOK, uc.u.SelectById(c.Param("id")))
}

func (uc usersController) postUser(c echo.Context) error {
	var user users.User
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := c.Validate(&user); err != nil {
		return err
	}
	fmt.Printf("%v\n", user)
	uc.u.Insert(user)
	return c.JSON(http.StatusOK, uc.u.SelectAll())
}
