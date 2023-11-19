package core

import "gorm.io/gorm"

type TxCallback = func(tx *gorm.DB) error

type Repository interface {
	Transaction(fc func(tx *gorm.DB) error) error
}
