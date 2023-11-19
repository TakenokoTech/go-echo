package repository

import (
	"go-echo/domain/users"
	"go-echo/utils"
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

func (u usersRepository) SelectAll() (users []users.User) {
	utils.
		GetTransaction(u.db, nil).
		Find(&users)
	return
}

func (u usersRepository) SelectById(id string) (user users.User) {
	utils.
		GetTransaction(u.db, nil).
		Where("id = ?", id).
		First(&user)
	return
}

func (u usersRepository) Insert(tx *gorm.DB, user users.User) {
	utils.
		GetTransaction(u.db, tx).
		Select("username").
		Create(&user)
}
