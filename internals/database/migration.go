package database

import (
	"errors"
	"github.com/islamghany/go-prod-rest/internals/comment"
	"gorm.io/gorm"
)

// migrate our database and create our comment table
func MigrateDB(db *gorm.DB) error {
	if res := db.AutoMigrate(&comment.Comment{}); res.Error != nil {
		return errors.New(res.Error())
	}
	return nil
}
