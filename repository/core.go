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

func (c coreRepository) Transaction(cb core.TxCallback) error {
	return c.db.Transaction(cb)
}
