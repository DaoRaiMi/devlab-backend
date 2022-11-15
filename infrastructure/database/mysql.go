package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConnection(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
