package main

import (
	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/schulke-214/hkp/api"
	"github.com/schulke-214/hkp/pkg/db"
)

func main() {
	db.AutoMigrate()

	router := gin.Default()
	api.RegisterCategories(router, "/categories")
	router.Run(":8080")
	defer db.DB.Close()
}
