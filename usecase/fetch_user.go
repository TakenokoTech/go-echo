package usecase

import (
	"go-echo/domain/core"
	"go-echo/domain/users"
)

type fetchUserUseCase struct {
	c core.Repository
	u users.Repository
}

func NewFetchUserUseCase(c core.Repository, u users.Repository) fetchUserUseCase {
	return fetchUserUseCase{c: c, u: u}
}

func (uc fetchUserUseCase) Execute() []users.User {
	return uc.u.SelectAll()
}
