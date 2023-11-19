package di

import (
	"github.com/labstack/echo/v4"
	"go-echo/controller"
	"go-echo/repository"
	"go-echo/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Setup(root *echo.Echo, v1 *echo.Group) {
	db := utils.ConnectSQL()
	gorm, _ := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	ur := repository.NewUsersRepository(gorm)
	controller.SetupIndexController(root)
	controller.SetupTestController(v1, ur)
	controller.SetupUsersController(v1, ur)
}
