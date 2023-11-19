package controller

import (
	"github.com/labstack/echo/v4"
	"go-echo/domain/core"
	"go-echo/domain/users"
	"go-echo/usecase"
	"net/http"
)

type usersController struct {
	c  core.Repository
	u  users.Repository
	au usecase.AddUserUseCase
	fu usecase.FetchUserUseCase
}

func SetupUsersController(
	g *echo.Group,
	c core.Repository,
	u users.Repository,
	au usecase.AddUserUseCase,
	fu usecase.FetchUserUseCase,
) {
	uc := usersController{c: c, u: u, au: au, fu: fu}
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
	if err := uc.au.Execute(user); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, uc.fu.Execute())
}
