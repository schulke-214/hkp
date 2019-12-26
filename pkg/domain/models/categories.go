package models

import (
	"github.com/jinzhu/gorm"
)

// Category is the core datatype this app is built around.
type Category struct {
	gorm.Model
	Name string
}
