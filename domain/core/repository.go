package core

import "gorm.io/gorm"

type Repository interface {
	Transaction(fc func(tx *gorm.DB) error) error
}
