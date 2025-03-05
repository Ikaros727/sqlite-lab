package repo

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

const (
	dsnFormat = "file:%s.sqlite"
)

func NewSQLite(dbName string) (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open(fmt.Sprintf(dsnFormat, dbName)), &gorm.Config{})
	if err != nil {
		return
	}
	return
}
