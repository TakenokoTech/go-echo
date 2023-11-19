package core

import "gorm.io/gorm"

func TransactionDummy(tx *gorm.DB) interface{} {
	return func(cb TxCallback) error { return cb(tx) }
}
