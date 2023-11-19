package repository

import (
	"go-echo/domain/core"
	"gorm.io/gorm"
)

type coreRepository struct {
	db *gorm.DB
}

var _ core.Repository = coreRepository{}

func NewCoreRepository(db *gorm.DB) core.Repository {
	return coreRepository{
		db: db,
	}
}

func (c coreRepository) Transaction(fc func(tx *gorm.DB) error) error {
	return c.db.Transaction(fc)
}
