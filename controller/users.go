package controller

import (
	"github.com/labstack/echo/v4"
	"go-echo/domain/core"
	"go-echo/domain/users"
	"gorm.io/gorm"
	"net/http"
)

type usersController struct {
	c core.Repository
	u users.Repository
}

func SetupUsersController(g *echo.Group, c core.Repository, u users.Repository) {
	uc := usersController{
		c: c,
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
	return uc.c.Transaction(func(tx *gorm.DB) error {
		var user users.User
		if err := c.Bind(&user); err != nil {
			return err
		}
		if err := c.Validate(&user); err != nil {
			return err
		}
		uc.u.Insert(tx, user)
		return c.JSON(http.StatusOK, uc.u.SelectAll())
	})
}
