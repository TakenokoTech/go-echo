package repository

import (
	"database/sql"
	"go-echo/domain/users"
)

type usersRepository struct {
	db *sql.DB
}

var _ users.Repository = usersRepository{}

func NewUsersRepository(db *sql.DB) users.Repository {
	return usersRepository{
		db: db,
	}
}

func (u usersRepository) a() {

}
