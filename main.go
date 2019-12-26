package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"github.com/schulke-214/hkp/api"
	"github.com/schulke-214/hkp/pkg/db"
)

func main() {
	router := gin.Default()
	api.RegisterCategories(router)
	router.Run(":8080")

	defer db.Connection.Close()
}
