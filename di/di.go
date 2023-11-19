package di

import (
	"github.com/labstack/echo/v4"
	"go-echo/controller"
	"go-echo/repository"
	"go-echo/usecase"
	"go-echo/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Setup(root *echo.Echo, v1 *echo.Group) {
	db := utils.ConnectSQL()
	g, _ := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	cr := repository.NewCoreRepository(g)
	ur := repository.NewUsersRepository(g)
	tr := repository.NewTodosRepository(g)
	au := usecase.NewAddUserUseCase(cr, ur)
	fu := usecase.NewFetchUserUseCase(cr, ur)
	controller.SetupIndexController(root)
	controller.SetupTestController(v1, ur)
	controller.SetupUsersController(v1, cr, ur, au, fu)
	controller.SetupTodosController(v1, tr)
}
