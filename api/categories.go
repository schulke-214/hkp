package api

import (
	"github.com/gin-gonic/gin"

	"github.com/schulke-214/hkp/pkg/db"
	"github.com/schulke-214/hkp/pkg/domain/models"
)

// RegisterCategories registers all category endpoints on a gin router under a given path
func RegisterCategories(router *gin.Engine, path string) {
	group := router.Group(path)

	group.GET("/lol", func(ctx *gin.Context) {

		cat := models.Category{
			Name: "ROFL",
		}

		db.DB.Create(&cat)

		ctx.JSON(200, cat)
	})
}
