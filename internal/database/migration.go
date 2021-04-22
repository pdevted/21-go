package database

import (
	"github.com/jinzhu/gorm"
	"github.com/pdevted/go-rest-api/internal/comment"
)

// MigrateDB - migrates database and creates comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
