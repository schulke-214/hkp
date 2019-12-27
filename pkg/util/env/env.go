package env

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	ginMode := os.Getenv("GIN_MODE")

	if ginMode != "" {
		gin.SetMode(ginMode)
	}
}
