package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // database driver for gorm

	// _ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/schulke-214/hkp/pkg/domain/models"
)

// DB is the database instance
var DB *gorm.DB

func init() {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("failed to connect database")
	}

	DB = db

	AutoMigrate()
}

// AutoMigrate provides bindings to db.AutoMigrate for all Models
func AutoMigrate() {
	DB.AutoMigrate(&models.Category{})
}
