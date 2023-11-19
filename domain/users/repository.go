package users

import "gorm.io/gorm"

type Repository interface {
	SelectAll() (users []User)
	SelectById(id string) (user User)
	Insert(tx *gorm.DB, user User)
}
