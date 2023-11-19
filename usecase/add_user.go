package usecase

import (
	"go-echo/domain/core"
	"go-echo/domain/users"
	"gorm.io/gorm"
)

type addUserUseCase struct {
	c core.Repository
	u users.Repository
}

func NewAddUserUseCase(c core.Repository, u users.Repository) addUserUseCase {
	return addUserUseCase{c: c, u: u}
}

func (uc addUserUseCase) Execute(user users.User) error {
	return uc.c.Transaction(func(tx *gorm.DB) error {
		uc.u.Insert(tx, user)
		return nil
	})
}
