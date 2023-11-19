package repository

import (
	"go-echo/domain/users"
	"gorm.io/gorm"
)

type usersRepository struct {
	db *gorm.DB
}

var _ users.Repository = usersRepository{}

func NewUsersRepository(db *gorm.DB) users.Repository {
	return usersRepository{
		db: db,
	}
}

func (u usersRepository) SelectAll() (user []users.User) {
	u.db.Find(&user)
	return
}

func (u usersRepository) SelectById(id string) (user users.User) {
	u.db.Where("id = ?", id).First(&user)
	return
}

func (u usersRepository) Insert(user users.User) {
	u.db.Select("username").Create(&user)
}
