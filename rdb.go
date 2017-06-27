package rdb

import (
	"github.com/jinzhu/gorm"
)

// DB is the gorm database which is used for every operation.
var DB *gorm.DB

// OpenDatabase opens a new database connection.
func OpenDatabase(dialect string, args ...interface{}) error {
	var err error
	DB, err = gorm.Open(dialect, args...)
	return err
}
