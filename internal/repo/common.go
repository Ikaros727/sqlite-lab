package repo

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"os"
)

const (
	dsnFormat = "file:database/%d/%s.opensql"
)

func NewSQLite(userID int64, database string) (db *gorm.DB, err error) {
	if err = os.MkdirAll(fmt.Sprintf("database/%d", userID), 0755); err != nil {
		return
	}

	db, err = gorm.Open(sqlite.Open(fmt.Sprintf(dsnFormat, userID, database)), &gorm.Config{})
	if err != nil {
		return
	}
	return
}
