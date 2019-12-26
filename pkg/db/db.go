package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // database driver for gorm

	"github.com/schulke-214/hkp/pkg/domain/models"
)

// Connection is the database connection
var Connection *gorm.DB

// DatabaseURI represents the database uri
type DatabaseURI struct {
	Host     string
	Dbname   string
	User     string
	Password string
}

// Serialize formats the database uri
func (uri *DatabaseURI) Serialize() string {
	return fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", uri.Host, uri.Dbname, uri.User, uri.Password)
}

// DatabaseURIFromEnv constructs the database uri from environment variables
func DatabaseURIFromEnv() *DatabaseURI {
	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")

	return &DatabaseURI{
		Host:     host,
		Dbname:   dbname,
		User:     user,
		Password: password,
	}
}

func init() {
	uri := DatabaseURIFromEnv()
	db, err := gorm.Open("postgres", uri.Serialize())

	if err != nil {
		panic("failed to connect to the database")
	}

	db.AutoMigrate(&models.Category{})

	Connection = db
}
