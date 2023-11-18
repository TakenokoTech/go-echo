package di

import (
	"github.com/labstack/echo/v4"
	"go-echo/controller"
	"go-echo/repository"
	"go-echo/utils"
)

func Setup(root *echo.Echo, v1 *echo.Group) {
	db := utils.ConnectSQL()
	ur := repository.NewUsersRepository(db)
	controller.SetupIndexController(root)
	controller.SetupTestController(v1, ur)
}
