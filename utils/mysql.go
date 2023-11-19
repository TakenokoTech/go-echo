package utils

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

func ConnectSQL() *sql.DB {
	tz, err := time.LoadLocation(os.Getenv("TZ"))
	if err != nil {
		panic(err)
	}
	c := mysql.Config{
		DBName:    os.Getenv("DB_NAME"),
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASS"),
		Addr:      os.Getenv("DB_ADDR"),
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       tz,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		panic(err)
	}
	return db
}

func GetTransaction(db *gorm.DB, tx *gorm.DB) *gorm.DB {
	if tx != nil {
		return tx
	} else {
		return db
	}
}
